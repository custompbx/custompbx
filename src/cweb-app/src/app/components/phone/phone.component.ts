import {Component, ElementRef, OnDestroy, OnInit, ViewChild} from '@angular/core';
import * as SIP from 'sip.js';
import {Verto} from 'vertojs/dist';

import {select, Store} from '@ngrx/store';
import {AppState, selectPhoneState} from '../../store/app.states';
import {Observable, Subscription} from 'rxjs';
import {AuthActionTypes, GetPhoneCreds, StorePhoneStatus, StoreTicker} from '../../store/phone/phone.actions';

@Component({
  selector: 'app-phone',
  templateUrl: './phone.component.html',
  styleUrls: ['./phone.component.css']
})
export class PhoneComponent implements OnInit, OnDestroy {
  public phoneData: Observable<any>;
  public phoneData$: Subscription;
  private creds: object;
  public showButtonsPad = true;
  public libName: string;
  public sipjsLib = 'sipjs';
  public vertoLib = 'verto';
  public debug: boolean;

  public data = {
    UA: <SIP.UserAgent>null,
    registerer: null,
    vertoUA: <Verto>null,
    session: <any>{}, // <SIP.InviteServerContext | SIP.InviteClientContext>null,
    answered: false,
    inCall: false,
    ringing: false,
    registered: false,
    onHold: false,
    number: '',
    timer: <any>null,
    totalTime: 0,
    startedAt: <Date>null,
    domain: '',
    uaParams: {
      uri: null,
      authorizationUsername: '',
      authorizationPassword: '',
      register: false,
      transportOptions: {
        server: '',
      },
      sessionDescriptionHandlerOptions: {
        constraints: {
          audio: true,
          video: false
        },
      },
      sessionDescriptionHandlerFactoryOptions: {
        iceGatheringTimeout: 1000,
        peerConnectionConfiguration: {
          iceServers: [{urls: 'stun:stun.l.google.com:19302'}],
        }
      },
      earlyMedia: true,
      delegate: {
        'onInvite': (i) => this.incoming.bind(this)(i),
        'onConnect': () => {
          console.log('connected');
        },
        'onDisconnect': () => {
          console.log('disconnected');
        },
        'onMessage': () => {
          console.log('message');
        },
      },
/*      contactParams: {
        'rinstance': this.makeid(9),
      },*/
    },

    vertoOptions: {
      transportConfig:
        {
          socketUrl: '',
          login: '',
          passwd: ''
        },
      rtcConfig: {
        iceServers: [{urls: 'stun:stun.l.google.com:19302'}],
      },
      debug: true,
      ice_timeout: 2000,
    },
    remoteTag: <HTMLAudioElement>null,
    localTag: <HTMLAudioElement>null,
    localStream: <MediaStream>null,
  };
  public audioContext = <AudioContext>null;
  public audioBuffer = {};
  public sourceList = {};
  public padButtons = ['1', '2', '3', '4', '5', '6', '7', '8', '9', '*', '0', '#'];

  @ViewChild('mediaTags', {static: false}) set mediaRef(ref: ElementRef) {
    this.data.remoteTag = ref.nativeElement.children.remoteTag;
    this.data.localTag = ref.nativeElement.children.localTag;
  }

  constructor(
    private store: Store<AppState>,
  ) {
    this.phoneData = this.store.pipe(select(selectPhoneState));
  }

