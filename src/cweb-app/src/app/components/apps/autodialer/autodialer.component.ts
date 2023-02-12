import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectApps, selectDirectoryState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {CdkDragDrop} from '@angular/cdk/drag-drop';
import {Iautodialer} from '../../../store/apps/apps.state.struct';
import {
  AddAutoDialerCompany,
  AddAutoDialerReducer,
  AddAutoDialerReducerMember,
  AddAutoDialerTeam,
  DelAutoDialerCompany,
  DelAutoDialerReducerMember,
  DelAutoDialerTeamMember,
  GetAutoDialerReducerMembers,
  GetAutoDialerReducers,
  StoreDropNewAutoDialerReducerMembers,
  StoreNewAutoDialerReducerMembers,
  StoreNewAutoDialerTeamMembers,
  UpdateAutoDialerCompany,
  UpdateAutoDialerReducerMember,
  GetAutoDialerTeamMembers,
  UpdateAutoDialerTeamMember,
  StoreDropNewAutoDialerTeamMembers,
  AddAutoDialerTeamMember,
  GetAutoDialerTeams,
  AddAutoDialerTeamMembers,
  AddAutoDialerList,
  GetAutoDialerLists,
  GetAutoDialerListMembers,
  UpdateAutoDialerListMember,
  DelAutoDialerListMember,
  StoreSetChangedAutodialerListMemberField, AddAutoDialerListMembers
} from '../../../store/apps/autodialer/autodialer.actions';
import {GetDirectoryUsers} from '../../../store/directory/directory.actions';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatOption} from '@angular/material/core';
import {CsvService} from '../../../services/csv-handler';
import {IfilterField, IsortField} from '../../cdr/cdr.component';
import {PageEvent} from '@angular/material/paginator';
import {Iitem} from '../../../store/config/config.state.struct';

