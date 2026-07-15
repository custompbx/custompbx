import {Component, DestroyRef, inject} from '@angular/core';
import {NavigationEnd, Router, RouterOutlet} from '@angular/router';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {filter} from 'rxjs/operators';
import {UserService} from './services/user.service';
import {ViewportService} from './services/viewport.service';

import {SidenavComponent} from "./components/sidenav/sidenav.component";
import {ServiceStatusComponent} from "./components/service-status/service-status.component";
import {HeaderComponent} from "./components/header/header.component";
import {ConversationsComponent} from "./components/conversations/conversations.component";
import {ToastContainerComponent} from './components/toast-container/toast-container.component';
import {ConfirmationDialogComponent} from './components/confirmation-dialog/confirmation-dialog.component';
import {IconSpriteService} from './services/icon-sprite.service';
import {WsDataService} from './services/ws-data.service';
import {ToastService} from './services/toast.service';

@Component({
  standalone: true,
  imports: [RouterOutlet, SidenavComponent, ServiceStatusComponent, HeaderComponent, ConversationsComponent, ToastContainerComponent, ConfirmationDialogComponent],
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent {

  login = true;
  public currentComponent: unknown;
  public showRightNav = false;
  public menuOpen = true;
  public readonly compact = this.viewport.compactNavigation;

  private readonly destroyRef = inject(DestroyRef);

  constructor(
    private router: Router,
    public userService: UserService,
    private viewport: ViewportService,
    iconSprite: IconSpriteService,
    wsData: WsDataService,
    toast: ToastService,
  ) {
    iconSprite.load();
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

    wsData.proceedMessageType('none')
      .pipe(
        filter(message => typeof message?.error === 'string' && message.error.length > 0),
        takeUntilDestroyed(this.destroyRef)
      )
      .subscribe(message => toast.error(message.error));
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

}
