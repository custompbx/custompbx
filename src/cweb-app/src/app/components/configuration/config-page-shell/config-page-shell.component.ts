import {ChangeDetectionStrategy, Component, input, model} from '@angular/core';
import {InnerHeaderComponent} from '../../inner-header/inner-header.component';
import {TabNavComponent} from '../../tab-nav/tab-nav.component';
import {ModuleNotExistsBannerComponent} from '../module-not-exists-banner/module-not-exists-banner.component';

@Component({
  selector: 'app-config-page-shell',
  standalone: true,
  imports: [InnerHeaderComponent, ModuleNotExistsBannerComponent, TabNavComponent],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <app-inner-header
      [name]="name()"
      [translationKey]="translationKey()"
      [loadCounter]="loadCounter()"
    />
    <app-module-not-exists-banner [list]="module()" />
    @if (module() && module().exists !== false) {
      <app-tab-nav
        [ariaLabel]="ariaLabel()"
        [ariaLabelKey]="ariaLabelKey()"
        [tabs]="tabs()"
        [tabKeys]="tabKeys()"
        [(selectedIndex)]="selectedIndex"
      />
    }
  `,
})
export class ConfigPageShellComponent {
  readonly name = input.required<string>();
  readonly translationKey = input('');
  readonly loadCounter = input(0);
  readonly module = input.required<any>();
  readonly tabs = input.required<readonly string[]>();
  readonly tabKeys = input<readonly string[]>([]);
  readonly ariaLabel = input('Configuration sections');
  readonly ariaLabelKey = input('');
  readonly selectedIndex = model(0);
}
