export const localeCodes = [
  'en', 'fr', 'de', 'es', 'pt-BR', 'it', 'tr', 'ru',
  'ar', 'fa', 'hi', 'zh-Hans', 'ja', 'ko'
] as const;

export type LocaleCode = typeof localeCodes[number];
export type TextDirection = 'ltr' | 'rtl';

export interface SupportedLocale {
  code: LocaleCode;
  label: string;
  nativeLabel: string;
  direction: TextDirection;
  flag: string;
}
