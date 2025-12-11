import {Component, inject, signal, computed, effect} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {Iitem, Iopal, State} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {
  AddOpalListener,
  AddOpalListenerParameter,
  AddOpalParameter,
  DelOpalListener, DelOpalListenerParameter, DelOpalParameter, GetOpalListenerParameters,
  UpdateOpalListener, StoreDropNewOpalListenerParameter, StoreDropNewOpalParameter, StoreNewOpalListenerParameter, StoreNewOpalParameter,
  StorePasteOpalListenerParameters, SwitchOpalListenerParameter,
  SwitchOpalParameter, UpdateOpalListenerParameter,
  UpdateOpalParameter
} from '../../../store/config/opal/config.actions.opal';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-opal',
  templateUrl: './opal.component.html',
  styleUrls: ['./opal.component.css']
})
export class OpalComponent {

  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);

  private configsObservable = this.store.pipe(select(selectConfigurationState));
  private configsSignal = toSignal(this.configsObservable, { initialValue: {} as State });

  public list = computed(() => this.configsSignal().opal || {} as Iopal);
  public loadCounter = computed(() => this.configsSignal().loadCounter || 0);
  private lastErrorMessage = computed(() => this.list().errorMessage || null);

  public newListenerName = signal<string>('');
  public selectedIndex: number = 0;
  private panelCloser: {[key: string]: boolean} = {};
  public toCopyListener: number = 0;

  public globalSettingsDispatchers = {
    addNewItemField: this.addNewOpalParam.bind(this),
    switchItem: this.switchOpalParam.bind(this),
    addItem: this.newOpalParam.bind(this),
    dropNewItem: this.dropNewOpalParam.bind(this),
    deleteItem: this.deleteOpalParam.bind(this),
    updateItem: this.updateOpalParam.bind(this),
    pasteItems: null,
  };

  public listenerSettingsDispatchers = {
    addNewItemField: this.addNewListenerParam.bind(this),
    switchItem: this.switchListenerParam.bind(this),
    addItem: this.newListenerParam.bind(this),
    dropNewItem: this.dropNewListenerParam.bind(this),
    deleteItem: this.deleteListenerParam.bind(this),
    updateItem: this.updateListenerParam.bind(this),
    pasteItems: this.pasteListenerParams.bind(this),
  };

  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  updateOpalParam(param: Iitem) {
    this.store.dispatch(new UpdateOpalParameter({param: param}));
  }

  switchOpalParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchOpalParameter({param: newParam}));
  }

  newOpalParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddOpalParameter({index: index, param: param}));
  }

  deleteOpalParam(param: Iitem) {
    this.store.dispatch(new DelOpalParameter({param: param}));
  }

  addNewOpalParam() {
    this.store.dispatch(new StoreNewOpalParameter(null));
  }

  dropNewOpalParam(index: number) {
    this.store.dispatch(new StoreDropNewOpalParameter({index: index}));
  }

  getOpalListenersParams(id: number) {
    this.panelCloser['listener' + id] = true;
    this.store.dispatch(new GetOpalListenerParameters({id: id}));
  }

  updateListenerParam(param: Iitem) {
    this.store.dispatch(new UpdateOpalListenerParameter({param: param}));
  }

  switchListenerParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchOpalListenerParameter({param: newParam}));
  }

  newListenerParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddOpalListenerParameter({id: parentId, index: index, param: param}));
  }

  deleteListenerParam(param: Iitem) {
    this.store.dispatch(new DelOpalListenerParameter({param: param}));
  }

  addNewListenerParam(parentId: number) {
    this.store.dispatch(new StoreNewOpalListenerParameter({id: parentId}));
  }

  dropNewListenerParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewOpalListenerParameter({id: parentId, index: index}));
  }

  onListenerSubmit() {
    this.store.dispatch(new AddOpalListener({name: this.newListenerName()}));
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

  isNewReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  copyListener(key: number) {
    const listeners = this.list().listeners;
    if (!listeners || !listeners[key]) {
      this.toCopyListener = 0;
      return;
    }
    this.toCopyListener = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteListenerParams(to: number) {
    this.store.dispatch(new StorePasteOpalListenerParameters({from_id: this.toCopyListener, to_id: to}));
  }

  openBottomSheetListener(id: number, newName: string, oldName: string, action: string): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete listener "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename listener "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelOpalListener({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateOpalListener({id: id, name: newName}));
      }
    });
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}
