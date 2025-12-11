import {Component, DestroyRef, inject, computed, effect} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {
  DelPostSwitchParameter,
  AddPostSwitchParameter,
  StoreNewPostSwitchParameter,
  StoreDropNewPostSwitchParameter,
  SwitchPostSwitchParameter,
  UpdatePostSwitchParameter,
  DelPostSwitchDefaultPtime,
  DelPostSwitchCliKeybinding,
  StoreNewPostSwitchCliKeybinding,
  SwitchPostSwitchDefaultPtime,
  StoreNewPostSwitchDefaultPtime,
  UpdatePostSwitchCliKeybinding,
  UpdatePostSwitchDefaultPtime,
  AddPostSwitchDefaultPtime,
  StoreDropNewPostSwitchDefaultPtime,
  StoreDropNewPostSwitchCliKeybinding,
  SwitchPostSwitchCliKeybinding,
  AddPostSwitchCliKeybinding
} from '../../../store/config/post-switch/config.actions.post-switch';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-post-load-switch',
  templateUrl: './post-load-switch.component.html',
  styleUrls: ['./post-load-switch.component.css']
})
export class PostLoadSwitchComponent {

  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);

  private configsObservable = this.store.pipe(select(selectConfigurationState));
  private configsSignal = toSignal(this.configsObservable, { initialValue: {} as State });

  public list = computed(() => this.configsSignal().post_load_switch || {} as IsimpleModule);
  public loadCounter = computed(() => this.configsSignal().loadCounter || 0);
  private lastErrorMessage = computed(() => this.list().errorMessage || null);

  public selectedIndex: number;
  public globalSettingsDispatchers: object;
  public cliKeybindingsDispatchers: object;
  public defaultPtimesDispatchers: object;
  public defaultPtimesMask: object;

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

    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewPostSwitchParam.bind(this),
      switchItem: this.switchPostSwitchParam.bind(this),
      addItem: this.newPostSwitchParam.bind(this),
      dropNewItem: this.dropNewPostSwitchParam.bind(this),
      deleteItem: this.deletePostSwitchParam.bind(this),
      updateItem: this.updatePostSwitchParam.bind(this),
      pasteItems: null,
    };
    this.cliKeybindingsDispatchers = {
      addNewItemField: this.addNewPostSwitchCliKeybinding.bind(this),
      switchItem: this.switchPostSwitchCliKeybinding.bind(this),
      addItem: this.newPostSwitchCliKeybinding.bind(this),
      dropNewItem: this.dropNewPostSwitchCliKeybinding.bind(this),
      deleteItem: this.deletePostSwitchCliKeybinding.bind(this),
      updateItem: this.updatePostSwitchCliKeybinding.bind(this),
      pasteItems: null,
    };
    this.defaultPtimesDispatchers = {
      addNewItemField: this.addNewPostSwitchDefaultPtime.bind(this),
      switchItem: this.switchPostSwitchDefaultPtime.bind(this),
      addItem: this.newPostSwitchDefaultPtime.bind(this),
      dropNewItem: this.dropNewPostSwitchDefaultPtime.bind(this),
      deleteItem: this.deletePostSwitchDefaultPtime.bind(this),
      updateItem: this.updatePostSwitchDefaultPtime.bind(this),
      pasteItems: null,
    };
    this.defaultPtimesMask = {
      name: {name: 'codec_name'},
      value: {name: 'codec_ptime'},
    };
  }

  updatePostSwitchParam(param: Iitem) {
    this.store.dispatch(new UpdatePostSwitchParameter({param: param}));
  }

  switchPostSwitchParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchPostSwitchParameter({param: newParam}));
  }

  newPostSwitchParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddPostSwitchParameter({index: index, param: param}));
  }

  deletePostSwitchParam(param: Iitem) {
    this.store.dispatch(new DelPostSwitchParameter({param: param}));
  }

  addNewPostSwitchParam() {
    this.store.dispatch(new StoreNewPostSwitchParameter(null));
  }

  dropNewPostSwitchParam(index: number) {
    this.store.dispatch(new StoreDropNewPostSwitchParameter({index: index}));
  }

  updatePostSwitchCliKeybinding(param: Iitem) {
    this.store.dispatch(new UpdatePostSwitchCliKeybinding({param: param}));
  }

  switchPostSwitchCliKeybinding(param: Iitem) {
    const newCliKeybinding = <Iitem>{...param};
    newCliKeybinding.enabled = !newCliKeybinding.enabled;
    this.store.dispatch(new SwitchPostSwitchCliKeybinding({param: newCliKeybinding}));
  }

  newPostSwitchCliKeybinding(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddPostSwitchCliKeybinding({index: index, param: param}));
  }

  deletePostSwitchCliKeybinding(param: Iitem) {
    this.store.dispatch(new DelPostSwitchCliKeybinding({param: param}));
  }

  addNewPostSwitchCliKeybinding() {
    this.store.dispatch(new StoreNewPostSwitchCliKeybinding(null));
  }

  dropNewPostSwitchCliKeybinding(index: number) {
    this.store.dispatch(new StoreDropNewPostSwitchCliKeybinding({index: index}));
  }

  updatePostSwitchDefaultPtime(param: { id: number, codec_name: string, codec_ptime: string }) {
    const para = <Iitem>{id: param.id, name: param.codec_name, value: param.codec_ptime};
    this.store.dispatch(new UpdatePostSwitchDefaultPtime({param: para}));
  }

  switchPostSwitchDefaultPtime(param: Iitem) {
    const newDefaultPtime = <Iitem>{...param};
    newDefaultPtime.enabled = !newDefaultPtime.enabled;
    this.store.dispatch(new SwitchPostSwitchDefaultPtime({param: newDefaultPtime}));
  }

  newPostSwitchDefaultPtime(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddPostSwitchDefaultPtime({index: index, param: param}));
  }

  deletePostSwitchDefaultPtime(param: Iitem) {
    this.store.dispatch(new DelPostSwitchDefaultPtime({param: param}));
  }

  addNewPostSwitchDefaultPtime() {
    this.store.dispatch(new StoreNewPostSwitchDefaultPtime(null));
  }

  dropNewPostSwitchDefaultPtime(index: number) {
    this.store.dispatch(new StoreDropNewPostSwitchDefaultPtime({index: index}));
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
