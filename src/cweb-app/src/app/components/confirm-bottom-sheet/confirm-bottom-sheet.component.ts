import {Component, Inject} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef} from '@angular/material/bottom-sheet';

@Component({
standalone: true,
    imports: [CommonModule, MaterialModule],
    selector: 'app-confirm-bottom-sheet',
    template: '<div [ngSwitch]="data.action">\n' +
        '  <h3 *ngSwitchCase="\'delete\'">{{data.case1Text}}</h3>\n' +
        '  <h3 *ngSwitchCase="\'rename\'">{{data.case2Text}}</h3>\n' +
        '</div>' +
        '<mat-nav-list>\n' +
        '  <a mat-list-item><button mat-button mat-line color="warn" (click)="confirmAction(true)">Confirm</button></a>\n' +
        '  <a mat-list-item><button mat-button mat-line color="primary" (click)="confirmAction(false)">Cancel</button></a>\n' +
        '</mat-nav-list>'
})
export class ConfirmBottomSheetComponent {
  constructor(
    private bottomSheetRef: MatBottomSheetRef<ConfirmBottomSheetComponent>,
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: any
  ) {
  }

  confirmAction(event: boolean): void {
    this.bottomSheetRef.dismiss(event);
  }
}
