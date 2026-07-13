import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { UsersPanelComponent } from './users-panel.component';

describe('UsersPanelComponent', () => {
  let component: UsersPanelComponent;
  let fixture: ComponentFixture<UsersPanelComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [ UsersPanelComponent ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UsersPanelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('uses the prominent warning status for a ringing user', () => {
    expect(component.getUserCardColor({in_call: true})).toBe('yellow');
  });
});
