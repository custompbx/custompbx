import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { AclComponent } from './acl.component';

describe('AclComponent', () => {
  let component: AclComponent;
  let fixture: ComponentFixture<AclComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ AclComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AclComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
