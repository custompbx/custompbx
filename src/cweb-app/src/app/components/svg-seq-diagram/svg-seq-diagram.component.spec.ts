import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { SvgSeqDiagramComponent } from './svg-seq-diagram.component';

describe('SvgSeqDiagramComponent', () => {
  let component: SvgSeqDiagramComponent;
  let fixture: ComponentFixture<SvgSeqDiagramComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ SvgSeqDiagramComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SvgSeqDiagramComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
