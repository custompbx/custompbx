import { Component, inject, DestroyRef, effect, computed, signal } from '@angular/core';
import { toSignal, takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { select, Store } from '@ngrx/store';
import { MatSnackBar } from '@angular/material/snack-bar';

import { AppState, selectAuthState, selectDaemonState } from '../../store/app.states';
import { Status } from '../../store/daemon/daemon.actions';
import { ReLogIn } from '../../store/auth/auth.actions';
import { CookiesStorageService } from '../../services/cookies-storage.service';
import { UserService } from '../../services/user.service';
import { WsDataService } from '../../services/ws-data.service';
import { MaterialModule } from "../../../material-module";
import { filter } from 'rxjs';

@Component({
  standalone: true,
  imports: [MaterialModule],
  selector: 'app-service-status',
  templateUrl: './service-status.component.html',
  styleUrls: ['./service-status.component.css']
})
export class ServiceStatusComponent {

  // Injectable services
  private ws = inject(WsDataService);
  private store = inject(Store<AppState>);
  private cookie = inject(CookiesStorageService);
  private _snackBar = inject(MatSnackBar);
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
  public noConnect = computed(() => !this.isConnected());
  public noESL = computed(() => !this.daemonState().eslConnection);
  public noDB = computed(() => !this.daemonState().dbConnection);
  public tokenFailed = signal(false);

  // Flag to ensure we only subscribe to daemon data once
  private dDataSubscribed: boolean = false;

  constructor() {
    // Effect to handle WebSocket connection status changes
    effect(() => {
      const connected = this.isConnected();

      if (connected) {
        this.trackDaemonState();
        if (this.cookie.getToken() && !this.isAuthenticated()) {
          const payload = {
            route: window.location.pathname.replace('cweb/', ''),
          };
          this.store.dispatch(new ReLogIn(payload));
        }
      } else {
        if (connected === false) {
          this._snackBar.open('No connection!', null, { duration: 5000 });
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
        this._snackBar.open(
          'No connection to FreeSwitch! DB: ' + (data.dbConnection ? 'OK' : 'FAIL') +
          '. ESL: ' + (data.eslConnection ? 'OK' : 'FAIL') + '.',
          null, { duration: 10000 }
        );
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
      return;
    }
    this.ws.waitDaemonData().pipe(
      takeUntilDestroyed(this.destroyRef)
    ).subscribe(payload => {
      this.store.dispatch(new Status(payload));
    });
    this.dDataSubscribed = true;
  }
}
