import {Component, OnDestroy, OnInit, inject, signal, computed, effect} from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';

import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {select, Store} from '@ngrx/store';
import {AppState, selectLogsState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {PageEvent} from '@angular/material/paginator';
import {State} from '../../store/logs/logs.reducers';
import {GetLogs} from '../../store/logs/logs.actions';
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
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent],
  selector: 'app-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.css'],
})
export class LogsComponent {

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private _snackBar = inject(MatSnackBar);

  // --- Reactive State from NgRx using toSignal ---
  private logsState = toSignal(
    this.store.pipe(select(selectLogsState)),
    {
      initialValue: {
        logsData: [],
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed State from NgRx State ---
  // The full logs list is needed for the component
  public list = computed(() => this.logsState());
  public loadCounter = computed(() => this.logsState().loadCounter);
  // Compute pageTotal dynamically based on the first item in logsData
  public pageTotal = computed(() => {
    const data = this.logsState().logsData;
    // Check if data exists and the first element has a 'total' property
    return data && data.length > 0 && data[0]['total'] ? data[0]['total'] : 0;
  });

  // --- Effect as a Class Property ---
  private snackbarEffect = effect(() => {
    const errorMessage = this.logsState().errorMessage;

    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
    // No explicit reset logic needed here, as the component relies on the store's state clearing the error.
  });

  // --- Local State as Signals/Properties ---
  public filters = signal<Array<IfilterField>>([]);

  protected paginationScale = [1000, 2000, 3000, 4000, 5000];
  // PageEvent is an object usually updated by an event handler, keep it as a signal for reactivity
  public pageEvent = signal<PageEvent>({
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  } as PageEvent);

  public operands: Array<string> = ['=', '>', '<', 'LIKE', '>=', '<=', '!=', 'NOT LIKE'];
  public sortObject = signal<IsortField>({fields: [], desc: false});
  // Use signal for filter object being edited/added
  public filter = signal<IfilterField>({field: null, operand: null, field_value: null});

  public columns: Array<string> = [
    'body',
    'created',
    'log_file',
    'log_func',
    'log_level',
    'log_line',
    'text_channel',
    'total',
    'user_data',
  ];
  public toEditFilter = signal<number | null>(null);
  public sortColumns: string;

  public colors = {
    1: 'log-alert',
    2: 'log-crit',
    3: 'log-err',
    4: 'log-warning',
    5: 'log-notice',
    6: 'log-info',
    7: 'log-debug',
    8: 'log-console',
  };

  handlePageEvent(event: PageEvent) {
    this.pageEvent.set(event);
    this.getRecords();
  }

  getRecords() {
    const currentPageEvent = this.pageEvent();
    this.store.dispatch(new GetLogs({
      db_request: {
        limit: currentPageEvent.pageSize,
        offset: currentPageEvent.pageIndex,
        filters: this.filters(), // Read signal value
        order: this.sortObject() // Read signal value
      }
    }));
  }

  removeFilter(filterToRemove: IfilterField): void {
    this.pageEvent.update(e => ({...e, pageIndex: 0})); // Reset index using update
    this.filters.update(currentFilters => {
      const index = currentFilters.indexOf(filterToRemove);
      if (index >= 0) {
        return currentFilters.filter(f => f !== filterToRemove);
      }
      return currentFilters;
    });
  }

  addSorter() {
    this.pageEvent.update(e => ({...e, pageIndex: 0})); // Reset index
    this.sortObject.update(currentSort => {
      if (this.sortColumns && currentSort.fields.indexOf(this.sortColumns) === -1) {
        return {
          ...currentSort,
          fields: [...currentSort.fields, this.sortColumns]
        };
      }
      return currentSort;
    });
  }

  clearSorting() {
    this.pageEvent.update(e => ({...e, pageIndex: 0})); // Reset index
    this.sortObject.set({fields: [], desc: false});
  }

  fileTypeByName(str: string): string {
    const res = str.split('.');
    return res[res.length - 1];
  }

  addFilter() {
    this.toEditFilter.set(null);
    this.pageEvent.update(e => ({...e, pageIndex: 0})); // Reset index

    // Read and trim value from signal
    const currentFilter = this.filter();
    const trimmedFilterValue = currentFilter.field_value ? currentFilter.field_value.trim() : null;

    if (currentFilter.field && currentFilter.operand && trimmedFilterValue) {
      // Add a copy of the filter object (using spread to ensure immutability for signal update)
      this.filters.update(f => [...f, {...currentFilter, field_value: trimmedFilterValue}]);
    }

    // Reset the filter signal for new input
    this.filter.set({field: null, operand: null, field_value: null});
  }

  editFilter(toEdit: number) {
    this.toEditFilter.set(toEdit);
    const filterToEdit = this.filters()[toEdit];
    // Copy values to the mutable filter signal
    this.filter.set({...filterToEdit});
  }

  saveFilter() {
    const editIndex = this.toEditFilter();
    if (editIndex !== null) {
      const savedFilter = this.filter();

      this.filters.update(currentFilters => {
        // Create an updated array with the new filter data
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
