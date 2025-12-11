import {Component, DestroyRef, inject, computed, effect} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {Iitem, IpostLoadModules, State} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {
  DelPostLoadModule,
  AddPostLoadModule,
  StoreNewPostLoadModule,
  StoreDropNewPostLoadModule,
  SwitchPostLoadModule,
  UpdatePostLoadModule
} from '../../../store/config/post_load_modules/config.actions.PostLoadModules';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, KeyValuePad2Component],
  selector: 'app-post-load-modules',
  templateUrl: './post-load-modules.component.html',
  styleUrls: ['./post-load-modules.component.css']
})
export class PostLoadModulesComponent {

  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);

  public globalSettingsMask: object = {};

  private configsObservable = this.store.pipe(select(selectConfigurationState));
  private configsSignal = toSignal(this.configsObservable, { initialValue: {} as State });

  public list = computed(() => this.configsSignal().post_load_modules || {} as IpostLoadModules);
  public loadCounter = computed(() => this.configsSignal().loadCounter || 0);
  private lastErrorMessage = computed(() => this.configsSignal().errorMessage || null);

  public selectedIndex: number = 0;
  public globalSettingsDispatchers: object;

  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  constructor() {
    this.selectedIndex = 0;

    this.globalSettingsMask = {name: {name: 'name'}};

    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewPostLoadModule.bind(this),
      switchItem: this.switchPostLoadModule.bind(this),
      addItem: this.newPostLoadModule.bind(this),
      dropNewItem: this.dropNewPostLoadModule.bind(this),
      deleteItem: this.deletePostLoadModule.bind(this),
      updateItem: this.updatePostLoadModule.bind(this),
      pasteItems: null,
    };
  }

  updatePostLoadModule(param: Iitem) {
    this.store.dispatch(new UpdatePostLoadModule({param: param}));
  }

  switchPostLoadModule(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchPostLoadModule({param: newParam}));
  }

  newPostLoadModule(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddPostLoadModule({index: index, param: param}));
  }

  deletePostLoadModule(param: Iitem) {
    this.store.dispatch(new DelPostLoadModule({param: param}));
  }

  addNewPostLoadModule() {
    this.store.dispatch(new StoreNewPostLoadModule(null));
  }

  dropNewPostLoadModule(index: number) {
    this.store.dispatch(new StoreDropNewPostLoadModule({index: index}));
  }

  checkDirty(condition: AbstractControl | null): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isReadyToSendThree(mainObject: AbstractControl | null, object2: AbstractControl | null, object3: AbstractControl | null): boolean {
    return (mainObject && mainObject.valid && mainObject.dirty)
      || ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty));
  }

  isvalueReadyToSend(valueObject: AbstractControl | null): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  isReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index: number, item: any) {
    return index;
  }

  isNewReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

}