  ngOnInit() {
    try {
      this.audioContext = new AudioContext();
    } catch (e) {
      alert('Web Audio API is not supported in this browser');
    }
    this.padButtons.forEach((b) => {
      if (b === '*') {
        b = 'star';
      } else if (b === '#') {
        b = 'pound';
      }
      this.preLoadSound('./assets/sounds/dtmf/Dtmf-' + b + '.wav', b);
    });
    this.preLoadSound('./assets/sounds/phone_2.wav', 'ring');
    this.phoneData$ = this.phoneData.subscribe((phone) => {
      if (phone.phoneCreds) {
        this.libName = phone.phoneCreds.webrtc_lib;
        this.data.domain = phone.phoneCreds.domain;
        this.data.uaParams.authorizationUsername = phone.phoneCreds.user_name || '';
        this.data.uaParams.authorizationPassword = phone.phoneCreds.password || '';
        this.data.uaParams.uri = SIP.UserAgent.makeURI('sip:' + phone.phoneCreds.user_name + '@' + phone.phoneCreds.domain);

        this.data.uaParams.transportOptions.server = phone.phoneCreds.ws;

        this.data.vertoOptions.transportConfig.passwd = phone.phoneCreds.password || '';
        this.data.vertoOptions.transportConfig.login = phone.phoneCreds.user_name + '@' + phone.phoneCreds.domain;
        this.data.vertoOptions.transportConfig.socketUrl = phone.phoneCreds.verto_ws;
        if (phone.phoneCreds.stun) {
          this.data.uaParams.sessionDescriptionHandlerFactoryOptions
            .peerConnectionConfiguration.iceServers = [{urls: phone.phoneCreds.stun}];
          this.data.vertoOptions.rtcConfig.iceServers = [{urls: phone.phoneCreds.stun}];
        }
      }
      if (phone.errorMessage) {
      }
      if (
        phone.lastActionName === AuthActionTypes.StoreGetPhoneCreds &&
        !this.data.UA
      ) {
        if (
          this.libName === this.sipjsLib &&
          this.data.uaParams.authorizationUsername !== '' &&
          this.data.uaParams.transportOptions.server.length > 0) {
          this.sipjsInit();
          if (this.data.UA) {
            this.store.dispatch(new StorePhoneStatus({phoneStatus: {isRunning: true}}));
          }
        } else if (
          this.libName === this.vertoLib &&
          this.data.vertoOptions.transportConfig.login !== '' &&
          this.data.vertoOptions.transportConfig.socketUrl !== '') {
          this.vertoInit();
          if (this.data.vertoUA) {
            this.store.dispatch(new StorePhoneStatus({phoneStatus: {isRunning: true}}));
          }
        }
      }
      if (phone.callTo) {
        this.panelCall(phone.callTo);
      }
      this.data.registered = phone.phoneStatus.registered;
      this.data.inCall = phone.phoneStatus.inCall;
      this.data.ringing = phone.phoneStatus.status === 'ringing';
      this.data.answered = phone.phoneStatus.status === 'answered';
      this.data.totalTime = phone.timer;
    });
    this.store.dispatch(new GetPhoneCreds(null));
  }

  ngOnDestroy() {
    this.hangup();
    if (this.data.UA) {
      this.data.UA.stop();
      this.store.dispatch(new StorePhoneStatus({phoneStatus: {isRunning: false}}));
    } else if (this.data.vertoUA) {
      this.data.vertoUA.logout();
      this.store.dispatch(new StorePhoneStatus({phoneStatus: {isRunning: false}}));
    }
  }

  headerColor(): string {
    switch (true) {
      case this.data.ringing:
      case this.data.inCall && !this.data.answered:
        return 'accent';
      case this.data.answered:
        return 'warn';
      default:
        return 'primary';
    }
  }

  restartUS() {
    if (this.data.UA) {
      this.data.UA.transport.disconnect();
      this.data.UA.stop();
      this.data.UA = null;
      this.store.dispatch(new StorePhoneStatus({phoneStatus: {isRunning: false}}));
    } else if (this.data.vertoUA) {
      this.data.vertoUA.logout();
      this.data.vertoUA = null;
      this.store.dispatch(new StorePhoneStatus({phoneStatus: {isRunning: false}}));
    }
    this.store.dispatch(new GetPhoneCreds(null));
  }

  sipjsInit() {
    this.data.UA = new SIP.UserAgent(this.data.uaParams);
    this.data.UA.start()
      .then(() => {
        console.log('connected');
      })
      .catch((error) => {
        console.error('failed to connect');
      });
  }

