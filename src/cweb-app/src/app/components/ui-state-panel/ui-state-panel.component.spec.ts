import {ComponentFixture, TestBed} from '@angular/core/testing';
import {UiStatePanelComponent} from './ui-state-panel.component';

describe('UiStatePanelComponent', () => {
  let fixture: ComponentFixture<UiStatePanelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({imports: [UiStatePanelComponent]}).compileComponents();
    fixture = TestBed.createComponent(UiStatePanelComponent);
  });

  it('uses alert semantics for danger states', () => {
    fixture.componentInstance.tone = 'danger';
    fixture.componentInstance.title = 'Unable to load';
    fixture.detectChanges();
    expect(fixture.nativeElement.querySelector('section').getAttribute('role')).toBe('alert');
  });
});
