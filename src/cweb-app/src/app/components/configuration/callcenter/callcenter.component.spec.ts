import { signal } from '@angular/core';
import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { CallcenterComponent } from './callcenter.component';

describe('CallcenterComponent', () => {
  let component: CallcenterComponent;
  let fixture: ComponentFixture<CallcenterComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [ CallcenterComponent ],
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

  it('uses shared native controls for the agent filter bar', () => {
    fixture.destroy();
    fixture = TestBed.createComponent(CallcenterComponent);
    component = fixture.componentInstance;
    (component as any).list = signal({
      exists: true,
      queues: {},
      agents: {total: 1, table: [{id: 1, name: 'agent-1'}]},
      changed: {agents: {}, tiers: {}},
    });
    component.detailTabIndex = 1;

    fixture.detectChanges();

    const filterBar = fixture.nativeElement.querySelector('.callcenter-filter-bar');
    expect(filterBar).toBeTruthy();
    expect(filterBar.querySelectorAll('select.cpbx-select').length).toBe(3);
    expect(filterBar.querySelector('input.cpbx-input')).toBeTruthy();
    expect(filterBar.querySelector('mat-form-field')).toBeNull();
  });
});
