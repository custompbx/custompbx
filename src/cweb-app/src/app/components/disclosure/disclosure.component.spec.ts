import {ComponentFixture, TestBed} from '@angular/core/testing';
import {DisclosureComponent} from './disclosure.component';

describe('DisclosureComponent', () => {
  let fixture: ComponentFixture<DisclosureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({imports: [DisclosureComponent]}).compileComponents();
    fixture = TestBed.createComponent(DisclosureComponent);
    fixture.detectChanges();
  });

  it('exposes an accessible collapsed trigger', () => {
    const trigger = fixture.nativeElement.querySelector('button') as HTMLButtonElement;

    expect(trigger.getAttribute('aria-expanded')).toBe('false');
    expect(fixture.nativeElement.querySelector('.cpbx-disclosure__content')).toBeNull();
  });

  it('renders content and emits opened and closed events', () => {
    const opened = jasmine.createSpy('opened');
    const closed = jasmine.createSpy('closed');
    fixture.componentInstance.opened.subscribe(opened);
    fixture.componentInstance.closed.subscribe(closed);
    const trigger = fixture.nativeElement.querySelector('button') as HTMLButtonElement;

    trigger.click();
    fixture.detectChanges();
    expect(fixture.componentInstance.expanded()).toBeTrue();
    expect(trigger.getAttribute('aria-expanded')).toBe('true');
    expect(opened).toHaveBeenCalledTimes(1);

    trigger.click();
    fixture.detectChanges();
    expect(fixture.componentInstance.expanded()).toBeFalse();
    expect(closed).toHaveBeenCalledTimes(1);
  });

  it('does not toggle while disabled', () => {
    fixture.componentRef.setInput('disabled', true);
    fixture.detectChanges();
    const trigger = fixture.nativeElement.querySelector('button') as HTMLButtonElement;

    trigger.click();
    fixture.detectChanges();

    expect(fixture.componentInstance.expanded()).toBeFalse();
  });
});
