import {computed, Component, effect, OnDestroy, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {Iitem, Iosp} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  AddOspProfile,
  AddOspProfileParameter,
  AddOspParameter,
  DelOspProfile, DelOspProfileParameter, DelOspParameter, GetOspProfileParameters,
  UpdateOspProfile, StoreDropNewOspProfileParameter, StoreDropNewOspParameter, StoreNewOspProfileParameter, StoreNewOspParameter,
  StorePasteOspProfileParameters, SwitchOspProfileParameter,
  SwitchOspParameter, UpdateOspProfileParameter,
  UpdateOspParameter
} from '../../../store/config/osp/config.actions.osp';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";
import {toSignal} from "@angular/core/rxjs-interop";

@Component({
standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-osp',
  templateUrl: './osp.component.html',
  styleUrls: ['./osp.component.css']
})
export class OspComponent implements OnInit, OnDestroy {

  private configState = toSignal(this.store.pipe(select(selectConfigurationState)), {initialValue: {} as any});
  public list = computed(() => this.configState().osp as Iosp);
  public loadCounter = computed(() => this.configState().loadCounter || 0);
  public errorMessage = computed(() => this.configState().osp?.errorMessage || null);
  public statusText = computed(() => this.loadCounter() > 0 ? 'Saving…' : null);
  public statusTone = computed(() => this.errorMessage() ? 'danger' : this.loadCounter() > 0 ? 'warning' : 'default');
  private newProfileName: string;
  public selectedIndex: number;
  private panelCloser = [];
  public settingsExpanded = false;
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
    effect(() => {
      const errorMessage = this.errorMessage();
      if (!errorMessage) {
        this.newProfileName = '';
      } else {
        this._snackBar.open('Error: ' + errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnInit() {
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewOspParam.bind(this),
      switchItem: this.switchOspParam.bind(this),
      addItem: this.newOspParam.bind(this),
      dropNewItem: this.dropNewOspParam.bind(this),
      deleteItem: this.deleteOspParam.bind(this),
      updateItem: this.updateOspParam.bind(this),
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
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateOspParam(param: Iitem) {
    this.store.dispatch(new UpdateOspParameter({param: param}));
  }

  switchOspParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchOspParameter({param: newParam}));
  }

  newOspParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddOspParameter({index: index, param: param}));
  }

  deleteOspParam(param: Iitem) {
    this.store.dispatch(new DelOspParameter({param: param}));
  }

  addNewOspParam() {
    this.store.dispatch(new StoreNewOspParameter(null));
  }

  dropNewOspParam(index: number) {
    this.store.dispatch(new StoreDropNewOspParameter({index: index}));
  }

  getOspProfilesParams(id) {
    this.panelCloser['profile' + id] = true;
    this.store.dispatch(new GetOspProfileParameters({id: id}));
  }

  expandAllPanels() {
    this.settingsExpanded = true;
    this.onlyValues(this.list()?.profiles).forEach((profile) => {
      if (profile?.id) {
        this.panelCloser['profile' + profile.id] = true;
        this.store.dispatch(new GetOspProfileParameters({id: profile.id}));
      }
    });
  }

  collapseAllPanels() {
    this.settingsExpanded = false;
    this.onlyValues(this.list()?.profiles).forEach((profile) => {
      if (profile?.id) {
        this.panelCloser['profile' + profile.id] = false;
      }
    });
  }

  updateProfileParam(param: Iitem) {
    this.store.dispatch(new UpdateOspProfileParameter({param: param}));
  }

  switchProfileParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchOspProfileParameter({param: newParam}));
  }

  newProfileParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddOspProfileParameter({id: parentId, index: index, param: param}));
  }

  deleteProfileParam(param: Iitem) {
    this.store.dispatch(new DelOspProfileParameter({param: param}));
  }

  addNewProfileParam(parentId: number) {
    this.store.dispatch(new StoreNewOspProfileParameter({id: parentId}));
  }

  dropNewProfileParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewOspProfileParameter({id: parentId, index: index}));
  }

  onProfileSubmit() {
    this.store.dispatch(new AddOspProfile({name: this.newProfileName}));
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
    if (!this.list()?.profiles[key]) {
      this.toCopyProfile = 0;
      return;
    }
    this.toCopyProfile = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteProfileParams(to: number) {
    this.store.dispatch(new StorePasteOspProfileParameters({from_id: this.toCopyProfile, to_id: to}));
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
        this.store.dispatch(new DelOspProfile({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateOspProfile({id: id, name: newName}));
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

