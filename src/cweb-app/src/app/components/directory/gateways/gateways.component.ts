import {Component, OnDestroy, OnInit, inject, signal, computed, effect} from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';

import {MaterialModule} from "../../../../material-module";
import {select, Store} from '@ngrx/store';
import {AppState, selectDirectoryState} from '../../../store/app.states';
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

  SwitchDirectoryUserGatewayParameter,
} from '../../../store/directory/directory.actions';
import {AbstractControl, FormsModule} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {Idetails, IdirectionItem} from '../../../store/directory/directory.reducers';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {ActivatedRoute, RouterLink} from '@angular/router';
import {KeyValuePadComponent} from "../../key-value-pad/key-value-pad.component";

@Component({
  standalone: true,
  imports: [MaterialModule, FormsModule, RouterLink, KeyValuePadComponent],
  selector: 'app-gateways',
  templateUrl: './gateways.component.html',
  styleUrls: ['./gateways.component.css']
})
export class GatewaysComponent implements OnDestroy {

  // --- Dependency Injection using inject() ---
  protected store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private directoryState = toSignal(
    this.store.pipe(select(selectDirectoryState)),
    {
      initialValue: {
        domains: {},
        users: {},
        userGateways: {},
        gatewayDetails: {} as Idetails,
        errorMessage: null,
        loadCounter: 0
      } as any // Replace 'any' with the actual DirectoryState interface if available
    }
  );

  // --- Computed State for Template Access ---
  public list = computed(() => this.directoryState().domains || {});
  protected userList = computed(() => this.directoryState().users || {});
  protected gatewayList = computed(() => this.directoryState().userGateways || {});
  public listDetails = computed(() => this.directoryState().gatewayDetails as Idetails || {} as Idetails);
  public loadCounter = computed(() => this.directoryState().loadCounter || 0);

  // --- Local State as Signals/Properties ---
  public newGatewayName = signal('');
  public domainId: number;
  public userId: number;
  public selectedIndex: number = 0; // Initialized directly
  protected toCopy: number;
  public gatewayParamDispatchers: object = { // Initialized directly
    addItemField: StoreNewDirectoryUserGatewayParameter,
    switchItem: SwitchDirectoryUserGatewayParameter,
    newItem: AddDirectoryUserGatewayParameter,
    dropNewItem: DropNewDirectoryUserGatewayParameter,
    deleteItem: DeleteDirectoryUserGatewayParameter,
    updateItem: UpdateDirectoryUserGatewayParameter,
    pasteItems: StorePasteDirectoryUserGatewayParameters,
  };

  private gatewayEffect = effect(() => {
    const users = this.directoryState();
    const lastErrorMessage = users.errorMessage;

    if (lastErrorMessage) {
      this._snackBar.open('Error: ' + lastErrorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    } else {
      // If the error message clears, reset state
      this.newGatewayName.set('');
      // Reset selectedIndex if it was due to a successful action
      this.selectedIndex = this.selectedIndex === 1 ? 0 : this.selectedIndex;
    }
  });

  ngOnDestroy() {
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
    // Access signal value with ()
    console.log({name: this.newGatewayName(), id: this.userId});
    this.store.dispatch(new AddDirectoryUserGateway({name: this.newGatewayName(), id: this.userId}));
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
    // Access computed signal userList()
    return Object.values(obj).filter((u: any) => u.parent?.id === Number(parentId));
  }

  copy(key) {
    // Access computed signal gatewayList()
    if (!this.gatewayList()[key]) {
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
