import {Pipe, PipeTransform} from '@angular/core';
import {LocaleService} from './locale.service';

@Pipe({
  name: 'localizedNumber',
  standalone: true,
  pure: false,
})
export class LocalizedNumberPipe implements PipeTransform {
  constructor(private readonly locale: LocaleService) {}

  transform(value: number | null | undefined, options?: Intl.NumberFormatOptions): string {
    if (value === null || value === undefined || Number.isNaN(value)) {
      return '';
    }
    return this.locale.formatNumber(value, options);
  }
}
