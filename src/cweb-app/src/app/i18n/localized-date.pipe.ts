import {Pipe, PipeTransform} from '@angular/core';
import {LocaleService} from './locale.service';

@Pipe({
  name: 'localizedDate',
  standalone: true,
  pure: false,
})
export class LocalizedDatePipe implements PipeTransform {
  constructor(private readonly locale: LocaleService) {}

  transform(
    value: Date | number | string | null | undefined,
    options?: Intl.DateTimeFormatOptions
  ): string {
    if (value === null || value === undefined || value === '') {
      return '';
    }
    return this.locale.formatDate(value, options);
  }
}
