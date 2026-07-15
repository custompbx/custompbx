import {ComponentFixture, TestBed} from '@angular/core/testing';
import {TabNavComponent} from './tab-nav.component';

describe('TabNavComponent', () => {
  let fixture: ComponentFixture<TabNavComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({imports: [TabNavComponent]}).compileComponents();
    fixture = TestBed.createComponent(TabNavComponent);
    fixture.componentRef.setInput('tabs', ['List', 'Add', 'Delete']);
    fixture.detectChanges();
  });

  it('exposes accessible tabs and updates the selected model on click', () => {
    const buttons = fixture.nativeElement.querySelectorAll('[role="tab"]') as NodeListOf<HTMLButtonElement>;
    expect(buttons.length).toBe(3);
    expect(buttons[0].getAttribute('aria-selected')).toBe('true');

    buttons[1].click();
    fixture.detectChanges();

    expect(fixture.componentInstance.selectedIndex()).toBe(1);
    expect(buttons[1].getAttribute('aria-selected')).toBe('true');
  });

  it('supports arrow navigation and skips disabled tabs', () => {
    fixture.componentRef.setInput('disabled', [false, true, false]);
    fixture.detectChanges();
    const first = fixture.nativeElement.querySelector('[role="tab"]') as HTMLButtonElement;

    first.dispatchEvent(new KeyboardEvent('keydown', {key: 'ArrowRight', bubbles: true}));
    fixture.detectChanges();

    expect(fixture.componentInstance.selectedIndex()).toBe(2);
  });

  it('supports Home and End keyboard navigation', () => {
    fixture.componentInstance.selectedIndex.set(1);
    fixture.detectChanges();
    const buttons = fixture.nativeElement.querySelectorAll('[role="tab"]') as NodeListOf<HTMLButtonElement>;

    buttons[1].dispatchEvent(new KeyboardEvent('keydown', {key: 'End', bubbles: true}));
    fixture.detectChanges();
    expect(fixture.componentInstance.selectedIndex()).toBe(2);

    buttons[2].dispatchEvent(new KeyboardEvent('keydown', {key: 'Home', bubbles: true}));
    fixture.detectChanges();
    expect(fixture.componentInstance.selectedIndex()).toBe(0);
  });
});
