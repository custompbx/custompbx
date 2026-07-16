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

  it('places export actions before the sequence diagram', () => {
    const actions = fixture.nativeElement.querySelector('.hep-details-toolbar') as HTMLElement;
    const body = fixture.nativeElement.querySelector('.hep-details-body') as HTMLElement;
    expect(actions).not.toBeNull();
    expect(body).not.toBeNull();
    expect(actions.compareDocumentPosition(body) & Node.DOCUMENT_POSITION_FOLLOWING).toBeTruthy();
  });
});
