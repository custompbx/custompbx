import {ChangeDetectionStrategy, Component, DestroyRef, HostListener, computed, effect, inject, signal} from '@angular/core';
import { toSignal, takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { select, Store } from '@ngrx/store';
import {ToastService} from '../../services/toast.service';

import { AppState, selectAuthState, selectDaemonState } from '../../store/app.states';
import { Status } from '../../store/daemon/daemon.actions';
import { ReLogIn } from '../../store/auth/auth.actions';
import { CookiesStorageService } from '../../services/cookies-storage.service';
import { UserService } from '../../services/user.service';
import { WsDataService } from '../../services/ws-data.service';
import { filter } from 'rxjs';
import {IconComponent} from '../icon/icon.component';
import {TranslocoPipe} from '@jsverse/transloco';

@Component({
  standalone: true,
  imports: [IconComponent, TranslocoPipe],
  selector: 'app-service-status',
  templateUrl: './service-status.component.html',
  styleUrls: ['./service-status.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ServiceStatusComponent {

  // Injectable services
  private ws = inject(WsDataService);
  private store = inject(Store<AppState>);
  private cookie = inject(CookiesStorageService);
  private _snackBar = inject(ToastService);
  public userService = inject(UserService);
  private destroyRef = inject(DestroyRef);

  // NgRx State converted to Signals
  private authState = toSignal(this.store.pipe(select(selectAuthState)), { initialValue: {} as any });
  private daemonState = toSignal(this.store.pipe(select(selectDaemonState)), { initialValue: {} as any });

  // WebSocket Connection Status Signal
  // We use filter(Boolean) to ensure we only get boolean true/false states.
  public isConnected = toSignal(this.ws.websocketService.status.pipe(
    filter(val => typeof val === 'boolean')
  ), { initialValue: null });

  // Component State Signals
  public isAuthenticated = computed(() => this.authState().isAuthenticated || false);
  public noConnect = computed(() => this.isConnected() === false);
  public noESL = computed(() => this.daemonStatusReceived() && !this.daemonState().eslConnection);
  public noDB = computed(() => this.daemonStatusReceived() && !this.daemonState().dbConnection);
  public tokenFailed = signal(false);
  public detailsOpen = signal(false);
  public daemonStatusReceived = signal(false);
  private hasConnectedOnce = false;
  private lastServiceWarning = '';

  public connectionState = computed<'connecting' | 'checking' | 'online' | 'degraded' | 'offline'>(() => {
    if (this.isConnected() === null) return 'connecting';
    if (this.isConnected() === false) return 'offline';
    if (!this.daemonStatusReceived()) return 'checking';
    if (this.noESL() || this.noDB()) return 'degraded';
    return 'online';
  });

  public statusSummary = computed(() => {
    switch (this.connectionState()) {
      case 'online': return 'All services online';
      case 'degraded': return 'Service connection issue';
      case 'offline': return 'Web app disconnected';
      case 'checking': return 'Checking services';
      default: return 'Connecting';
    }
  });

  public websocketLabel = computed(() => {
    if (this.isConnected() === null) return 'Connecting';
    return this.isConnected() ? 'Connected' : 'Disconnected';
  });

  public eslLabel = computed(() => this.serviceLabel(this.daemonState().eslConnection));
  public databaseLabel = computed(() => this.serviceLabel(this.daemonState().dbConnection));

  // Flag to ensure we only subscribe to daemon data once
  private dDataSubscribed: boolean = false;

  constructor() {
    // Effect to handle WebSocket connection status changes
    effect(() => {
      const connected = this.isConnected();

      if (connected) {
        this.hasConnectedOnce = true;
        this.trackDaemonState();
        if (this.cookie.getToken() && !this.isAuthenticated()) {
          const payload = {
            route: window.location.pathname.replace('cweb/', ''),
          };
          this.store.dispatch(new ReLogIn(payload));
        }
      } else {
        this.daemonStatusReceived.set(false);
        this.lastServiceWarning = '';
        if (this.hasConnectedOnce && connected === false) {
          this._snackBar.open('WebSocket connection lost. Reconnecting...', null, {
            duration: 7000,
            tone: 'error',
          });
        }
      }
    });

    // Effect to handle Daemon state changes (ESL/DB connection and Token Failure)
    effect(() => {
      const data = this.daemonState();
      // Ensure data object is meaningful before proceeding
      if (!data || Object.keys(data).length === 0) return;

      // SnackBar for FreeSwitch/DB connection issues
      if (this.noESL() || this.noDB()) {
        const warning = 'Backend service issue. Database: ' + (data.dbConnection ? 'connected' : 'unavailable') +
          '. FreeSWITCH ESL: ' + (data.eslConnection ? 'connected' : 'unavailable') + '.';
        if (warning !== this.lastServiceWarning) {
          this.lastServiceWarning = warning;
          this._snackBar.open(warning, null, { duration: 10000, tone: 'warning' });
        }
      } else {
        this.lastServiceWarning = '';
      }

      // Token failed logic
      if (!data.tokenFailed) {
        this.tokenFailed.set(false);
      }

      if (data.dbConnection && data.tokenFailed && !this.tokenFailed()) {
        this.tokenFailed.set(true);
        this.userService.logOut();
        // this.store.dispatch(new StoreFlushTokenState(null)); // Kept commented as in original
      }
    });
  }

  /**
   * Subscribes to the WebSocket to receive and dispatch daemon status updates.
   * This is only run once upon the first successful connection.
   */
  trackDaemonState() {
    if (this.dDataSubscribed) {
      this.ws.websocketService.send('get_status', {});
      return;
    }
    this.ws.waitDaemonData().pipe(
      takeUntilDestroyed(this.destroyRef)
    ).subscribe(payload => {
      this.daemonStatusReceived.set(true);
      this.store.dispatch(new Status(payload));
    });
    this.dDataSubscribed = true;
  }

  toggleDetails(): void {
    this.detailsOpen.update(open => !open);
  }

  @HostListener('document:keydown.escape')
  closeDetails(): void {
    this.detailsOpen.set(false);
  }

  private serviceLabel(available: boolean): string {
    if (!this.isConnected() || !this.daemonStatusReceived()) return 'Unknown';
    return available ? 'Connected' : 'Unavailable';
  }
}
