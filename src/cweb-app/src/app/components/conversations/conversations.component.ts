import {
  Component, effect,
  ElementRef, HostListener, inject,
  ViewChild, signal, computed, DestroyRef, Signal
} from '@angular/core';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {debounceTime, filter, map, Subject, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {
  AppState,
  selectConversations, selectDirectoryState,
  selectPhoneState,
  selectSettingsState
} from '../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {StoreCommand} from '../../store/phone/phone.actions';
import {WsDataService} from '../../services/ws-data.service';
import {PersistentSubscription} from '../../store/dataFlow/dataFlow.actions';
import {GetWebUsers} from '../../store/settings/settings.actions';
import {
  GetConversationPrivateCalls,
  GetConversationPrivateMessages,
  GetNewConversationMessage,
  SendConversationPrivateCall,
  SendConversationPrivateMessage,
  StoreCurrentUser,
  StoreGetNewConversationMessage
} from '../../store/conversations/conversations.actions';
import {UserService} from '../../services/user.service';
import {StartPhone, ToggleShowPhone} from '../../store/header/header.actions';
import {GetDirectoryUsers} from '../../store/directory/directory.actions';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {FormsModule} from "@angular/forms";
import {InnerHeaderComponent} from "../inner-header/inner-header.component";
import {IwebUser} from "../../store/settings/settings.reducers";
import {FormatTimerPipe} from "../../pipes/format-timer.pipe";

const scrollTop = 64;

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, FormatTimerPipe],
  selector: 'app-conversations',
  templateUrl: './conversations.component.html',
  styleUrls: ['./conversations.component.css']
})
export class ConversationsComponent {

  @ViewChild('scrollContainer') scrollContainer!: ElementRef<HTMLElement>;

