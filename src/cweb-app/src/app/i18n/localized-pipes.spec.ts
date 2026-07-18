import {LocalizedDatePipe} from './localized-date.pipe';
import {LocalizedNumberPipe} from './localized-number.pipe';
import {LocaleService} from './locale.service';

describe('localized pipes', () => {
  const locale = {
    formatDate: jasmine.createSpy('formatDate').and.returnValue('date'),
    formatNumber: jasmine.createSpy('formatNumber').and.returnValue('number'),
  } as unknown as LocaleService;

  it('formats dates and leaves empty values empty', () => {
    const pipe = new LocalizedDatePipe(locale);
    expect(pipe.transform(null)).toBe('');
    expect(pipe.transform('2026-07-17')).toBe('date');
  });

  it('formats numbers and leaves empty values empty', () => {
    const pipe = new LocalizedNumberPipe(locale);
    expect(pipe.transform(null)).toBe('');
    expect(pipe.transform(1234)).toBe('number');
  });
});
