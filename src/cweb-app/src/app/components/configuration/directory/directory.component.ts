import {Component, computed, effect, OnDestroy, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {Iitem, Idirectory} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmationService} from '../../../services/confirmation.service';
import {ToastService} from '../../../services/toast.service';
import {ActivatedRoute} from '@angular/router';
import {
  AddDirectoryProfile,
  AddDirectoryProfileParameter,
  AddDirectoryParameter,
  DelDirectoryProfile, DelDirectoryProfileParameter, DelDirectoryParameter, GetDirectoryProfileParameters,
  UpdateDirectoryProfile, StoreDropNewDirectoryProfileParameter, StoreDropNewDirectoryParameter,
  StoreNewDirectoryProfileParameter, StoreNewDirectoryParameter,
  StorePasteDirectoryProfileParameters, SwitchDirectoryProfileParameter,
  SwitchDirectoryParameter, UpdateDirectoryProfileParameter,
  UpdateDirectoryParameter
} from '../../../store/config/directory/config.actions.directory';
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
  selector: 'app-directory',
  templateUrl: './directory.component.html',
  styleUrls: ['./directory.component.css']
})
export class DirectoryComponent implements OnInit, OnDestroy {

  private configState = toSignal(this.store.pipe(select(selectConfigurationState)), {initialValue: {} as any});
  public list = computed(() => this.configState().directory as Idirectory);
  public loadCounter = computed(() => this.configState().loadCounter || 0);
  public newProfileName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private panelCloser = [];
  public toCopyProfile: number;
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
      const message = this.list()?.errorMessage || null;
      if (message && message !== this.lastErrorMessage) {
        this._snackBar.open('Error: ' + message + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
      this.lastErrorMessage = message;
    });
  }

  ngOnInit() {
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewDirectoryParam.bind(this),
      switchItem: this.switchDirectoryParam.bind(this),
      addItem: this.newDirectoryParam.bind(this),
      dropNewItem: this.dropNewDirectoryParam.bind(this),
      deleteItem: this.deleteDirectoryParam.bind(this),
      updateItem: this.updateDirectoryParam.bind(this),
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

  updateDirectoryParam(param: Iitem) {
    this.store.dispatch(new UpdateDirectoryParameter({param: param}));
  }

  switchDirectoryParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchDirectoryParameter({param: newParam}));
  }

  newDirectoryParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddDirectoryParameter({index: index, param: param}));
  }

  deleteDirectoryParam(param: Iitem) {
    this.store.dispatch(new DelDirectoryParameter({param: param}));
  }

  addNewDirectoryParam() {
    this.store.dispatch(new StoreNewDirectoryParameter(null));
  }

  dropNewDirectoryParam(index: number) {
    this.store.dispatch(new StoreDropNewDirectoryParameter({index: index}));
  }

  getDirectoryProfilesParams(id) {
    this.panelCloser['profile' + id] = true;
    this.store.dispatch(new GetDirectoryProfileParameters({id: id}));
  }

  updateProfileParam(param: Iitem) {
    this.store.dispatch(new UpdateDirectoryProfileParameter({param: param}));
  }

  switchProfileParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchDirectoryProfileParameter({param: newParam}));
  }

  newProfileParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddDirectoryProfileParameter({id: parentId, index: index, param: param}));
  }

  deleteProfileParam(param: Iitem) {
    this.store.dispatch(new DelDirectoryProfileParameter({param: param}));
  }

  addNewProfileParam(parentId: number) {
    this.store.dispatch(new StoreNewDirectoryProfileParameter({id: parentId}));
  }

  dropNewProfileParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewDirectoryProfileParameter({id: parentId, index: index}));
  }

  onProfileSubmit() {
    this.store.dispatch(new AddDirectoryProfile({name: this.newProfileName}));
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
    if (!this.list()?.profiles?.[key]) {
      this.toCopyProfile = 0;
      return;
    }
    this.toCopyProfile = key;
    this._snackBar.copied();
  }

  pasteProfileParams(to: number) {
    this.store.dispatch(new StorePasteDirectoryProfileParameters({from_id: this.toCopyProfile, to_id: to}));
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
        this.store.dispatch(new DelDirectoryProfile({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateDirectoryProfile({id: id, name: newName}));
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

