import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { OdbcCdrComponent } from './odbc-cdr.component';

describe('AclComponent', () => {
  let component: OdbcCdrComponent;
  let fixture: ComponentFixture<OdbcCdrComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ OdbcCdrComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(OdbcCdrComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
