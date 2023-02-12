import {Component, OnDestroy, OnInit} from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {Observable, Subscription} from 'rxjs';
import {Icallcenter, Iitem} from '../../../store/config/config.state.struct';
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
import {AbstractControl} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {IfilterField, IsortField} from '../../cdr/cdr.component';

@Component({
  selector: 'app-callcenter',
  templateUrl: './callcenter.component.html',
  styleUrls: ['./callcenter.component.css']
})
export class CallcenterComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Icallcenter;
  private newQueueName: string;
  private newAgentName: string;
  private agentName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private queueId: number;
  private panelCloser = [];
  private toEdit = {};
  private toEditTier = {};
  private showDel = {};
  public loadCounter: number;
  private toCopyqueue: number;

  public operands: Array<string> = ['=', '>', '<', 'CONTAINS'];
  public sortObject: IsortField = <IsortField>{fields: [], desc: false};
  public filter: IfilterField = <IfilterField>{};
  public filters: Array<IfilterField> = [];
  private paginationScale = [25, 50, 100, 250];
  public pageEvent: PageEvent = <PageEvent>{
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  };
  public columns: Array<string>;
  public sortColumns: string;
  public tiersSortObject: IsortField = <IsortField>{fields: [], desc: false};
  public tiersFilter: IfilterField = <IfilterField>{};
  public tiersFilters: Array<IfilterField> = [];
  public tiersColumns: Array<string>;
  public tiersSortColumns: string;
  public tiersPageEvent: PageEvent = <PageEvent>{
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  };

  public membersSortObject: IsortField = <IsortField>{fields: [], desc: false};
  public membersFilter: IfilterField = <IfilterField>{};
  public membersFilters: Array<IfilterField> = [];
  public membersColumns: Array<string>;
  public membersSortColumns: string;
  public membersPageEvent: PageEvent = <PageEvent>{
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  };

  public queueParamDispatchers: object;
  public globalSettingsDispatchers: object;
  public toEditAgentFilter: number = <number>null;
  public toEditTierFilter: number = <number>null;
  public toEditMemberFilter: number = <number>null;

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.configs = this.store.pipe(select(selectConfigurationState));
  }

  ngOnInit() {
    this.configs$ = this.configs.subscribe((configs) => {
      console.log(configs);
      this.loadCounter = configs.loadCounter;
      this.list = configs.callcenter;
      this.toEdit = {};
      this.toEditTier = {};
      if (this.list && this.list.agents && this.list.agents.table && this.list.agents.table.length > 0) {
        this.columns = [];
        console.log(this.list.agents.table[0]);
        Object.keys(this.list.agents.table[0]).forEach( key => {
          this.columns.push(key);
        });
      }
      if (this.list && this.list.tiers && this.list.tiers.table && this.list.tiers.table.length > 0) {
        this.tiersColumns = [];
        console.log(this.list.tiers.table[0]);
        Object.keys(this.list.tiers.table[0]).forEach( key => {
          this.tiersColumns.push(key);
        });
      }
      if (this.list && this.list.members && this.list.members.table && this.list.members.table.length > 0) {
        this.membersColumns = [];
        console.log(this.list.members.table[0]);
        Object.keys(this.list.members.table[0]).forEach( key => {
          this.membersColumns.push(key);
        });
      }

      this.lastErrorMessage = configs.callcenter && configs.callcenter.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newQueueName = '';
        this.newAgentName = '';
        this.agentName = '';
        // this.queueId = null;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
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

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  checkDirty(condition: AbstractControl): boolean {
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

  isReadyToSendOne(nameObject: AbstractControl): boolean {
    return nameObject && nameObject.dirty && nameObject.valid;
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

  leaveDelIco(id) {
    setTimeout(function () {
      this.showDel[id] = false;
    }.bind(this), 300);
  }

  copyQueue(key) {
    if (!this.list.queues[key]) {
      this.toCopyqueue = 0;
      return;
    }
    this.toCopyqueue = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  tabChanged(event) {
    this.panelCloser = [];
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

  mainTabChanged(event) {
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
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddCallcenterSettings({index: index, param: param}));
  }

  StoreNewCallcenterSettings() {
    this.store.dispatch(new StoreNewCallcenterSettings(null));
  }

  StoreDropNewCallcenterSettings(index: number) {
    this.store.dispatch(new StoreDropNewCallcenterSettings({index: index}));
  }


  GetCallcenterQueuesParams(id) {
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
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

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

  GetCallcenterAgents() {
    this.store.dispatch(new GetCallcenterAgents({
      db_request: {limit: this.pageEvent.pageSize,
        offset: this.pageEvent.pageIndex, filters: this.filters, order: this.sortObject}
    }));
  }

  AddCallcenterAgent() {
    this.store.dispatch(new AddCallcenterAgent({name: this.newAgentName}));
  }

  UpdateCallcenterAgent(id, key, value) {
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

  openBottomSheetQueue(id, newName, oldName, action): void {
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

  DelCallcenterAgent(id, name): void {
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
    const index = this.sortObject.fields.indexOf(this.sortColumns);

    if (index === -1) {
      this.sortObject.fields.push(this.sortColumns);
    }
  }

  clearSorting() {
    this.sortObject.fields = [];
  }

  GetCallcenterTiers() {
    this.store.dispatch(new GetCallcenterTiers({
      db_request: {limit: this.tiersPageEvent.pageSize,
        offset: this.tiersPageEvent.pageIndex, filters: this.tiersFilters, order: this.tiersSortObject}
    }));
  }

  AddCallcenterTier() {
    this.store.dispatch(new AddCallcenterTier({id: this.queueId, name: this.agentName}));
  }

  UpdateCallcenterTier(id, key, value) {
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

  DelCallcenterTier(id, name): void {
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
    const index = this.tiersSortObject.fields.indexOf(this.sortColumns);

    if (index === -1) {
      this.tiersSortObject.fields.push(this.sortColumns);
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
    const index = this.membersSortObject.fields.indexOf(this.sortColumns);

    if (index === -1) {
      this.membersSortObject.fields.push(this.sortColumns);
    }
  }

  membersClearSorting() {
    this.membersSortObject.fields = [];
  }

  GetCallcenterMembers() {
    this.store.dispatch(new GetCallcenterMembers({
      db_request: {limit: this.membersPageEvent.pageSize,
        offset: this.membersPageEvent.pageIndex, filters: this.membersFilters, order: this.membersSortObject}
    }));
  }

  DelCallcenterMember(uuid): void {
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
    this.filters.push(<IfilterField>this.filter);
    this.filter = <IfilterField>{};
  }

  editFilter(toEdit: number) {
    this.toEditAgentFilter = toEdit;
    this.filter.field = this.filters[toEdit].field;
    this.filter.operand = this.filters[toEdit].operand;
    this.filter.field_value = this.filters[toEdit].field_value;
  }

  saveFilter() {
    this.filters[this.toEditAgentFilter].field = this.filter.field;
    this.filters[this.toEditAgentFilter].operand = this.filter.operand ;
    this.filters[this.toEditAgentFilter].field_value = this.filter.field_value;
    this.toEditAgentFilter = null;
    this.filter.field = null;
    this.filter.operand = null;
    this.filter.field_value = null;
  }

  tiersAddFilter() {
    this.toEditTierFilter = null;
    if (this.tiersFilter.field_value) {
      this.tiersFilter.field_value.trim();
    }
    this.tiersFilters.push(<IfilterField>this.tiersFilter);
    this.tiersFilter = <IfilterField>{};
  }

  editTierFilter(toEdit: number) {
    this.toEditTierFilter = toEdit;
    this.tiersFilter.field = this.tiersFilters[toEdit].field;
    this.tiersFilter.operand = this.tiersFilters[toEdit].operand;
    this.tiersFilter.field_value = this.tiersFilters[toEdit].field_value;
  }

  saveTierFilter() {
    this.tiersFilters[this.toEditTierFilter].field = this.tiersFilter.field;
    this.tiersFilters[this.toEditTierFilter].operand = this.tiersFilter.operand ;
    this.tiersFilters[this.toEditTierFilter].field_value = this.tiersFilter.field_value;
    this.toEditTierFilter = null;
    this.tiersFilter.field = null;
    this.tiersFilter.operand = null;
    this.tiersFilter.field_value = null;
  }

  membersAddFilter() {
    this.toEditMemberFilter = null;
    if (this.membersFilter.field_value) {
      this.membersFilter.field_value.trim();
    }
    this.membersFilters.push(<IfilterField>this.membersFilter);
    this.membersFilter = <IfilterField>{};
  }

  editMemberFilter(toEdit: number) {
    this.toEditMemberFilter = toEdit;
    this.membersFilter.field = this.membersFilters[toEdit].field;
    this.membersFilter.operand = this.membersFilters[toEdit].operand;
    this.membersFilter.field_value = this.membersFilters[toEdit].field_value;
  }

  saveMemberFilter() {
    this.membersFilters[this.toEditMemberFilter].field = this.membersFilter.field;
    this.membersFilters[this.toEditMemberFilter].operand = this.membersFilter.operand ;
    this.membersFilters[this.toEditMemberFilter].field_value = this.membersFilter.field_value;
    this.toEditMemberFilter = null;
    this.membersFilter.field = null;
    this.membersFilter.operand = null;
    this.membersFilter.field_value = null;
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}
