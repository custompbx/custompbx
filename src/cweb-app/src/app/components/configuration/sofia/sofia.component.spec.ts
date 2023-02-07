import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { SofiaComponent } from './sofia.component';

describe('SofiaComponent', () => {
  let component: SofiaComponent;
  let fixture: ComponentFixture<SofiaComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ SofiaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SofiaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
