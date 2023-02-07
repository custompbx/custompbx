import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Iitem, Iunicall} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  AddUnicallSpan,
  AddUnicallSpanParameter,
  AddUnicallParameter,
  DelUnicallSpan, DelUnicallSpanParameter, DelUnicallParameter, GetUnicallSpanParameters,
  UpdateUnicallSpan, StoreDropNewUnicallSpanParameter, StoreDropNewUnicallParameter, StoreNewUnicallSpanParameter, StoreNewUnicallParameter,
  StorePasteUnicallSpanParameters, SwitchUnicallSpanParameter,
  SwitchUnicallParameter, UpdateUnicallSpanParameter,
  UpdateUnicallParameter
} from '../../../store/config/unicall/config.actions.unicall';

@Component({
  selector: 'app-unicall',
  templateUrl: './unicall.component.html',
  styleUrls: ['./unicall.component.css']
})
export class UnicallComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Iunicall;
  private newSpanId: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private spanId: number;
  private panelCloser = [];
  public loadCounter: number;
  private toCopySpan: number;
  public globalSettingsDispatchers: object;
  public spanSettingsDispatchers: object;

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
      this.list = configs.unicall;
      this.lastErrorMessage = configs.unicall && configs.unicall.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newSpanId = '';
        this.spanId = (this.list && this.list.spans && this.list.spans[this.spanId]) ? this.spanId : 0;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewUnicallParam.bind(this),
      switchItem: this.switchUnicallParam.bind(this),
      addItem: this.newUnicallParam.bind(this),
      dropNewItem: this.dropNewUnicallParam.bind(this),
      deleteItem: this.deleteUnicallParam.bind(this),
      updateItem: this.updateUnicallParam.bind(this),
      pasteItems: null,
    };
    this.spanSettingsDispatchers = {
      addNewItemField: this.addNewSpanParam.bind(this),
      switchItem: this.switchSpanParam.bind(this),
      addItem: this.newSpanParam.bind(this),
      dropNewItem: this.dropNewSpanParam.bind(this),
      deleteItem: this.deleteSpanParam.bind(this),
      updateItem: this.updateSpanParam.bind(this),
      pasteItems: this.pasteSpanParams.bind(this),
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateUnicallParam(param: Iitem) {
    this.store.dispatch(new UpdateUnicallParameter({param: param}));
  }

  switchUnicallParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchUnicallParameter({param: newParam}));
  }

  newUnicallParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddUnicallParameter({index: index, param: param}));
  }

  deleteUnicallParam(param: Iitem) {
    this.store.dispatch(new DelUnicallParameter({param: param}));
  }

  addNewUnicallParam() {
    this.store.dispatch(new StoreNewUnicallParameter(null));
  }

  dropNewUnicallParam(index: number) {
    this.store.dispatch(new StoreDropNewUnicallParameter({index: index}));
  }

  getUnicallSpansParams(id) {
    this.panelCloser['span' + id] = true;
    this.store.dispatch(new GetUnicallSpanParameters({id: id}));
  }

  updateSpanParam(param: Iitem) {
    this.store.dispatch(new UpdateUnicallSpanParameter({param: param}));
  }

  switchSpanParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchUnicallSpanParameter({param: newParam}));
  }

  newSpanParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddUnicallSpanParameter({id: parentId, index: index, param: param}));
  }

  deleteSpanParam(param: Iitem) {
    this.store.dispatch(new DelUnicallSpanParameter({param: param}));
  }

  addNewSpanParam(parentId: number) {
    this.store.dispatch(new StoreNewUnicallSpanParameter({id: parentId}));
  }

  dropNewSpanParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewUnicallSpanParameter({id: parentId, index: index}));
  }

  onSpanSubmit() {
    this.store.dispatch(new AddUnicallSpan({name: this.newSpanId}));
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

  copySpan(key) {
    if (!this.list.spans[key]) {
      this.toCopySpan = 0;
      return;
    }
    this.toCopySpan = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteSpanParams(to: number) {
    this.store.dispatch(new StorePasteUnicallSpanParameters({from_id: this.toCopySpan, to_id: to}));
  }

  openBottomSheetSpan(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete span "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename span "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelUnicallSpan({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateUnicallSpan({id: id, name: newName}));
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

