import { Component, OnInit } from '@angular/core';
import {NavigationEnd, Router} from '@angular/router';
import {UserService} from './services/user.service';
import {filter} from 'rxjs/operators';
import {select, Store} from "@ngrx/store";
import {AppState, selectApp} from "./store/app.states";
import {Observable, Subscription} from "rxjs";
import {ToggleShowConversations} from "./store/app/app.actions";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  login = true;
  public currentComponent;
  public showRightNav = false;
  public appState: Observable<any>;
  public state$: Subscription;

  constructor(
    private router: Router,
    public userService: UserService,
    private store: Store<AppState>,
  ) {
    this.appState = this.store.pipe(select(selectApp));
    router.events.pipe(
      filter(event => event instanceof NavigationEnd)
    ).subscribe((event: NavigationEnd) => {
      if (event.url === 'login' || event.url === '' || event.url === '/login' || event.url === '/') {
        this.login = true;
      } else {
        this.login = false;
      }
      /*      if (this.userService.isAuthenticated && event.url === '/login') {
              this.router.navigateByUrl('dashboard');
            }*/
      this.store.dispatch(ToggleShowConversations({showConversations: false}))
    });
    this.state$ = this.appState.subscribe((s) => {
      this.showRightNav = s.showConversations;
    });
  }

  ngOnInit() {
  }

  onRouterOutletActivate(event: any) {
    this.currentComponent = event;
  }
}
