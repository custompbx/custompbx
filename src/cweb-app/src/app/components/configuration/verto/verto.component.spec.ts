import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { VertoComponent } from './verto.component';

describe('SofiaComponent', () => {
  let component: VertoComponent;
  let fixture: ComponentFixture<VertoComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ VertoComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(VertoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
