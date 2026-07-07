import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ConfirmBottomSheetComponent } from './confirm-bottom-sheet.component';
import {MAT_BOTTOM_SHEET_DATA, MatBottomSheetRef} from '@angular/material/bottom-sheet';

describe('ConfirmBottomSheetComponent', () => {
  let component: ConfirmBottomSheetComponent;
  let fixture: ComponentFixture<ConfirmBottomSheetComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [ ConfirmBottomSheetComponent ],
      providers: [
        {provide: MatBottomSheetRef, useValue: {dismiss: jasmine.createSpy('dismiss')}},
        {provide: MAT_BOTTOM_SHEET_DATA, useValue: {action: 'delete', case1Text: 'Delete item?'}},
      ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ConfirmBottomSheetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
