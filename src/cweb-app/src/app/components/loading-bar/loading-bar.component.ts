import {ChangeDetectionStrategy, Component, input} from '@angular/core';

@Component({
  selector: 'app-loading-bar',
  standalone: true,
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <div class="cpbx-progress" role="progressbar" [attr.aria-label]="label()">
      <span class="cpbx-progress__bar"></span>
    </div>
  `,
})
export class LoadingBarComponent {
  readonly label = input('Loading');
}
