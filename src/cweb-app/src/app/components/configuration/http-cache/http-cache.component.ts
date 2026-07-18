import {computed, Component, effect, OnDestroy, OnInit} from '@angular/core';

import {Ihttpcache, Iitem, IvertoParameterItem} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmationService} from '../../../services/confirmation.service';
import {ToastService} from '../../../services/toast.service';
import {ActivatedRoute} from '@angular/router';
import {
  DelHttpCacheParameter,
  AddHttpCacheParameter,
  StoreNewHttpCacheParameter,
  StoreDropNewHttpCacheParameter,
  SwitchHttpCacheParameter,
  UpdateHttpCacheParameter,
  GetHttpCacheProfileParameters,
  AddHttpCacheProfile,
  UpdateHttpCacheProfileDomain,
  StoreDropNewHttpCacheProfileDomain,
  StoreNewHttpCacheProfileDomain,
  DelHttpCacheProfileDomain,
  AddHttpCacheProfileDomain,
  SwitchHttpCacheProfileDomain,
  DelHttpCacheProfile,
  RenameHttpCacheProfile, UpdateHttpCacheProfileAws, UpdateHttpCacheProfileAzure
} from '../../../store/config/http_cache/config.actions.http_cache';

import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";
import {KeyValuePadPositionComponent} from "../../key-value-pad-position/key-value-pad-position.component";
import {TabNavComponent} from '../../tab-nav/tab-nav.component';
import {DisclosureComponent} from '../../disclosure/disclosure.component';
import {toSignal} from "@angular/core/rxjs-interop";
import {TranslocoPipe} from '@jsverse/transloco';

@Component({
standalone: true,
imports:  [FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component, KeyValuePadPositionComponent, TabNavComponent, DisclosureComponent, TranslocoPipe],
    selector: 'app-http-cache',
    templateUrl: './http-cache.component.html',
    styleUrls: ['./http-cache.component.css']
})
export class HttpCacheComponent implements OnInit, OnDestroy {

  private configState = toSignal(this.store.pipe(select(selectConfigurationState)), {initialValue: {} as any});
  public list = computed(() => this.configState().http_cache as Ihttpcache);
  public loadCounter = computed(() => this.configState().loadCounter || 0);
  public errorMessage = computed(() => this.configState().http_cache?.errorMessage || null);
  public statusText = computed(() => this.loadCounter() > 0 ? 'Saving…' : null);
  public statusTone = computed(() => this.errorMessage() ? 'danger' : this.loadCounter() > 0 ? 'warning' : 'default');
  public selectedIndex: number;
  public globalSettingsDispatchers: object;
  public ProfileDomainsDispatchers: object;
  public ProfileDomainsMask: object;
  private panelCloser = [];
  private newProfileName: string;

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
      addNewItemField: this.addNewHttpCacheParam.bind(this),
      switchItem: this.switchHttpCacheParam.bind(this),
      addItem: this.newHttpCacheParam.bind(this),
      dropNewItem: this.dropNewHttpCacheParam.bind(this),
      deleteItem: this.deleteHttpCacheParam.bind(this),
      updateItem: this.updateHttpCacheParam.bind(this),
      pasteItems: null,
    };
    this.ProfileDomainsDispatchers = {
      addNewItemField: this.addNewProfileDomain.bind(this),
      switchItem: this.switchProfileDomain.bind(this),
      addItem: this.newProfileDomain.bind(this),
      dropNewItem: this.dropNewProfileDomain.bind(this),
      deleteItem: this.deleteProfileDomain.bind(this),
      updateItem: this.updateProfileDomain.bind(this),
      pasteItems: null,
      dropActionItem: null,
    };
    this.ProfileDomainsMask = {name: {name: 'name'}};
  }

  ngOnDestroy() {
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateHttpCacheParam(param: Iitem) {
    this.store.dispatch(new UpdateHttpCacheParameter({param: param}));
  }

  switchHttpCacheParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchHttpCacheParameter({param: newParam}));
  }

  newHttpCacheParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddHttpCacheParameter({index: index, param: param}));
  }

  deleteHttpCacheParam(param: Iitem) {
    this.store.dispatch(new DelHttpCacheParameter({param: param}));
  }

  addNewHttpCacheParam() {
    this.store.dispatch(new StoreNewHttpCacheParameter(null));
  }

  dropNewHttpCacheParam(index: number) {
    this.store.dispatch(new StoreDropNewHttpCacheParameter({index: index}));
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

  GetHttpCacheProfileParameters(id) {
    this.panelCloser['profile' + id] = true;
    this.store.dispatch(new GetHttpCacheProfileParameters({id: id}));
  }

  expandAllPanels() {
    this.onlyValues(this.list()?.profiles).forEach((profile) => {
      if (profile?.id) {
        this.panelCloser['profile' + profile.id] = true;
        this.store.dispatch(new GetHttpCacheProfileParameters({id: profile.id}));
      }
    });
  }

  collapseAllPanels() {
    this.onlyValues(this.list()?.profiles).forEach((profile) => {
      if (profile?.id) {
        this.panelCloser['profile' + profile.id] = false;
      }
    });
  }

  updateProfileDomain(param: IvertoParameterItem) {
    this.store.dispatch(new UpdateHttpCacheProfileDomain({param: param}));
  }

  switchProfileDomain(param: IvertoParameterItem) {
    const newParam = <IvertoParameterItem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchHttpCacheProfileDomain({param: newParam}));
  }

  newProfileDomain(parentId: number, index: number, name: string, value: string, secure: string) {
    const param = <IvertoParameterItem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;
    param.secure = secure;

    this.store.dispatch(new AddHttpCacheProfileDomain({id: parentId, index: index, param: param}));
  }

  deleteProfileDomain(param: IvertoParameterItem) {
    this.store.dispatch(new DelHttpCacheProfileDomain({param: param}));
  }

  addNewProfileDomain(parentId: number) {
    this.store.dispatch(new StoreNewHttpCacheProfileDomain({id: parentId}));
  }

  dropNewProfileDomain(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewHttpCacheProfileDomain({id: parentId, index: index}));
  }

  updateItemAws(item) {
    item.expires = Number(item.expires);
    this.store.dispatch(new UpdateHttpCacheProfileAws({aws_s3: item}));
  }
  updateItemAzure(item) {
    this.store.dispatch(new UpdateHttpCacheProfileAzure({azure: item}));
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }
  firstElement(obj: object) {
    if (!obj) {
      return {};
    }
    const res = Object.values(obj);
    if (res.length === 0) {
      return {};
    }
    return res[0];
  }

  onProfileSubmit() {
    this.store.dispatch(new AddHttpCacheProfile({name: this.newProfileName}));
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
        this.store.dispatch(new DelHttpCacheProfile({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameHttpCacheProfile({id: id, name: newName}));
      }
    });
  }

  getFirstElement(obj): any {
    if (!obj) {
      return {};
    }
    const keys = Object.keys(obj);
    if (keys.length === 0) {
      return {}; // return empty object if the input object is empty
    }
    const firstKey = keys[0];
    return obj[firstKey];
  }
}
