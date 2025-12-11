import { Component, ElementRef, OnDestroy, OnInit, ViewChild, inject, effect, signal, computed, ChangeDetectionStrategy } from '@angular/core';
import { MaterialModule } from "../../../material-module";
import * as SIP from 'sip.js';
// Assuming 'vertojs/dist' is correctly aliased in the build config or path is correct
import { Verto } from 'vertojs/dist';

import { select, Store } from '@ngrx/store';
import { AppState, selectPhoneState } from '../../store/app.states';
import { Observable } from 'rxjs';
import { AuthActionTypes, GetPhoneCreds, StorePhoneStatus, StoreTicker } from '../../store/phone/phone.actions';
import { toSignal } from '@angular/core/rxjs-interop';
import {State} from "../../store/phone/phone.reducers";
import {FormatTimerPipe} from "../../pipes/format-timer.pipe";
import {FormsModule} from "@angular/forms";

@Component({
  standalone: true,
  imports: [MaterialModule, FormatTimerPipe, FormsModule],
  selector: 'app-phone',
  templateUrl: './phone.component.html',
  styleUrls: ['./phone.component.css'],
  // Use OnPush strategy as we are using Signals for local state management
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PhoneComponent implements OnInit, OnDestroy {
  // --- Dependency Injection ---
  private store = inject(Store<AppState>);

  // --- Reactive View State (Signals) ---
  public inCall = signal(false);
  public ringing = signal(false);
  public answered = signal(false);
  public registered = signal(false);
  public onHold = signal(false);
  public number = signal('');
  public totalTime = signal(0);
  public authorizationUsername = signal('');

  // NEW: Signals to track external library connectivity state
  public isSipJsTransportConnected = signal(false);
  public isVertoClientLogged = signal(false);

  // Local mutable state for timers and start time (needs manual mutation, but triggers signal updates)
  private timerInstance: any = null;
  private startedAt: Date | null = null;

  // --- NgRx State Signal ---
  private phoneState = toSignal(
    this.store.pipe(select(selectPhoneState)),
    {
      initialValue: {
        phoneCreds: null,
        errorMessage: null,
        lastActionName: null,
        command: null,
        phoneStatus: { isRunning: false, registered: false, inCall: false, status: '' },
        timer: 0
      } as State
    }
  );

  // --- Non-Reactive/Config State ---
  public showButtonsPad = true;
  public sipjsLib = 'sipjs';
  public vertoLib = 'verto';
  public debug: boolean;
  public libName: string = '';
  public domain: string = '';

  // Core data structure holding external library instances and fixed configurations
  public data = {
    UA: <SIP.UserAgent>null,
    registerer: null,
    vertoUA: <Verto>null,
    session: <any>{}, // Current active session (SIP.Session, Verto Call)
    uaParams: {
      uri: null,
      authorizationUsername: '',
      authorizationPassword: '',
      register: false,
      transportOptions: { server: '' },
      sessionDescriptionHandlerOptions: { constraints: { audio: true, video: false } },
      sessionDescriptionHandlerFactoryOptions: {
        iceGatheringTimeout: 1000,
        peerConnectionConfiguration: { iceServers: [{ urls: 'stun:stun.l.google.com:19302' }] }
      },
      earlyMedia: true,
      delegate: {
        'onInvite': (i) => this.incoming.bind(this)(i),
        'onConnect': () => { console.log('connected'); this.isSipJsTransportConnected.set(true); },
        'onDisconnect': () => { console.log('disconnected'); this.isSipJsTransportConnected.set(false); },
        'onMessage': () => { console.log('message'); },
      },
    },
    vertoOptions: {
      transportConfig: { socketUrl: '', login: '', passwd: '' },
      rtcConfig: { iceServers: [{ urls: 'stun:stun.l.google.com:19302' }] },
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

  @ViewChild('mediaTags', { static: false }) set mediaRef(ref: ElementRef) {
    if (ref) {
      this.data.remoteTag = ref.nativeElement.children.remoteTag;
      this.data.localTag = ref.nativeElement.children.localTag;
    }
  }

  ngOnInit() {
    // Initialization for Web Audio API and sounds
    try {
      this.audioContext = new AudioContext();
    } catch (e) {
      console.error('Web Audio API is not supported in this browser', e);
    }
    this.padButtons.forEach((b) => {
      const name = b === '*' ? 'star' : b === '#' ? 'pound' : b;
      this.preLoadSound('./assets/sounds/dtmf/Dtmf-' + name + '.wav', name);
    });
    this.preLoadSound('./assets/sounds/phone_2.wav', 'ring');

    this.store.dispatch(new GetPhoneCreds(null));
  }

  /**
   * NEW: Computed signal to check if the active library is fully connected/logged in.
   * This replaces the non-reactive template logic (e.g., calling isConnected()).
   */
  public isClientConnected = computed(() => {
/*    if (this.libName === this.sipjsLib) {
      // Check the reactive SIP transport signal
      return this.isSipJsTransportConnected();
    } else if (this.libName === this.vertoLib) {
      // Check the reactive Verto logged signal
      return this.isVertoClientLogged();
    }
    return false;*/
    return this.isSipJsTransportConnected() || this.isVertoClientLogged();
  });

  /**
   * Effect replaces the manual subscription and handles all reactive logic
   * driven by the NgRx store state.
   */
  phoneDataEffect = effect(() => {
    const phone = this.phoneState();

    // 1. Update Credential and Connection Data
    if (phone.phoneCreds) {
      this.libName = phone.phoneCreds.webrtc_lib;
      this.domain = phone.phoneCreds.domain; // Use flat property
      this.data.uaParams.authorizationUsername = phone.phoneCreds.user_name || '';
      this.authorizationUsername.set(this.data.uaParams.authorizationUsername);
      this.data.uaParams.authorizationPassword = phone.phoneCreds.password || '';
      this.data.uaParams.uri = SIP.UserAgent.makeURI('sip:' + phone.phoneCreds.user_name + '@' + phone.phoneCreds.domain);
      this.data.uaParams.transportOptions.server = phone.phoneCreds.ws;

      this.data.vertoOptions.transportConfig.passwd = phone.phoneCreds.password || '';
      this.data.vertoOptions.transportConfig.login = phone.phoneCreds.user_name + '@' + phone.phoneCreds.domain;
      this.data.vertoOptions.transportConfig.socketUrl = phone.phoneCreds.verto_ws;

      if (phone.phoneCreds.stun) {
        this.data.uaParams.sessionDescriptionHandlerFactoryOptions
          .peerConnectionConfiguration.iceServers = [{ urls: phone.phoneCreds.stun }];
        this.data.vertoOptions.rtcConfig.iceServers = [{ urls: phone.phoneCreds.stun }];
      }
    }

    // 2. Initialization Logic (on credential load)
    if (
      phone.lastActionName === AuthActionTypes.StoreGetPhoneCreds &&
      !this.data.UA && !this.data.vertoUA && phone.phoneCreds
    ) {
      const creds = phone.phoneCreds;
      if (
        this.libName === this.sipjsLib &&
        creds.user_name && creds.ws
      ) {
        this.sipjsInit();
        if (this.data.UA) {
          this.store.dispatch(new StorePhoneStatus({ phoneStatus: { isRunning: true } }));
        }
      } else if (
        this.libName === this.vertoLib &&
        creds.user_name && creds.verto_ws
      ) {
        this.vertoInit();
        if (this.data.vertoUA) {
          this.store.dispatch(new StorePhoneStatus({ phoneStatus: { isRunning: true } }));
        }
      }
    }

    // 3. Handle Commands
    if (phone.command?.callTo) {
      this.panelCall(phone.command.callTo);
    }
    if (phone.command?.hangup) {
      this.hangup();
    }
    if (phone.command?.answer) {
      this.answer();
    }
    if (phone.command?.register) {
      this.register();
    }

    // 4. Update Local Status Signals from NgRx
    this.registered.set(phone.phoneStatus.registered);
    this.inCall.set(phone.phoneStatus.inCall);
    this.ringing.set(phone.phoneStatus.status === 'ringing');
    this.answered.set(phone.phoneStatus.status === 'answered');
    this.totalTime.set(phone.timer);
  });

  ngOnDestroy() {
    this.hangup();
    if (this.data.UA) {
      this.data.UA.stop();
      this.isSipJsTransportConnected.set(false); // Cleanup signal
      this.store.dispatch(new StorePhoneStatus({ phoneStatus: { isRunning: false } }));
    } else if (this.data.vertoUA) {
      this.data.vertoUA.logout();
      this.isVertoClientLogged.set(false); // Cleanup signal
      this.store.dispatch(new StorePhoneStatus({ phoneStatus: { isRunning: false } }));
    }
  }

  // --- Computed Property for Header Color ---
  headerColor = computed((): string => {
    switch (true) {
      case this.ringing():
      case this.inCall() && !this.answered():
        return 'accent';
      case this.answered():
        return 'warn';
      default:
        return 'primary';
    }
  });

  restartUS() {
    if (this.data.UA) {
      this.data.UA.transport.disconnect();
      this.data.UA.stop();
      this.data.UA = null;
      this.isSipJsTransportConnected.set(false); // Cleanup signal
      this.store.dispatch(new StorePhoneStatus({ phoneStatus: { isRunning: false } }));
    } else if (this.data.vertoUA) {
      this.data.vertoUA.logout();
      this.data.vertoUA = null;
      this.isVertoClientLogged.set(false); // Cleanup signal
      this.store.dispatch(new StorePhoneStatus({ phoneStatus: { isRunning: false } }));
    }
    this.store.dispatch(new GetPhoneCreds(null));
  }

  sipjsInit() {
    this.data.UA = new SIP.UserAgent(this.data.uaParams);
    this.data.UA.start()
      .then(() => {
        console.log('UA started.');
      })
      .catch((error) => {
        console.error('failed to connect', error);
      });
  }

  sipjsCaller(dest: string) {
    const target = SIP.UserAgent.makeURI('sip:' + dest);
    const inviteOptions = {
      earlyMedia: true,
      requestDelegate: {
        onAccept: (response) => { },
        onReject: (response) => { },
        onInvite: () => { },
        onProgress: (e) => { },
        onTrying: () => { }
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
    if (this.number() === '' || this.inCall()) {
      return;
    }
    this.resetTimer();
    this.startTimer();
    if (this.libName === this.sipjsLib) {
      this.sipjsCaller(this.number() + '@' + this.domain);
    } else if (this.vertoLib === this.vertoLib) {
      this.vertoCall(this.number());
    }
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: true } }));
  }

  answer() {
    this.sourceList['ring']?.stop();
    if (this.inCall() || !this.data.session || !this.ringing()) {
      return;
    }
    this.resetTimer();
    this.startTimer();
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: true, status: 'answered' } }));
    if (this.libName === this.sipjsLib) {
      this.data.session.accept();
    } else if (this.libName === this.vertoLib) {
      navigator.mediaDevices.getUserMedia({ audio: true })
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
      this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: false, status: '' } }));
      return;
    }
    try {
      this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: false, status: '' } }));
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
        this.data.session.hangup();
      }
    } catch (e) {
      console.error(e);
    }

    this.stopTimer();

    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: false, status: '' } }));
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
    if (this.inCall() && this.data.session && this.answered()) {
      if (this.libName === this.sipjsLib) {
        if (this.data.session.state !== SIP.SessionState.Established) {
          return;
        }
        this.data.session.invite({ sessionDescriptionHandlerModifiers: [SIP.Web.holdModifier] });
        // NOTE: SIP.Web.holdModifier is usually for hold, and the empty array for unhold.
        // I am correcting the logic based on the original structure but flag the potential issue.
        this.setupRemoteMedia();
      } else {
        this.data.session.hold();
      }
      this.onHold.set(true);
    }
  }

  unhold() {
    if (this.inCall() && this.data.session && this.answered()) {
      if (this.libName === this.sipjsLib) {
        if (this.data.session.state !== SIP.SessionState.Established) {
          return;
        }
        this.data.session.invite({ sessionDescriptionHandlerModifiers: [] });
      } else {
        this.data.session.unhold();
      }
      this.onHold.set(false);
    }
  }

  incoming(context) {
    if (this.inCall()) {
      return;
    }
    this.playSound('ring');
    this.data.session = context;

    this.data.session.stateChange.addListener((state) => {
      switch (state) {
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
    });
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { status: 'ringing' } }));
  }

  eventRegistered() {
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { registered: this.data.registerer.state === SIP.RegistererState.Registered } }));
  }

  eventUnregistered() {
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { registered: this.data.registerer.state === SIP.RegistererState.Registered } }));
  }

  eventProgress() {
    this.setupRemoteMedia();
    console.log('eventProgress');
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: true } }));
  }

  eventAccepted() {
    console.log('ACCEPTED');
    this.sourceList['ring']?.stop();
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: true, status: 'answered' } }));
    this.setupRemoteMedia.bind(this)();
  }

  eventFailed() {
    this.sourceList['ring']?.stop();
    console.log('eventFailed');
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: false, status: '' } }));
    this.stopTimer();
  }

  eventTerminated() {
    this.sourceList['ring']?.stop();
    console.log('eventTerminated');
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: false, status: '' } }));
    this.stopTimer();
  }

  eventBye() {
    this.sourceList['ring']?.stop();
    console.log('eventBye');
    if (this.inCall()) {
     // this.hangup();
    }
    this.onHold.set(false);
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: false, status: '' } }));
    this.stopStream();
    this.stopTimer();
  }

  setupRemoteMedia = () => {
    console.log('setupRemoteMedia');
    try {
      const remoteStream = new MediaStream();
      this.data.session.sessionDescriptionHandler.peerConnection.getReceivers().forEach((receiver) => {
        if (receiver.track) {
          remoteStream.addTrack(receiver.track);
        }
      });
      if (this.data.remoteTag) {
        this.data.remoteTag.srcObject = remoteStream;
        this.data.remoteTag.play().catch(e => console.log('Remote playback failed:', e));
      }
    } catch (e) {
      console.error('Failed to setup remote media:', e);
    }
  }

  phoneButton(arg) {
    let buttonName = arg;
    if (arg === '*') {
      buttonName = 'star';
    } else if (arg === '#') {
      buttonName = 'hash';
    }
    this.playSoundWithGain(buttonName, 0.2);
    if (this.answered()) {
      this.dtmf(arg);
    } else {
      this.addNumber(arg);
    }
  }

  addNumber(arg): void {
    // Update the signal directly
    this.number.update(currentNum => currentNum + String(arg));
  }

  dtmf(arg): void {
    this.data.session.dtmf(arg);
  }

  startTimer(): void {
    this.startedAt = new Date();
    // Use the component's private timer instance
    this.timerInstance = setInterval(() => this.countdown(), 1000);
  }

  stopTimer(): void {
    clearInterval(this.timerInstance);
    this.timerInstance = null;
    this.startedAt = null;
  }

  resetTimer(): void {
    this.startedAt = null;
    clearInterval(this.timerInstance);
    this.timerInstance = null;
  }

  countdown(): void {
    // Send the update to NgRx, which will eventually update the totalTime signal via the effect
    this.store.dispatch(StoreTicker({ date: this.startedAt?.toString() }));
  }

  removeLastDigit() {
    if (this.number() === '') {
      return;
    }
    // Update the signal directly
    this.number.update(currentNum => currentNum.slice(0, -1));
  }

  panelCall(user: string) {
    if (user.length < 3 || this.inCall()) {
      return;
    }
    this.resetTimer();
    this.startTimer();
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: true } }));
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
    this.data.vertoUA.login()
      .then(() => {
        this.isVertoClientLogged.set(true); // SET SIGNAL on successful login
        console.log('Verto Login Success.');
      })
      .catch((e) => {
        this.isVertoClientLogged.set(false); // SET SIGNAL on failed login
        console.log('Access denied', e);
      });

    this.data.vertoUA.subscribeEvent('invite', this.vertoIncoming.bind(this));
  }

  vertoIncoming(call) {
    if (this.inCall()) {
      return;
    }
    this.playSound('ring');
    this.data.session = call;
    this.data.session.subscribeEvent('track', this.vertoTrack.bind(this));
    this.data.session.subscribeEvent('answer', this.vertoAccepted.bind(this));
    this.data.session.subscribeEvent('bye', this.eventBye.bind(this));
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { status: 'ringing' } }));
  }

  vertoTrack(track) {
    if (track.kind !== 'audio') {
      return;
    }
    const remoteStream = new MediaStream();
    remoteStream.addTrack(track);
    if (this.data.remoteTag) {
      this.data.remoteTag.srcObject = remoteStream;
    }
  }

  vertoAccepted() {
    console.log('ACCEPTED');
    this.sourceList['ring']?.stop();
    this.store.dispatch(new StorePhoneStatus({ phoneStatus: { inCall: true, status: 'answered' } }));
  }

  vertoCall(user: string) {
    let direction = this.number();
    if (user) {
      direction = user;
    }
    navigator.mediaDevices.getUserMedia({ audio: true })
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

  // NOTE: isVeroLogged() removed as it is non-reactive. Use isClientConnected() instead.

  // Web Audio API methods (non-reactive)
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
    if (gain < 0 || gain > 1) {
      gain = 1;
    }

    const source = this.audioContext.createBufferSource();
    source.buffer = this.audioBuffer[name];

    const gainNode = this.audioContext.createGain();
    gainNode.gain.value = gain;
    gainNode.connect(this.audioContext.destination);
    source.connect(gainNode);

    source.start(0);
    this.sourceList[name] = source;
  }

  playSound(name) {
    this.playSoundWithGain(name, 1);
  }
}
