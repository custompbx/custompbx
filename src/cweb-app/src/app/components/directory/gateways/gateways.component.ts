import {Component, Inject, OnDestroy, OnInit} from '@angular/core';
import {select, Store} from '@ngrx/store';
import {AppState, selectDirectoryState} from '../../../store/app.states';
import {Observable, Subscription} from 'rxjs';
import {
  GetDirectoryUserGatewayDetails,
  StoreNewDirectoryUserGatewayParameter,
  DropNewDirectoryUserGatewayParameter,
  AddDirectoryUserGatewayParameter,
  DeleteDirectoryUserGatewayParameter,
  UpdateDirectoryUserGatewayParameter,
  AddDirectoryUserGateway,
  DeleteDirectoryUserGateway,
  UpdateDirectoryUserGatewayName,
  UpdateDirectoryUserGatewayVariable,
  AddDirectoryUserGatewayVariable,
  StoreNewDirectoryUserGatewayVariable,
  DeleteDirectoryUserGatewayVariable,
  DropNewDirectoryUserGatewayVariable,
  SwitchDirectoryUserGatewayVariable,
  ImportDirectory,
  StorePasteDirectoryUserGatewayVariables,
  StorePasteDirectoryUserGatewayParameters,

  SwitchDirectoryDomainParameter, SwitchDirectoryUserGatewayParameter,
} from '../../../store/directory/directory.actions';
import {AbstractControl} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {Idetails, IdirectionItem} from '../../../store/directory/directory.reducers';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-gateways',
  templateUrl: './gateways.component.html',
  styleUrls: ['./gateways.component.css']
})
export class GatewaysComponent implements OnInit, OnDestroy {
  public users: Observable<any>;
  public users$: Subscription;
  public list: any;
  private userList: any;
  private gatewayList: any;
  public listDetails: Idetails;
  private newGatewayName: string;
  private domainId: number;
  private userId: number;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  private toCopy: number;
  public gatewayParamDispatchers: object;

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
      this.gatewayList = users.userGateways;
      this.listDetails = users.gatewayDetails;
      this.lastErrorMessage = users.errorMessage;
      if (!this.lastErrorMessage) {
        this.newGatewayName = '';
        // this.domainId = 0;
        this.selectedIndex = this.selectedIndex === 1 ? 0 : this.selectedIndex;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.gatewayParamDispatchers = {
      addItemField: StoreNewDirectoryUserGatewayParameter,
      switchItem: SwitchDirectoryUserGatewayParameter,
      newItem: AddDirectoryUserGatewayParameter,
      dropNewItem: DropNewDirectoryUserGatewayParameter,
      deleteItem: DeleteDirectoryUserGatewayParameter,
      updateItem: UpdateDirectoryUserGatewayParameter,
      pasteItems: StorePasteDirectoryUserGatewayParameters,
    };
  }

  ngOnDestroy() {
    this.users$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  getDetails(id) {
    this.store.dispatch(new GetDirectoryUserGatewayDetails({id: id}));
  }

  importDirectory() {
    this.store.dispatch(new ImportDirectory(null));
  }

  clearDetails(id) {
    //  this.store.dispatch(new ClearDetails(id));
  }

  addVarField(id) {
    this.store.dispatch(new StoreNewDirectoryUserGatewayVariable({id: id}));
  }

  dropNewVar(id: number, index: number) {
    this.store.dispatch(new DropNewDirectoryUserGatewayVariable({id: id, index: index}));
  }

  newVar(id: number, index: number, name: string, value: string, direction: string) {
    this.store.dispatch(new AddDirectoryUserGatewayVariable({id, index, name, value, direction}));
  }

  deleteVar(variable: IdirectionItem) {
    this.store.dispatch(new DeleteDirectoryUserGatewayVariable({id: variable.id}));
  }

  updateVar(variable) {
    this.store.dispatch(new UpdateDirectoryUserGatewayVariable({
      id: variable.id, name: variable.name, value: variable.value, direction: variable.direction
    }));
  }

  switchVar(variable: IdirectionItem) {
    this.store.dispatch(new SwitchDirectoryUserGatewayVariable({id: variable.id, enabled: !variable.enabled}));
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

  onUserSubmit() {
    console.log({name: this.newGatewayName, id: this.userId});
    this.store.dispatch(new AddDirectoryUserGateway({name: this.newGatewayName, id: this.userId}));
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

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  onlyValuesByParent(obj: object, parentId: number): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj).filter(u => u.parent?.id === Number(parentId));
  }

  copy(key) {
    if (!this.gatewayList[key]) {
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

  pasteVars(to: number) {
    this.store.dispatch(new StorePasteDirectoryUserGatewayVariables({from_id: this.toCopy, to_id: to}));
  }

  openBottomSheet(id, newName, oldName, action): void {
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
        this.store.dispatch(new DeleteDirectoryUserGateway({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateDirectoryUserGatewayName({id: id, name: newName}));
      }
    });
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

}
