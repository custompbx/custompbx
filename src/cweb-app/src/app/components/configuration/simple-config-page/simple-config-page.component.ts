import {ChangeDetectionStrategy, Component, computed, input} from '@angular/core';
import {TranslocoPipe} from '@jsverse/transloco';
import {InnerHeaderComponent} from '../../inner-header/inner-header.component';
import {KeyValuePad2Component} from '../../key-value-pad-2/key-value-pad-2.component';
import {ModuleNotExistsBannerComponent} from '../module-not-exists-banner/module-not-exists-banner.component';

@Component({
  selector: 'app-simple-config-page',
  standalone: true,
  imports: [InnerHeaderComponent, KeyValuePad2Component, ModuleNotExistsBannerComponent, TranslocoPipe],
  templateUrl: './simple-config-page.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SimpleConfigPageComponent {
  readonly name = input.required<string>();
  readonly loadCounter = input(0);
  readonly module = input<any>();
  readonly collectionKey = input('settings');
  readonly sectionTitleKey = input('configuration.settings');
  readonly contentTitleKey = input('common.parameters');
  readonly dispatchersCallbacks = input.required<any>();
  readonly fieldsMask = input<any>();
  readonly pending = input(false);
  readonly sortable = input(false);

  protected readonly collection = computed(() => this.module()?.[this.collectionKey()]);
  protected readonly newItems = computed(() => this.collection()?.new);
}