  sipjsCaller(dest) {
    const target = SIP.UserAgent.makeURI('sip:' + dest);
    const inviteOptions = {
      earlyMedia: true,
      requestDelegate: {
        onAccept: (response) => {},
        onReject: (response) => {},
        onInvite: () => {},
        onProgress: (e) => {},
        onTrying: () => {}
      },
      sessionDescriptionHandlerOptions: {
        constraints: {
          audio: true,
          video: false
        },
      },
      sessionDescriptionHandlerFactoryOptions: {
        iceGatheringTimeout: 1000,
      },
    };

    this.data.session = new SIP.Inviter(this.data.UA, target, inviteOptions);
    this.data.session.stateChange.addListener((state) => {
      switch (state) {
        case SIP.SessionState.Initial:
          break;
        case SIP.SessionState.Establishing:
          this.eventProgress();
          break;
        case SIP.SessionState.Established:
          this.eventAccepted();
          break;
        case SIP.SessionState.Terminating:
        case SIP.SessionState.Terminated:
          this.eventTerminated();
          break;
        default:
          throw new Error('Unknown session state.');
      }
    });

    this.data.session.invite(inviteOptions)
      .then((request) => {
        // call.talkToNumberField.html( this.data.session.outgoingInviteRequest.message.to.uri.normal.user);
      })
      .catch((error) => {
        this.eventTerminated();
      });
  }

