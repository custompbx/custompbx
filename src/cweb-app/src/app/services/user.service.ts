import {Injectable, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Store, select} from '@ngrx/store';
import {AppState, selectAuthState} from '../store/app.states';
import {LogOut} from '../store/auth/auth.actions';

@Injectable({
  providedIn: 'root'
})
export class UserService implements OnDestroy {

  public getState: Observable<any>;
  public getState$: Subscription;
  public isAuthenticated: false;
  public user = null;
  public errorMessage = null;

  constructor(
    private store: Store<AppState>
  ) {
    this.getState = this.store.pipe(select(selectAuthState));
    this.getState$ = this.getState.subscribe((state) => {
      this.isAuthenticated = state.isAuthenticated;
      this.user = state.user;
      this.errorMessage = state.errorMessage;
    });
  }

  ngOnDestroy() {
    this.getState$.unsubscribe();
  }

  public logOut(): void {
    this.store.dispatch(new LogOut);
  }
}
