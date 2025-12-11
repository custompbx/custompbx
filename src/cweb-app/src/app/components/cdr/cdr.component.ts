import {Component, inject, signal, computed, effect, OnInit} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {select, Store} from '@ngrx/store';
import {AppState, selectCDRState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {PageEvent} from '@angular/material/paginator';
import {State} from '../../store/cdr/cdr.reducers';
import {GetCDR} from '../../store/cdr/cdr.actions';
import {GetWebSettings, SaveWebSettings} from '../../store/settings/settings.actions';

import {MaterialModule} from "../../../material-module";
import {ResizeInputDirective} from "../../directives/resize-input.directive";
import {FormsModule} from "@angular/forms";
import {InnerHeaderComponent} from "../inner-header/inner-header.component";

export interface IfilterField {
  field: string|null;
  operand: string|null;
  field_value: string|null;
}

export interface IsortField {
  fields: Array<string>;
  desc: boolean;
}

@Component({
  standalone: true,
  imports: [MaterialModule, FormsModule, InnerHeaderComponent, ResizeInputDirective],
  selector: 'app-cdr',
  templateUrl: './cdr.component.html',
  styleUrls: ['./cdr.component.css']
})
export class CdrComponent { // Removed OnDestroy

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private _snackBar = inject(MatSnackBar);

  // --- Reactive State from NgRx using toSignal ---
  private cdrState = toSignal(
    this.store.pipe(select(selectCDRState)),
    {
      initialValue: {
        cdrData: [],
        settings: {},
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.cdrState());
  public loadCounter = computed(() => this.cdrState().loadCounter);

  // --- Local Component State as Signals/Properties ---
  public filters = signal<Array<IfilterField>>([]);
  public pageTotal = 0; // Kept as property updated by effect
  protected paginationScale = [25, 50, 100, 250];

  public pageEvent = signal<PageEvent>({
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  } as PageEvent);

  public operands: Array<string> = ['=', '>', '<', 'LIKE', '>=', '<=', '!=', 'NOT LIKE'];
  public sortObject = signal<IsortField>({fields: [], desc: false});
  public filter = signal<IfilterField>({field: null, operand: null, field_value: null});
  public toEditFilter = signal<number | null>(null);

  public columns = signal<Array<string>>([]);
  public sortColumns: string | null = null; // Can remain a property for simple input binding

  // Settings fields
  public settings = signal<{[key: string]: any}>({});
  public moduleOptions = ['auto', 'cdr_pg_csv', 'odbc_cdr'];
  public fieldModule = 'cdr_module';
  public fieldTable = 'cdr_table';
  public fieldFileServeColumn = 'cdr_file_serve_column';
  public fieldFileServerPath = 'cdr_file_server_path';

  // --- Effect for Side Effects (Data updates and error handling) ---
  private stateUpdateEffect = effect(() => {
    const cdr = this.cdrState();
    const errorMessage = cdr.errorMessage;
    const data = cdr.cdrData;
    const stateSettings = cdr.settings;

    // 1. Handle Snackbar Error
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }

    // 2. Handle data/column setup
    if (data.length > 0) {
      const firstRecord = data[0];

      // Update pageTotal property
      this.pageTotal = firstRecord['total'] || 0;

      // Update columns signal
      const newColumns = Object.keys(firstRecord).filter(key => key !== 'total');
      this.columns.set(newColumns);
    } else {
      this.pageTotal = 0;
      this.columns.set([]);
    }

    // 3. Update Settings signal
    if (stateSettings) {
      this.settings.set({
        [this.fieldModule]: stateSettings[this.fieldModule],
        [this.fieldTable]: stateSettings[this.fieldTable],
        [this.fieldFileServeColumn]: stateSettings[this.fieldFileServeColumn],
        [this.fieldFileServerPath]: stateSettings[this.fieldFileServerPath],
      });
    }
  });

  handlePageEvent(event: PageEvent) {
    this.pageEvent.set(event);
    this.getRecords();
  }

  getRecords() {
    // Read current settings and page event
    const currentSettings = this.settings();
    const currentPageEvent = this.pageEvent();

    if (!currentSettings[this.fieldModule] || !currentSettings[this.fieldFileServeColumn]) {
      this.getSettings();
    }

    this.store.dispatch(new GetCDR({
      db_request: {
        limit: currentPageEvent.pageSize,
        offset: currentPageEvent.pageIndex,
        filters: this.filters(), // Read signal value
        order: this.sortObject() // Read signal value
      }
    }));
  }

  removeFilter(filterToRemove: IfilterField): void {
    this.pageEvent.update(e => ({...e, pageIndex: 0}));
    this.filters.update(currentFilters => {
      // Find and remove the filter
      return currentFilters.filter(f => f !== filterToRemove);
    });
  }

  addSorter() {
    this.pageEvent.update(e => ({...e, pageIndex: 0}));
    const sortCol = this.sortColumns;
    if (!sortCol) return;

    this.sortObject.update(currentSort => {
      if (currentSort.fields.indexOf(sortCol) === -1) {
        return {
          ...currentSort,
          fields: [...currentSort.fields, sortCol]
        };
      }
      return currentSort;
    });
  }

  clearSorting() {
    this.pageEvent.update(e => ({...e, pageIndex: 0}));
    this.sortObject.set({fields: [], desc: false});
  }

  tabChanged(event: number) {
    if (event === 1) {
      this.getSettings();
    }
  }

  getSettings() {
    const webSettings: object = {};
    webSettings[this.fieldModule] = '';
    webSettings[this.fieldTable] = '';
    webSettings[this.fieldFileServeColumn] = '';
    webSettings[this.fieldFileServerPath] = '';
    this.store.dispatch(new GetWebSettings({web_settings: webSettings}));
  }

  saveSettings() {
    const currentSettings = this.settings();
    const webSettings: object = {};
    webSettings[this.fieldModule] = currentSettings[this.fieldModule];
    webSettings[this.fieldTable] = currentSettings[this.fieldTable];
    webSettings[this.fieldFileServeColumn] = currentSettings[this.fieldFileServeColumn];
    webSettings[this.fieldFileServerPath] = currentSettings[this.fieldFileServerPath];
    this.store.dispatch(new SaveWebSettings({web_settings: webSettings}));
  }

  fileTypeByName(str: string): string {
    const res = str.split('.');
    return res[res.length - 1];
  }

  addFilter() {
    this.toEditFilter.set(null);
    this.pageEvent.update(e => ({...e, pageIndex: 0}));

    const currentFilter = this.filter();
    const trimmedFilterValue = currentFilter.field_value ? currentFilter.field_value.trim() : null;

    if (currentFilter.field && currentFilter.operand && trimmedFilterValue) {
      this.filters.update(f => [...f, {...currentFilter, field_value: trimmedFilterValue}]);
    }

    this.filter.set({field: null, operand: null, field_value: null});
  }

  editFilter(toEdit: number) {
    this.toEditFilter.set(toEdit);
    const filterToEdit = this.filters()[toEdit];
    this.filter.set({...filterToEdit});
  }

  saveFilter() {
    const editIndex = this.toEditFilter();
    if (editIndex !== null) {
      const savedFilter = this.filter();

      this.filters.update(currentFilters => {
        const updatedFilters = [...currentFilters];
        updatedFilters[editIndex] = {
          ...savedFilter,
          field_value: savedFilter.field_value ? savedFilter.field_value.trim() : null
        };
        return updatedFilters;
      });

      this.toEditFilter.set(null);
      this.filter.set({field: null, operand: null, field_value: null});
    }
  }
}
