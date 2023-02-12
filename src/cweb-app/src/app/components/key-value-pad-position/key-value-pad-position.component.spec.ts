import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { KeyValuePadPositionComponent } from './key-value-pad-position.component';

describe('KeyValuePadComponent', () => {
  let component: KeyValuePadPositionComponent;
  let fixture: ComponentFixture<KeyValuePadPositionComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ KeyValuePadPositionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KeyValuePadPositionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
