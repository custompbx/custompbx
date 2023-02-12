import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {State} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {
  AutoloadModule,
  FromScratchConfModule, ImportAllModules,
  ImportConfModule, ImportXMLModuleConfig, LoadModule,
  ReloadModule,
  SwitchModule, TruncateModuleConfig,
  UnloadModule
} from '../../../store/config/config.actions';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';

@Component({
  selector: 'app-modules',
  templateUrl: './modules.component.html',
  styleUrls: ['./modules.component.css']
})
export class ModulesComponent implements OnInit, OnDestroy {

  public autoload: object;
  public configs: Observable<any>;
  public configs$: Subscription;
  public list: State;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public XMLBody: string;
  public editorInited: boolean;

  constructor(
    private store: Store<AppState>,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
    private bottomSheet: MatBottomSheet,
  ) {
    this.selectedIndex = 0;
    this.configs = this.store.pipe(select(selectConfigurationState));
    this.autoload = {};
  }

  ngOnInit() {
    this.configs$ = this.configs.subscribe((configs) => {
      this.loadCounter = configs.loadCounter;
      this.list = configs;
      this.lastErrorMessage = configs.errorMessage;

      if (this.list.post_load_modules && this.list.post_load_modules.modules) {
        Object.keys(this.list.post_load_modules.modules).forEach(
          e => {
            this.autoload[this.list.post_load_modules.modules[e].name] = this.list.post_load_modules.modules[e];
          }
        );
      }

      if (!this.lastErrorMessage) {

      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  reloadModule(module: number) {
    this.store.dispatch(new ReloadModule({id: module}));
  }

  unloadModule(module: number) {
    this.store.dispatch(new UnloadModule({id: module}));
  }

  loadModule(module: number) {
    this.store.dispatch(new LoadModule({id: module}));
  }

  switchModule(module: number, enabled: boolean) {
    this.store.dispatch(new SwitchModule({id: module, enabled: !enabled}));
  }

  importConfigModule(module: string) {
    this.store.dispatch(new ImportConfModule({name: module}));
  }

  createConfigModule(module: string) {
    this.store.dispatch(new FromScratchConfModule({name: module}));
  }

  importConfigAllModules() {
    this.store.dispatch(new ImportAllModules({}));
  }

  autoLoadModule(module: number) {
    this.store.dispatch(new AutoloadModule({id: module}));
  }

  ImportXMLModuleConfig() {
    this.store.dispatch(new ImportXMLModuleConfig({file: this.XMLBody}));
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  isModuleConf(item): boolean {
    switch (typeof item) {
      case null:
      case 'object':
        return true;
      default:
        return false;
    }
  }

  escapeModuleName(str: string): string {
    return str.replace(/_/g, '-');
  }

  openBottomSheetModule(id, name, action): void {
    const config = {
      data:
        {
          name: name,
          action: action,
          case1Text: 'Are you sure you want to delete config of module "' + name + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new TruncateModuleConfig({id: id}));
      }
    });
  }

  initEditor() {
    this.editorInited = true;
  }

  onlySortedValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    const arr = Object.values(obj).sort(
      function (a, b) {
        if (a.position > b.position) {
          return 1;
        }
        if (a.position < b.position) {
          return -1;
        }
        return 0;
      }
    );
    return arr;
  }

}
