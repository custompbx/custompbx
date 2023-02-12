import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ContextsComponent } from './contexts.component';

describe('ContextsComponent', () => {
  let component: ContextsComponent;
  let fixture: ComponentFixture<ContextsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ ContextsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ContextsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
