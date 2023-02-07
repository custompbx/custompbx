import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Iitem, Iverto, IvertoParameterItem} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  AddVertoProfile,
  AddVertoProfileParam,
  AddVertoSetting,
  DelVertoProfile, DelVertoProfileParam, DelVertoSetting, GetVertoProfileParams, MoveVertoProfileParameter,
  RenameVertoProfile, StoreDropNewVertoProfileParam, StoreDropNewVertoSetting, StoreNewVertoProfileParam, StoreNewVertoSetting,
  StorePasteVertoProfileParams, SwitchVertoProfileParam,
  SwitchVertoSetting, UpdateVertoProfileParam,
  UpdateVertoSetting
} from '../../../store/config/verto/config.actions.verto';
import {CdkDragDrop} from '@angular/cdk/drag-drop';

@Component({
  selector: 'app-verto',
  templateUrl: './verto.component.html',
  styleUrls: ['./verto.component.css']
})
export class VertoComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Iverto;
  private newProfileName: string;
  private newGatewayName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private profileId: number;
  private panelCloser = [];
  public loadCounter: number;
  private toCopyProfile: number;
  public globalSettingsDispatchers: object;
  public profileParamsDispatchers: object;
  public profileParamsMask: object;

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
      this.list = configs.verto;
      this.lastErrorMessage = configs.verto && configs.verto.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newProfileName = '';
        this.newGatewayName = '';
        this.profileId = (this.list && this.list.profiles && this.list.profiles[this.profileId]) ? this.profileId : 0;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewVertoParam.bind(this),
      switchItem: this.switchVertoParam.bind(this),
      addItem: this.newVertoParam.bind(this),
      dropNewItem: this.dropNewVertoParam.bind(this),
      deleteItem: this.deleteVertoParam.bind(this),
      updateItem: this.updateVertoParam.bind(this),
      pasteItems: null,
    };
    this.profileParamsDispatchers = {
      addNewItemField: this.addNewProfileParam.bind(this),
      switchItem: this.switchProfileParam.bind(this),
      addItem: this.newProfileParam.bind(this),
      dropNewItem: this.dropNewProfileParam.bind(this),
      deleteItem: this.deleteProfileParam.bind(this),
      updateItem: this.updateProfileParam.bind(this),
      pasteItems: this.pasteProfileParams.bind(this),
      dropActionItem: this.dropAction.bind(this),
    };
    this.profileParamsMask = {
      name: {name: 'name'},
      value: {name: 'value'},
      extraField1: {name: 'secure', style: {'max-width': '55px'}, depend: 'name', value: 'bind-local'}};
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateVertoParam(param: Iitem) {
    this.store.dispatch(new UpdateVertoSetting({param: param}));
  }

  switchVertoParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchVertoSetting({param: newParam}));
  }

  newVertoParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddVertoSetting({index: index, param: param}));
  }

  deleteVertoParam(param: Iitem) {
    this.store.dispatch(new DelVertoSetting({param: param}));
  }

  addNewVertoParam() {
    this.store.dispatch(new StoreNewVertoSetting(null));
  }

  dropNewVertoParam(index: number) {
    this.store.dispatch(new StoreDropNewVertoSetting({index: index}));
  }

  getVertoProfilesParams(id) {
    this.panelCloser['profile' + id] = true;
    this.store.dispatch(new GetVertoProfileParams({id: id}));
  }

  updateProfileParam(param: IvertoParameterItem) {
    this.store.dispatch(new UpdateVertoProfileParam({param: param}));
  }

  switchProfileParam(param: IvertoParameterItem) {
    const newParam = <IvertoParameterItem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchVertoProfileParam({param: newParam}));
  }

  newProfileParam(parentId: number, index: number, name: string, value: string, secure: string) {
    const param = <IvertoParameterItem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;
    param.secure = secure;

    this.store.dispatch(new AddVertoProfileParam({id: parentId, index: index, param: param}));
  }

  deleteProfileParam(param: IvertoParameterItem) {
    this.store.dispatch(new DelVertoProfileParam({param: param}));
  }

  addNewProfileParam(parentId: number) {
    this.store.dispatch(new StoreNewVertoProfileParam({id: parentId}));
  }

  dropNewProfileParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewVertoProfileParam({id: parentId, index: index}));
  }

  onProfileSubmit() {
    this.store.dispatch(new AddVertoProfile({name: this.newProfileName}));
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

  copyProfile(key) {
    if (!this.list.profiles[key]) {
      this.toCopyProfile = 0;
      return;
    }
    this.toCopyProfile = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteProfileParams(to: number) {
    this.store.dispatch(new StorePasteVertoProfileParams({from_id: this.toCopyProfile, to_id: to}));
  }

  openBottomSheetProfile(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete profile "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename profile "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelVertoProfile({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameVertoProfile({id: id, name: newName}));
      }
    });
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }
    this.store.dispatch(new MoveVertoProfileParameter({
      previous_index: parent[event.previousIndex].position,
      current_index: parent[event.currentIndex].position,
      id: parent[event.previousIndex].id
    }));
  }

}
