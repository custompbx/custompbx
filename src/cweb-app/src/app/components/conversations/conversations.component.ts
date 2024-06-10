import {
  Component,
  ElementRef, HostListener,
  OnDestroy,
  OnInit,
  ViewChild
} from '@angular/core';
import {debounceTime, Observable, Subject, Subscription, take} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {
  AppState,
  selectConversations, selectDirectoryState, selectHeader,
  selectPhoneState,
  selectSettingsState
} from '../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {StoreCommand} from '../../store/phone/phone.actions';
import {WsDataService} from "../../services/ws-data.service";
import {SubscriptionList} from "../../store/dataFlow/dataFlow.actions";
import {GetWebUsers} from "../../store/settings/settings.actions";
import {
  GetConversationPrivateCalls,
  GetConversationPrivateMessages,
  GetNewConversationMessage,
  SendConversationPrivateCall,
  SendConversationPrivateMessage,
  StoreCurrentUser,
  StoreGetNewConversationMessage
} from "../../store/conversations/conversations.actions";
import {Iuser} from "../../store/auth/auth.reducers";
import {UserService} from "../../services/user.service";
import {StartPhone, ToggleShowPhone} from "../../store/header/header.actions";
import {filter, map, switchMap, tap} from "rxjs/operators";
import {GetDirectoryUsers} from "../../store/directory/directory.actions";
import {ToggleShowConversations} from "../../store/app/app.actions";

const scrollTop = 64;

@Component({
  selector: 'app-conversations',
  templateUrl: './conversations.component.html',
  styleUrls: ['./conversations.component.css']
})
export class ConversationsComponent implements OnInit, OnDestroy {

  public webUsers: Observable<any>;
  public pmessages: Observable<any>;
  public pmessages$: Subscription;
  public webUsers$: Subscription;
  public userList: { [key: string]: any };
  private lastErrorMessage: string;
  public loadCounter: number;
  public timersIntervalUpdater: any;
  public phone: Observable<any>;
  public phone$: Subscription;
  public phoneStatus: boolean;
  public phoneUser: string;
  public fixedTopGapScrolled: number = scrollTop;
  public searchUser = '';
  public newMsg = '';
  public currentChat = null;
  public currentVoice = null;
  public toChat = false;
  public messages = {};
  public calls = {};
  public lastCallsAmount = 0;
  public showItems = {};
  public getState$: Subscription;
  public user: Iuser;
  public totalTime: 0;
  public inCall: boolean;
  public inConversationsCall: boolean;
  public isMouseOverChat: boolean;
  private wheelEvent$ = new Subject<WheelEvent>();
  public isUpdatingChat: boolean;
  private previousScrollItemIndex: number | null = null;
  public directory$: Subscription;
  public directory: Observable<any>;
  public directoryDomains: object;
  private directoryUsers: object;
  public isInbound: boolean;
  public isRinging: boolean;
  public isRegistered: boolean;

