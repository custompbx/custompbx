import {Component, effect, inject, Input} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';

import {Iuser} from '../../store/auth/auth.reducers';
import {UserService} from '../../services/user.service';
import {IsActiveMatchOptions, NavigationEnd, Router, RouterLink} from '@angular/router';
import {filter, map, startWith} from 'rxjs/operators';
import {IconComponent} from '../icon/icon.component';
import {TranslocoPipe} from '@jsverse/transloco';

@Component({
  standalone: true,
  imports: [RouterLink, IconComponent, TranslocoPipe],
    selector: 'app-sidenav',
    templateUrl: './sidenav.component.html',
    styleUrls: ['./sidenav.component.css']
})
export class SidenavComponent {

  @Input() collapsed = false;

  public menuItems = Array<IMenuItemExpand>();
  public user: Iuser;

  private userService = inject(UserService);
  private router = inject(Router);
  private readonly currentUrl = toSignal(
    this.router.events.pipe(
      filter((event): event is NavigationEnd => event instanceof NavigationEnd),
      map((event) => event.urlAfterRedirects),
      startWith(this.router.url),
    ),
    {initialValue: this.router.url},
  );

  private menuUpdateEffect = effect(() => {
    // a) Read the user signal (userSignal is public in the UserService)
    const user = this.userService.userSignal();
    // b) Perform the side effect logic
    this.user = user;
    this.menuItems = this.getMenuItems(user?.group_id);
  });

  isRouteActive(route: string): boolean {
    this.currentUrl();
    if (!route) {
      return false;
    }
    return this.router.isActive(this.normalizeRoute(route), <IsActiveMatchOptions>{paths: 'subset', queryParams: 'subset', fragment: 'ignored', matrixParams: 'ignored'});
  }

  isMenuSectionActive(item: IMenuItemExpand): boolean {
    return item?.subMenu?.some((subItem) => this.isRouteActive(subItem.route)) ?? false;
  }

  private normalizeRoute(route: string): string {
    return route.startsWith('/') ? route : `/${route}`;
  }
  getMenuItems(id: number | null | undefined): Array<IMenuItemExpand> {
    if (!id) {
      return [];
    }
    switch (id) {
      case 1:
        return [
          {
            route: '',
            icon: 'assessment',
            name: 'navigation.monitoring',
            subMenu: [
              {
                route: '/dashboard',
                icon: 'navigate_next',
                name: 'navigation.dashboard',
              },
              {
                route: 'monitoring/users-panel',
                icon: 'navigate_next',
                name: 'navigation.usersPanel',
              },
            ],
          },
          {
            route: '',
            icon: 'folder_shared',
            name: 'navigation.directory',
            subMenu: [
              {
                route: 'directory/domains',
                icon: 'navigate_next',
                name: 'navigation.domains',
              },
              {
                route: 'directory/users',
                icon: 'navigate_next',
                name: 'navigation.users',
              },
              {
                route: 'directory/groups',
                icon: 'navigate_next',
                name: 'navigation.groups',
              },
              {
                route: 'directory/gateways',
                icon: 'navigate_next',
                name: 'navigation.gateways',
              },
            ],
          },
          {
            route: '',
            icon: 'settings',
            name: 'navigation.configuration',
            subMenu: [
              {
                route: '/configuration/modules',
                icon: 'navigate_next',
                name: 'navigation.modules',
              },
              {
                route: '/configuration/acl',
                icon: 'navigate_next',
                name: 'navigation.acl',
              },
              {
                route: '/configuration/callcenter',
                icon: 'navigate_next',
                name: 'navigation.callcenter',
              },
              {
                route: '/configuration/sofia',
                icon: 'navigate_next',
                name: 'navigation.sofia',
              },
              {
                route: '/configuration/verto',
                icon: 'navigate_next',
                name: 'navigation.verto',
              },
              {
                route: '/configuration/post-load-switch',
                icon: 'navigate_next',
                name: 'navigation.switch',
              },
            ],
          },
          {
            route: '',
            icon: 'dialpad',
            name: 'navigation.dialplan',
            subMenu: [
              {
                route: '/dialplan/contexts',
                icon: 'navigate_next',
                name: 'navigation.contexts',
              },
            ]
          },
          {
            route: '/cdr',
            icon: 'find_in_page',
            name: 'navigation.cdr',
            subMenu: null,
          },
          {
            route: '/logs',
            icon: 'insert_drive_file',
            name: 'navigation.logs',
            subMenu: null,
          },
          {
            route: '/fs-cli',
            icon: 'chevron_right',
            name: 'navigation.fsCli',
            subMenu: null,
          },
          {
            route: '/hep',
            icon: 'storage',
            name: 'navigation.hep',
            subMenu: null,
          },
          {
            route: '/instances',
            icon: 'settings_input_component',
            name: 'navigation.instances',
            subMenu: null,
          },
          {
            route: '/global-variables',
            icon: 'playlist_add_check',
            name: 'navigation.globalVariables',
            subMenu: null,
          }
        ];
      case 2:
        return [
            {
            route: 'monitoring/users-panel',
            icon: 'navigate_next',
            name: 'navigation.usersPanel',
              subMenu: null,
          },
          {
            route: '/cdr',
            icon: 'find_in_page',
            name: 'navigation.cdr',
            subMenu: null,
          },
        ];
      case 3:
        return [];
    }
  }
}

export interface IMenuItem {
  icon: string;
  route: string;
  name: string;
}

export interface IMenuItemExpand {
  icon: string;
  route: string;
  name: string;
  subMenu: Array<IMenuItem> | null;
}