@Component({
  selector: 'app-autodialer',
  templateUrl: './autodialer.component.html',
  styleUrls: ['./autodialer.component.css']
})
export class AutodialerComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Iautodialer;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public newCompanyName: string;
  public newReducerName: string;
  public newTeamName: string;
  public newListName: string;
  public domains: Observable<any>;
  public domains$: Subscription;
  public domainsList: any;
  public domainId: number;
  public reducerId: number;
  public users$: Subscription;
  public users: Observable<any>;
  public userList: any;
  public test: any;
  public csvHeader: string;
   // need to add to store later
  public teamUserArr: { [id: number]: Array<number> };
  private paginationScale = [25, 50, 100, 250];
  public operands: Array<string> = ['=', '>', '<', 'CONTAINS'];
  public pageEvent: PageEvent = <PageEvent>{
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  };
  public filter: IfilterField = <IfilterField>{};
  public toEditFilter: number = <number>null;
  private toEdit = {};
  private showDel = {};

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
    private csv: CsvService,
  ) {
    this.selectedIndex = 0;
    this.configs = this.store.pipe(select(selectApps));
    this.domains = this.store.pipe(select(selectDirectoryState));
    this.csvHeader = 'to_number, from_number, retries, name, custom_vars\n';
    // this.csvData = this.csvHeader;
  }

  ngOnInit() {
    this.configs$ = this.configs.subscribe((apps) => {
      console.log(apps);
      this.loadCounter = apps.loadCounter;
      this.list = apps.autodialer;
      this.lastErrorMessage = apps.autodialer && apps.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.onlyValues(this.list.AutoDialerTeams).forEach(t => {
          this.teamUserArr = {...this.teamUserArr, [t.id]: []};
        });
        this.onlyValues(this.list.AutoDialerTeamMembers).forEach(m => {
          if (!this.teamUserArr[m.parent?.id]) {
            this.teamUserArr = {...this.teamUserArr, [m.parent?.id]: []};
          }
          this.teamUserArr[m.parent?.id].push(m.user?.id);
        });
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
      this.toEdit = {};
    });
    this.domains$ = this.domains.subscribe((domains) => {
      this.loadCounter = domains.loadCounter;
      this.domainsList = domains.domains;
      this.userList = domains.users;
      this.lastErrorMessage = domains.errorMessage;
      if (!this.lastErrorMessage) {
        this.selectedIndex = 0;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
      this.route.snapshot.data.reconnectUpdater.unsubscribe();
    }
  }

  mainTabChanged(event) {
    if (event === 2) {

    }
  }

  public onStepChange(event: any, company): void {
    switch (event.selectedIndex) {
      case 1:
        console.log(event.selectedIndex, company?.domain?.id);
        this.GetAutoDialerTeams(company?.domain?.id);
        this.GetAutoDialerTeamMembers(company?.team?.id);
        this.store.dispatch(new GetDirectoryUsers(null));
        break;
      case 2:
        this.GetAutoDialerLists(company?.domain?.id);
    }
    console.log(event.selectedIndex);
  }

  getAutoDialerReducers(id) {
    this.store.dispatch(new GetAutoDialerReducers({data: {domain: {id: id}}}));
  }

  addAutoDialerCompany() {
    if (!this.newCompanyName || !this.domainId) {
      return;
    }
    this.store.dispatch(new AddAutoDialerCompany({data: {name: this.newCompanyName, domain: {id: this.domainId}}}));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
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

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  ObjLength(item: object): number {
    return Object.keys(item).length;
  }

  onlySortedValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    const arr = Object.values(obj).sort(
      function (a, b) {
        if (a.position > b.position) {
          return 1;
        }
        if (a.position < b.position) {
          return -1;
        }
        return 0;
      }
    );
    return arr;
  }

  trackByIdFn(index, item) {
    return item.id;
  }

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }

    this.store.dispatch(
      new UpdateAutoDialerReducerMember({
        fields: ['position'],
        data: {
          id: parent[event.previousIndex].id,
          position: parent[event.currentIndex].position,
          parent: parent[event.previousIndex].parent
        },
      }),
    );
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  openBottomSheetCompany(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete company "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename company "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelAutoDialerCompany({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateAutoDialerCompany({data: {id: id, name: newName}, fields: ['name']}));
      }
    });
  }

  addAutoDialerCompanyReducer(id) {
    if (!this.newReducerName) {
      return;
    }
    this.store.dispatch(new AddAutoDialerReducer({data: {name: this.newReducerName, domain: {id: id}}}));
  }

  changeReducer(cid, id: number) {
    if (!cid || !id) {
      return;
    }
    let reducer = {id: Number(id)};
    if (Number(id) === 0) {
      reducer = null;
    }

    this.store.dispatch(new UpdateAutoDialerCompany({data: {id: cid, reducer: reducer}, fields: ['reducer']}));
  }

  GetAutoDialerReducerMembers(id) {
    if (!id) {
      return;
    }
    this.store.dispatch(new GetAutoDialerReducerMembers({id: id}));
  }

  UpdateAutoDialerReducerMember(param) {
    this.store.dispatch(
      new UpdateAutoDialerReducerMember({data: param, fields: []})
    );
  }

  SwitchAutoDialerReducerMember(param) {
    this.store.dispatch(
      new UpdateAutoDialerReducerMember({data: {id: param.id, enabled: !param.enabled}, fields: ['enabled']})
    );
  }

  deleteAutoDialerCompanyReducerMember(id) {
    if (!id.id) {
      return;
    }
    this.store.dispatch(new DelAutoDialerReducerMember({id: id.id, parent: id.parent.id}));
  }

  addNewAutoDialerReducerMembers(id) {
    this.store.dispatch(new StoreNewAutoDialerReducerMembers({id}));
  }

  dropNewAutoDialerReducerMembers(id: number, index: number) {
    this.store.dispatch(new StoreDropNewAutoDialerReducerMembers({id, index: index}));
  }

  addAutoDialerReducerMember(id: number, index: number, application: string, data: string) {
    this.store.dispatch(new AddAutoDialerReducerMember({id, index, data: {enabled: true, application, data, parent: {id}}}));
  }

  SwitchAutoDialerCompanyPredictive(company) {
    this.store.dispatch(
      new UpdateAutoDialerCompany({data: {id: company.id, predictive: !company.predictive}, fields: ['predictive']})
    );
  }

  addAutoDialerCompanyTeam(id) {
    if (!this.newTeamName) {
      return;
    }
    this.store.dispatch(new AddAutoDialerTeam({data: {name: this.newTeamName, domain: {id: id}}}));
  }

  changeTeam(cid, id: number) {
    if (!cid || !id) {
      return;
    }
    let team = {id: Number(id)};
    if (Number(id) === 0) {
      team = null;
    }

    this.store.dispatch(new UpdateAutoDialerCompany({data: {id: cid, team: team}, fields: ['team']}));
  }

  GetAutoDialerTeamMembers(id) {
    if (!id) {
      return;
    }
    this.store.dispatch(new GetAutoDialerTeamMembers({id: id}));
  }

  UpdateAutoDialerTeamMembers(id, arr) {
    this.store.dispatch(
      new AddAutoDialerTeamMembers({id: id, ids: arr})
    );
  }

  UpdateAutoDialerTeamMember(param) {
    this.store.dispatch(
      new UpdateAutoDialerTeamMember({data: param, fields: []})
    );
  }

  SwitchAutoDialerTeamMember(param) {
    this.store.dispatch(
      new UpdateAutoDialerTeamMember({data: {id: param.id, enabled: !param.enabled}, fields: ['enabled']})
    );
  }

  deleteAutoDialerCompanyTeamMember(id) {
    if (!id.id) {
      return;
    }
    this.store.dispatch(new DelAutoDialerTeamMember({id: id.id, parent: id.parent.id}));
  }

  addNewAutoDialerTeamMembers(id) {
    this.store.dispatch(new StoreNewAutoDialerTeamMembers({id}));
  }

  dropNewAutoDialerTeamMembers(id: number, index: number) {
    this.store.dispatch(new StoreDropNewAutoDialerTeamMembers({id, index: index}));
  }

  addAutoDialerTeamMember(id: number, index: number, application: string) {
    this.store.dispatch(new AddAutoDialerTeamMember({id, index, data: {enabled: true, application, parent: {id}}}));
  }

  addAutoDialerTeamMembers(id: number, index: number, application: string) {
    this.store.dispatch(new AddAutoDialerTeamMembers({id, index, data: {enabled: true, application, parent: {id}}}));
  }

  GetAutoDialerTeams(id: number) {
    this.store.dispatch(new GetAutoDialerTeams({id}));
  }

  onlyValuesByParent(obj: object, parentId: number): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj).filter(u => u.parent.id === parentId);
  }

  toggleAllSelection(matSelect) {
    const isSelected: boolean = matSelect.options
      .filter((item: MatOption) => item.value === 0)
      .map((item: MatOption) => item.selected)
      // Get the first element (there should only be 1 option with the value 0 in the select)
      [0];

    if (isSelected) {
      matSelect.options.forEach((item: MatOption) => item.select());
    } else {
      matSelect.options.forEach((item: MatOption) => item.deselect());
    }
  }

  addAutoDialerCompanyList(id) {
    if (!this.newListName) {
      return;
    }
    this.store.dispatch(new AddAutoDialerList({data: {name: this.newListName, domain: {id: id}}}));
  }

  changeList(id, cid: number) {
    if (!cid || !id) {
      return;
    }
    let list = {id: Number(cid)};
    if (Number(cid) === 0) {
      list = null;
    }

    this.store.dispatch(new UpdateAutoDialerCompany({data: {id: id, list: list}, fields: ['list']}));
  }

  GetAutoDialerLists(id: number) {
    this.store.dispatch(new GetAutoDialerLists({id}));
  }

  UpdateAutoDialerList(id, csvData) {
    const arrNames = ['to_number', 'from_number', 'retries', 'name', 'custom_vars'];
    const propertyNames = csvData.slice(0, csvData.indexOf('\n')).split(/,|;/);
    if (!propertyNames.every((v, i) => v.trim() === arrNames[i])) {
      csvData = this.csvHeader + csvData;
    }
    const res = this.csv.importDataFromCSV(csvData);
    console.log(res);

    if (res.length === 0) {
      this._snackBar.open('Nothing to update', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
      return;
    }
    if (!res[0].to_number) {
      this._snackBar.open('Nothing to update', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
      return;
    }

    this.store.dispatch(new AddAutoDialerListMembers({id, data: res}));
  }

  GetAutoDialerListMembers(id: number, tableMeta) {
    this.store.dispatch(new GetAutoDialerListMembers({
      id,
      db_request: {
        limit: this.pageEvent.pageSize,
        offset: this.pageEvent.pageIndex, filters: tableMeta.filters, order: tableMeta.sortObject
      }
    }));
  }

  addFilter(filters: Array<any>) {
    this.toEditFilter = null;
    if (this.filter.field_value) {
      this.filter.field_value.trim();
    }
    filters.push(<IfilterField>this.filter);
    this.filter = <IfilterField>{};
  }

  editFilter(toEdit: number, filters) {
    this.toEditFilter = toEdit;
    this.filter.field = filters[toEdit].field;
    this.filter.operand = filters[toEdit].operand;
    this.filter.field_value = filters[toEdit].field_value;
  }

  saveFilter(filters) {
    filters[this.toEditFilter].field = this.filter.field;
    filters[this.toEditFilter].operand = this.filter.operand;
    filters[this.toEditFilter].field_value = this.filter.field_value;
    this.toEditFilter = null;
    this.filter.field = null;
    this.filter.operand = null;
    this.filter.field_value = null;
  }

  removeFilter(filter: IfilterField, filters): void {
    const index = filters.indexOf(filter);

    if (index >= 0) {
      filters.splice(index, 1);
    }
  }

  addSorter(tableMeta) {
    const index = tableMeta.sortObject.fields.indexOf(tableMeta.sortColumns);

    if (index === -1) {
      tableMeta.sortObject.fields.push(tableMeta.sortColumns);
    }
  }

  clearSorting(tableMeta) {
    tableMeta.sortObject.fields = [];
  }

  toInput(id: number, columnName: any) {
    this.toEdit = {[id]: columnName};
  }

  leaveDelIco(id) {
    setTimeout(function () {
      this.showDel[id] = false;
    }.bind(this), 300);
  }

  markChanged(id: number, columnName: any) {
    this.store.dispatch(new StoreSetChangedAutodialerListMemberField({rowId: id, fieldName: columnName}));
  }

  DelAutoDialerListMember(id, name): void {
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
      this.store.dispatch(new DelAutoDialerListMember({id: id}));
    });
  }

  UpdateAutoDialerListMember(id, key, value) {
    if (id === 0 || key === '') {
      return;
    }

    const param = <Iitem>{
      id: id,
      name: key,
      value: value,
    };
    this.store.dispatch(new UpdateAutoDialerListMember({param}));
  }

}

