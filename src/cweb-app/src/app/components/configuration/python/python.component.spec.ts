import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { PythonComponent } from './python.component';

describe('SofiaComponent', () => {
  let component: PythonComponent;
  let fixture: ComponentFixture<PythonComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [ PythonComponent ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PythonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
