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

    expect(labels).toContain('header.toggleNavigation');
    expect(labels).toContain('CustomPBX dashboard');
    expect(labels).toContain('header.openConversations');
    expect(labels).toContain('header.togglePhone');
    expect(labels).toContain('header.openUserMenu');
    expect(fixture.nativeElement.querySelector('.user-menu .user-name')).not.toBeNull();
    expect(fixture.nativeElement.querySelector(':scope > .user-name')).toBeNull();
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

  it('closes header menus when clicking outside', () => {
    const localeMenu = fixture.nativeElement.querySelector('.locale-menu') as HTMLDetailsElement;
    const userMenu = fixture.nativeElement.querySelector('.user-menu') as HTMLDetailsElement;
    localeMenu.open = true;
    userMenu.open = true;

    component.closeMenusOnOutsideClick({target: document.body} as unknown as Event);

    expect(localeMenu.open).toBeFalse();
    expect(userMenu.open).toBeFalse();
  });

  it('keeps the clicked menu open and closes the other menu', () => {
    const localeMenu = fixture.nativeElement.querySelector('.locale-menu') as HTMLDetailsElement;
    const userMenu = fixture.nativeElement.querySelector('.user-menu') as HTMLDetailsElement;
    localeMenu.open = true;
    userMenu.open = true;

    component.closeMenusOnOutsideClick({
      target: userMenu.querySelector('summary')
    } as unknown as Event);

    expect(localeMenu.open).toBeFalse();
    expect(userMenu.open).toBeTrue();
  });

  it('closes all header menus on Escape', () => {
    const localeMenu = fixture.nativeElement.querySelector('.locale-menu') as HTMLDetailsElement;
    const userMenu = fixture.nativeElement.querySelector('.user-menu') as HTMLDetailsElement;
    localeMenu.open = true;
    userMenu.open = true;

    component.closeMenusOnEscape();

    expect(localeMenu.open).toBeFalse();
    expect(userMenu.open).toBeFalse();
  });
});
