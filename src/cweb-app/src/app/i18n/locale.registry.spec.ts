import {
  isLocaleCode,
  LANGUAGE_MENU_LOCALES,
  localeFromBrowser,
  localeFromLegacy,
  normalizeLocale,
  SUPPORTED_LOCALES
} from './locale.registry';

describe('locale registry', () => {
  it('contains the complete supported locale set', () => {
    expect(SUPPORTED_LOCALES.length).toBe(14);
    expect(SUPPORTED_LOCALES.map(locale => locale.code)).toContain('zh-Hans');
    expect(SUPPORTED_LOCALES.map(locale => locale.code)).toContain('pt-BR');
  });

  it('rejects arbitrary catalog paths', () => {
    expect(isLocaleCode('../../secrets')).toBeFalse();
    expect(normalizeLocale('../../secrets')).toBe('en');
  });

  it('maps legacy language values', () => {
    expect(localeFromLegacy(0)).toBe('en');
    expect(localeFromLegacy(1)).toBe('ru');
    expect(localeFromLegacy(99)).toBe('en');
  });

  it('matches regional browser locales to supported catalogs', () => {
    expect(localeFromBrowser('en-US')).toBe('en');
    expect(localeFromBrowser('pt-PT')).toBe('pt-BR');
    expect(localeFromBrowser('zh-CN')).toBe('zh-Hans');
    expect(localeFromBrowser('xx-ZZ')).toBeNull();
  });

  it('marks only Arabic and Persian as RTL', () => {
    expect(SUPPORTED_LOCALES.filter(locale => locale.direction === 'rtl').map(locale => locale.code)).toEqual(['ar', 'fa']);
  });

  it('pins English first and alphabetizes the remaining language menu entries', () => {
    expect(LANGUAGE_MENU_LOCALES[0].code).toBe('en');
    expect(LANGUAGE_MENU_LOCALES.slice(1).map(locale => locale.label)).toEqual([
      'Arabic',
      'Chinese',
      'French',
      'German',
      'Hindi',
      'Italian',
      'Japanese',
      'Korean',
      'Persian',
      'Portuguese',
      'Russian',
      'Spanish',
      'Turkish',
    ]);
  });
});
