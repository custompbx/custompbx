import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Iitem, Iopal} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
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

@Component({
  selector: 'app-opal',
  templateUrl: './opal.component.html',
  styleUrls: ['./opal.component.css']
})
export class OpalComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Iopal;
  private newListenerName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private panelCloser = [];
  public loadCounter: number;
  private toCopyListener: number;
  public globalSettingsDispatchers: object;
  public listenerSettingsDispatchers: object;

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.configs = this.store.pipe(select(selectConfigurationState));
  }

  ngOnInit() {
    this.configs$ = this.configs.subscribe((configs) => {
      this.loadCounter = configs.loadCounter;
      this.list = configs.opal;
      this.lastErrorMessage = configs.opal && configs.opal.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newListenerName = '';
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewOpalParam.bind(this),
      switchItem: this.switchOpalParam.bind(this),
      addItem: this.newOpalParam.bind(this),
      dropNewItem: this.dropNewOpalParam.bind(this),
      deleteItem: this.deleteOpalParam.bind(this),
      updateItem: this.updateOpalParam.bind(this),
      pasteItems: null,
    };
    this.listenerSettingsDispatchers = {
      addNewItemField: this.addNewListenerParam.bind(this),
      switchItem: this.switchListenerParam.bind(this),
      addItem: this.newListenerParam.bind(this),
      dropNewItem: this.dropNewListenerParam.bind(this),
      deleteItem: this.deleteListenerParam.bind(this),
      updateItem: this.updateListenerParam.bind(this),
      pasteItems: this.pasteListenerParams.bind(this),
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

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

  getOpalListenersParams(id) {
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
    this.store.dispatch(new AddOpalListener({name: this.newListenerName}));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isReadyToSendThree(mainObject: AbstractControl, object2: AbstractControl, object3: AbstractControl): boolean {
    return (mainObject && mainObject.valid && mainObject.dirty)
      || ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty));
  }

  isvalueReadyToSend(valueObject: AbstractControl): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  copyListener(key) {
    if (!this.list.listeners[key]) {
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

  openBottomSheetListener(id, newName, oldName, action): void {
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

