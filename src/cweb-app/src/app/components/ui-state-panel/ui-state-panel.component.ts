import {Component, Input} from '@angular/core';
import {IconComponent} from '../icon/icon.component';

@Component({
  standalone: true,
  selector: 'app-ui-state-panel',
  imports: [IconComponent],
  template: `
    <section class="state-panel" [attr.role]="tone === 'danger' ? 'alert' : 'status'">
      <app-icon [class]="tone" [name]="icon"></app-icon>
      <h2>{{title}}</h2>
      @if (message) { <p>{{message}}</p> }
      <div class="state-actions"><ng-content></ng-content></div>
    </section>
  `,
  styles: [`
    :host { display: block; }
    app-icon { font-size: 40px; color: var(--cpbx-text-muted); }
    app-icon.danger { color: var(--cpbx-danger); }
    app-icon.success { color: var(--cpbx-success); }
    h2 { margin: 12px 0 6px; font-size: 18px; }
    p { margin: 0 auto 16px; max-width: 560px; color: var(--cpbx-text-muted); }
    .state-actions { display: flex; justify-content: center; gap: 8px; }
  `]
})
export class UiStatePanelComponent {
  @Input() icon = 'info';
  @Input() title = '';
  @Input() message = '';
  @Input() tone: 'default' | 'success' | 'danger' = 'default';
}
