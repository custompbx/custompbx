import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { HepComponent } from './hep.component';

describe('SystemComponent', () => {
  let component: HepComponent;
  let fixture: ComponentFixture<HepComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [
        HepComponent,
      ]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HepComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should compile', () => {
    expect(component).toBeTruthy();
  });
});
