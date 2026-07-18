import {Component, OnDestroy, OnInit, inject, signal, computed, effect} from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';

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
  GetDirectoryUsers,
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
import {ConfirmationService} from '../../../services/confirmation.service';
import {Idetails, IdirectionItem} from '../../../store/directory/directory.reducers';
import {ToastService} from '../../../services/toast.service';
import {ActivatedRoute, RouterLink} from '@angular/router';
import {KeyValuePadComponent} from "../../key-value-pad/key-value-pad.component";
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {TabNavComponent} from '../../tab-nav/tab-nav.component';
import {DisclosureComponent} from '../../disclosure/disclosure.component';
import {CpbxSelectDirective} from '../../../directives/cpbx-select.directive';
import {IconComponent} from '../../icon/icon.component';
import {KeyValuePad2Component} from '../../key-value-pad-2/key-value-pad-2.component';
import {TranslocoPipe, TranslocoService} from '@jsverse/transloco';

@Component({
  standalone: true,
  imports: [FormsModule, RouterLink, KeyValuePadComponent, KeyValuePad2Component, InnerHeaderComponent, CpbxSelectDirective, TabNavComponent, DisclosureComponent, IconComponent, TranslocoPipe],
  selector: 'app-gateways',
  templateUrl: './gateways.component.html',
  styleUrls: ['./gateways.component.css']
})
export class GatewaysComponent implements OnInit, OnDestroy {

  // --- Dependency Injection using inject() ---
  protected store = inject(Store<AppState>);
  private bottomSheet = inject(ConfirmationService);
  private _snackBar = inject(ToastService);
  private route = inject(ActivatedRoute);
  private i18n = inject(TranslocoService);

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
  public gatewayVariableMask = {
    name: {name: 'name'},
    value: {name: 'value'},
    extraField1: {name: 'direction', size: 'sm' as const, options: ['', 'inbound', 'outbound']},
  };
  public gatewayVariableDispatchers = {
    addNewItemField: (id: number) => this.addVarField(id),
    switchItem: (variable: IdirectionItem) => this.switchVar(variable),
    addItem: (id: number, index: number, name: string, value: string, direction: string) =>
      this.newVar(id, index, name, value, direction),
    dropNewItem: (id: number, index: number) => this.dropNewVar(id, index),
    deleteItem: (variable: IdirectionItem) => this.deleteVar(variable),
    updateItem: (variable: IdirectionItem) => this.updateVar(variable),
    pasteItems: (id: number) => this.pasteVars(id),
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

  ngOnInit() {
    if (Object.keys(this.userList()).length === 0) {
      this.store.dispatch(new GetDirectoryUsers(null));
    }
  }

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
    this._snackBar.open(this.i18n.translate('common.copied'), null, {
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
          case1Text: this.i18n.translate('directory.confirmDeleteGateway', {name: oldName}),
          case2Text: this.i18n.translate('directory.confirmRenameGateway', {oldName, newName}),
        }
    };
    const sheet = this.bottomSheet.open(config);
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
