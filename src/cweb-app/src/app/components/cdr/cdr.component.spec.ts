import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { CdrComponent } from './cdr.component';

describe('SystemComponent', () => {
  let component: CdrComponent;
  let fixture: ComponentFixture<CdrComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [
        CdrComponent,
      ]
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CdrComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should compile', () => {
    expect(component).toBeTruthy();
  });
});
