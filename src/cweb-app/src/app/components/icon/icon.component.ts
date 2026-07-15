import {ChangeDetectionStrategy, Component, computed, input} from '@angular/core';

@Component({
  selector: 'app-icon',
  standalone: true,
  template: `
    <svg
      class="cpbx-icon"
      focusable="false"
      [attr.aria-hidden]="label() ? null : 'true'"
      [attr.aria-label]="label() || null"
      [attr.role]="label() ? 'img' : null"
    >
      <use [attr.href]="href()"></use>
    </svg>
  `,
  styles: [`
    :host { display: inline-flex; flex: 0 0 auto; line-height: 0; }
    .cpbx-icon {
      display: block;
      width: 1em;
      height: 1em;
      fill: none;
      stroke: currentColor;
      stroke-width: 1.8;
      stroke-linecap: round;
      stroke-linejoin: round;
      vector-effect: non-scaling-stroke;
    }
  `],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class IconComponent {
  readonly name = input.required<string>();
  readonly label = input('');
  readonly href = computed(() => `#${this.name()}`);
}
