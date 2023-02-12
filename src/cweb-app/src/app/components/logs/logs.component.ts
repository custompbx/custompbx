import {Component, OnDestroy, OnInit} from '@angular/core';
import {select, Store} from '@ngrx/store';
import {AppState, selectLogsState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {PageEvent} from '@angular/material/paginator';
import {Observable, Subscription} from 'rxjs';
import {State} from '../../store/logs/logs.reducers';
import {GetLogs} from '../../store/logs/logs.actions';

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
  selector: 'app-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.css'],
})
export class LogsComponent implements OnInit, OnDestroy {

  public filters: Array<IfilterField> = [];

  public pageTotal = 0;
  private paginationScale = [1000, 2000, 3000, 4000, 5000];
  public pageEvent: PageEvent = <PageEvent>{
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  };

  public logs: Observable<any>;
  public logs$: Subscription;
  public list: State;
  public loadCounter: number;
  public operands: Array<string> = ['=', '>', '<', 'LIKE', '>=', '<=', '!=', 'NOT LIKE'];
  public sortObject: IsortField = <IsortField>{fields: [], desc: false};
  public filter: IfilterField = <IfilterField>{};

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
  public toEditFilter: number = <number>null;
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

  constructor(
    private store: Store<AppState>,
    private _snackBar: MatSnackBar,
  ) {
    this.logs = this.store.pipe(select(selectLogsState));
  }

  ngOnInit() {
    this.logs$ = this.logs.subscribe((logs) => {
      this.loadCounter = logs.loadCounter;
      this.list = logs;
      if (this.list.logsData.length > 0) {
        this.pageTotal = this.list.logsData[0]['total'];
      }
      if (!logs.errorMessage) {

      } else {
        this._snackBar.open('Error: ' + logs.errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.logs$.unsubscribe();
  }

  getRecords() {
    this.store.dispatch(new GetLogs({
      db_request: {limit: this.pageEvent.pageSize, offset: this.pageEvent.pageIndex, filters: this.filters, order: this.sortObject}
    }));
  }

  removeFilter(filter: IfilterField): void {
    this.pageEvent.pageIndex = 0;
    const index = this.filters.indexOf(filter);

    if (index >= 0) {
      this.filters.splice(index, 1);
    }
  }

  addSorter() {
    this.pageEvent.pageIndex = 0;
    const index = this.sortObject.fields.indexOf(this.sortColumns);

    if (index === -1) {
      this.sortObject.fields.push(this.sortColumns);
    }
  }

  clearSorting() {
    this.pageEvent.pageIndex = 0;
    this.sortObject.fields = [];
  }

  fileTypeByName(str: string): string {
    const res = str.split('.');
    return res[res.length - 1];
  }

  addFilter() {
    this.toEditFilter = null;
    this.pageEvent.pageIndex = 0;
    this.filter.field_value.trim();
    this.filters.push(<IfilterField>this.filter);
    this.filter = <IfilterField>{};
  }

  editFilter(toEdit: number) {
    this.toEditFilter = toEdit;
    this.filter.field = this.filters[toEdit].field;
    this.filter.operand = this.filters[toEdit].operand;
    this.filter.field_value = this.filters[toEdit].field_value;
  }

  saveFilter() {
    this.filters[this.toEditFilter].field = this.filter.field;
    this.filters[this.toEditFilter].operand = this.filter.operand ;
    this.filters[this.toEditFilter].field_value = this.filter.field_value;
    this.toEditFilter = null;
    this.filter.field = null;
    this.filter.operand = null;
    this.filter.field_value = null;
  }

}