  // Injectable services
  private userService = inject(UserService);
  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);
  private ws = inject(WsDataService);
  private destroyRef = inject(DestroyRef);

  // NgRx State converted to Signals
  private webUsersState = toSignal(this.store.pipe(select(selectSettingsState)), {initialValue: {} as any});
  private phoneState = toSignal(this.store.pipe(select(selectPhoneState)), {initialValue: {} as any});
  private conversationsState = toSignal(this.store.pipe(select(selectConversations)), {initialValue: {} as any});
  private directoryState = toSignal(this.store.pipe(select(selectDirectoryState)), {initialValue: {} as any});

  // Derived/Computed Signals
  public user = this.userService.userSignal; // Already a signal
  public userList: Signal<{ [id: number]: IwebUser }> = computed(() => this.webUsersState().webUsers || {});
  public loadCounter = computed(() => this.webUsersState().loadCounter || 0);
  public webUsersErrorMessage = computed(() => this.webUsersState().errorMessage);

  // Phone State
  public phoneStatus = computed(() => this.phoneState().phoneStatus?.isRunning || false);
  public phoneUser = computed(() => this.phoneState().phoneCreds?.user_name || '');
  public isRegistered = computed(() => this.phoneState().phoneStatus?.registered || false);
  public totalTime = computed(() => this.phoneState().timer || 0);
  public inCall = computed(() => this.phoneState().phoneStatus?.inCall || false);
  public isRinging = computed(() => this.phoneState().phoneStatus?.status === 'ringing');

  // Conversations State
  public messagesSignal = computed(() => this.conversationsState().conversations || {});
  public callsSignal = computed(() => this.conversationsState().calls || {});
  public conversationsErrorMessage = computed(() => this.conversationsState().errorMessage);
  public scrollDown = computed(() => this.conversationsState().scrollDown || false);
  private eventData = computed(() => this.conversationsState().event?.data);

  // Directory State
  public directoryDomains = computed(() => this.directoryState().domains || {});
  private directoryUsersSignal = computed(() => this.directoryState().users || {});
  public directoryErrorMessage = computed(() => this.directoryState().errorMessage);

  // Component State (Now converted to Signals)
  public currentChat = signal<number | null>(null);
  public currentVoice = signal<number | null>(null);
  public lastCallsAmount = signal(0);
  public showItems = signal<{ [key: number]: any[] }>({});
  public toChat = signal(false);

  // Standard properties (UI/event handling)
  public searchUser = '';
  public newMsg = '';
  public fixedTopGapScrolled: number = scrollTop;
  public isInbound: boolean = false;
  public inConversationsCall: boolean = false;
  public isMouseOverChat: boolean = false;
  public isUpdatingChat: boolean = false;
  private previousScrollItemIndex: number | null = null;
  private wheelEvent$ = new Subject<WheelEvent>();
  private wheelSubscription: Subscription;

  // Effects for Side Effects and Derived UI Updates
  private menuUpdateEffect = effect(() => {
    const currentUser = this.user();
    this.store.dispatch(StoreCurrentUser({user: currentUser}));
  });

  private errorHandlingEffect = effect(() => {
    const errors = [
      this.webUsersErrorMessage(),
      this.conversationsErrorMessage(),
      this.directoryErrorMessage()
    ].filter(e => !!e);

    if (errors.length > 0) {
      this._snackBar.open('Error: ' + errors[0] + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  private phoneStatusEffect = effect(() => {
    // Logic from original phone$ subscription
    const inCall = this.inCall();
    const isRinging = this.isRinging();
    const isRegistered = this.isRegistered();

    // Reset isInbound flag
    if (this.isInbound && ((!isRinging && !inCall) || (inCall && this.currentChat() !== this.currentVoice()))) {
      this.isInbound = false;
    }

    // Reset inConversationsCall if main call drops
    if (!inCall && this.inConversationsCall) {
      this.inConversationsCall = false;
    }
  });


  private conversationUpdateEffect = effect(() => {
    const messages = this.messagesSignal();
    const calls = this.callsSignal();
    const currentScrollDown = this.scrollDown();
    const eventData = this.eventData();
    const currentChatId = this.currentChat();

    if (currentChatId !== null && messages[currentChatId]) {
      const chatMessages = messages[currentChatId];
      const chatCalls = calls[currentChatId] || [];
      const currentLastCallsAmount = this.lastCallsAmount();

      if (currentScrollDown) {
        this.scrollToBottom();
      } else {
        setTimeout(() => {
          this.restoreScrollPosition();
        }, 0);
      }
      this.isUpdatingChat = false;

      // Logic for merging messages and calls based on count/infinite scroll
      let newShowItems: any[] = [];
      if (chatMessages.length === 0) {
        newShowItems = chatCalls;
      } else if (chatMessages.length <= 20) {
        newShowItems = [
          ...chatMessages,
          ...chatCalls
        ].sort((a, b) => new Date(a.created_at).getTime() - new Date(b.created_at).getTime());
      } else if (chatMessages.length > 20) {
        const firstMes = chatMessages[0];
        const lastCall = chatCalls.length >= 20 ? chatCalls[chatCalls.length - 1] : null;

        if (lastCall && currentLastCallsAmount !== chatCalls.length && lastCall.created_at < firstMes.created_at) {
          this.lastCallsAmount.set(chatCalls.length);
          this.store.dispatch(GetConversationPrivateCalls({id: currentChatId, up_to_time: lastCall.created_at}));
        } else {
          newShowItems = [
            ...chatMessages,
            ...chatCalls.filter((a: any) => new Date(a.created_at).getTime() >= new Date(firstMes.created_at).getTime())
          ].sort((a, b) => new Date(a.created_at).getTime() - new Date(b.created_at).getTime());
        }
      }

      // Update showItems signal
      this.showItems.update(items => ({
        ...items,
        [currentChatId]: newShowItems
      }));
    }

    // Handle incoming new-call event (from mes.event)
    if (eventData?.type === 'new-call') {
      const data = eventData.data;
      if (this.user().id === data.rid) {
        if (currentChatId === data.sid) {
          this.voiceCall();
          this.isInbound = true;
        }
        if (!this.isRegistered()) {
          setTimeout(() => this.store.dispatch(StoreCommand({register: true})), 100);
        }
      }
    }
  });

  constructor() {
    this.store.dispatch(GetNewConversationMessage(null));
    this.setupWebSocketAndInitialData();

    // Setup wheel event handling subscription and ensure cleanup
    this.wheelSubscription = this.wheelEvent$
      .pipe(
        takeUntilDestroyed(this.destroyRef),
        debounceTime(100),
        filter(() => !!this.scrollContainer),
        filter(() => this.isMouseOverChat),
        filter(() => !this.isUpdatingChat),
        filter(() => !this.toChat()),
        map(() => {
          const element = this.scrollContainer.nativeElement;
          const currentChatId = this.currentChat();
          const items = this.showItems()[currentChatId!];

          if (currentChatId !== null && element && this.hasVerticalScrollbar(element) && element.scrollTop === 0 && items) {
            const message = items[0];
            // Capture scroll position before dispatching
            for (let i = 0; i < this.scrollContainer.nativeElement.children[0].children.length; i++) {
              const child = this.scrollContainer.nativeElement.children[0].children[i] as HTMLElement;
              if (child.offsetTop + child.offsetHeight > scrollTop) {
                this.previousScrollItemIndex = parseInt(child.getAttribute('data-index') || '0', 10);
                break;
              }
            }
            this.store.dispatch(GetConversationPrivateMessages({id: currentChatId, up_to_time: message.created_at}));
            this.isUpdatingChat = true;
          }
        })
      )
      .subscribe();
  }

  private setupWebSocketAndInitialData() {
    const persistentActions = [new GetWebUsers(null).type, StoreGetNewConversationMessage.type];

    const initializeData = () => {
      this.store.dispatch(new PersistentSubscription({values: persistentActions}));
      this.store.dispatch(new GetWebUsers(null));
      if (Object.entries(this.directoryUsersSignal() || {}).length === 0) {
        this.store.dispatch(new GetDirectoryUsers(null));
      }
    };

    if (this.ws.isConnected) {
      initializeData();
    }

    this.ws.websocketService.status.pipe(
      filter(connected => connected)
    ).subscribe(connected => {
      initializeData();
    });
  }

  @HostListener('wheel', ['$event'])
  onScroll(event: WheelEvent) {
    this.wheelEvent$.next(event);
  }

  @HostListener('mouseover', ['$event'])
  onMouseOver(event: MouseEvent) {
    if (this.scrollContainer && this.scrollContainer.nativeElement.contains(event.target as Node)) {
      this.isMouseOverChat = true;
    }
  }

  @HostListener('mouseout', ['$event'])
  onMouseOut(event: MouseEvent) {
    if (this.scrollContainer && this.scrollContainer.nativeElement.contains(event.target as Node)) {
      this.isMouseOverChat = false;
    }
  }

  connectToUser() {
    const currentChatId = this.currentChat();
    const currentVoiceId = this.currentVoice();

    if (this.isRinging()) {
      this.store.dispatch(StoreCommand({answer: true}));
      return;
    }
    if (this.inConversationsCall) {
      this.hangup();
      return;
    }
    if (currentVoiceId === null || !this.userList()[currentVoiceId]) {
      return;
    }

    const targetUser = this.userList()[currentVoiceId];
    if (!targetUser.sip_id?.Valid) {
      return;
    }

    const sipUser = this.directoryUsersSignal()[targetUser.sip_id.Int64];
    if (!sipUser) {
      return;
    }
    const domainName = this.directoryDomains()[sipUser.parent.id]?.name;
    const fullName = sipUser.name + '@' + domainName;

    this.store.dispatch(SendConversationPrivateCall({id: currentChatId!}));
    this.store.dispatch(StoreCommand({callTo: fullName}));
    this.inConversationsCall = true;
  }

  getLogins(filterString: string = ''): any[] {
    const userArray = Object.values(this.userList() || {});
    const filteredUsers = userArray.filter(user => user.login.includes(filterString));
    const sortedUsers = filteredUsers.sort((a, b) => a.id - b.id);
    return sortedUsers;
  }

  scrollToBottom() {
    if (this.scrollContainer && this.scrollContainer.nativeElement) {
      setTimeout(() => {
        this.scrollContainer.nativeElement.scrollTop = this.scrollContainer.nativeElement.scrollHeight;
      }, 0);
    }
  }

  sendMsg() {
    const currentChatId = this.currentChat();
    if (!this.newMsg || currentChatId === null) {
      return;
    }
    this.store.dispatch(SendConversationPrivateMessage({id: currentChatId, text: this.newMsg}));
    this.newMsg = '';
    this.scrollToBottom();
  }

  enterChat(user: { id: number }) {
    this.isInbound = false;
    this.previousScrollItemIndex = null;
    this.isUpdatingChat = false;

    if (this.currentChat() !== user.id) {
      this.currentChat.set(user.id);
      this.currentVoice.set(null);
      this.lastCallsAmount.set(0); // Reset for new chat
      this.store.dispatch(GetConversationPrivateMessages({id: user.id}));
      this.store.dispatch(GetConversationPrivateCalls({id: user.id}));
    }
    this.scrollToBottom();
  }

  convertDate(timestamp: string): string {
    const f = new Date(timestamp);
    const year = f.getFullYear().toString();
    const month = (f.getUTCMonth() + 1).toString().padStart(2, '0'); // month is 0-indexed
    const day = f.getDate().toString().padStart(2, '0');
    const date = `${year}-${month}-${day}`;

    const hours = f.getHours().toString().padStart(2, '0');
    const minutes = f.getMinutes().toString().padStart(2, '0');
    const time = `${hours}:${minutes}`;
    let res = date + ' ' + time;

    if (f.toDateString() === new Date().toDateString()) {
      res = time;
    }
    return res;
  }

  voiceCall() {
    this.store.dispatch(StartPhone(null));
    this.store.dispatch(ToggleShowPhone({show: false}));
    this.currentVoice.set(this.currentChat());
    this.toChat.set(true);
  }

  backToChat() {
    this.toChat.set(false);
    if (this.inCall()) {
      this.store.dispatch(ToggleShowPhone({show: true}));
    }
    setTimeout(() => this.scrollToBottom(), 0);
  }

  hasVerticalScrollbar(element: HTMLElement): boolean {
    return element.scrollHeight > element.clientHeight;
  }

  restoreScrollPosition() {
    if (this.previousScrollItemIndex === null || !this.scrollContainer) {
      return;
    }
    const children = this.scrollContainer.nativeElement.children[0]?.children;
    if (!children) return;

    for (let i = 0; i < children.length; i++) {
      const child = children[i] as HTMLElement;
      if (parseInt(child.getAttribute('data-index') || '0', 10) === this.previousScrollItemIndex) {
        // Scroll slightly above the element to restore context
        this.scrollContainer.nativeElement.scrollTop = child.offsetTop - child.offsetHeight;
        break;
      }
    }
    this.previousScrollItemIndex = null; // Clear after restoring
  }

  hangup() {
    this.store.dispatch(StoreCommand({hangup: true}));
  }
}
