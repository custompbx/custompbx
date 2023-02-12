import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { CallcenterComponent } from './callcenter.component';

describe('CallcenterComponent', () => {
  let component: CallcenterComponent;
  let fixture: ComponentFixture<CallcenterComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ CallcenterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CallcenterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
