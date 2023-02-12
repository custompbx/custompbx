import {Component, OnInit} from '@angular/core';
import {WsDataService} from '../../services/ws-data.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import {select, Store} from '@ngrx/store';
import {AppState, selectAuthState, selectDaemonState} from '../../store/app.states';
import {Observable} from 'rxjs';
import {Status, StoreFlushTokenState} from '../../store/daemon/daemon.actions';
import {ReLogIn} from '../../store/auth/auth.actions';
import {CookiesStorageService} from '../../services/cookies-storage.service';
import {UserService} from '../../services/user.service';

@Component({
  selector: 'app-service-status',
  templateUrl: './service-status.component.html',
  styleUrls: ['./service-status.component.css']
})
export class ServiceStatusComponent implements OnInit {

  public noConnect: boolean;
  public noESL: boolean;
  public noDB: boolean;
  private tokenFailed: boolean;
  public daemon$: Observable<any>;
  private isAuthenticated: false;
  private authState: Observable<any>;
  private dDataSubscribed: boolean;

  constructor(
    private ws: WsDataService,
    private store: Store<AppState>,
    private cookie: CookiesStorageService,
    private _snackBar: MatSnackBar,
    public userService: UserService,
  ) {
    this.noConnect = false;
    this.noESL = false;
    this.noDB = false;
    this.tokenFailed = false;

    this.authState = this.store.pipe(select(selectAuthState));
  }

  ngOnInit() {
    this.ws.websocketService.status
      .subscribe((isConnected) => {
        if (typeof (isConnected) !== 'boolean') {
          return;
        }
        if (isConnected) {
          this.noConnect = false;
          this.trackDaemonState();
          if (this.cookie.getToken() && !this.isAuthenticated) {
            const payload = {
              route: window.location.pathname.replace('cweb/', ''),
            };
            this.store.dispatch(new ReLogIn(payload));
          } else {

          }
        } else {
          this._snackBar.open('No connection!', null, {duration: 5000});
          this.noConnect = true;
        }
      });
    this.authState.subscribe((state) => {
      this.isAuthenticated = state.isAuthenticated;
    });
  }

  trackDaemonState() {
    if (this.dDataSubscribed) {
      return;
    }
    this.ws.waitDaemonData().subscribe( payload => {
        this.store.dispatch(new Status(payload));
      }
    );
    this.daemon$ = this.store.pipe(select(selectDaemonState));
    this.daemon$.subscribe((data) => {
      this.noESL = !data.eslConnection;
      this.noDB = !data.dbConnection;
      if (!data.eslConnection || !data.dbConnection) {
        this._snackBar.open('No connection to FreeSwitch! DB: ' + data.dbConnection + '. ESL: ' + data.eslConnection + '.',
          null, {duration: 10000});
      }
      if (!data.tokenFailed) {
        this.tokenFailed = false;
      }
      if (data.dbConnection && data.tokenFailed && !this.tokenFailed) {
        this.tokenFailed = true;
        this.userService.logOut();
        // this.store.dispatch(new StoreFlushTokenState(null));
      }
    });
    this.dDataSubscribed = true;
  }
}
