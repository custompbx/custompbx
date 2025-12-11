import {Injectable, inject, computed} from '@angular/core';
import {Store, select} from '@ngrx/store';
import {AppState, selectAuthState} from '../store/app.states';
import {LogOut} from '../store/auth/auth.actions';
import {toSignal} from "@angular/core/rxjs-interop";

interface AuthState {
  isAuthenticated: boolean;
  user: any; // Use a proper user type if available
  errorMessage: any; // Use a proper error type if available
  token: any;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private store = inject(Store<AppState>);

  private authStateSignal = toSignal(
    this.store.pipe(select(selectAuthState)),
    { initialValue: { isAuthenticated: false, user: null, errorMessage: null , token: null} as AuthState }
  );

  public isAuthenticatedSignal = computed(() => this.authStateSignal().isAuthenticated);
  public userSignal = computed(() => this.authStateSignal().user);
  public errorMessageSignal = computed(() => this.authStateSignal().errorMessage);

  public get isAuthenticated(): boolean {
    return this.isAuthenticatedSignal();
  }

  public get user(): any {
    return this.userSignal();
  }

  public get errorMessage(): any {
    return this.errorMessageSignal();
  }

  public logOut(): void {
    this.store.dispatch(new LogOut);
  }
}
