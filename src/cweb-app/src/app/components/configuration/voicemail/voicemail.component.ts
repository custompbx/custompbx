import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Ivoicemail, Iitem} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  AddVoicemailProfile,
  AddVoicemailProfileParameter,
  AddVoicemailSetting,
  DelVoicemailProfile,
  DelVoicemailProfileParameter,
  DelVoicemailSetting,
  GetVoicemailProfileParameters,
  UpdateVoicemailProfile,
  StoreDropNewVoicemailProfileParameter,
  StoreDropNewVoicemailSetting,
  StoreNewVoicemailProfileParameter,
  StoreNewVoicemailSetting,
  StorePasteVoicemailProfileParameters,
  SwitchVoicemailProfileParameter,
  SwitchVoicemailSetting,
  UpdateVoicemailProfileParameter,
  UpdateVoicemailSetting, GetVoicemailSettings,
} from '../../../store/config/voicemail/config.actions.voicemail';

@Component({
  selector: 'app-voicemail',
  templateUrl: './voicemail.component.html',
  styleUrls: ['./voicemail.component.css']
})
export class VoicemailComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Ivoicemail;
  private newProfileName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  private toCopyProfile: number;
  public globalSettingsDispatchers: object;
  public profileSettingsDispatchers: object;
  public chatPermissionSettingsDispatchers: object;

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
      this.list = configs.voicemail;
      this.lastErrorMessage = configs.voicemail && configs.voicemail.errorMessage || null;
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
      addNewItemField: this.addNewVoicemailSetting.bind(this),
      switchItem: this.switchVoicemailSetting.bind(this),
      addItem: this.newVoicemailSetting.bind(this),
      dropNewItem: this.dropNewVoicemailSetting.bind(this),
      deleteItem: this.deleteVoicemailSetting.bind(this),
      updateItem: this.updateVoicemailSetting.bind(this),
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

  getVoicemailSetting() {
    this.store.dispatch(new GetVoicemailSettings(null));
  }

  updateVoicemailSetting(param: Iitem) {
    this.store.dispatch(new UpdateVoicemailSetting({param: param}));
  }

  switchVoicemailSetting(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchVoicemailSetting({param: newParam}));
  }

  newVoicemailSetting(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddVoicemailSetting({index: index, param: param}));
  }

  deleteVoicemailSetting(param: Iitem) {
    this.store.dispatch(new DelVoicemailSetting({param: param}));
  }

  addNewVoicemailSetting() {
    this.store.dispatch(new StoreNewVoicemailSetting(null));
  }

  dropNewVoicemailSetting(index: number) {
    this.store.dispatch(new StoreDropNewVoicemailSetting({index: index}));
  }

  getVoicemailProfilesParams(id) {
    this.store.dispatch(new GetVoicemailProfileParameters({id: id}));
  }

  updateProfileParam(param: Iitem) {
    this.store.dispatch(new UpdateVoicemailProfileParameter({param: param}));
  }

  switchProfileParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchVoicemailProfileParameter({param: newParam}));
  }

  newProfileParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddVoicemailProfileParameter({id: parentId, index: index, param: param}));
  }

  deleteProfileParam(param: Iitem) {
    this.store.dispatch(new DelVoicemailProfileParameter({param: param}));
  }

  addNewProfileParam(parentId: number) {
    this.store.dispatch(new StoreNewVoicemailProfileParameter({id: parentId}));
  }

  dropNewProfileParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewVoicemailProfileParameter({id: parentId, index: index}));
  }

  onProfileSubmit() {
    this.store.dispatch(new AddVoicemailProfile({name: this.newProfileName}));
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
    this.store.dispatch(new StorePasteVoicemailProfileParameters({from_id: this.toCopyProfile, to_id: to}));
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
        this.store.dispatch(new DelVoicemailProfile({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateVoicemailProfile({id: id, name: newName}));
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
