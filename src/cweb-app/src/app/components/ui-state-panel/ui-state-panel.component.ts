import {Component, Input} from '@angular/core';
import {MaterialModule} from '../../../material-module';

@Component({
  standalone: true,
  selector: 'app-ui-state-panel',
  imports: [MaterialModule],
  template: `
    <section class="state-panel" [attr.role]="tone === 'danger' ? 'alert' : 'status'">
      <mat-icon [class]="tone">{{icon}}</mat-icon>
      <h2>{{title}}</h2>
      @if (message) { <p>{{message}}</p> }
      <div class="state-actions"><ng-content></ng-content></div>
    </section>
  `,
  styles: [`
    :host { display: block; }
    mat-icon { width: 40px; height: 40px; font-size: 40px; color: var(--cpbx-text-muted); }
    mat-icon.danger { color: var(--cpbx-danger); }
    mat-icon.success { color: var(--cpbx-success); }
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
