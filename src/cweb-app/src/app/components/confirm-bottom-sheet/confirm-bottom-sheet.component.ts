import {Component, Inject} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef} from '@angular/material/bottom-sheet';

@Component({
standalone: true,
    imports: [CommonModule, MaterialModule],
    selector: 'app-confirm-bottom-sheet',
    templateUrl: './confirm-bottom-sheet.component.html',
    styleUrls: ['./confirm-bottom-sheet.component.css']
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
