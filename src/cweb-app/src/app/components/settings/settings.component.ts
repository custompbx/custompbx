import {Component, OnDestroy, OnInit} from '@angular/core';
import {select, Store} from '@ngrx/store';
import {AppState, selectDirectoryState, selectSettingsState} from '../../store/app.states';
import {Observable, Subscription} from 'rxjs';
import {
  AddWebUser,
  ClearWebUserAvatar,
  DeleteWebUser,
  GetSettings,
  GetWebUsers,
  RenameWebUser,
  SetSettings,
  SwitchWebUser,
  UpdateWebUserAvatar,
  UpdateWebUserLang,
  UpdateWebUserPassword,
  UpdateWebUserSipUser,
  UpdateWebUserStun,
  UpdateWebUserVertoWs,
  UpdateWebUserWebRTCLib,
  UpdateWebUserWs,
  GetUserTokens,
  RemoveUserToken,
  AddUserToken,
  UpdateWebUserGroup,
  GetWebDirectoryUsersTemplates,
  AddWebDirectoryUsersTemplate,
  GetWebDirectoryUsersTemplateParameters,
  GetWebDirectoryUsersTemplateVariables,
  DelWebDirectoryUsersTemplate,
  AddWebDirectoryUsersTemplateParameter,
  UpdateWebDirectoryUsersTemplateParameter,
  StoreNewWebDirectoryUsersTemplateParameter,
  DelWebDirectoryUsersTemplateParameter,
  StoreDelNewWebDirectoryUsersTemplateParameter,
  UpdateWebDirectoryUsersTemplate,
  StoreDelNewWebDirectoryUsersTemplateVariable,
  DelWebDirectoryUsersTemplateVariable,
  StoreNewWebDirectoryUsersTemplateVariable, UpdateWebDirectoryUsersTemplateVariable, AddWebDirectoryUsersTemplateVariable,
} from '../../store/settings/settings.actions';
import {ConfirmBottomSheetComponent} from '../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {AbstractControl, UntypedFormGroup} from '@angular/forms';
import {GetDirectoryDomains, GetDirectoryUsers} from '../../store/directory/directory.actions';
import {Isettings, IwebUser} from '../../store/settings/settings.reducers';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit, OnDestroy {
  public formData: Observable<any>;
  public directory: Observable<any>;
  public formData$: Subscription;
  public directory$: Subscription;
  public settings: Isettings = <Isettings>{};
  public users: { [id: number]: IwebUser };
  public usersTemplates: {};
  public usersTemplateParameters: {};
  public newParameters: {};
  public usersTemplateVariables: {};
  public newVariables: {};
  public usersGroups: {};
  public directoryDomains: object;
  private directoryUsers: object;
  public loadCounter: number;
  private newUserName: string;
  private newUserFormSent: UntypedFormGroup;
  private userLangs: Array<string>;
  private wssUris: Array<string>;
  private vertoWsUris: Array<string>;
  public login: string;
  public password: string;
  public groupId: 0;
  public acceptFile: 'image/*';
  public columns: Array<string>;
  public templateName: string;
  public domainId: number;

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
  ) {
    this.userLangs = ['EN', 'RU'];
    this.formData = this.store.pipe(select(selectSettingsState));
    this.directory = this.store.pipe(select(selectDirectoryState));
  }

  ngOnInit() {
    this.formData$ = this.formData.subscribe((settings) => {
      this.loadCounter = settings.loadCounter;
      this.settings = settings.settingsData;
      this.users = settings.webUsers;
      this.usersTemplates = settings.webUsersTemplates;
      this.usersTemplateParameters = settings.webUsersTemplateParameters;
      this.usersTemplateVariables = settings.webUsersTemplateVariables;
      this.newParameters = settings.newWUTPs;
      this.newVariables = settings.newWUTVs;
      this.usersGroups = settings.webGroups;
      this.wssUris = settings.wssUris;
      this.vertoWsUris = settings.vertoWsUris;
      this.columns = ['token', 'created', 'purpose'];
      if (this.newUserFormSent) {
        this.newUserFormSent.reset();
        this.newUserFormSent = null;
      }
      if (settings.errorMessage) {
        this._snackBar.open('Error: ' + settings.errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.directory$ = this.directory.subscribe((users) => {
      this.directoryDomains = users.domains;
      this.directoryUsers = users.users;
      if (users.errorMessage) {
        this._snackBar.open('Error: ' + users.errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  getConnectionsSettings() {
    this.store.dispatch(new GetSettings(null));
  }

  getWebUsers() {
    if (Object.entries(this.directoryUsers || {}).length === 0 && this.directoryUsers.constructor === Object) {
      this.store.dispatch(new GetDirectoryUsers(null));
    }
    this.store.dispatch(new GetWebUsers(null));
  }

  onSubmit(f) {
    f.form.markAsPristine();
    this.store.dispatch(new SetSettings(this.settings));
  }

  onNewUserSubmit(f) {
    this.newUserFormSent = f.form;
    this.store.dispatch(new AddWebUser(f.value));
  }

  updatePassword(id, pass) {
    this.store.dispatch(new UpdateWebUserPassword({password: pass, id: id}));
  }

  updateLang(id, pass) {
    this.store.dispatch(new UpdateWebUserLang({param_id: pass, id: id}));
  }

  updateGroup(id, pass) {
    this.store.dispatch(new UpdateWebUserGroup({group_id: pass, id: id}));
  }

  updateSipUser(id, pass) {
    this.store.dispatch(new UpdateWebUserSipUser({param_id: pass, id: id}));
  }

  updateWs(id, ws) {
    this.store.dispatch(new UpdateWebUserWs({value: ws, id: id}));
  }

  updateVertoWs(id, ws) {
    this.store.dispatch(new UpdateWebUserVertoWs({value: ws, id: id}));
  }

  updateWebRTCLib(id, lib) {
    this.store.dispatch(new UpdateWebUserWebRTCLib({value: lib, id: id}));
  }

  updateStun(id, stun) {
    this.store.dispatch(new UpdateWebUserStun({value: stun, id: id}));
  }

  switchWebUser(object) {
    this.store.dispatch(new SwitchWebUser({Enabled: !object.enabled, id: object.id}));
  }
  AddUserToken(id) {
    this.store.dispatch(new AddUserToken({id: id}));
  }
  ngOnDestroy() {
    this.formData$.unsubscribe();
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  trackByFnId(index, item) {
    return item.value.id;
  }

  trackById(index, item) {
    return item.id;
  }

  objectLength (obj: object): number {
    return Object.keys(obj).length;
  }

  isvalueReadyToSend(valueObject: AbstractControl): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  getLangIndex(index: number): number {
    if (this.userLangs.length <= index ) {
      return 0;
    }
    return index;
  }

  openBottomSheet(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete user "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename user "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DeleteWebUser({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameWebUser({id: id, name: newName}));
      }
    });
  }

  chooseAvatar(event, id) {
    if (!event.files[0] || !event.files[0].size || event.files[0].size > 512000) {
      this._snackBar.open('Error: bad file size!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
      return;
    }
    const file_reader = new FileReader();
    file_reader.onload = function(evt) {
      this.store.dispatch(new UpdateWebUserAvatar({file: file_reader.result, id: id}));
    }.bind(this);
    file_reader.readAsDataURL(event.files[0]);
  }

  clearAvatar(id) {
      this.store.dispatch(new ClearWebUserAvatar({file: '', id: id}));
  }

  getUserTokens(id) {
      this.store.dispatch(new GetUserTokens({id: id}));
  }

  openBottomSheetRemoveToken(id): void {
    const config = {
      data:
        {
          case1Text: 'Are you sure you want to remove this token?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      this.store.dispatch(new RemoveUserToken({id: id}));
    });
  }
  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  getWebUsersTemplates() {
    this.store.dispatch(new GetDirectoryDomains(null));
    this.store.dispatch(new GetWebDirectoryUsersTemplates(null));
  }
  AddWebUsersTemplates() {
    this.store.dispatch(
      new AddWebDirectoryUsersTemplate({data: {name: this.templateName, domain: {id: this.domainId}}})
    );
  }
  delWebUsersTemplates(id: number) {
    this.store.dispatch(new DelWebDirectoryUsersTemplate({id: id}));
  }
  getWebUsersTemplateDetails(id: number) {
    this.store.dispatch(new GetWebDirectoryUsersTemplateParameters({id: id}));
    this.store.dispatch(new GetWebDirectoryUsersTemplateVariables({id: id}));
  }
  UpdateWebDirectoryUsersTemplate(template) {
    this.store.dispatch(new UpdateWebDirectoryUsersTemplate({data: template}));
  }

  DelWebDirectoryUsersTemplate(name, id): void {
    const config = {
      data:
        {
          action: 'delete',
          case1Text: 'Are you sure you want to delete template "' + name + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      this.store.dispatch(new DelWebDirectoryUsersTemplate({id: id}));
    });
  }

  UpdateWebDirectoryUsersTemplateParameter(param: object, key) {
    if (key) {
      param = {...param};
      param[key] = !param[key];
    }

    this.store.dispatch(
      new UpdateWebDirectoryUsersTemplateParameter({data: param})
    );
  }
  AddWebDirectoryUsersTemplateParameter(param: object, id: number, index: number) {
    param = {...param, parent: {id}};
    this.store.dispatch(
      new AddWebDirectoryUsersTemplateParameter({data: param, index: index})
    );
  }
  DelWebDirectoryUsersTemplateParameter(id: number) {
    this.store.dispatch(
      new DelWebDirectoryUsersTemplateParameter({id: id})
    );
  }
  StoreNewWebDirectoryUsersTemplateParameter(id: number) {
    this.store.dispatch(
      new StoreNewWebDirectoryUsersTemplateParameter({id: id})
    );
  }
  StoreDelNewWebDirectoryUsersTemplateParameter(id: number, index: number) {
    this.store.dispatch(
      new StoreDelNewWebDirectoryUsersTemplateParameter({id: id, index: index})
    );
  }

  UpdateWebDirectoryUsersTemplateVariable(variable: object, key) {
    if (key) {
      variable = {...variable};
      variable[key] = !variable[key];
    }

    this.store.dispatch(
      new UpdateWebDirectoryUsersTemplateVariable({data: variable})
    );
  }
  AddWebDirectoryUsersTemplateVariable(variable: object, id: number, index: number) {
    variable = {...variable, parent: {id}};
    this.store.dispatch(
      new AddWebDirectoryUsersTemplateVariable({data: variable, index: index})
    );
  }
  DelWebDirectoryUsersTemplateVariable(id: number) {
    this.store.dispatch(
      new DelWebDirectoryUsersTemplateVariable({id: id})
    );
  }
  StoreNewWebDirectoryUsersTemplateVariable(id: number) {
    this.store.dispatch(
      new StoreNewWebDirectoryUsersTemplateVariable({id: id})
    );
  }
  StoreDelNewWebDirectoryUsersTemplateVariable(id: number, index: number) {
    this.store.dispatch(
      new StoreDelNewWebDirectoryUsersTemplateVariable({id: id, index: index})
    );
  }

  onlyValuesByParent(obj: object, parentId: number): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj).filter(u => u.parent.id === Number(parentId));
  }

}
