import {Component, OnDestroy, OnInit, ViewChild, inject, signal, computed, effect} from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';

import {MaterialModule} from "../../../../material-module";
import {select, Store} from '@ngrx/store';
import {AppState, selectDirectoryState} from '../../../store/app.states';
import {
  GetDirectoryUserDetails,
  AddDirectoryUserParameter,
  AddDirectoryUserVariable,
  StoreDeleteNewDirectoryUserParameter,
  StoreDeleteNewDirectoryUserVariable,
  StoreAddNewDirectoryUserParameter,
  StoreAddNewDirectoryUserVariable,
  DeleteDirectoryUserVariable,
  DeleteDirectoryUserParameter,
  UpdateDirectoryUserParameter,
  UpdateDirectoryUserVariable,
  UpdateDirectoryUserCache,
  UpdateDirectoryUserCidr,
  AddDirectoryUser,
  DeleteDirectoryUser,
  UpdateDirectoryUserName,
  SwitchDirectoryUser,
  SwitchDirectoryUserParameter,
  SwitchDirectoryUserVariable,
  StorePasteDirectoryUserVariables,
  StorePasteDirectoryUserParameters,
  ImportXMLDomainUser,
  UpdateDirectoryUserNumberAlias,
  GetWebDirectoryUsersTemplatesList,
  GetWebDirectoryUsersTemplateForm, CreateWebDirectoryUsersByTemplate
} from '../../../store/directory/directory.actions';
import {AbstractControl, FormsModule} from '@angular/forms';
import { MatBottomSheet } from '@angular/material/bottom-sheet';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {KeyValuePadComponent} from "../../key-value-pad/key-value-pad.component";
import {CodeEditorComponent} from "../../code-editor/code-editor.component";
import {State} from "../../../store/directory/directory.reducers";

