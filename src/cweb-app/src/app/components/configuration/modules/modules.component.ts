import {Component, DestroyRef, inject, signal, computed, effect} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
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
import {RouterLink} from '@angular/router';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {FormsModule} from "@angular/forms";
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {CodeEditorComponent} from "../../code-editor/code-editor.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, RouterLink, CodeEditorComponent],
  selector: 'app-modules',
  templateUrl: './modules.component.html',
  styleUrls: ['./modules.component.css']
})
export class ModulesComponent {

  private store = inject(Store<AppState>);
  private _snackBar = inject(MatSnackBar);
  private bottomSheet = inject(MatBottomSheet);

  private configsObservable = this.store.pipe(select(selectConfigurationState));
  private configsSignal = toSignal(this.configsObservable, { initialValue: {} as State });

  public list = computed(() => this.configsSignal());
  public loadCounter = computed(() => this.list().loadCounter || 0);
  private lastErrorMessage = computed(() => this.list().errorMessage || '');

  public autoload = computed(() => {
    const list = this.list();
    const newAutoload: { [key: string]: any } = {};
    if (list.post_load_modules?.modules) {
      Object.keys(list.post_load_modules.modules).forEach(
        e => {
          const moduleEntry = list.post_load_modules!.modules[e];
          if (moduleEntry?.name) {
            newAutoload[moduleEntry.name] = moduleEntry;
          }
        }
      );
    }
    return newAutoload;
  });

  public selectedIndex: number = 0;
  public XMLBody = signal<string>('');
  public editorInited = signal<boolean>(false);

  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

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
    this.store.dispatch(new ImportXMLModuleConfig({file: this.XMLBody()}));
  }

  isModuleConf(item: any): boolean {
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

  openBottomSheetModule(id: number, name: string, action: string): void {
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
    this.editorInited.set(true);
  }
}
