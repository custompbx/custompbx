import {
  Attribute,
  Component, ComponentRef,
  OnInit,
  ViewChild,
  ViewContainerRef
} from '@angular/core';
import {NavigationEnd, Router} from '@angular/router';
import {filter} from 'rxjs/operators';

@Component({
  selector: 'app-lazy-wrapper',
  templateUrl: './lazy-wrapper.component.html',
})
export class LazyWrapperComponent implements OnInit {
  @ViewChild('lazyContent', { read: ViewContainerRef, static: true }) lazyContentContainer: ViewContainerRef;

  private type:
    'phone' |
    'cdr' |
    'logs' |
    'settings' |
    'fs-cli' |
    'instances' |
    'global-variables' |
    'hep' |
    'domains' |
    'users' |
    'groups' |
    'gateways' |
    'modules' |
    'acl' |
    'callcenter' |
    'sofia' |
    'verto' |
    'cdr-pg-csv' |
    'odbc-cdr' |
    'lcr' |
    'shout' |
    'redis' |
    'nibblebill' |
    'avmd' |
    'cdr-mongodb' |
    'db' |
    'http-cache' |
    'memcache' |
    'opus' |
    'python' |
    'tts-commandline' |
    'alsa' |
    'amr' |
    'amrwb' |
    'cepstral' |
    'cidlookup' |
    'curl' |
    'dialplan-directory' |
    'easyroute' |
    'erlang-event' |
    'event-multicast' |
    'fax' |
    'lua' |
    'mongo' |
    'msrp' |
    'oreka' |
    'perl' |
    'pocketsphinx' |
    'sangoma-codec' |
    'sndfile' |
    'xml-cdr' |
    'xml-rpc' |
    'zeroconf' |
    'post-load-switch' |
    'distributor' |
    'directory' |
    'fifo' |
    'opal' |
    'osp' |
    'unicall' |
    'contexts' |
    'users-panel' |
    'conference' |
    'post-load-modules' |
    'voicemail' |
    'autodialer' |
    'not-found';
  private prePath: string;

  private componentRef = null;

  constructor(
              @Attribute('type') private atrType,
              private router: Router,
              private cfr: ViewContainerRef
  ) {
    this.prePath = '';
    if (atrType) {
      this.type = atrType;
    } else {
      router.events.pipe(
        filter(event => event instanceof NavigationEnd)
      ).subscribe((event: NavigationEnd) => {
        const pathItems = event.url.split('/').filter(Boolean);
        if (pathItems.length > 1) {
          this.prePath = pathItems[0] + '/';
        }
        const pathName = pathItems.slice(-1)[0];
        switch (pathName) {
          case 'cdr':
          case 'logs':
          case 'settings':
          case 'fs-cli':
          case 'instances':
          case 'global-variables':
          case 'domains':
          case 'hep':
          case 'users':
          case 'groups':
          case 'gateways':
          case 'modules':
          case 'acl':
          case 'callcenter':
          case 'sofia':
          case 'verto':
          case 'cdr-pg-csv':
          case 'odbc-cdr':
          case 'lcr':
          case 'shout':
          case 'redis':
          case 'nibblebill':
          case 'avmd':
          case 'cdr-mongodb':
          case 'db':
          case 'http-cache':
          case 'memcache':
          case 'opus':
          case 'python':
          case 'tts-commandline':
          case 'alsa':
          case 'amr':
          case 'amrwb':
          case 'cepstral':
          case 'cidlookup':
          case 'curl':
          case 'dialplan-directory':
          case 'easyroute':
          case 'erlang-event':
          case 'event-multicast':
          case 'fax':
          case 'lua':
          case 'mongo':
          case 'msrp':
          case 'oreka':
          case 'perl':
          case 'pocketsphinx':
          case 'sangoma-codec':
          case 'sndfile':
          case 'xml-cdr':
          case 'xml-rpc':
          case 'zeroconf':
          case 'post-load-switch':
          case 'distributor':
          case 'directory':
          case 'fifo':
          case 'opal':
          case 'osp':
          case 'unicall':
          case 'contexts':
          case 'users-panel':
          case 'conference':
          case 'post-load-modules':
          case 'voicemail':
          case 'autodialer':
            this.type = pathName;
            break;
          default:
            this.type = 'not-found';
            console.log(pathItems);
        }
      });
    }
  }

  async ngOnInit() {
    if (!this.type) {
      return;
    }
    const lazyContentComponent = await import(`../${this.prePath}${this.type}/${this.type}.component`);
    const componentClass = lazyContentComponent[`${this.capitalize(this.type)}Component`];
    this.componentRef = this.cfr.createComponent(componentClass);
  }

  private capitalize(value: string): string {
    return `${value.charAt(0).toUpperCase()}${value.slice(1).toLowerCase().replace(/([-_].)/g, function (x) { return x[1].toUpperCase(); })}`;
  }

  public getChildComponentFactory(): ComponentRef<any> {
    return this.componentRef;
  }

}