@Component({
  standalone: true,
  imports: [MaterialModule, FormsModule, InnerHeaderComponent, KeyValuePadComponent, CodeEditorComponent],
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit, OnDestroy {

  // --- Dependency Injection using inject() ---
  protected store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private directoryState = toSignal(
    this.store.pipe(select(selectDirectoryState)),
    {
      // Initialize with a known structure to prevent runtime errors
      initialValue: {
        domains: {}, // Corresponds to `this.list`
        users: {},   // Corresponds to `this.userList`
        webUsersTemplates: {}, // Corresponds to `this.usersTemplates`
        templatesItems: {}, // Corresponds to `this.templatesItems`
        errorMessage: null,
        loadCounter: 0
      } as State
    }
  );

  // --- Computed State for Template Access ---
  public list = computed(() => this.directoryState().domains || {}); // All domains
  public userList = computed(() => this.directoryState().users || {}); // All user details
  public usersTemplates = computed(() => this.directoryState().webUsersTemplates || {});
  public templatesItems = computed(() => this.directoryState().templatesItems || {});
  public loadCounter = computed(() => this.directoryState().loadCounter || 0);

  // --- Local State as Signals/Properties ---
  public newUserName = signal('');
  public bulkUsers: number;
  public domainId: number;
  public selectedIndex: number = 0;
  public toCopy: number;
  public userParamDispatchers: object;
  public userVarDispatchers: object;
  public domainIds: Array<number> = [];
  public XMLBody: string;
  public editorInited: boolean;

  // --- Effect for Side Effects (Replaces Subscription logic) ---
  private userEffect = effect(() => {
    const errorMessage = this.directoryState().errorMessage;

    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    } else {
      this.newUserName.set('');
    }
  });

  ngOnInit() {
    this.userParamDispatchers = {
      addItemField: StoreAddNewDirectoryUserParameter,
      switchItem: SwitchDirectoryUserParameter,
      newItem: AddDirectoryUserParameter,
      dropNewItem: StoreDeleteNewDirectoryUserParameter,
      deleteItem: DeleteDirectoryUserParameter,
      updateItem: UpdateDirectoryUserParameter,
      pasteItems: StorePasteDirectoryUserParameters,
    };
    this.userVarDispatchers = {
      addItemField: StoreAddNewDirectoryUserVariable,
      switchItem: SwitchDirectoryUserVariable,
      newItem: AddDirectoryUserVariable,
      dropNewItem: StoreDeleteNewDirectoryUserVariable,
      deleteItem: DeleteDirectoryUserVariable,
      updateItem: UpdateDirectoryUserVariable,
      pasteItems: StorePasteDirectoryUserVariables,
    };
  }

  pasteVars(to: number) {
    this.store.dispatch(new StorePasteDirectoryUserVariables({from_id: this.toCopy, to_id: to}));
  }

  pasteParams(to: number) {
    this.store.dispatch(new StorePasteDirectoryUserParameters({from_id: this.toCopy, to_id: to}));
  }


  ngOnDestroy() {
    // Subscription to `this.users$` is no longer needed/present due to toSignal.
    // We only keep the router data cleanup logic if necessary.
    if (this.route.snapshot?.data?.reconnectUpdater) {
      this.route.snapshot.data.reconnectUpdater.unsubscribe();
    }
  }

  getDetails(id) {
    this.store.dispatch(new GetDirectoryUserDetails({id: id}));
  }

  clearDetails(id) {
    //  this.store.dispatch(new ClearDetails(id));
  }

  switchUser(object) {
    this.store.dispatch(new SwitchDirectoryUser({id: object.id, enabled: !object.enabled}));
  }

  updateCache(userId: number, valueObject: AbstractControl) {
    const value = valueObject.value;
    valueObject.reset();
    this.store.dispatch(new UpdateDirectoryUserCache({value: value, id: userId}));
  }

  updateCidr(userId: number, valueObject: AbstractControl) {
    const value = valueObject.value;
    valueObject.reset();
    this.store.dispatch(new UpdateDirectoryUserCidr({value: value, id: userId}));
  }

  updateNumberAlias(userId: number, valueObject: AbstractControl) {
    const value = valueObject.value;
    valueObject.reset();
    this.store.dispatch(new UpdateDirectoryUserNumberAlias({value: value, id: userId}));
  }

  isvalueReadyToSend(valueObject: AbstractControl): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  onUserSubmit() {
    // Read the signal value using newUserName()
    this.store.dispatch(new AddDirectoryUser({name: this.newUserName(), id: this.domainId, bulk: this.bulkUsers}));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  trackByFnId(index, item) {
    return item.id;
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  ImportXMLUser() {
    this.store.dispatch(new ImportXMLDomainUser({file: this.XMLBody, id: this.domainId}));
  }

  copy(key) {
    // Read computed signal value
    if (!this.userList()[key]) {
      this.toCopy = 0;
      return;
    }
    this.toCopy = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
      // horizontalPosition: 'right',
      // verticalPosition: 'top',
    });
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
        this.store.dispatch(new DeleteDirectoryUser({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateDirectoryUserName({id: id, name: newName}));
      }
    });
  }

  // Domain list filtered by domainIds
  public filteredDomains = computed(() => {
    const list = this.list(); // Read the domains list from state
    const domainIds = this.domainIds; // Read the local filter array

    if (domainIds.length === 0) {
      return list;
    }

    const res: object = {};
    if (!list) {
      return res;
    }

    // Note: The logic in the original domainFilter iterates over Object.keys(list)
    // which is fine, but we can stick to the computed signal approach for reactivity.
    Object.keys(list).forEach(
      key => {
        if (domainIds.includes(Number(key))) {
          res[key] = list[key];
        }
      }
    );

    return res;
  });

  // Replaced domainFilter() with filteredDomains computed signal,
  // but kept the original function signature for compatibility if it's called outside the template.
  domainFilter (): object {
    return this.filteredDomains();
  }

  objectLenght(obj: object): number {
    return Object.keys(obj).length;
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  onlyValuesByParent(obj: object, parentId: number): Array<any> {
    if (!obj) {
      return [];
    }
    // Access userList computed signal here
    return Object.values(obj).filter((u: any) => u.parent && u.parent.id === parentId);
  }

  initEditor() {
    this.editorInited = true;
  }

  trackById(index, item) {
    return item.id;
  }

  GetWebDirectoryUsersTemplatesList() {
    this.store.dispatch(new GetWebDirectoryUsersTemplatesList(null));
  }

  GetWebDirectoryUsersTemplateForm(id) {
    this.store.dispatch(new GetWebDirectoryUsersTemplateForm({id}));
  }

  onTemplateUserSubmit(data) {
    this.store.dispatch(new CreateWebDirectoryUsersByTemplate({data}));
  }

}
