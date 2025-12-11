import { Component, OnInit } from '@angular/core';
import {NavigationEnd, Router, RouterOutlet} from '@angular/router';
import {UserService} from './services/user.service';
import {filter} from 'rxjs/operators';
import {MaterialModule} from "../material-module";

import {SidenavComponent} from "./components/sidenav/sidenav.component";
import {ServiceStatusComponent} from "./components/service-status/service-status.component";
import {HeaderComponent} from "./components/header/header.component";
import {ConversationsComponent} from "./components/conversations/conversations.component";

@Component({
  standalone: true,
  imports: [MaterialModule, RouterOutlet, SidenavComponent, ServiceStatusComponent, HeaderComponent, ConversationsComponent],
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  login = true;
  public currentComponent;
  public showRightNav = false;

  constructor(
    private router: Router,
    public userService: UserService,
  ) {
    router.events
      .pipe(filter(event => event instanceof NavigationEnd))
      .subscribe((event: NavigationEnd) => {
          this.login = event.urlAfterRedirects.startsWith('/login');
          this.showRightNav = false;
      });
  }

  ngOnInit() {
  }

  onRouterOutletActivate(event: any) {
    this.currentComponent = event;
  }

  toggleRightSideNav($event) {
    this.showRightNav = !this.showRightNav;
  }
}
