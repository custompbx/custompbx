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
  selectConversations, selectHeader,
  selectPhoneState,
  selectSettingsState
} from '../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {StoreMakePhoneCall} from '../../store/phone/phone.actions';
import {WsDataService} from "../../services/ws-data.service";
import {SubscriptionList} from "../../store/dataFlow/dataFlow.actions";
import {GetWebUsers} from "../../store/settings/settings.actions";
import {
  GetConversationPrivateMessages, GetNewConversationMessage, SendConversationPrivateMessage, StoreCurrentUser,
  StoreGetNewConversationMessage
} from "../../store/conversations/conversations.actions";
import {Iuser} from "../../store/auth/auth.reducers";
import {UserService} from "../../services/user.service";
import {StartPhone, ToggleShowPhone} from "../../store/header/header.actions";
import {filter, map, switchMap, tap} from "rxjs/operators";

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
  public getState$: Subscription;
  public user: Iuser;
  public totalTime: 0;
  public inCall: boolean;
  public isMouseOverChat: boolean;
  private wheelEvent$ = new Subject<WheelEvent>();
  public isUpdatingChat: boolean;
  private previousScrollItemIndex: number | null = null;

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
  }

  ngOnInit() {
    this.store.dispatch(GetNewConversationMessage(null))
    if (this.ws.isConnected) {
      this.store.dispatch(new SubscriptionList({values: [new GetWebUsers(null).type, StoreGetNewConversationMessage.type]}));
      this.store.dispatch(new GetWebUsers(null));
    }

    this.ws.websocketService.status.subscribe(connected => {
      if (connected) {
        this.store.dispatch(new SubscriptionList({values: [new GetWebUsers(null).type, StoreGetNewConversationMessage.type]}));
        this.store.dispatch(new GetWebUsers(null));
      }
      this.getState$ = this.userService.getState.subscribe((state) => {
        this.user = state.user;
        this.store.dispatch(StoreCurrentUser({user: this.user}));
      });
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
    this.phone$ = this.phone.subscribe((phone) => {
      this.phoneStatus = phone.phoneStatus.isRunning;
      if (phone.phoneCreds) {
        this.phoneUser = phone.phoneCreds.user_name || '';
      }
      this.totalTime = phone.timer;
      this.inCall = phone.phoneStatus.inCall;
    });
    this.pmessages$ = this.pmessages.subscribe((mes) => {
      this.messages = mes ? mes.conversations : {};
      this.lastErrorMessage =  mes ? mes.errorMessage : null;
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
          setTimeout(() => {this.restoreScrollPosition()},0)
        }
        this.isUpdatingChat = false;
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
          if (element && this.hasVerticalScrollbar(element) && element.scrollTop === 0 && this.messages[this.currentChat]) {
            const message = this.messages[this.currentChat][0];
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

  connectToUser(user: object, domainName: string) {
    if (!this.phoneStatus || this.phoneUser === user['name'] || user['name'] === '') {
      return;
    }
    const fullName = user['name'] + '@' + domainName;

    if (user['sip_register'] || user['verto_register']) {
      this.store.dispatch(new StoreMakePhoneCall({user: user['name']}));
    }
  }

  getLogins(filterString: string = ''): any[] {
    const userArray = Object.values(this.userList);
    const filteredUsers = userArray.filter(user => user.login.includes(filterString));
    const sortedUsers = filteredUsers.sort((a, b) => a.id - b.id);
    //return sortedUsers.map(user => ({ id: user.id, login: user.login }));
    return sortedUsers
  }

  scrollToBottom() {
    if (this.scrollContainer) {
      setTimeout(() => {this.scrollContainer.nativeElement.scrollTop = this.scrollContainer.nativeElement.scrollHeight;}, 0);
    }
  }

  sendMsg(){
/*    if (!this.messages[this.currentChat]) {
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
    this.isUpdatingChat = false;
    if (this.currentChat !== user.id) {
      this.currentChat = user.id;
      this.store.dispatch(GetConversationPrivateMessages({id: this.currentChat}));
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

    if (f.toDateString() == new Date().toDateString()){
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
          this.scrollContainer.nativeElement.scrollTop = child.offsetTop - child.scrollHeight*2;
          break;
        }
      }
  }
}
