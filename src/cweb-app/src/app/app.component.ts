import {Component, DestroyRef, inject} from '@angular/core';
import {NavigationEnd, Router, RouterOutlet} from '@angular/router';
import {BreakpointObserver} from '@angular/cdk/layout';
import {takeUntilDestroyed, toSignal} from '@angular/core/rxjs-interop';
import {filter, map} from 'rxjs/operators';
import {UserService} from './services/user.service';
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
export class AppComponent {

  login = true;
  public currentComponent: unknown;
  public showRightNav = false;
  public menuOpen = true;
  public readonly compact = toSignal(
    this.breakpointObserver.observe('(max-width: 1023px)').pipe(map(result => result.matches)),
    {initialValue: false}
  );

  private readonly destroyRef = inject(DestroyRef);

  constructor(
    private router: Router,
    public userService: UserService,
    private breakpointObserver: BreakpointObserver,
  ) {
    router.events
      .pipe(
        filter((event): event is NavigationEnd => event instanceof NavigationEnd),
        takeUntilDestroyed(this.destroyRef)
      )
      .subscribe(event => {
          this.login = event.urlAfterRedirects.startsWith('/login');
          this.showRightNav = false;
          if (this.compact()) this.menuOpen = false;
      });
  }

  onRouterOutletActivate(event: unknown): void {
    this.currentComponent = event;
  }

  toggleRightSideNav(): void {
    this.showRightNav = !this.showRightNav;
  }

  toggleNavigation(): void {
    this.menuOpen = !this.menuOpen;
  }

  onNavigationOpenedChange(opened: boolean): void {
    // Desktop keeps the drawer mounted and only changes its width. Updating the
    // collapsed state from a desktop close event (including the login screen)
    // made the first authenticated view start collapsed unexpectedly.
    if (this.compact()) this.menuOpen = opened;
  }
}
