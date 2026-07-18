import {LocaleCode, SupportedLocale, localeCodes} from './locale.model';

export const DEFAULT_LOCALE: LocaleCode = 'en';

export const SUPPORTED_LOCALES: readonly SupportedLocale[] = [
  {code: 'en', label: 'English', nativeLabel: 'English', direction: 'ltr', flag: '🇬🇧'},
  {code: 'fr', label: 'French', nativeLabel: 'Français', direction: 'ltr', flag: '🇫🇷'},
  {code: 'de', label: 'German', nativeLabel: 'Deutsch', direction: 'ltr', flag: '🇩🇪'},
  {code: 'es', label: 'Spanish', nativeLabel: 'Español', direction: 'ltr', flag: '🇪🇸'},
  {code: 'pt-BR', label: 'Portuguese', nativeLabel: 'Português (Brasil)', direction: 'ltr', flag: '🇧🇷'},
  {code: 'it', label: 'Italian', nativeLabel: 'Italiano', direction: 'ltr', flag: '🇮🇹'},
  {code: 'tr', label: 'Turkish', nativeLabel: 'Türkçe', direction: 'ltr', flag: '🇹🇷'},
  {code: 'ru', label: 'Russian', nativeLabel: 'Русский', direction: 'ltr', flag: '🇷🇺'},
  {code: 'ar', label: 'Arabic', nativeLabel: 'العربية', direction: 'rtl', flag: '🇸🇦'},
  {code: 'fa', label: 'Persian', nativeLabel: 'فارسی', direction: 'rtl', flag: '🇮🇷'},
  {code: 'hi', label: 'Hindi', nativeLabel: 'हिन्दी', direction: 'ltr', flag: '🇮🇳'},
  {code: 'zh-Hans', label: 'Chinese', nativeLabel: '简体中文', direction: 'ltr', flag: '🇨🇳'},
  {code: 'ja', label: 'Japanese', nativeLabel: '日本語', direction: 'ltr', flag: '🇯🇵'},
  {code: 'ko', label: 'Korean', nativeLabel: '한국어', direction: 'ltr', flag: '🇰🇷'},
] as const;

const localeSet = new Set<string>(localeCodes);

export function isLocaleCode(value: unknown): value is LocaleCode {
  return typeof value === 'string' && localeSet.has(value);
}

export function normalizeLocale(value: unknown): LocaleCode {
  return isLocaleCode(value) ? value : DEFAULT_LOCALE;
}

export function localeFromLegacy(lang: unknown): LocaleCode {
  return Number(lang) === 1 ? 'ru' : DEFAULT_LOCALE;
}

export function localeFromBrowser(value: string): LocaleCode | null {
  if (isLocaleCode(value)) {
    return value;
  }
  const normalized = value.toLowerCase();
  if (normalized === 'pt' || normalized.startsWith('pt-')) {
    return 'pt-BR';
  }
  if (normalized === 'zh' || normalized === 'zh-cn' || normalized === 'zh-sg' || normalized.startsWith('zh-hans')) {
    return 'zh-Hans';
  }
  const base = normalized.split('-')[0];
  return isLocaleCode(base) ? base : null;
}
