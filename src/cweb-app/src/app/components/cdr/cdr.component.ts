// @ts-ignore
import {Component, OnDestroy, OnInit} from '@angular/core';
import {select, Store} from '@ngrx/store';
import {AppState, selectCDRState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {PageEvent} from '@angular/material/paginator';
import {Observable, Subscription} from 'rxjs';
import {State} from '../../store/cdr/cdr.reducers';
import {GetCDR} from '../../store/cdr/cdr.actions';
import {GetWebSettings, SaveWebSettings} from '../../store/settings/settings.actions';

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
  selector: 'app-cdr',
  templateUrl: './cdr.component.html',
  styleUrls: ['./cdr.component.css'],
})
export class CdrComponent implements OnInit, OnDestroy {

  public filters: Array<IfilterField> = [];

  public pageTotal = 0;
  private paginationScale = [25, 50, 100, 250];
  public pageEvent: PageEvent = <PageEvent>{
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  };

  public cdrs: Observable<any>;
  public cdrs$: Subscription;
  public list: State;
  public loadCounter: number;
  public operands: Array<string> = ['=', '>', '<', 'LIKE', '>=', '<=', '!=', 'NOT LIKE'];
  public sortObject: IsortField = <IsortField>{fields: [], desc: false};
  public filter: IfilterField = <IfilterField>{};
  public toEditFilter: number = <number>null;

  public columns: Array<string>;
  public sortColumns: string;

  public settings = {};
  public moduleOptions = ['auto', 'cdr_pg_csv', 'odbc_cdr'];
  public fieldModule = 'cdr_module';
  public fieldTable = 'cdr_table';
  public fieldFileServeColumn = 'cdr_file_serve_column';
  public fieldFileServerPath = 'cdr_file_server_path';

  constructor(
    private store: Store<AppState>,
    private _snackBar: MatSnackBar,
  ) {
    this.cdrs = this.store.pipe(select(selectCDRState));
  }

  ngOnInit() {
    this.cdrs$ = this.cdrs.subscribe((cdr) => {
      this.loadCounter = cdr.loadCounter;
      this.list = cdr;
      if (this.list.cdrData.length > 0) {
        this.columns = [];
        this.pageTotal = this.list.cdrData[0]['total'];
        Object.keys(this.list.cdrData[0]).forEach( key => {
          if (key === 'total' ) {
            return;
          }
          this.columns.push(key);
        });
      }
      this.settings[this.fieldModule] = cdr.settings[this.fieldModule];
      this.settings[this.fieldTable] = cdr.settings[this.fieldTable];
      this.settings[this.fieldFileServeColumn] = cdr.settings[this.fieldFileServeColumn];
      this.settings[this.fieldFileServerPath] = cdr.settings[this.fieldFileServerPath];
      if (!cdr.errorMessage) {

      } else {
        this._snackBar.open('Error: ' + cdr.errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.cdrs$.unsubscribe();
  }

  getRecords() {
    if (!this.settings[this.fieldModule] || !this.settings[this.fieldFileServeColumn]) {
      this.getSettings();
    }
    this.store.dispatch(new GetCDR({
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

  tabChanged(event) {
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
    const webSettings: object = {};
    webSettings[this.fieldModule] = this.settings[this.fieldModule];
    webSettings[this.fieldTable] = this.settings[this.fieldTable];
    webSettings[this.fieldFileServeColumn] = this.settings[this.fieldFileServeColumn];
    webSettings[this.fieldFileServerPath] = this.settings[this.fieldFileServerPath];
    this.store.dispatch(new SaveWebSettings({web_settings: webSettings}));
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