  call() {
    if (this.data.number === '' || this.data.inCall) {
      return;
    }
    this.resetTimer();
    this.startTimer();
    if (this.libName === this.sipjsLib) {
      this.sipjsCaller(this.data.number + '@' + this.data.domain);
    } else if (this.vertoLib === this.vertoLib) {
      this.vertoCall(this.data.number);
    }
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: true}}));
  }

  answer() {
    this.sourceList['ring']?.stop();
    if (this.data.inCall || !this.data.session || !this.data.ringing) {
      return;
    }
    this.resetTimer();
    this.startTimer();
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: true, status: 'answered'}}));
    if (this.libName === this.sipjsLib) {
      this.data.session.accept();
    } else if (this.libName === this.vertoLib) {
      navigator.mediaDevices.getUserMedia({audio: true})
        .catch((e) => console.log('getUserMedia fail: ', e))
        .then((e: MediaStream) => {
          if (e) {
            this.data.localStream = e;
            this.data.session.answer(e.getTracks());
          }
        });
    }
  }

  hangup() {
    this.sourceList['ring']?.stop();
    if (!this.data.session) {
      this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: false, status: ''}}));
      return;
    }
    try {
      this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: false, status: ''}}));
      if (this.libName === this.sipjsLib) {
        switch (this.data.session.state) {
          case SIP.SessionState.Initial:
          case SIP.SessionState.Establishing:
            if (this.data.session instanceof SIP.Inviter) {
              this.data.session.cancel();
            } else {
              this.data.session.reject();
            }
            break;
          case SIP.SessionState.Established:
            this.data.session.bye();
            break;
        }
      } else if (this.libName === this.vertoLib) {
        // TODO: need to check status (on ice gathering!)
        this.data.session.hangup();
      }
    } catch (e) {
      console.error(e);
    }

    this.stopTimer();

    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: false, status: ''}}));
    this.data.session = null;
  }

  register() {
    if (!this.data.UA) {
      return;
    }
    if (!this.data.registerer || this.data.registerer.state === SIP.RegistererState.Terminated) {
      this.data.registerer = new SIP.Registerer(this.data.UA);
    }
    if (this.data.registerer.state === SIP.RegistererState.Registered) {
      this.data.registerer.unregister();
      return;
    }
    this.data.registerer.stateChange.addListener((newState) => {
      switch (newState) {
        case SIP.RegistererState.Registered:
          this.eventRegistered();
          break;
        case SIP.RegistererState.Unregistered:
          this.eventUnregistered();
          break;
        case SIP.RegistererState.Terminated:
          console.log('Terminated');
          break;
      }
    });

    this.data.registerer.register()
      .then(() => {
        this.eventRegistered();
      })
      .catch((error) => {
        this.eventUnregistered();
      });
  }

  hold() {
    if (this.data.inCall && this.data.session && this.data.answered) {
      if (this.libName === this.sipjsLib) {
        if (this.data.session.state !== SIP.SessionState.Established) {
          return;
        }
        this.data.session.invite({sessionDescriptionHandlerModifiers: []});
        this.setupRemoteMedia();
      } else {
        this.data.session.hold();
      }
      this.data.onHold = true;
    }
  }

  unhold() {
      if (this.data.inCall && this.data.session && this.data.answered) {
        if (this.libName === this.sipjsLib) {
          if (this.data.session.state !== SIP.SessionState.Established) {
            return;
          }
          this.data.session.invite({sessionDescriptionHandlerModifiers: [SIP.Web.holdModifier]});
          this.data.onHold = true;
        } else {
          this.data.session.unhold();
        }
      this.data.onHold = false;
    }
  }

  incoming(context) {
    if (this.data.inCall) {
      return;
    }
    this.playSound('ring');
    this.data.session = context;

    this.data.session.stateChange.addListener((state) => {
        switch (state) {
          case SIP.SessionState.Initial:
            break;
          case SIP.SessionState.Establishing:
            break;
          case SIP.SessionState.Established:
            this.eventAccepted();
            break;
          case SIP.SessionState.Terminating:
          case SIP.SessionState.Terminated:
            this.eventBye();
            break;
          default:
            this.sourceList['ring']?.stop();
        }
      }
    );
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {status: 'ringing'}}));
  }

  eventRegistered() {
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {registered: true}}));
  }

  eventUnregistered() {
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {registered: false}}));
    console.log('unRegistered');
  }

  eventProgress() {
    this.setupRemoteMedia();
    console.log('eventProgress');
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: true}}));
  }

  eventAccepted() {
    console.log('ACCEPTED');
    this.sourceList['ring']?.stop();
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: true, status: 'answered'}}));
    this.setupRemoteMedia.bind(this)();
  }

  eventFailed() {
    this.sourceList['ring']?.stop();
    console.log('eventFailed');
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: false, status: ''}}));
    this.stopTimer();
  }

  eventTerminated() {
    this.sourceList['ring']?.stop();
    console.log('eventTerminated');
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: false, status: ''}}));
    this.stopTimer();
  }

  eventBye() {
    this.sourceList['ring']?.stop();
    console.log('eventBye');
    if (this.data.inCall) {
      this.hangup();
    }
    this.data.onHold = false;
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: false, status: ''}}));
    this.stopStream();
    this.stopTimer();
  }

  setupRemoteMedia = () => {
    console.log('setupRemoteMedia');
    const remoteStream = new MediaStream();
    this.data.session.sessionDescriptionHandler.peerConnection.getReceivers().forEach((receiver) => {
      if (receiver.track) {
        remoteStream.addTrack(receiver.track);
      }
    });
    this.data.remoteTag.srcObject = remoteStream;
    this.data.remoteTag.play().catch(e => console.log(e));
  }

  phoneButton(arg) {
    let buttonName = arg;
    if (arg === '*') {
      buttonName = 'star';
    } else if (arg === '#') {
      buttonName = 'hash';
    }
    this.playSoundWithGain(buttonName, 0.2);
    if (this.data.answered) {
      this.dtmf(arg);
    } else {
      this.addNumber(arg);
    }
  }

  addNumber(arg): void {
/*    if (isNaN(arg)) {
      return;
    }
    let num = this.data.number;
    this.data.number = <number>Number(String(num) + String(arg));*/
    this.data.number = this.data.number + String(arg);
  }

  dtmf(arg): void {
    this.data.session.dtmf(arg);
  }

  startTimer(): void {
    this.data.startedAt = new Date();
    this.data.timer = setInterval(() => this.countdown(), 1000);
  }

  stopTimer(): void {
    clearInterval(this.data.timer);
    this.data.timer = null;
    this.data.startedAt = null;
  }

  resetTimer(): void {
    this.data.startedAt = null;
    clearInterval(this.data.timer);
    this.data.timer = null;
  }

  countdown(): void {
    this.store.dispatch(StoreTicker({date: this.data.startedAt?.toString()}))
  }

  removeLastDigit() {
    if (this.data.number === '') {
      return;
    }
    this.data.number = this.data.number.slice(0, -1);
  }

  panelCall(user: string) {
    if (user.length < 3 || this.data.inCall) {
      return;
    }
    this.resetTimer();
    this.startTimer();
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: true}}));
    if (this.libName === this.sipjsLib) {
      this.sipjsCaller(user);
    } else if (this.libName === this.vertoLib) {
      this.vertoCall(user);
    }
  }

  switchHideButtonsPad() {
    this.showButtonsPad = !this.showButtonsPad;
  }

  vertoInit() {
    this.data.vertoUA = new Verto(this.data.vertoOptions);
    this.data.vertoUA.login().catch((e) =>
      console.log('Access denied')
    );
    this.data.vertoUA.subscribeEvent('invite', this.vertoIncoming.bind(this));
  }

  vertoIncoming(call) {
    if (this.data.inCall) {
      return;
    }
    this.playSound('ring');
    this.data.session = call;
    this.data.session.subscribeEvent('track', this.vertoTrack.bind(this));
    this.data.session.subscribeEvent('answer', this.vertoAccepted.bind(this));
    this.data.session.subscribeEvent('bye', this.eventBye.bind(this));
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {status: 'ringing'}}));
  }

  vertoTrack(track) {
    if (track.kind !== 'audio') {
      return;
    }
    const remoteStream = new MediaStream();
    remoteStream.addTrack(track);
    this.data.remoteTag.srcObject = remoteStream;
  }

  vertoAccepted() {
    console.log('ACCEPTED');
    this.sourceList['ring']?.stop();
    this.store.dispatch(new StorePhoneStatus({phoneStatus: {inCall: true, status: 'answered'}}));
  }

  vertoCall(user: string) {
    let direction = this.data.number;
    if (user) {
      direction = user;
    }
    navigator.mediaDevices.getUserMedia({audio: true})
      .catch((e) => console.log('getUserMedia fail: ', e))
      .then((e: MediaStream) => {
        if (e) {
          this.data.localStream = e;
          this.data.session = this.data.vertoUA.call(e.getTracks(), direction.toString());
          this.data.session.subscribeEvent('track', this.vertoTrack.bind(this));
          this.data.session.subscribeEvent('answer', this.vertoAccepted.bind(this));
          this.data.session.subscribeEvent('bye', this.eventBye.bind(this));
        }
      });
  }

  stopStream() {
    if (!this.data.localStream) {
      return;
    }
    this.data.localStream.getAudioTracks().forEach(function (track) {
      track.stop();
    });

    this.data.localStream.getVideoTracks().forEach(function (track) {
      track.stop();
    });

    this.data.localStream = null;
  }

  makeid(length) {
    let result = '';
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    const charactersLength = characters.length;
    for (let i = 0; i < length; i++) {
      result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
  }

  isVeroLogged(): boolean {
    if (!this.data.vertoUA) {
      return false;
    }
    return this.data.vertoUA.isLogged();
  }

  preLoadSound(url, name) {
    const request = new XMLHttpRequest();
    request.open('GET', url, true);
    request.responseType = 'arraybuffer';
    request.onload = () => {
      this.audioContext.decodeAudioData(request.response, (buffer) => {
        this.audioBuffer[name] = buffer;
      }, null);
    };
    request.send();
  }

  playSoundWithGain(name, gain) {
    if (!this.audioBuffer[name]) {
      return;
    }
    if (this.sourceList[name]) {
      this.sourceList[name]?.stop();
    }
    if (gain < 0 || gain > 1 ) {
      gain = 1;
    }

    const source = this.audioContext.createBufferSource();
    source.buffer = this.audioBuffer[name];
    if (gain !== 1) {
      const gainNode = this.audioContext.createGain();
      gainNode.gain.value = gain;
      gainNode.connect(this.audioContext.destination);
      source.connect(gainNode);
    } else {
      source.connect(this.audioContext.destination);
    }
    source.start(0);
    this.sourceList[name] = source;
  }

  playSound(name) {
    this.playSoundWithGain(name, 1);
  }
}
