import {Component, Inject, OnDestroy, OnInit} from '@angular/core';
import {select, Store} from '@ngrx/store';
import {AppState, selectHEPState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {PageEvent} from '@angular/material/paginator';
import {Observable, Subscription} from 'rxjs';
import {State} from '../../store/hep/hep.reducers';
import {GetHEP, GetHEPDetails} from '../../store/hep/hep.actions';
import {MAT_BOTTOM_SHEET_DATA, MatBottomSheet, MatBottomSheetRef} from '@angular/material/bottom-sheet';
import * as svg from 'save-svg-as-png';

export interface IfilterField {
  field: string | null;
  operand: string | null;
  field_value: string | null;
}

export interface IsortField {
  fields: Array<string>;
  desc: boolean;
}

@Component({
  selector: 'app-hep',
  templateUrl: './hep.component.html',
  styleUrls: ['./hep.component.css'],
})
export class HepComponent implements OnInit, OnDestroy {

  public filters: Array<IfilterField> = [];

  public pageTotal = 0;
  private paginationScale = [100, 200, 500, 1000];
  public pageEvent: PageEvent = <PageEvent>{
    length: 0,
    pageIndex: 0,
    pageSize: this.paginationScale[0],
  };

  public heps: Observable<any>;
  public heps$: Subscription;
  public list: State;
  public loadCounter: number;
  public operands: Array<string> = ['=', '>', '<', 'LIKE', '>=', '<=', '!=', 'NOT LIKE'];
  public sortObject: IsortField = <IsortField>{fields: [], desc: false};
  public filter: IfilterField = <IfilterField>{};
  public columns: Array<string> = [
    'hep_timestamp',
    'sip_first_method',
    'sip_call_id',
    'sip_from_user',
    'sip_from_host',
    // 'sip_to_user',
    // 'sip_to_host',
    'sip_uri_user',
    'sip_uri_host',
    'hep_dst_ip',
    'hep_dst_port',
    'hep_src_ip',
    'hep_src_port',
    'sip_user_agent',
    'hep_node_id'
  ];
  public toEditFilter: number = <number>null;

  public sortColumns: string;
  public selectedIndex: number;
  public toView: { [index: string]: boolean } = {};
  public showMsg = null;

  constructor(
    private store: Store<AppState>,
    private _snackBar: MatSnackBar,
    private _bottomSheet: MatBottomSheet
  ) {
    this.heps = this.store.pipe(select(selectHEPState));
  }

  ngOnInit() {
    this.heps$ = this.heps.subscribe((hep) => {
      this.loadCounter = hep.loadCounter;
      this.list = hep;
      if (this.list.hepData.length > 0) {
        this.pageTotal = this.list.hepData[0]['total'];
      }
      if (!hep.errorMessage) {

      } else {
        this._snackBar.open('Error: ' + hep.errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.heps$.unsubscribe();
  }

  getRecords() {
    this.toView = {};
    this.store.dispatch(new GetHEP({
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
    this.filters[this.toEditFilter].operand = this.filter.operand;
    this.filters[this.toEditFilter].field_value = this.filter.field_value;
    this.toEditFilter = null;
    this.filter.field = null;
    this.filter.operand = null;
    this.filter.field_value = null;
  }

  chooseCallId(row) {
    this.toView[row['sip_call_id']] = !this.toView[row['sip_call_id']];
  }

  mainTabChanged(event) {
    if (event === 1) {
      const res = this.selectedCallIds();
      if (res.length === 0) {
        return;
      }
      this.store.dispatch(new GetHEPDetails({values: res}));
    }
  }

  selectedCallIds() {
    return Object.keys(this.toView).filter(callId => this.toView[callId]);
  }

  gotMsg(msg) {
    this.showMsg = msg;
  }

  openExportBottomSheet(): void {
    this._bottomSheet.open(BottomSheetExportComponent, {
      data: this.list.hepDetails,
    });
  }
}

@Component({
  selector: 'app-bottom-sheet-export',
  template: `
    <mat-nav-list>
      <a mat-list-item (click)="saveTextAsPng()">
        <span mat-line>PNG</span>
        <span mat-line>Export as png picture</span>
      </a>

      <a mat-list-item (click)="saveTextAsFile()">
        <span mat-line>TXT</span>
        <span mat-line>Export as txt file</span>
      </a>
    </mat-nav-list>
  `,
})
export class BottomSheetExportComponent {
  constructor(
    private _bottomSheetRef: MatBottomSheetRef<BottomSheetExportComponent>,
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: any
  ) {
  }

  openLink(event: MouseEvent): void {
    this._bottomSheetRef.dismiss();
    event.preventDefault();
  }

  saveTextAsFile() {
    this._bottomSheetRef.dismiss();

    let txt = '';
    let filename = '';
    const type = 'text/plain';
    if (this.data.length > 0) {
      filename = this.data[0].hep_timestamp + '.txt';
      this.data.forEach((msg) => {
        txt += '\n' + msg.hep_payload;
      });
    } else {
      filename = 'empty.txt';
    }

    const a = document.createElement('a');
    a.href = URL.createObjectURL(
      new Blob([txt], {
        type: type
      })
    );
    a.setAttribute('download', filename);
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
  }

  saveTextAsPng() {
    this._bottomSheetRef.dismiss();
    svg.saveSvgAsPng(document.getElementById('idOfMySvgGraphic'), 'callflow.png', {scale: 1, modifyCss: 0});
  }
}
