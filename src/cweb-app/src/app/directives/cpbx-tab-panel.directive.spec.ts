import {Component, signal} from '@angular/core';
import {ComponentFixture, TestBed} from '@angular/core/testing';
import {CpbxTabPanelDirective} from './cpbx-tab-panel.directive';

@Component({
  standalone: true,
  imports: [CpbxTabPanelDirective],
  template: `<section appCpbxTabPanel [active]="active()">Panel</section>`,
})
class TabPanelHostComponent {
  readonly active = signal(false);
}

describe('CpbxTabPanelDirective', () => {
  let fixture: ComponentFixture<TabPanelHostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({imports: [TabPanelHostComponent]}).compileComponents();
    fixture = TestBed.createComponent(TabPanelHostComponent);
    fixture.detectChanges();
  });

  it('keeps inactive panels out of the layout', async () => {
    const panel = fixture.nativeElement.querySelector('section');
    expect(panel.hidden).toBeTrue();
    fixture.componentInstance.active.set(true);
    await fixture.whenStable();
    expect(panel.hidden).toBeFalse();
  });
});
