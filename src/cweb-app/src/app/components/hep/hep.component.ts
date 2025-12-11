import {Component, OnDestroy, OnInit, inject, signal, computed, effect} from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';

import {MaterialModule} from "../../../material-module";
import {select, Store} from '@ngrx/store';
import {AppState, selectHEPState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {PageEvent} from '@angular/material/paginator';
import {State} from '../../store/hep/hep.reducers';
import {GetHEP, GetHEPDetails} from '../../store/hep/hep.actions';
import {MAT_BOTTOM_SHEET_DATA, MatBottomSheet, MatBottomSheetRef} from '@angular/material/bottom-sheet';
import * as svg from 'save-svg-as-png';
import {FormsModule} from "@angular/forms";
import {InnerHeaderComponent} from "../inner-header/inner-header.component";
import {SvgSeqDiagramComponent} from "../svg-seq-diagram/svg-seq-diagram.component";
import {NgClass} from "@angular/common";

export interface IfilterField {
  field: string | null;
  operand: string | null;
  field_value: string | null;
}

export interface IsortField {
  fields: Array<string>;
  desc: boolean;
}

// Defining the BottomSheetExportComponent here to ensure it's defined before it's used
@Component({
  standalone: true,
  imports: [MaterialModule],
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
  `
})
export class BottomSheetExportComponent {
  // Use inject for dependencies
  private _bottomSheetRef = inject(MatBottomSheetRef<BottomSheetExportComponent>);
  public data = inject(MAT_BOTTOM_SHEET_DATA); // Data is a plain value

  saveTextAsFile() {
    this._bottomSheetRef.dismiss();

    let txt = '';
    let filename = '';
    const type = 'text/plain';

    // Check if data is array and has elements
    if (Array.isArray(this.data) && this.data.length > 0) {
      filename = this.data[0].hep_timestamp + '.txt';
      this.data.forEach((msg: any) => {
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
    // Assuming svg.saveSvgAsPng is available
    if (typeof svg !== 'undefined' && svg.saveSvgAsPng) {
      // 'idOfMySvgGraphic' must be the ID of the SVG element in the parent HepComponent template
      svg.saveSvgAsPng(document.getElementById('idOfMySvgGraphic'), 'callflow.png', {scale: 1, modifyCss: 0});
    } else {
      console.error('save-svg-as-png library not loaded or function missing.');
    }
  }
}

@Component({
  standalone: true,
  imports: [MaterialModule, FormsModule, InnerHeaderComponent, SvgSeqDiagramComponent, NgClass],
  selector: 'app-hep',
  templateUrl: './hep.component.html',
  styleUrls: ['./hep.component.css']
})
export class HepComponent implements OnInit, OnDestroy {

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private _snackBar = inject(MatSnackBar);
  private _bottomSheet = inject(MatBottomSheet);

  // --- Reactive State from NgRx using toSignal ---
  private hepState = toSignal(
    this.store.pipe(select(selectHEPState)),
    {
      initialValue: {
        hepData: [],
        hepDetails: [],
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed State from NgRx State ---
  public list = computed(() => this.hepState());
  public loadCounter = computed(() => this.hepState().loadCounter);
  // Compute pageTotal dynamically based on the first item in hepData
  public pageTotal = computed(() => {
    const data = this.hepState().hepData;
    // Use optional chaining and nullish coalescing for safe access
    return data && data.length > 0 && data[0]?.['total'] ? data[0]['total'] : 0;
  });

  // --- Effect for Side Effects (Replaces Subscription logic for error handling) ---
  private snackbarEffect = effect(() => {
    const errorMessage = this.hepState().errorMessage;

    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  // --- Local State as Signals/Properties ---
  public filters = signal<Array<IfilterField>>([]);
  protected paginationScale = [100, 200, 500, 1000];

  // PageEvent signal
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
  public toEditFilter = signal<number | null>(null);
  public sortColumns: string | null = null;
  public selectedIndex: number = 0; // Initialize index

  // Use a signal for the dictionary of call IDs to view
  public toView = signal<{ [index: string]: boolean }>({});
  public showMsg = signal<any | null>(null); // Use signal for showMsg

  // Constructor and Subscription logic are removed.

  ngOnInit() {
    // Initialization logic can remain here if necessary.
  }

  ngOnDestroy() {
    // The subscription cleanup is handled automatically by toSignal.
  }

  handlePageEvent(event: PageEvent) {
    this.pageEvent.set(event);
    this.getRecords();
  }

  getRecords() {
    this.toView.set({}); // Reset toView signal
    const currentPageEvent = this.pageEvent();

    this.store.dispatch(new GetHEP({
      db_request: {
        limit: currentPageEvent.pageSize,
        offset: currentPageEvent.pageIndex,
        filters: this.filters(), // Read signal value
        order: this.sortObject() // Read signal value
      }
    }));
  }

  removeFilter(filterToRemove: IfilterField): void {
    this.pageEvent.update(e => ({...e, pageIndex: 0})); // Reset index
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

  chooseCallId(row: any) {
    const callId = row['sip_call_id'];
    this.toView.update(currentView => ({
      ...currentView,
      [callId]: !currentView[callId] // Toggle the state
    }));
  }

  mainTabChanged(event: number) {
    if (event === 1) {
      const res = this.selectedCallIds();
      if (res.length === 0) {
        return;
      }
      this.store.dispatch(new GetHEPDetails({values: res}));
    }
  }

  selectedCallIds(): Array<string> {
    const currentView = this.toView();
    // Return keys where the value is true
    return Object.keys(currentView).filter(callId => currentView[callId]);
  }

  gotMsg(msg: any) {
    this.showMsg.set(msg);
  }

  openExportBottomSheet(): void {
    // Pass hepDetails from the computed list
    this._bottomSheet.open(BottomSheetExportComponent, {
      data: this.list().hepDetails,
    });
  }
}
