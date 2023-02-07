import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {select, Store} from '@ngrx/store';
import {AppState, selectDirectoryState} from '../../../store/app.states';
import {Observable, Subscription} from 'rxjs';
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
import {AbstractControl} from '@angular/forms';
import { MatBottomSheet } from '@angular/material/bottom-sheet';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit, OnDestroy {

  public users: Observable<any>;
  public users$: Subscription;
  public list: any;
  public userList: any;
  public newUserName: string;
  public bulkUsers: number;
  public domainId: number;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public toCopy: number;
  public userParamDispatchers: object;
  public userVarDispatchers: object;
  public domainIds: Array<number> = [];
  public XMLBody: string;
  public editorInited: boolean;
  public usersTemplates: {[id: number]: object};
  public templatesItems: {
    [id: number]:
      {
        id: number;
        name: string;
        parameters: Array<{ id: number; name: string; value: string; description: string; disabled: boolean; editable: boolean; }>;
        variables: Array<{ id: number; name: string; value: string; description: string; disabled: boolean; editable: boolean; }>;
      };
  };

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.users = this.store.pipe(select(selectDirectoryState));
  }

  ngOnInit() {
    this.users$ = this.users.subscribe((users) => {
      this.loadCounter = users.loadCounter;
      this.list = users.domains;
      this.userList = users.users;
      this.usersTemplates = users.webUsersTemplates;
      this.templatesItems = users.templatesItems;
      this.lastErrorMessage = users.errorMessage;
      if (!this.lastErrorMessage) {
        this.newUserName = '';
        // this.domainId = 0;
        // this.selectedIndex = this.selectedIndex === 1 ? 0 : this.selectedIndex;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
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
    this.users$.unsubscribe();
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
    this.store.dispatch(new AddDirectoryUser({name: this.newUserName, id: this.domainId, bulk: this.bulkUsers}));
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
    if (!this.userList[key]) {
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

  domainFilter (): object {
    if (this.domainIds.length === 0) {
      return this.list;
    }

    const res: object = {};
    if (!this.list) {
      return res;
    }
    Object.keys(this.list).forEach(
      key => {
        if (this.domainIds.includes(Number(key))) {
          res[key] = this.list[key];
        }
      }
    );

    return res;
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
    return Object.values(obj).filter(u => u.parent.id === parentId);
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
