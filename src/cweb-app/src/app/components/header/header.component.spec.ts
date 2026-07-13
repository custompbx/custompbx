import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { HeaderComponent } from './header.component';
import { Store } from '@ngrx/store';
import { StartPhone, ToggleShowPhone } from '../../store/header/header.actions';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [ HeaderComponent ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('renders key action buttons with accessible labels', () => {
    const labels = Array.from<HTMLElement>(fixture.nativeElement.querySelectorAll('[aria-label]'))
      .map((element) => element.getAttribute('aria-label'));

    expect(labels).toContain('Toggle navigation');
    expect(labels).toContain('CustomPBX dashboard');
    expect(labels).toContain('Open conversations');
    expect(labels).toContain('Toggle phone');
    expect(labels).toContain('Open user menu');
  });

  it('does not require a loaded user before first render', () => {
    expect(() => fixture.detectChanges()).not.toThrow();
  });

  it('starts and shows the phone on the first click', () => {
    const store = TestBed.inject(Store);
    const dispatch = spyOn(store, 'dispatch');

    component.showHidePhone();

    const dispatchedTypes = dispatch.calls.allArgs()
      .map(([action]) => (action as unknown as {type: string}).type);
    expect(dispatchedTypes).toEqual([StartPhone.type, ToggleShowPhone.type]);
  });
});
