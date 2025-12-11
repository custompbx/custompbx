import {Component, OnDestroy, OnInit} from '@angular/core';

import {MaterialModule} from "../../../../material-module";
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {AbstractControl, FormsModule} from '@angular/forms';
import {IodbcCdr, Iitem, IodbcField, Itable} from '../../../store/config/config.state.struct';
import {MatSnackBar} from '@angular/material/snack-bar';
import {
  AddOdbcCdrField,
  AddOdbcCdrParameter, AddOdbcCdrTable, DeleteOdbcCdrField, DeleteOdbcCdrParameter, DeleteOdbcCdrTable,
  StoreDropNewOdbcCdrField,
  StoreDropNewOdbcCdrParameter,
  StoreNewOdbcCdrField,
  StoreNewOdbcCdrParameter, SwitchOdbcCdrField,
  SwitchOdbcCdrParameter,
  UpdateOdbcCdrField,
  UpdateOdbcCdrParameter, UpdateOdbcCdrTable, GetOdbcCdrField
} from '../../../store/config/odbc_cdr/config.actions.odbc-cdr';
import {ActivatedRoute} from '@angular/router';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";

@Component({
standalone: true,
imports:  [MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent],
    selector: 'app-odbc-cdr',
    templateUrl: './odbc-cdr.component.html',
    styleUrls: ['./odbc-cdr.component.css']
})
export class OdbcCdrComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: IodbcCdr;
  private newTableName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private panelCloser = [];
  public loadCounter: number;
  public tableId: number;
  public globalSettingsDispatchers: object;
  public fieldsDispatchers: object;
  public tableFieldMask: object;

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
      this.loadCounter = configs.loadCounter;
      this.list = configs.odbc_cdr;
      this.lastErrorMessage = configs.odbc_cdr && configs.odbc_cdr.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newTableName = '';
        this.selectedIndex = 0;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewParam.bind(this),
      switchItem: this.switchParam.bind(this),
      addItem: this.newParam.bind(this),
      dropNewItem: this.dropNewParam.bind(this),
      deleteItem: this.deleteParam.bind(this),
      updateItem: this.updateParam.bind(this),
      pasteItems: null,
    };
    this.fieldsDispatchers = {
      addNewItemField: this.addNewField.bind(this),
      switchItem: this.switchField.bind(this),
      addItem: this.newField.bind(this),
      dropNewItem: this.dropNewField.bind(this),
      deleteItem: this.deleteField.bind(this),
      updateItem: this.updateField.bind(this),
      pasteItems: null,
    };
    this.tableFieldMask = {name: {name: 'name'}, value: {name: 'chan_var_name'}};
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  addNewParam() {
    this.store.dispatch(new StoreNewOdbcCdrParameter(null));
  }

  dropNewParam(index: number) {
    this.store.dispatch(new StoreDropNewOdbcCdrParameter({index: index}));
  }

  addNewField(tableId) {
    this.store.dispatch(new StoreNewOdbcCdrField({id: tableId}));
  }

  dropNewField(tableId, index: number) {
    this.store.dispatch(new StoreDropNewOdbcCdrField({index: index, id: tableId}));
  }

  newParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddOdbcCdrParameter({index: index, param: param}));
  }

  onTableSubmit() {
    const table = <Itable>{};
    table.enabled = true;
    table.name = this.newTableName;

    this.store.dispatch(new AddOdbcCdrTable({table: table}));
  }

  newField(tableId, index: number, variable: string, column: string) {
    const field = <IodbcField>{};
    field.enabled = true;
    field.name = variable;
    field.chan_var_name = column;

    this.store.dispatch(new AddOdbcCdrField({index: index, odbc_cdr_field: field, id: tableId}));
  }

  GetOdbcCdrField(id) {
    this.panelCloser['table' + id] = true;
    this.store.dispatch(new GetOdbcCdrField({id: id}));
  }

  updateParam(param: Iitem) {
    this.store.dispatch(new UpdateOdbcCdrParameter({param: param}));
  }

  updateTable(table: Itable) {
    this.store.dispatch(new UpdateOdbcCdrTable({table: table}));
  }

  updateField(field: IodbcField) {
    this.store.dispatch(new UpdateOdbcCdrField({odbc_cdr_field: field}));
  }

  switchParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchOdbcCdrParameter({param: newParam}));
  }

  switchTable(table: Itable) {
    const newTable = <Itable>{...table};
    newTable.enabled = !newTable.enabled;
    this.store.dispatch(new SwitchOdbcCdrField({table: newTable}));
  }

  switchField(field: IodbcField) {
    const newField = <IodbcField>{...field};
    newField.enabled = !newField.enabled;
    this.store.dispatch(new SwitchOdbcCdrField({odbc_cdr_field: newField}));
  }

  deleteParam(param: Iitem) {
    this.store.dispatch(new DeleteOdbcCdrParameter({param: param}));
  }

  deleteField(field: IodbcField) {
    this.store.dispatch(new DeleteOdbcCdrField({odbc_cdr_field: field}));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
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

  clearDetails(id) {
    //  this.store.dispatch(new ClearDetails(id));
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

  openBottomSheetTable(id, newName, oldName, action): void {
    const config = {
      data:
        {
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete table "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename gateway "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DeleteOdbcCdrTable({table: {id: Number(id)}}));
      } else if (action === 'rename') {
        console.log(this.list.tables[id]);
        const table = <Itable>{};
        table.id = this.list.tables[id].id;
        table.enabled = this.list.tables[id].enabled;
        table.name = newName;
        table.log_leg = this.list.tables[id].log_leg;
        this.updateTable(table);
      }
    });
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}
