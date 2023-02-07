import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { KeyValuePadComponent } from './key-value-pad.component';

describe('KeyValuePadComponent', () => {
  let component: KeyValuePadComponent;
  let fixture: ComponentFixture<KeyValuePadComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ KeyValuePadComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KeyValuePadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
