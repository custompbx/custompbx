import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ConfirmBottomSheetComponent } from './confirm-bottom-sheet.component';

describe('ConfirmBottomSheetComponent', () => {
  let component: ConfirmBottomSheetComponent;
  let fixture: ComponentFixture<ConfirmBottomSheetComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ ConfirmBottomSheetComponent ]
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
