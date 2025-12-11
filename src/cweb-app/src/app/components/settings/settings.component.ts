import { Component, inject, effect, computed } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { CommonModule } from '@angular/common';
import { MaterialModule } from "../../../material-module";
import { select, Store } from '@ngrx/store';
import { AppState, selectDirectoryState, selectSettingsState } from '../../store/app.states';
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
import { MatBottomSheet } from '@angular/material/bottom-sheet';
import { MatSnackBar } from '@angular/material/snack-bar';
import { AbstractControl, FormsModule, UntypedFormGroup } from '@angular/forms';
import { GetDirectoryDomains, GetDirectoryUsers } from '../../store/directory/directory.actions';
import { Isettings, IwebUser } from '../../store/settings/settings.reducers';
import { InnerHeaderComponent } from "../inner-header/inner-header.component";
import {ResizeInputDirective} from "../../directives/resize-input.directive";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ResizeInputDirective],
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
// Removed OnInit and OnDestroy as toSignal handles subscriptions and cleanup
export class SettingsComponent {

  // Injectable services
  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);

  // --- NgRx State converted to Signals (toSignal handles subscriptions automatically) ---

  // Initial value is used to ensure signals always return a value
  private settingsState = toSignal(this.store.pipe(select(selectSettingsState)), { initialValue: {} as any });
  private directoryState = toSignal(this.store.pipe(select(selectDirectoryState)), { initialValue: {} as any });

  // --- Computed Signals (Derived State) ---
  public loadCounter = computed(() => this.settingsState().loadCounter || 0);
  public settings = computed<Isettings>(() => this.settingsState().settingsData || {} as Isettings);
  public users = computed<{ [id: number]: IwebUser }>(() => this.settingsState().webUsers || {});
  public usersTemplates = computed(() => this.settingsState().webUsersTemplates || {});
  public usersTemplateParameters = computed(() => this.settingsState().webUsersTemplateParameters || {});
  public newParameters = computed(() => this.settingsState().newWUTPs || {});
  public usersTemplateVariables = computed(() => this.settingsState().webUsersTemplateVariables || {});
  public newVariables = computed(() => this.settingsState().newWUTVs || {});
  public usersGroups = computed(() => this.settingsState().webGroups || {});
  public wssUris = computed(() => this.settingsState().wssUris || []);
  public vertoWsUris = computed(() => this.settingsState().vertoWsUris || []);
  public settingsError = computed(() => this.settingsState().errorMessage || null);

  public directoryDomains = computed(() => this.directoryState().domains || {});
  // Renamed from private directoryUsers to public directoryUsersSignal for clarity
  public directoryUsersSignal = computed(() => this.directoryState().users || {});
  public directoryError = computed(() => this.directoryState().errorMessage || null);

  // --- Local Component State & Constants ---
  public userLangs: Array<string> = ['EN', 'RU'];
  public login: string = '';
  public password: string = '';
  public groupId: 0 = 0;
  public acceptFile: 'image/*' = 'image/*';
  public columns: Array<string> = ['token', 'created', 'purpose'];
  public templateName: string = '';
  public domainId: number = 0;
  public newUserFormSent: UntypedFormGroup | null = null; // Stays as a mutable property

  // --- Effects for Side Effects (Error Handling, Form Reset) ---

  // Handles errors from the settings state and form reset logic
  private settingsSideEffect = effect(() => {
    const settingsError = this.settingsError();
    if (settingsError) {
      this._snackBar.open('Error: ' + settingsError + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }

    // Logic to reset the form after a successful user submission
    // We observe the users signal to detect changes that likely indicate success.
    const users = this.users();
    if (this.newUserFormSent) {
      this.newUserFormSent.reset();
      this.newUserFormSent = null;
    }
  });

  // Handles errors from the directory state
  private directorySideEffect = effect(() => {
    const directoryError = this.directoryError();
    if (directoryError) {
      this._snackBar.open('Error: ' + directoryError + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  // Constructor is now clean as dependencies are injected using inject()
  constructor() { }

  getConnectionsSettings() {
    this.store.dispatch(new GetSettings(null));
  }

  getWebUsers() {
    // Accessing signal value with parentheses
    if (Object.entries(this.directoryUsersSignal()).length === 0 && this.directoryUsersSignal().constructor === Object) {
      this.store.dispatch(new GetDirectoryUsers(null));
    }
    this.store.dispatch(new GetWebUsers(null));
  }

  onSubmit(f: any) {
    f.form.markAsPristine();
    this.store.dispatch(new SetSettings(this.settings())); // Accessing signal value with parentheses
  }

  onNewUserSubmit(f: any) {
    this.newUserFormSent = f.form;
    this.store.dispatch(new AddWebUser(f.value));
  }

  updatePassword(id: number, pass: string) {
    this.store.dispatch(new UpdateWebUserPassword({password: pass, id: id}));
  }

  updateLang(id: number, pass: string) {
    this.store.dispatch(new UpdateWebUserLang({param_id: pass, id: id}));
  }

  updateGroup(id: number, pass: number) {
    this.store.dispatch(new UpdateWebUserGroup({group_id: pass, id: id}));
  }

  updateSipUser(id: number, pass: string) {
    this.store.dispatch(new UpdateWebUserSipUser({param_id: pass, id: id}));
  }

  updateWs(id: number, ws: string) {
    this.store.dispatch(new UpdateWebUserWs({value: ws, id: id}));
  }

  updateVertoWs(id: number, ws: string) {
    this.store.dispatch(new UpdateWebUserVertoWs({value: ws, id: id}));
  }

  updateWebRTCLib(id: number, lib: string) {
    this.store.dispatch(new UpdateWebUserWebRTCLib({value: lib, id: id}));
  }

  updateStun(id: number, stun: string) {
    this.store.dispatch(new UpdateWebUserStun({value: stun, id: id}));
  }

  switchWebUser(object: any) {
    this.store.dispatch(new SwitchWebUser({Enabled: !object.enabled, id: object.id}));
  }

  AddUserToken(id: number) {
    this.store.dispatch(new AddUserToken({id: id}));
  }

  // ngOnDestroy is removed

  checkDirty(condition: AbstractControl | null): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  objectLength (obj: object): number {
    return Object.keys(obj).length;
  }

  isvalueReadyToSend(valueObject: AbstractControl | null): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  getLangIndex(index: number): number {
    if (this.userLangs.length <= index ) {
      return 0;
    }
    return index;
  }

  openBottomSheet(id: number, newName: string, oldName: string, action: string): void {
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

  chooseAvatar(event: any, id: number) {
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

  clearAvatar(id: number) {
    this.store.dispatch(new ClearWebUserAvatar({file: '', id: id}));
  }

  getUserTokens(id: number) {
    this.store.dispatch(new GetUserTokens({id: id}));
  }

  openBottomSheetRemoveToken(id: number): void {
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

  UpdateWebDirectoryUsersTemplate(template: any) {
    this.store.dispatch(new UpdateWebDirectoryUsersTemplate({data: template}));
  }

  DelWebDirectoryUsersTemplate(name: string, id: number): void {
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

  UpdateWebDirectoryUsersTemplateParameter(param: any, key: string | null) {
    if (key) {
      param = {...param};
      param[key] = !param[key];
    }

    this.store.dispatch(
      new UpdateWebDirectoryUsersTemplateParameter({data: param})
    );
  }

  AddWebDirectoryUsersTemplateParameter(param: any, id: number, index: number) {
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

  UpdateWebDirectoryUsersTemplateVariable(variable: any, key: string | null) {
    if (key) {
      variable = {...variable};
      variable[key] = !variable[key];
    }

    this.store.dispatch(
      new UpdateWebDirectoryUsersTemplateVariable({data: variable})
    );
  }

  AddWebDirectoryUsersTemplateVariable(variable: any, id: number, index: number) {
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
    return Object.values(obj).filter((u: any) => u.parent.id === Number(parentId));
  }
}
