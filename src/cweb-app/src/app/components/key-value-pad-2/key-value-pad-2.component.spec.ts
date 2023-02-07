import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { KeyValuePad2Component } from './key-value-pad-2.component';

describe('KeyValuePadComponent', () => {
  let component: KeyValuePad2Component;
  let fixture: ComponentFixture<KeyValuePad2Component>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ KeyValuePad2Component ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KeyValuePad2Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
