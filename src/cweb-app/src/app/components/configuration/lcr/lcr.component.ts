import {Component, OnDestroy, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {Observable, Subscription} from 'rxjs';
import {Iitem, Ilcr} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  AddLcrProfile,
  AddLcrProfileParameter,
  AddLcrParameter,
  DelLcrProfile, DelLcrProfileParameter, DelLcrParameter, GetLcrProfileParameters,
  UpdateLcrProfile, StoreDropNewLcrProfileParameter, StoreDropNewLcrParameter, StoreNewLcrProfileParameter, StoreNewLcrParameter,
  StorePasteLcrProfileParameters, SwitchLcrProfileParameter,
  SwitchLcrParameter, UpdateLcrProfileParameter,
  UpdateLcrParameter
} from '../../../store/config/lcr/config.actions.lcr';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";

@Component({
standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent],
  selector: 'app-lcr',
  templateUrl: './lcr.component.html',
  styleUrls: ['./lcr.component.css']
})
export class LcrComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Ilcr;
  private newProfileName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private panelCloser = [];
  public loadCounter: number;
  private toCopyProfile: number;
  public globalSettingsDispatchers: object;
  public profileSettingsDispatchers: object;

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
      this.list = configs.lcr;
      this.lastErrorMessage = configs.lcr && configs.lcr.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newProfileName = '';
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewLcrParam.bind(this),
      switchItem: this.switchLcrParam.bind(this),
      addItem: this.newLcrParam.bind(this),
      dropNewItem: this.dropNewLcrParam.bind(this),
      deleteItem: this.deleteLcrParam.bind(this),
      updateItem: this.updateLcrParam.bind(this),
      pasteItems: null,
    };
    this.profileSettingsDispatchers = {
      addNewItemField: this.addNewProfileParam.bind(this),
      switchItem: this.switchProfileParam.bind(this),
      addItem: this.newProfileParam.bind(this),
      dropNewItem: this.dropNewProfileParam.bind(this),
      deleteItem: this.deleteProfileParam.bind(this),
      updateItem: this.updateProfileParam.bind(this),
      pasteItems: this.pasteProfileParams.bind(this),
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateLcrParam(param: Iitem) {
    this.store.dispatch(new UpdateLcrParameter({param: param}));
  }

  switchLcrParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchLcrParameter({param: newParam}));
  }

  newLcrParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddLcrParameter({index: index, param: param}));
  }

  deleteLcrParam(param: Iitem) {
    this.store.dispatch(new DelLcrParameter({param: param}));
  }

  addNewLcrParam() {
    this.store.dispatch(new StoreNewLcrParameter(null));
  }

  dropNewLcrParam(index: number) {
    this.store.dispatch(new StoreDropNewLcrParameter({index: index}));
  }

  getLcrProfilesParams(id) {
    this.panelCloser['profile' + id] = true;
    this.store.dispatch(new GetLcrProfileParameters({id: id}));
  }

  updateProfileParam(param: Iitem) {
    this.store.dispatch(new UpdateLcrProfileParameter({param: param}));
  }

  switchProfileParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchLcrProfileParameter({param: newParam}));
  }

  newProfileParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddLcrProfileParameter({id: parentId, index: index, param: param}));
  }

  deleteProfileParam(param: Iitem) {
    this.store.dispatch(new DelLcrProfileParameter({param: param}));
  }

  addNewProfileParam(parentId: number) {
    this.store.dispatch(new StoreNewLcrProfileParameter({id: parentId}));
  }

  dropNewProfileParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewLcrProfileParameter({id: parentId, index: index}));
  }

  onProfileSubmit() {
    this.store.dispatch(new AddLcrProfile({name: this.newProfileName}));
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
    this.store.dispatch(new StorePasteLcrProfileParameters({from_id: this.toCopyProfile, to_id: to}));
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
        this.store.dispatch(new DelLcrProfile({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateLcrProfile({id: id, name: newName}));
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
