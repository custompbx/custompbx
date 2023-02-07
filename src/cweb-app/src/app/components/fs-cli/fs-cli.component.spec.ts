import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { FsCliComponent } from './fs-cli.component';

describe('FscliComponent', () => {
  let component: FsCliComponent;
  let fixture: ComponentFixture<FsCliComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ FsCliComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FsCliComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
