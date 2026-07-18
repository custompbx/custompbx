import {Directive, HostBinding, input} from '@angular/core';

@Directive({
  selector: '[appCpbxTabPanel]',
  standalone: true,
})
export class CpbxTabPanelDirective {
  readonly active = input.required<boolean>();

  @HostBinding('class.cpbx-tab-panel')
  readonly panelClass = true;

  @HostBinding('class.cpbx-tab-panel--nested')
  readonly nestedClass = true;

  @HostBinding('attr.role')
  readonly role = 'tabpanel';

  @HostBinding('attr.hidden')
  get hidden(): '' | null {
    return this.active() ? null : '';
  }
}
