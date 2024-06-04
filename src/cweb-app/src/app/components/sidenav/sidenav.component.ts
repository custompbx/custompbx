import {Component, OnInit} from '@angular/core';
import {Iuser} from '../../store/auth/auth.reducers';
import {UserService} from '../../services/user.service';
import {Subscription} from 'rxjs';
import {IsActiveMatchOptions, Router} from '@angular/router';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.css']
})
export class SidenavComponent implements OnInit {

  public menuItems = Array<IMenuItemExpand>();
  public user: Iuser;
  public getState$: Subscription;

  constructor(
    private userService: UserService,
    private router: Router,
  ) {
    this.user = this.userService.user;
  }

  ngOnInit() {
    this.getState$ = this.userService.getState.subscribe((state) => {
      this.user = state.user;
      this.menuItems = this.getMenuItems(this.user?.group_id);
    });
  }

  isRouteActive(route: string): boolean {
    if (!route) {
      return false
    }
    return this.router.isActive(route, <IsActiveMatchOptions>{paths: 'subset', queryParams: 'subset', fragment: 'ignored', matrixParams: 'ignored'});
  }
  getMenuItems(id): Array<IMenuItemExpand> {
    if (!id) {
      return [];
    }
    switch (this.user.group_id) {
      case 1:
        return [
          {
            route: '',
            icon: 'assessment',
            name: 'Monitoring',
            subMenu: [
              {
                route: '/dashboard',
                icon: 'navigate_next',
                name: 'Dashboard',
              },
              {
                route: 'monitoring/users-panel',
                icon: 'navigate_next',
                name: 'Users Panel',
              },
            ],
          },
          {
            route: '',
            icon: 'folder_shared',
            name: 'Directory',
            subMenu: [
              {
                route: 'directory/domains',
                icon: 'navigate_next',
                name: 'Domains',
              },
              {
                route: 'directory/users',
                icon: 'navigate_next',
                name: 'Users',
              },
              {
                route: 'directory/groups',
                icon: 'navigate_next',
                name: 'Groups',
              },
              {
                route: 'directory/gateways',
                icon: 'navigate_next',
                name: 'Gateways',
              },
            ],
          },
          {
            route: '',
            icon: 'settings',
            name: 'Configuration',
            subMenu: [
              {
                route: '/configuration/modules',
                icon: 'navigate_next',
                name: 'Modules',
              },
              {
                route: '/configuration/acl',
                icon: 'navigate_next',
                name: 'Acl',
              },
              {
                route: '/configuration/callcenter',
                icon: 'navigate_next',
                name: 'Callcenter',
              },
              {
                route: '/configuration/sofia',
                icon: 'navigate_next',
                name: 'Sofia',
              },
              {
                route: '/configuration/verto',
                icon: 'navigate_next',
                name: 'Verto',
              },
              {
                route: '/configuration/post-load-switch',
                icon: 'navigate_next',
                name: 'Switch',
              },
            ],
          },
          {
            route: '',
            icon: 'dialpad',
            name: 'Dialplan',
            subMenu: [
              {
                route: '/dialplan/contexts',
                icon: 'navigate_next',
                name: 'Contexts',
              },
            ]
          },
          {
            route: '/cdr',
            icon: 'find_in_page',
            name: 'CDR',
            subMenu: null,
          },
          {
            route: '/logs',
            icon: 'insert_drive_file',
            name: 'Logs',
            subMenu: null,
          },
          {
            route: '/fs-cli',
            icon: 'chevron_right',
            name: 'FS_CLI',
            subMenu: null,
          },
          {
            route: '/hep',
            icon: 'storage',
            name: 'HEP',
            subMenu: null,
          },
          {
            route: '/instances',
            icon: 'settings_input_component',
            name: 'Instances',
            subMenu: null,
          },
          {
            route: '/global-variables',
            icon: 'playlist_add_check',
            name: 'Global Variables',
            subMenu: null,
          }/*,
          {
            route: '',
            icon: 'apps',
            name: 'Apps',
            subMenu: [
              {
                route: '/apps/autodialer',
                icon: 'navigate_next',
                name: 'Autodialer',
              },
            ],
          }*/
        ];
      case 2:
        return [
            {
            route: 'monitoring/users-panel',
            icon: 'navigate_next',
            name: 'Users Panel',
              subMenu: null,
          },
          {
            route: '/cdr',
            icon: 'find_in_page',
            name: 'CDR',
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
  subMenu: Array<IMenuItem>;
}