  @ViewChild('scrollContainer') scrollContainer: ElementRef;

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
    private ws: WsDataService,
    private userService: UserService,
  ) {
    this.phone = this.store.pipe(select(selectPhoneState));
    this.webUsers = this.store.pipe(select(selectSettingsState));
    this.pmessages = this.store.pipe(select(selectConversations));
    this.directory = this.store.pipe(select(selectDirectoryState));
  }

  ngOnInit() {
    this.store.dispatch(GetNewConversationMessage(null))
    if (this.ws.isConnected) {
      this.store.dispatch(new SubscriptionList({values: [new GetWebUsers(null).type, StoreGetNewConversationMessage.type]}));
      this.store.dispatch(new GetWebUsers(null));
      if (Object.entries(this.directoryUsers || {}).length === 0) {
        this.store.dispatch(new GetDirectoryUsers(null));
      }
    }

    this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new SubscriptionList({values: [new GetWebUsers(null).type, StoreGetNewConversationMessage.type]}));
        this.store.dispatch(new GetWebUsers(null));
        if (Object.entries(this.directoryUsers || {}).length === 0) {
          this.store.dispatch(new GetDirectoryUsers(null));
        }
      }
    });
    this.getState$ = this.userService.getState.subscribe((state) => {
      this.user = state.user;
      this.store.dispatch(StoreCurrentUser({user: this.user}));
    });

    this.webUsers$ = this.webUsers.subscribe((users) => {
      this.loadCounter = users.loadCounter;
      this.userList = users.webUsers;
      this.lastErrorMessage = users.errorMessage;
      if (this.lastErrorMessage) {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    // PHONE
    this.phone$ = this.phone.subscribe((phone) => {
      this.phoneStatus = phone.phoneStatus.isRunning;
      if (phone.phoneCreds) {
        this.phoneUser = phone.phoneCreds.user_name || '';
      }
      this.isRegistered = phone.phoneStatus.registered;
      this.totalTime = phone.timer;
      if (this.isInbound && (
        (this.isRinging && phone.phoneStatus.status === '') ||
        (this.inCall && this.inCall !== phone.phoneStatus.inCall))
      ) {
        this.isInbound = false;
      }
      this.inCall = phone.phoneStatus.inCall;
      this.isRinging = phone.phoneStatus.status === 'ringing';
      if (!this.inCall && this.inConversationsCall) {
        this.inConversationsCall = false;
      }
    });

    this.pmessages$ = this.pmessages.subscribe((mes) => {
      this.messages = mes ? mes.conversations : {};
      this.calls = mes ? mes.calls : {};
      this.lastErrorMessage = mes ? mes.errorMessage : null;
      if (this.lastErrorMessage) {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
      if (this.messages[this.currentChat]) {
        if (mes.scrollDown) {
          this.scrollToBottom();
        } else {
          setTimeout(() => {
            this.restoreScrollPosition()
          }, 0)
        }
        this.isUpdatingChat = false;
        if (this.messages[this.currentChat].length === 0) {
          this.showItems[this.currentChat] = mes.calls[this.currentChat] || [];
        } else if (this.messages[this.currentChat].length <= 20) {
          this.showItems[this.currentChat] = [...this.messages[this.currentChat], ...(mes.calls[this.currentChat] || [])].sort((a, b) => new Date(a.created_at).getTime() - new Date(b.created_at).getTime());
        } else if (this.messages[this.currentChat].length > 20) {
          const firstMes = this.messages[this.currentChat][0]
          const lastCall = this.calls[this.currentChat].length >= 20 ? this.calls[this.currentChat][this.calls[this.currentChat].length - 1] : null;
          if (lastCall && this.lastCallsAmount !== this.calls[this.currentChat].length && lastCall.created_at < firstMes.created_at) {
            this.lastCallsAmount = this.calls[this.currentChat].length;
            this.store.dispatch(GetConversationPrivateCalls({id: this.currentChat, up_to_time: lastCall.created_at}));
          } else {
            this.showItems[this.currentChat] = [...this.messages[this.currentChat], ...(mes.calls[this.currentChat] || []).filter((a) => a.created_at >= firstMes.created_at)].sort((a, b) => new Date(a.created_at).getTime() - new Date(b.created_at).getTime());
          }
        }
      }
      if (mes.event.type === 'new-call' && this.user.id === mes.event.data.rid) {
        if (this.currentChat === mes.event.data.sid) {
          this.voiceCall()
          this.isInbound = true;
        } else {
          this.snackHandler(this.userList[mes.event.data.sid]?.login + ': Incoming call.', mes.event.data.sid);
        }
        if (!this.isRegistered) {
          setTimeout(() => this.store.dispatch(StoreCommand({register: true})), 100);
        }
      }
      if (mes.event.type === 'new-message' && this.user.id === mes.event.data.rid && this.currentChat !== mes.event.data.sid) {
        this.snackHandler(this.userList[mes.event.data.sid]?.login + ': ' + mes.event.data.text.slice(0, 25), mes.event.data.sid);
      }
    });

    //directory users
    this.directory$ = this.directory.subscribe((users) => {
      this.directoryDomains = users.domains;
      this.directoryUsers = users.users;
      if (users.errorMessage) {
        this._snackBar.open('Error: ' + users.errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });

    this.wheelEvent$
      .pipe(
        debounceTime(100), // Debounce to avoid too many events in a short time
        filter(() => !!this.scrollContainer), // Ensure the scroll container is available
        filter(() => this.isMouseOverChat), // Ensure the mouse is over the chat
        filter(() => !this.isUpdatingChat),
        map(() => {
          const element = this.scrollContainer.nativeElement;
          if (element && this.hasVerticalScrollbar(element) && element.scrollTop === 0 && this.showItems[this.currentChat]) {
            const message = this.showItems[this.currentChat][0];
            for (let i = 0; i < this.scrollContainer.nativeElement.children[0].children.length; i++) {
              const child = this.scrollContainer.nativeElement.children[0].children[i];
              if (child.offsetTop + child.offsetHeight > scrollTop) {
                this.previousScrollItemIndex = parseInt(child.getAttribute('data-index'), 10);
                break;
              }
            }
            this.store.dispatch(GetConversationPrivateMessages({id: this.currentChat, up_to_time: message.created_at}));
          }
        })
      )
      .subscribe();
  }

  @HostListener('wheel', ['$event'])
  onScroll(event: WheelEvent) {
    this.wheelEvent$.next(event);
  }

  @HostListener('mouseover', ['$event'])
  onMouseOver(event: MouseEvent) {
    if (this.scrollContainer && this.scrollContainer.nativeElement.contains(event.target)) {
      this.isMouseOverChat = true;
    }
  }

  @HostListener('mouseout', ['$event'])
  onMouseOut(event: MouseEvent) {
    if (this.scrollContainer && this.scrollContainer.nativeElement.contains(event.target)) {
      this.isMouseOverChat = false;
    }
  }

  ngOnDestroy() {
    this.timersIntervalUpdater = null;
    this.webUsers$.unsubscribe();
    this.phone$.unsubscribe();
    this.pmessages$.unsubscribe();
    this.getState$.unsubscribe();
    this.wheelEvent$.complete();
    if (this.route.snapshot?.data?.reconnectUpdater) {
      this.route.snapshot.data.reconnectUpdater.unsubscribe();
    }
  }

  connectToUser() {
    if (this.isRinging) {
      this.store.dispatch(StoreCommand({answer: true}));
      return;
    }
    if (this.inConversationsCall) {
      this.hangup()
      return;
    }
    if (!this.userList[this.currentVoice]) {
      return;
    }
    if (!this.userList[this.currentVoice].sip_id?.Valid) {
      return;
    }
    const sipUser = this.directoryUsers[this.userList[this.currentVoice].sip_id.Int64];
    const domainName = this.directoryDomains[sipUser.parent.id]?.name;
    const fullName = sipUser.name + '@' + domainName;
    if (!this.directoryUsers[this.userList[this.currentVoice].sip_id.Int64]) {
      return;
    }
    this.store.dispatch(SendConversationPrivateCall({id: this.currentChat}));
    this.store.dispatch(StoreCommand({callTo: fullName}));
    this.inConversationsCall = true;
  }

  getLogins(filterString: string = ''): any[] {
    const userArray = Object.values(this.userList || {});
    const filteredUsers = userArray.filter(user => user.login.includes(filterString));
    const sortedUsers = filteredUsers.sort((a, b) => a.id - b.id);
    //return sortedUsers.map(user => ({ id: user.id, login: user.login }));
    return sortedUsers
  }

  scrollToBottom() {
    if (this.scrollContainer && this.scrollContainer.nativeElement) {
      setTimeout(() => {
        this.scrollContainer.nativeElement.scrollTop = this.scrollContainer.nativeElement.scrollHeight;
      }, 0);
    }
  }

  sendMsg() {
    /*    if (!this.showItems[this.currentChat]) {
          return
        }*/
    if (!this.newMsg) {
      return
    }
    this.store.dispatch(SendConversationPrivateMessage({id: this.currentChat, text: this.newMsg}));
    this.newMsg = '';
    this.scrollToBottom();
  }

  enterChat(user) {
    this.isInbound = false;
    this.previousScrollItemIndex = null;
    this.isUpdatingChat = false;
    if (this.currentChat !== user.id) {
      this.currentChat = user.id;
      this.currentVoice = null;
      this.store.dispatch(GetConversationPrivateMessages({id: this.currentChat}));
      this.store.dispatch(GetConversationPrivateCalls({id: this.currentChat}));
    }
    this.scrollToBottom();
  }

  convertDate(timestamp) {
    const f = new Date(timestamp);
    const year = f.getFullYear().toString();
    const month = f.getUTCMonth().toString().padStart(2, '0');
    const day = f.getUTCDay().toString().padStart(2, '0');
    let date = `${year}-${month}-${day}`;

    const hours = f.getHours().toString().padStart(2, '0');
    const minutes = f.getMinutes().toString().padStart(2, '0');
    const time = `${hours}:${minutes}`;
    let res = date + " " + time;

    if (f.toDateString() == new Date().toDateString()) {
      res = time
    }
    return res
  }

  voiceCall() {
    this.store.dispatch(StartPhone(null))
    this.store.dispatch(ToggleShowPhone({show: false}))
    this.currentVoice = this.currentChat;
    this.toChat = true;
  }

  backToChat() {
    this.toChat = false;
    if (this.inCall) {
      this.store.dispatch(ToggleShowPhone({show: true}))
    }
    setTimeout(() => this.scrollToBottom(), 0)
  }

  hasVerticalScrollbar(element: HTMLElement): boolean {
    return element.scrollHeight > element.clientHeight;
  }

  restoreScrollPosition() {
    if (this.previousScrollItemIndex === null) {
      return;
    }
    const children = this.scrollContainer.nativeElement.children[0].children;
    for (let i = 0; i < children.length; i++) {
      let child = children[i];
      if (parseInt(child.getAttribute('data-index'), 10) === this.previousScrollItemIndex) {
        this.scrollContainer.nativeElement.scrollTop = child.offsetTop - child.scrollHeight * 2;
        break;
      }
    }
  }

  hangup() {
    this.store.dispatch(StoreCommand({hangup: true}));
  }
  snackHandler(msg, id) {
    let ref = this._snackBar.open(msg, 'Show', {
      duration: 5000,
      panelClass: ['mat-blue'],
      horizontalPosition: 'end',
      verticalPosition: 'bottom',
    });
    ref.onAction().subscribe(() => {
      this.store.dispatch(ToggleShowConversations({showConversations: true}));
      this.enterChat({id});
    });
  }
}
