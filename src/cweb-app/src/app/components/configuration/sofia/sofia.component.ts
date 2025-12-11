import {Component, DestroyRef, inject, computed, effect} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {
  Ialias, IdirectionItem, Idomain, Idomains, Iitem, Iprofile, Isofia, State
} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {
  AddSofiaGlobalSettings,
  AddSofiaProfile,
  AddSofiaProfileAlias,
  AddSofiaProfileDomain,
  AddSofiaProfileGateway,
  AddSofiaProfileGatewayParam,
  AddSofiaProfileGatewayVar,
  AddSofiaProfileParam,
  DelSofiaGlobalSettings,
  DelSofiaProfile,
  DelSofiaProfileAlias,
  DelSofiaProfileDomain,
  DelSofiaProfileGateway,
  DelSofiaProfileGatewayParam,
  DelSofiaProfileGatewayVar,
  DelSofiaProfileParam,
  GetSofiaGlobalSettings,
  GetSofiaProfileAliases,
  GetSofiaProfileDomains,
  GetSofiaProfileGatewayParameters,
  GetSofiaProfileGateways,
  GetSofiaProfilesParams,
  RenameSofiaProfile,
  RenameSofiaProfileGateway,
  SofiaProfileCommand,
  StoreDropNewSofiaGlobalSettings,
  StoreDropNewSofiaProfileAlias,
  StoreDropNewSofiaProfileDomain,
  StoreDropNewSofiaProfileGatewayParam,
  StoreDropNewSofiaProfileGatewayVar,
  StoreDropNewSofiaProfileParam,
  StoreNewSofiaGlobalSettings,
  StoreNewSofiaProfileAlias,
  StoreNewSofiaProfileDomain,
  StoreNewSofiaProfileGatewayParam,
  StoreNewSofiaProfileGatewayVar,
  StoreNewSofiaProfileParam,
  StorePasteSofiaProfileGatewayParams,
  StorePasteSofiaProfileGatewayVars,
  StorePasteSofiaProfileParams,
  SwitchSofiaGlobalSettings,
  SwitchSofiaProfile,
  SwitchSofiaProfileAlias,
  SwitchSofiaProfileDomain,
  SwitchSofiaProfileGatewayParam,
  SwitchSofiaProfileGatewayVar,
  SwitchSofiaProfileParam,
  UpdateSofiaGlobalSettings,
  UpdateSofiaProfileAlias,
  UpdateSofiaProfileDomain,
  UpdateSofiaProfileGatewayParam,
  UpdateSofiaProfileGatewayVar,
  UpdateSofiaProfileParam
} from '../../../store/config/sofia/config.actions.sofia';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-sofia',
  templateUrl: './sofia.component.html',
  styleUrls: ['./sofia.component.css']
})
export class SofiaComponent {

  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);

  private configsObservable = this.store.pipe(select(selectConfigurationState));
  private configsSignal = toSignal(this.configsObservable, { initialValue: {} as State });

  public list = computed(() => this.configsSignal().sofia || {} as Isofia);
  public loadCounter = computed(() => this.configsSignal().loadCounter || 0);
  private lastErrorMessage = computed(() => this.list().errorMessage || null);

  public newProfileName: string = '';
  public newGatewayName: string = '';
  public selectedIndex: number = 0;
  public profileId: number = 0;
  public panelCloser = {};
  public choosedGateway = [];
  public toCopyProfile: number = 0;
  public toCopyGateway: number = 0;
  public toCopyProfileGateway: number = 0;

  public globalSettingsDispatchers: object;
  public profileParamsDispatchers: object;
  public gatewayParamsDispatchers: object;
  public gatewayVarsDispatchers: object;
  public profileDomainsDispatchers: object;

  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    } else {
      // Logic from ngOnInit subscription: Reset names and ensure profileId is valid
      this.newProfileName = '';
      this.newGatewayName = '';
      const profiles = this.list().profiles;
      if (profiles && profiles[this.profileId]) {
        // If profileId is still valid, keep it. Otherwise, default to 0.
      } else {
        this.profileId = 0;
      }
    }
  });

  constructor() {
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewGlobalParam.bind(this),
      switchItem: this.switchGlobalParam.bind(this),
      addItem: this.newGlobalParam.bind(this),
      dropNewItem: this.dropNewGlobalParam.bind(this),
      deleteItem: this.deleteGlobalParam.bind(this),
      updateItem: this.updateGlobalParam.bind(this),
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
    };
    this.gatewayParamsDispatchers = {
      addNewItemField: this.addNewProfileGatewayParam.bind(this),
      switchItem: this.switchProfileGatewayParam.bind(this),
      addItem: this.newProfileGatewayParam.bind(this),
      dropNewItem: this.dropNewProfileGatewayParam.bind(this),
      deleteItem: this.deleteProfileGatewayParam.bind(this),
      updateItem: this.updateProfileGatewayParam.bind(this),
      pasteItems: this.pasteGatewayParams.bind(this),
    };
    this.gatewayVarsDispatchers = {
      addNewItemField: this.addNewProfileGatewayVar.bind(this),
      switchItem: this.switchProfileGatewayVar.bind(this),
      addItem: this.newProfileGatewayVar.bind(this),
      dropNewItem: this.dropNewProfileGatewayVar.bind(this),
      deleteItem: this.deleteProfileGatewayVar.bind(this),
      updateItem: this.updateProfileGatewayVar.bind(this),
      pasteItems: this.pasteGatewayVars.bind(this),
    };
    this.profileDomainsDispatchers = {
      addNewItemField: this.addNewProfileDomain.bind(this),
      switchItem: this.switchProfileDomain.bind(this),
      addItem: this.newProfileDomain.bind(this),
      dropNewItem: this.dropNewProfileDomain.bind(this),
      deleteItem: this.deleteProfileDomain.bind(this),
      updateItem: this.updateProfileDomain.bind(this),
      pasteItems: null,
    };
  }

  getGlobalSettings() {
    this.store.dispatch(new GetSofiaGlobalSettings(null));
  }

  clearGlobalSettings() {
    // This function was empty in the original code, keeping it as a stub.
  }

  updateGlobalParam(param: Iitem) {
    this.store.dispatch(new UpdateSofiaGlobalSettings({param: param}));
  }

  switchGlobalParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchSofiaGlobalSettings({param: newParam}));
  }

  newGlobalParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddSofiaGlobalSettings({index: index, param: param}));
  }

  deleteGlobalParam(param: Iitem) {
    this.store.dispatch(new DelSofiaGlobalSettings({param: param}));
  }

  addNewGlobalParam() {
    this.store.dispatch(new StoreNewSofiaGlobalSettings(null));
  }

  dropNewGlobalParam(index: number) {
    this.store.dispatch(new StoreDropNewSofiaGlobalSettings({index: index}));
  }

  getSofiaProfilesParams(id: number) {
    this.panelCloser['profile' + id] = true;
    this.store.dispatch(new GetSofiaProfilesParams({id: id}));
  }

  updateProfileParam(param: Iitem) {
    this.store.dispatch(new UpdateSofiaProfileParam({param: param}));
  }

  switchProfileParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchSofiaProfileParam({param: newParam}));
  }

  newProfileParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddSofiaProfileParam({id: parentId, index: index, param: param}));
  }

  deleteProfileParam(param: Iitem) {
    this.store.dispatch(new DelSofiaProfileParam({param: param}));
  }

  addNewProfileParam(parentId: number) {
    this.store.dispatch(new StoreNewSofiaProfileParam({id: parentId}));
  }

  dropNewProfileParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewSofiaProfileParam({id: parentId, index: index}));
  }

  tabChanged(event: number) {
    this.panelCloser = {};
    if (event === 1 || event === 4) {
      this.getSofiaProfilesGateways();
    }
  }

  mainTabChanged(event: number) {
    if (event === 2) {
      this.getSofiaProfilesGateways();
    }
  }

  getSofiaProfilesGateways() {
    const profiles = this.list().profiles;
    if (profiles) {
      const ids = Object.keys(profiles).map(Number);
      ids.forEach((id) => this.store.dispatch(new GetSofiaProfileGateways({id: id, keep_subscription: true})));
    }
  }

  getSofiaProfilesGatewayParams(id: number) {
    this.panelCloser['gateway' + id] = true;
    this.store.dispatch(new GetSofiaProfileGatewayParameters({id: id}));
    // this.store.dispatch(new GetSofiaProfileGatewayVariables({id: id})); // Original code commented this out, keeping it that way.
  }

  updateProfileGatewayParam(param: Iitem) {
    this.store.dispatch(new UpdateSofiaProfileGatewayParam({param: param}));
  }

  switchProfileGatewayParam(param: Iitem) {
    const newParam = {...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchSofiaProfileGatewayParam({param: newParam}));
  }

  newProfileGatewayParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddSofiaProfileGatewayParam({id: parentId, index: index, param: param}));
  }

  deleteProfileGatewayParam(param: Iitem) {
    this.store.dispatch(new DelSofiaProfileGatewayParam({param: param}));
  }

  addNewProfileGatewayParam(grandPrentId: number, parentId: number) {
    this.store.dispatch(new StoreNewSofiaProfileGatewayParam({profileId: grandPrentId, id: parentId}));
  }

  dropNewProfileGatewayParam(grandPrentId: number, parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewSofiaProfileGatewayParam({profileId: grandPrentId, id: parentId, index: index}));
  }


  updateProfileGatewayVar(param: Iitem) {
    this.store.dispatch(new UpdateSofiaProfileGatewayVar({variable: param}));
  }

  switchProfileGatewayVar(variable: Iitem) {
    const newVar = {...variable};
    newVar.enabled = !newVar.enabled;
    this.store.dispatch(new SwitchSofiaProfileGatewayVar({variable: newVar}));
  }

  newProfileGatewayVar(parentId: number, index: number, name: string, value: string, direction: string) {
    const variable = <IdirectionItem>{};
    variable.enabled = true;
    variable.name = name;
    variable.value = value;
    variable.direction = direction;

    this.store.dispatch(new AddSofiaProfileGatewayVar({id: parentId, index: index, variable: variable}));
  }

  deleteProfileGatewayVar(param: Iitem) {
    this.store.dispatch(new DelSofiaProfileGatewayVar({variable: param}));
  }

  addNewProfileGatewayVar(grandPrentId: number, parentId: number) {
    this.store.dispatch(new StoreNewSofiaProfileGatewayVar({profileId: grandPrentId, id: parentId}));
  }

  dropNewProfileGatewayVar(grandPrentId: number, parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewSofiaProfileGatewayVar({profileId: grandPrentId, id: parentId, index: index}));
  }

  getSofiaProfileDomain(id: number) {
    this.panelCloser['domain' + id] = true;
    this.store.dispatch(new GetSofiaProfileDomains({id: id}));
  }

  updateProfileDomain(domain: Idomain) {
    const newDomain = {...domain};
    // Ensure boolean conversion for alias and parse fields
    newDomain.alias = (typeof domain.alias === 'string') ? (<string><any>domain.alias).toLowerCase() === 'true' : domain.alias;
    newDomain.parse = (typeof domain.parse === 'string') ? (<string><any>domain.parse).toLowerCase() === 'true' : domain.parse;
    this.store.dispatch(new UpdateSofiaProfileDomain({sofia_domain: newDomain}));
  }

  switchProfileDomain(domain: Idomain) {
    const newPDomain = {...domain};
    newPDomain.enabled = !newPDomain.enabled;
    this.store.dispatch(new SwitchSofiaProfileDomain({sofia_domain: newPDomain}));
  }

  newProfileDomain(parentId: number, index: number, name: string, alias: string | boolean, parse: string | boolean) {
    const domain = <Idomain>{};
    domain.enabled = true;
    domain.name = name;
    // Ensure boolean conversion for new item
    domain.alias = (typeof alias === 'string') ? <string>String(alias).toLowerCase() === 'true' : !!alias;
    domain.parse = (typeof parse === 'string') ? <string>String(parse).toLowerCase() === 'true' : !!parse;

    this.store.dispatch(new AddSofiaProfileDomain({id: parentId, index: index, sofia_domain: domain}));
  }

  deleteProfileDomain(domain: Idomains) {
    this.store.dispatch(new DelSofiaProfileDomain({sofia_domain: domain}));
  }

  addNewProfileDomain(parentId: number) {
    this.store.dispatch(new StoreNewSofiaProfileDomain({profileId: parentId}));
  }

  dropNewProfileDomain(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewSofiaProfileDomain({profileId: parentId, index: index}));
  }

  getSofiaProfileAlias(id: number) {
    this.panelCloser['alias' + id] = true;
    this.store.dispatch(new GetSofiaProfileAliases({id: id}));
  }

  updateProfileAlias(alias: Ialias) {
    const newAlias = {...alias};
    this.store.dispatch(new UpdateSofiaProfileAlias({sofia_alias: newAlias}));
  }

  switchProfileAlias(alias: Ialias) {
    const newPAlias = <Ialias>{...alias};
    newPAlias.enabled = !newPAlias.enabled;
    this.store.dispatch(new SwitchSofiaProfileAlias({sofia_alias: newPAlias}));
  }

  newProfileAlias(parentId: number, index: number, name: string) {
    const alias = <Ialias>{};
    alias.enabled = true;
    alias.name = name;

    this.store.dispatch(new AddSofiaProfileAlias({id: parentId, index: index, sofia_alias: alias}));
  }

  deleteProfileAlias(alias: Ialias) {
    this.store.dispatch(new DelSofiaProfileAlias({sofia_alias: alias}));
  }

  addNewProfileAlias(grandPrentId: number, parentId: number) {
    this.store.dispatch(new StoreNewSofiaProfileAlias({profileId: grandPrentId, id: parentId}));
  }

  dropNewProfileAlias(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewSofiaProfileAlias({profileId: parentId, index: index}));
  }

  onProfileSubmit() {
    this.store.dispatch(new AddSofiaProfile({name: this.newProfileName}));
  }

  onGatewaySubmit() {
    this.store.dispatch(new AddSofiaProfileGateway({name: this.newGatewayName, id: this.profileId}));
  }

  profileComand(id: number, subId: number, command: string) {
    this.store.dispatch(new SofiaProfileCommand({name: command, id: id, id_int: Number(subId)}));
  }

  switchProfile(profile: Iprofile) {
    this.store.dispatch(new SwitchSofiaProfile({id: profile.id, enabled: !profile.enabled}));
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
    this.store.dispatch(new StorePasteSofiaProfileParams({from_id: this.toCopyProfile, to_id: to}));
  }

  copyGateway(key: number, id: number) {
    const profile = this.list().profiles?.[key];
    if (!profile || !profile.gateways || !profile.gateways[id]) {
      this.toCopyProfileGateway = 0;
      this.toCopyGateway = 0;
      return;
    }
    this.toCopyProfileGateway = key;
    this.toCopyGateway = id;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteGatewayParams(profileId: number, to: number) {
    this.store.dispatch(new StorePasteSofiaProfileGatewayParams({
      from_profile: this.toCopyProfileGateway,
      from_id: this.toCopyGateway,
      to_id: to,
      id: profileId
    }));
  }

  pasteGatewayVars(profileId: number, to: number) {
    this.store.dispatch(new StorePasteSofiaProfileGatewayVars({
      from_profile: this.toCopyProfileGateway,
      from_id: this.toCopyGateway,
      to_id: to,
      id: profileId
    }));
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
        this.store.dispatch(new DelSofiaProfile({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameSofiaProfile({id: id, name: newName}));
      }
    });
  }

  openBottomSheetGateway(id: number, newName: string, oldName: string, action: 'delete' | 'rename'): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete gateway "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename gateway "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelSofiaProfileGateway({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameSofiaProfileGateway({id: id, name: newName}));
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
