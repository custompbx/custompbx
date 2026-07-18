import {computed, Component, effect, OnDestroy, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {Iitem, Ilcr} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmationService} from '../../../services/confirmation.service';
import {ToastService} from '../../../services/toast.service';
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
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";
import {TabNavComponent} from '../../tab-nav/tab-nav.component';
import {DisclosureComponent} from '../../disclosure/disclosure.component';
import {TranslocoPipe} from '@jsverse/transloco';
import {toSignal} from "@angular/core/rxjs-interop";

@Component({
standalone: true,
  imports: [CommonModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component, TabNavComponent, DisclosureComponent, TranslocoPipe],
  selector: 'app-lcr',
  templateUrl: './lcr.component.html',
  styleUrls: ['./lcr.component.css']
})
export class LcrComponent implements OnInit, OnDestroy {

  private configState = toSignal(this.store.pipe(select(selectConfigurationState)), {initialValue: {} as any});
  public list = computed(() => this.configState().lcr as Ilcr);
  public loadCounter = computed(() => this.configState().loadCounter || 0);
  public errorMessage = computed(() => this.configState().lcr?.errorMessage || null);
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
    private bottomSheet: ConfirmationService,
    private _snackBar: ToastService,
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

  expandAllPanels() {
    this.settingsExpanded = true;
    this.onlyValues(this.list()?.profiles).forEach((profile) => {
      if (profile?.id) {
        this.panelCloser['profile' + profile.id] = true;
        this.store.dispatch(new GetLcrProfileParameters({id: profile.id}));
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
    if (!this.list()?.profiles[key]) {
      this.toCopyProfile = 0;
      return;
    }
    this.toCopyProfile = key;
    this._snackBar.copied();
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
    const sheet = this.bottomSheet.open(config);
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
