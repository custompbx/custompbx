import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Iitem, Ififo, IfifoMember} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  AddFifoFifo,
  AddFifoFifoMember,
  AddFifoParameter,
  DelFifoFifo, DelFifoFifoMember, DelFifoParameter, GetFifoFifoMembers,
  UpdateFifoFifo, StoreDropNewFifoFifoMember, StoreDropNewFifoParameter, StoreNewFifoFifoMember, StoreNewFifoParameter,
  StorePasteFifoFifoMembers, SwitchFifoFifoMember,
  SwitchFifoParameter, UpdateFifoFifoMember,
  UpdateFifoParameter, UpdateFifoFifoImportance
} from '../../../store/config/fifo/config.actions.fifo';

@Component({
  selector: 'app-fifo',
  templateUrl: './fifo.component.html',
  styleUrls: ['./fifo.component.css']
})
export class FifoComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Ififo;
  private newFifoName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private panelCloser = [];
  public loadCounter: number;
  private toCopyFifo: number;
  public globalSettingsDispatchers: object;
  public fifoSettingsDispatchers: object;
  public memberMask: object;

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
      this.list = configs.fifo;
      this.lastErrorMessage = configs.fifo && configs.fifo.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newFifoName = '';
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewFifoParam.bind(this),
      switchItem: this.switchFifoParam.bind(this),
      addItem: this.newFifoParam.bind(this),
      dropNewItem: this.dropNewFifoParam.bind(this),
      deleteItem: this.deleteFifoParam.bind(this),
      updateItem: this.updateFifoParam.bind(this),
      pasteItems: null,
    };
    this.fifoSettingsDispatchers = {
      addNewItemField: this.addNewFifoMember.bind(this),
      switchItem: this.switchFifoMember.bind(this),
      addItem: this.newFifoMember.bind(this),
      dropNewItem: this.dropNewFifoMember.bind(this),
      deleteItem: this.deleteFifoMember.bind(this),
      updateItem: this.updateFifoMember.bind(this),
      pasteItems: this.pasteFifoMembers.bind(this),
    };
    this.memberMask = {name: {name: 'timeout'}, value: {name: 'simo'}, extraField1: {name: 'lag'}, extraField2: {name: 'body'}};
  }
  UpdateConferenceCallerControl
  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateImportance(id: number, valueObject: AbstractControl) {
    const value = valueObject.value;
    valueObject.reset();
    this.store.dispatch(new UpdateFifoFifoImportance({value: value, id: id}));
  }

  updateFifoParam(param: Iitem) {
    this.store.dispatch(new UpdateFifoParameter({param: param}));
  }

  switchFifoParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchFifoParameter({param: newParam}));
  }

  newFifoParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddFifoParameter({index: index, param: param}));
  }

  deleteFifoParam(param: Iitem) {
    this.store.dispatch(new DelFifoParameter({param: param}));
  }

  addNewFifoParam() {
    this.store.dispatch(new StoreNewFifoParameter(null));
  }

  dropNewFifoParam(index: number) {
    this.store.dispatch(new StoreDropNewFifoParameter({index: index}));
  }

  getFifoFifosParams(id) {
    this.panelCloser['fifo' + id] = true;
    this.store.dispatch(new GetFifoFifoMembers({id: id}));
  }

  updateFifoMember(param: IfifoMember) {
    this.store.dispatch(new UpdateFifoFifoMember({fifo_fifo_member: param}));
  }

  switchFifoMember(param: IfifoMember) {
    const newParam = <IfifoMember>{...param};
    newParam.enabled = !newParam.enabled;

    this.store.dispatch(new SwitchFifoFifoMember({fifo_fifo_member: newParam}));
  }

  newFifoMember(parentId: number, index: number, name: string, value: string, lag: string, body: string) {
    const param = <IfifoMember>{};
    param.enabled = true;
    param.timeout = name;
    param.simo = value;
    param.lag = lag;
    param.body = body;

    this.store.dispatch(new AddFifoFifoMember({id: parentId, index: index, fifo_fifo_member: {...param}}));
  }

  deleteFifoMember(param) {
    this.store.dispatch(new DelFifoFifoMember({fifo_fifo_member: param}));
  }

  addNewFifoMember(parentId: number) {
    this.store.dispatch(new StoreNewFifoFifoMember({id: parentId}));
  }

  dropNewFifoMember(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewFifoFifoMember({id: parentId, index: index}));
  }

  onFifoSubmit() {
    this.store.dispatch(new AddFifoFifo({name: this.newFifoName}));
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

  copyFifo(key) {
    if (!this.list.fifos[key]) {
      this.toCopyFifo = 0;
      return;
    }
    this.toCopyFifo = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteFifoMembers(to: number) {
    this.store.dispatch(new StorePasteFifoFifoMembers({from_id: this.toCopyFifo, to_id: to}));
  }

  openBottomSheetFifo(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete fifo "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename fifo "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelFifoFifo({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateFifoFifo({id: id, name: newName}));
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

