import {Component, inject, computed, effect, DestroyRef} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {Ivoicemail, Iitem, State} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
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
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-voicemail',
  templateUrl: './voicemail.component.html',
  styleUrls: ['./voicemail.component.css']
})
export class VoicemailComponent {

  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);

  private configsObservable = this.store.pipe(select(selectConfigurationState));
  // Assuming a structure similar to the previous component's state
  private configsSignal = toSignal(this.configsObservable, { initialValue: {} as State });

  public list = computed(() => this.configsSignal().voicemail || {} as Ivoicemail);
  public loadCounter = computed(() => this.configsSignal().loadCounter || 0);
  private lastErrorMessage = computed(() => this.list().errorMessage || null);

  public newProfileName: string = '';
  public selectedIndex: number = 0;
  public toCopyProfile: number = 0;
  public globalSettingsDispatchers: object;
  public profileSettingsDispatchers: object;
  public chatPermissionSettingsDispatchers: object; // Not used in component but kept as per original

  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    } else {
      // Logic from ngOnInit subscription: Reset profile name
      this.newProfileName = '';
    }
  });

  constructor() {
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
    // chatPermissionSettingsDispatchers is defined but not initialized in the original constructor/ngOnInit
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

  getVoicemailProfilesParams(id: number) {
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
    return index; // or item.id
  }

  isNewReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  copyProfile(key: number) {
    const profiles = this.list().profiles;
    if (!profiles || !profiles[key]) {
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

  openBottomSheetProfile(id: number, newName: string, oldName: string, action: 'delete' | 'rename'): void {
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

  onlyValues(obj: object | null): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}
