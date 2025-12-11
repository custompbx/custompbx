import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { AppState } from '../../store/app.states';
import { LogIn } from '../../store/auth/auth.actions';
import {UserService} from '../../services/user.service';
import {Iuser} from '../../store/auth/auth.reducers';
import {MaterialModule} from "../../../material-module";

import {FormsModule} from "@angular/forms";



@Component({
standalone: true,
    imports: [MaterialModule, FormsModule],
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  user: Iuser = {};

  constructor(
    public userService: UserService,
    private store: Store<AppState>
  ) {
  }

  ngOnInit() {
  }

  onSubmit(): void {
    const payload = {
      login: this.user.login,
      password: this.user.password
    };
    this.store.dispatch(new LogIn(payload));
  }

}
