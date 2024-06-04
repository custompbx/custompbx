import { Component, OnInit } from '@angular/core';
import {NavigationEnd, Router} from '@angular/router';
import {UserService} from './services/user.service';
import {filter} from 'rxjs/operators';

@Component({
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
      this.showRightNav = false;
    });
  }

  ngOnInit() {
  }

  onRouterOutletActivate(event: any) {
    this.currentComponent = event;
  }

  toggleRightSideNav($event) {
    this.showRightNav = !this.showRightNav
  }
}
