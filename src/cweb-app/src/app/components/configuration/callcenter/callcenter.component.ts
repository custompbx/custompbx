import {Component, inject, computed, OnInit, effect, signal} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {ActivatedRoute} from '@angular/router';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {PageEvent} from '@angular/material/paginator';
import {
  AddCallcenterAgent,
  AddCallcenterQueue,
  AddCallcenterQueueParam,
  AddCallcenterSettings,
  AddCallcenterTier,
  DelCallcenterAgent,
  DelCallcenterMember,
  DelCallcenterQueue,
  DelCallcenterQueueParam,
  DelCallcenterSettings,
  DelCallcenterTier,
  GetCallcenterAgents,
  GetCallcenterMembers,
  GetCallcenterQueuesParams,
  GetCallcenterSettings,
  GetCallcenterTiers,
  ImportCallcenterAgentsAndTiers,
  RenameCallcenterQueue,
  SendCallcenterCommand,
  StoreDropNewCallcenterQueueParam,
  StoreDropNewCallcenterSettings,
  StoreNewCallcenterQueueParam,
  StoreNewCallcenterSettings,
  StorePasteCallcenterQueueParams, StoreSetChangedCallcenterTableField,
  SwitchCallcenterQueueParam,
  SwitchCallcenterSettings,
  UpdateCallcenterAgent,
  UpdateCallcenterQueueParam,
  UpdateCallcenterSettings,
  UpdateCallcenterTier,
} from '../../../store/config/callcenter/config.actions.callcenter';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {IfilterField, IsortField} from '../../cdr/cdr.component';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {AppAutoFocusDirective} from "../../../directives/auto-focus.directive";
import {Icallcenter, Iitem, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, AppAutoFocusDirective, KeyValuePad2Component],
  selector: 'app-callcenter',
  templateUrl: './callcenter.component.html',
  styleUrls: ['./callcenter.component.css']
})
export class CallcenterComponent implements OnInit {

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private configState = toSignal(
    this.store.pipe(select(selectConfigurationState)),
    {
      initialValue: {
        callcenter: {} as Icallcenter,
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().callcenter);
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().callcenter?.errorMessage || null);

  // Computed properties for table columns
  public columns = computed(() => {
    const agents = this.list().agents;
    return agents?.table?.length ? Object.keys(agents.table[0]) : [];
  });
  public tiersColumns = computed(() => {
    const tiers = this.list().tiers;
    return tiers?.table?.length ? Object.keys(tiers.table[0]) : [];
  });
  public membersColumns = computed(() => {
    const members = this.list().members;
    return members?.table?.length ? Object.keys(members.table[0]) : [];
  });

  // --- Local Component State (Signals) ---
  public showDel = signal<{[key: number]: boolean}>({});

  // --- Local Component State (Variables) ---
  public newQueueName: string = '';
  public newAgentName: string = '';
  public agentName: string = '';
  public selectedIndex: number = 0;
  protected queueId: number | null = null;
  protected panelCloser: {[key: string]: boolean} = {};
  protected toEdit: {[key: number]: any} = {};
  protected toEditTier: {[key: number]: any} = {};
  protected toCopyqueue: number = 0;

  public operands: Array<string> = ['=', '>', '<', 'CONTAINS'];
  // Agent table filtering/sorting
  public sortObject: IsortField = {fields: [], desc: false};
  public filter: IfilterField = {} as IfilterField;
  public filters: Array<IfilterField> = [];
  // Tier table filtering/sorting
  public tiersSortObject: IsortField = {fields: [], desc: false};
  public tiersFilter: IfilterField = {} as IfilterField;
  public tiersFilters: Array<IfilterField> = [];
  // Member table filtering/sorting
  public membersSortObject: IsortField = {fields: [], desc: false};
  public membersFilter: IfilterField = {} as IfilterField;
  public membersFilters: Array<IfilterField> = [];

  public paginationScale = [25, 50, 100, 250];
  public pageEvent: PageEvent = {
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  } as PageEvent;
  public tiersPageEvent: PageEvent = {
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  } as PageEvent;
  public membersPageEvent: PageEvent = {
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  } as PageEvent;

  public sortColumns: string | null = null;
  public tiersSortColumns: string | null = null;
  public membersSortColumns: string | null = null;

  public queueParamDispatchers: object;
  public globalSettingsDispatchers: object;
  public toEditAgentFilter: number | null = null;
  public toEditTierFilter: number | null = null;
  public toEditMemberFilter: number | null = null;

  // --- Effect for Side Effects (Error handling) ---
  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });

      // Reset new item names on error, as per original logic
      this.newQueueName = '';
      this.newAgentName = '';
      this.agentName = '';
    }
  });

  constructor() {
    // Dependencies are handled by inject()
  }

  ngOnInit() {
    // Initialize dispatchers
    this.queueParamDispatchers = {
      addNewItemField: this.StoreNewCallcenterQueueParam.bind(this),
      switchItem: this.SwitchCallcenterQueueParam.bind(this),
      addItem: this.AddCallcenterQueueParam.bind(this),
      dropNewItem: this.StoreDropNewCallcenterQueueParam.bind(this),
      deleteItem: this.DelCallcenterQueueParam.bind(this),
      updateItem: this.UpdateCallcenterQueueParam.bind(this),
      pasteItems: this.pasteProfileParams.bind(this),
    };
    this.globalSettingsDispatchers = {
      addNewItemField: this.StoreNewCallcenterSettings.bind(this),
      switchItem: this.SwitchCallcenterSettings.bind(this),
      addItem: this.AddCallcenterSettings.bind(this),
      dropNewItem: this.StoreDropNewCallcenterSettings.bind(this),
      deleteItem: this.DelCallcenterSettings.bind(this),
      updateItem: this.UpdateCallcenterSettings.bind(this),
      pasteItems: null,
    };
  }

  // ngOnDestroy is removed as toSignal handles cleanup automatically.

  checkDirty(condition: AbstractControl | null): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  toInput(id: number, columnName: any) {
    this.toEdit = {[id]: columnName};
  }

  toInputTier(id: number, columnName: any) {
    this.toEditTier = {[id]: columnName};
  }

  markChanged(id: number, columnName: any) {
    this.store.dispatch(new StoreSetChangedCallcenterTableField({tableName: 'agents', rowId: id, fieldName: columnName}));
  }

  markChangedTier(id: number, columnName: any) {
    this.store.dispatch(new StoreSetChangedCallcenterTableField({tableName: 'tiers', rowId: id, fieldName: columnName}));
  }

  isReadyToSendThree(mainObject: AbstractControl | null, object2: AbstractControl | null, object3: AbstractControl | null): boolean {
    return !!((mainObject && mainObject.valid && mainObject.dirty)
      || ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty)));
  }

  isvalueReadyToSend(valueObject: AbstractControl | null): boolean {
    return !!(valueObject && valueObject.dirty && valueObject.valid);
  }

  isReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return !!(nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid);
  }

  isReadyToSendOne(nameObject: AbstractControl | null): boolean {
    return !!(nameObject && nameObject.dirty && nameObject.valid);
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index: number, item: any): number {
    return index; // or item.id
  }

  trackByFnId(index: number, item: any): number {
    return item.id;
  }

  isNewReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return !!(nameObject && valueObject && nameObject.valid && valueObject.valid);
  }

  showDelIco(id: number) {
      this.showDel.update(current => {
        const updated = { ...current };
        updated[id] = true;
        return updated;
      });
  }

  leaveDelIco(id: number) {
    setTimeout(() => {
      // Update the signal immutably
      this.showDel.update(current => {
        const updated = { ...current };
        updated[id] = false;
        return updated;
      });
    }, 300);
  }

  copyQueue(key: number) {
    if (!this.list().queues[key]) {
      this.toCopyqueue = 0;
      return;
    }
    this.toCopyqueue = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  tabChanged(event: number) {
    this.panelCloser = {};
    if (event === 1/* || event === 4*/) {
      this.GetCallcenterAgents();
    } else if (event === 2) {
      this.GetCallcenterTiers();
    } else if (event === 3) {
      this.GetCallcenterMembers();
    } else if (event === 5) {
      this.GetCallcenterSettings();
    }
  }

  mainTabChanged(event: number) {
    if (event === 2) {
      // this.GetCallcenterAgents();
    }
  }

  pasteProfileParams(to: number) {
    this.store.dispatch(new StorePasteCallcenterQueueParams({from_id: this.toCopyqueue, to_id: to}));
  }

  GetCallcenterSettings() {
    this.store.dispatch(new GetCallcenterSettings(null));
  }

  clearGlobalSettings() {
  }

  UpdateCallcenterSettings(param: Iitem) {
    this.store.dispatch(new UpdateCallcenterSettings({param: param}));
  }

  SwitchCallcenterSettings(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchCallcenterSettings({param: newParam}));
  }

  DelCallcenterSettings(param: Iitem) {
    this.store.dispatch(new DelCallcenterSettings({param: param}));
  }

  AddCallcenterSettings(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };
    this.store.dispatch(new AddCallcenterSettings({index: index, param: param}));
  }

  StoreNewCallcenterSettings() {
    this.store.dispatch(new StoreNewCallcenterSettings(null));
  }

  StoreDropNewCallcenterSettings(index: number) {
    this.store.dispatch(new StoreDropNewCallcenterSettings({index: index}));
  }


  GetCallcenterQueuesParams(id: number) {
    this.panelCloser['profile' + id] = true;
    this.store.dispatch(new GetCallcenterQueuesParams({id: id}));
  }

  UpdateCallcenterQueueParam(param: Iitem) {
    this.store.dispatch(new UpdateCallcenterQueueParam({param: param}));
  }

  SwitchCallcenterQueueParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchCallcenterQueueParam({param: newParam}));
  }

  AddCallcenterQueueParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddCallcenterQueueParam({id: parentId, index: index, param: param}));
  }

  DelCallcenterQueueParam(param: Iitem) {
    this.store.dispatch(new DelCallcenterQueueParam({param: param}));
  }

  StoreNewCallcenterQueueParam(parentId: number) {
    this.store.dispatch(new StoreNewCallcenterQueueParam({id: parentId}));
  }

  StoreDropNewCallcenterQueueParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewCallcenterQueueParam({id: parentId, index: index}));
  }

  AddCallcenterQueue() {
    this.store.dispatch(new AddCallcenterQueue({name: this.newQueueName}));
  }

  ImportCallcenterAgentsAndTiers() {
    this.store.dispatch(new ImportCallcenterAgentsAndTiers(null));
  }

  handlePageEvent(e: PageEvent) {
    this.pageEvent = e;
    this.GetCallcenterAgents();
  }

  GetCallcenterAgents() {
    this.store.dispatch(new GetCallcenterAgents({
      db_request: {
        limit: this.pageEvent.pageSize,
        offset: this.pageEvent.pageIndex * this.pageEvent.pageSize,
        filters: this.filters,
        order: this.sortObject
      }
    }));
  }

  AddCallcenterAgent() {
    this.store.dispatch(new AddCallcenterAgent({name: this.newAgentName}));
  }

  UpdateCallcenterAgent(id: number, key: string, value: any) {
    if (id === 0 || key === '') {
      return;
    }

    const param = <Iitem>{
      id: id,
      name: key,
      value: value,
    };
    this.store.dispatch(new UpdateCallcenterAgent({param: param}));
  }

  openBottomSheetQueue(id: number, newName: string, oldName: string, action: string): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete queue "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename queue "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelCallcenterQueue({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameCallcenterQueue({id: id, name: newName}));
      }
    });
  }

  DelCallcenterAgent(id: number, name: string): void {
    const config = {
      data:
        {
          action: 'delete',
          case1Text: 'Are you sure you want to delete agent "' + name + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      this.store.dispatch(new DelCallcenterAgent({id: id}));
    });
  }

  removeFilter(filter: IfilterField): void {
    const index = this.filters.indexOf(filter);

    if (index >= 0) {
      this.filters.splice(index, 1);
    }
  }

  addSorter() {
    if (this.sortColumns) {
      const index = this.sortObject.fields.indexOf(this.sortColumns);

      if (index === -1) {
        this.sortObject.fields.push(this.sortColumns);
      }
    }
  }

  clearSorting() {
    this.sortObject.fields = [];
  }

  handleTiersPageEvent(e: PageEvent) {
    this.tiersPageEvent = e;
    this.GetCallcenterTiers();
  }

  GetCallcenterTiers() {
    this.store.dispatch(new GetCallcenterTiers({
      db_request: {
        limit: this.tiersPageEvent.pageSize,
        offset: this.tiersPageEvent.pageIndex * this.tiersPageEvent.pageSize,
        filters: this.tiersFilters,
        order: this.tiersSortObject
      }
    }));
  }

  AddCallcenterTier() {
    if (this.queueId === null) {
      return; // Cannot add tier without a queue ID
    }
    this.store.dispatch(new AddCallcenterTier({id: this.queueId, name: this.agentName}));
  }

  UpdateCallcenterTier(id: number, key: string, value: any) {
    if (id === 0 || key === '') {
      return;
    }

    const param = <Iitem>{
      id: id,
      name: key,
      value: value,
    };
    this.store.dispatch(new UpdateCallcenterTier({param: param}));
  }

  DelCallcenterTier(id: number, name: string): void {
    const config = {
      data:
        {
          action: 'delete',
          case1Text: 'Are you sure you want to delete tier with id "' + name + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      this.store.dispatch(new DelCallcenterTier({id: id}));
    });
  }

  tiersRemoveFilter(filter: IfilterField): void {
    const index = this.tiersFilters.indexOf(filter);

    if (index >= 0) {
      this.tiersFilters.splice(index, 1);
    }
  }

  tiersAddSorter() {
    if (this.tiersSortColumns) {
      const index = this.tiersSortObject.fields.indexOf(this.tiersSortColumns);

      if (index === -1) {
        this.tiersSortObject.fields.push(this.tiersSortColumns);
      }
    }
  }

  tiersClearSorting() {
    this.tiersSortObject.fields = [];
  }

  queueCommand(id: number, subId: number, command: string) {
    this.store.dispatch(new SendCallcenterCommand({name: command, id: id}));
  }

  membersRemoveFilter(filter: IfilterField): void {
    const index = this.membersFilters.indexOf(filter);

    if (index >= 0) {
      this.membersFilters.splice(index, 1);
    }
  }

  membersAddSorter() {
    if (this.membersSortColumns) {
      const index = this.membersSortObject.fields.indexOf(this.membersSortColumns);

      if (index === -1) {
        this.membersSortObject.fields.push(this.membersSortColumns);
      }
    }
  }

  membersClearSorting() {
    this.membersSortObject.fields = [];
  }

  handleMembersPageEvent(e: PageEvent) {
    this.membersPageEvent = e;
    this.GetCallcenterMembers();
  }

  GetCallcenterMembers() {
    this.store.dispatch(new GetCallcenterMembers({
      db_request: {
        limit: this.membersPageEvent.pageSize,
        offset: this.membersPageEvent.pageIndex * this.membersPageEvent.pageSize,
        filters: this.membersFilters,
        order: this.membersSortObject
      }
    }));
  }

  DelCallcenterMember(uuid: string): void {
    const config = {
      data:
        {
          action: 'delete',
          case1Text: 'Are you sure you want to delete member with uuid "' + uuid + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      this.store.dispatch(new DelCallcenterMember({uuid: uuid}));
    });
  }

  addFilter() {
    this.toEditAgentFilter = null;
    if (this.filter.field_value) {
      this.filter.field_value.trim();
    }
    this.filters.push(<IfilterField>{...this.filter});
    this.filter = {} as IfilterField;
  }

  editFilter(toEdit: number) {
    this.toEditAgentFilter = toEdit;
    this.filter.field = this.filters[toEdit].field;
    this.filter.operand = this.filters[toEdit].operand;
    this.filter.field_value = this.filters[toEdit].field_value;
  }

  saveFilter() {
    this.filters[this.toEditAgentFilter!].field = this.filter.field;
    this.filters[this.toEditAgentFilter!].operand = this.filter.operand ;
    this.filters[this.toEditAgentFilter!].field_value = this.filter.field_value;
    this.toEditAgentFilter = null;
    this.filter = {} as IfilterField;
  }

  tiersAddFilter() {
    this.toEditTierFilter = null;
    if (this.tiersFilter.field_value) {
      this.tiersFilter.field_value.trim();
    }
    this.tiersFilters.push(<IfilterField>{...this.tiersFilter});
    this.tiersFilter = {} as IfilterField;
  }

  editTierFilter(toEdit: number) {
    this.toEditTierFilter = toEdit;
    this.tiersFilter.field = this.tiersFilters[toEdit].field;
    this.tiersFilter.operand = this.tiersFilters[toEdit].operand;
    this.tiersFilter.field_value = this.tiersFilters[toEdit].field_value;
  }

  saveTierFilter() {
    this.tiersFilters[this.toEditTierFilter!].field = this.tiersFilter.field;
    this.tiersFilters[this.toEditTierFilter!].operand = this.tiersFilter.operand ;
    this.tiersFilters[this.toEditTierFilter!].field_value = this.tiersFilter.field_value;
    this.toEditTierFilter = null;
    this.tiersFilter = {} as IfilterField;
  }

  membersAddFilter() {
    this.toEditMemberFilter = null;
    if (this.membersFilter.field_value) {
      this.membersFilter.field_value.trim();
    }
    this.membersFilters.push(<IfilterField>{...this.membersFilter});
    this.membersFilter = {} as IfilterField;
  }

  editMemberFilter(toEdit: number) {
    this.toEditMemberFilter = toEdit;
    this.membersFilter.field = this.membersFilters[toEdit].field;
    this.membersFilter.operand = this.membersFilters[toEdit].operand;
    this.membersFilter.field_value = this.membersFilters[toEdit].field_value;
  }

  saveMemberFilter() {
    this.membersFilters[this.toEditMemberFilter!].field = this.membersFilter.field;
    this.membersFilters[this.toEditMemberFilter!].operand = this.membersFilter.operand ;
    this.membersFilters[this.toEditMemberFilter!].field_value = this.membersFilter.field_value;
    this.toEditMemberFilter = null;
    this.membersFilter = {} as IfilterField;
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}
