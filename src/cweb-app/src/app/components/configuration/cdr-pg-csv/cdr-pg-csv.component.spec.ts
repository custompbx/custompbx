import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { CdrPgCsvComponent } from './cdr-pg-csv.component';

describe('AclComponent', () => {
  let component: CdrPgCsvComponent;
  let fixture: ComponentFixture<CdrPgCsvComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ CdrPgCsvComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CdrPgCsvComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
