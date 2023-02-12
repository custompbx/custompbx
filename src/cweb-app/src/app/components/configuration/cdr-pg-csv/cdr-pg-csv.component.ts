import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {AbstractControl} from '@angular/forms';
import {IcdrPgCsv, Ifield, Iitem} from '../../../store/config/config.state.struct';
import {MatSnackBar} from '@angular/material/snack-bar';
import {
  AddCdrPgCsvField,
  AddCdrPgCsvParam, DeleteCdrPgCsvField, DeleteCdrPgCsvParameter,
  StoreDropNewCdrPgCsvField,
  StoreDropNewCdrPgCsvParam,
  StoreNewCdrPgCsvField,
  StoreNewCdrPgCsvParam, SwitchCdrPgCsvField, SwitchCdrPgCsvParameter,
  UpdateCdrPgCsvField, UpdateCdrPgCsvParameter,
} from '../../../store/config/cdr_pg_csv/config.actions.cdr-pg-csv';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-cdr-pg-csv',
  templateUrl: './cdr-pg-csv.component.html',
  styleUrls: ['./cdr-pg-csv.component.css']
})
export class CdrPgCsvComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: IcdrPgCsv;
  private newItemName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public globalSettingsDispatchers: object;
  public schemaDispatchers: object;
  public schemaFieldMask: object;

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
      this.list = configs.cdr_pg_csv;
      this.lastErrorMessage = configs.cdr_pg_csv && configs.cdr_pg_csv.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newItemName = '';
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
    this.schemaDispatchers = {
      addNewItemField: this.addNewField.bind(this),
      switchItem: this.switchField.bind(this),
      addItem: this.newField.bind(this),
      dropNewItem: this.dropNewField.bind(this),
      deleteItem: this.deleteField.bind(this),
      updateItem: this.updateField.bind(this),
      pasteItems: null,
    };
    this.schemaFieldMask = {name: {name: 'var'}, value: {name: 'column'}};
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  addNewParam() {
    this.store.dispatch(new StoreNewCdrPgCsvParam(null));
  }

  dropNewParam(index: number) {
    this.store.dispatch(new StoreDropNewCdrPgCsvParam({index: index}));
  }

  addNewField() {
    this.store.dispatch(new StoreNewCdrPgCsvField(null));
  }

  dropNewField(index: number) {
    this.store.dispatch(new StoreDropNewCdrPgCsvField({index: index}));
  }

  newParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddCdrPgCsvParam({index: index, param: param}));
  }

  newField(index: number, variable: string, column: string) {
    const field = <Ifield>{};
    field.enabled = true;
    field.var = variable;
    field.column = column;

    this.store.dispatch(new AddCdrPgCsvField({index: index, field: field}));
  }

  updateParam(param: Iitem) {
    this.store.dispatch(new UpdateCdrPgCsvParameter({param: param}));
  }

  updateField(field: Ifield) {
    this.store.dispatch(new UpdateCdrPgCsvField({field: field}));
  }

  switchParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchCdrPgCsvParameter({param: newParam}));
  }

  switchField(field: Ifield) {
    const newField = <Ifield>{...field};
    newField.enabled = !newField.enabled;
    this.store.dispatch(new SwitchCdrPgCsvField({field: newField}));
  }

  deleteParam(param: Iitem) {
    this.store.dispatch(new DeleteCdrPgCsvParameter({param: param}));
  }

  deleteField(field: Ifield) {
    this.store.dispatch(new DeleteCdrPgCsvField({field: field}));
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

  clearDetails(id) {
    //  this.store.dispatch(new ClearDetails(id));
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }
}
