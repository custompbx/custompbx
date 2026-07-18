import { Component, OnDestroy, OnInit, inject, computed, effect } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { select, Store } from '@ngrx/store';
import { AppState, selectConfigurationState } from '../../../store/app.states';
import { AbstractControl, FormsModule } from '@angular/forms';
import {IcdrPgCsv, Ifield, Iitem, State} from '../../../store/config/config.state.struct';
import {ToastService} from '../../../services/toast.service';
import {
  AddCdrPgCsvField,
  AddCdrPgCsvParam, DeleteCdrPgCsvField, DeleteCdrPgCsvParameter,
  StoreDropNewCdrPgCsvField,
  StoreDropNewCdrPgCsvParam,
  StoreNewCdrPgCsvField,
  StoreNewCdrPgCsvParam, SwitchCdrPgCsvField, SwitchCdrPgCsvParameter,
  UpdateCdrPgCsvField, UpdateCdrPgCsvParameter,
} from '../../../store/config/cdr_pg_csv/config.actions.cdr-pg-csv';
import { ActivatedRoute } from '@angular/router';
import { InnerHeaderComponent } from "../../inner-header/inner-header.component";
import { ModuleNotExistsBannerComponent } from "../module-not-exists-banner/module-not-exists-banner.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";
import {TranslocoPipe} from '@jsverse/transloco';

@Component({
  standalone: true,
  imports: [FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component, TranslocoPipe],
  selector: 'app-cdr-pg-csv',
  templateUrl: './cdr-pg-csv.component.html',
  styleUrls: ['./cdr-pg-csv.component.css']
})
export class CdrPgCsvComponent implements OnInit, OnDestroy {

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private _snackBar = inject(ToastService);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private configState = toSignal(
    this.store.pipe(select(selectConfigurationState)),
    {
      initialValue: {
        cdr_pg_csv: {} as IcdrPgCsv,
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from Signals ---
  public list = computed(() => this.configState().cdr_pg_csv);
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().cdr_pg_csv?.errorMessage || null);

  // --- Local Component State (Variables) ---
  private newItemName: string = '';
  public selectedIndex: number = 0;
  public globalSettingsDispatchers: object = {};
  public schemaDispatchers: object = {};
  public schemaFieldMask: object = {};

  private errorEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
      this.newItemName = '';
      this.selectedIndex = 0;
    }
  });

  ngOnInit() {
    // Dispatchers setup remains the same
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
    this.schemaFieldMask = {name: {name: 'var'}, value: {name: 'column'}, extraField1: {name: 'quote', size: 'sm'}};
  }

  ngOnDestroy() {
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

  newField(index: number, variable: string, column: string, quote: string) {
    const field = <Ifield>{};
    field.enabled = true;
    field.var = variable;
    field.column = column;
    field.quote = quote;

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
}
