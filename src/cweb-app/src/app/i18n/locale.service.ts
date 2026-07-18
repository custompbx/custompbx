import {computed, effect, inject, Injectable, signal} from '@angular/core';
import {TranslocoService} from '@jsverse/transloco';
import {UserService} from '../services/user.service';
import {Store} from '@ngrx/store';
import {AppState} from '../store/app.states';
import {UpdateWebUserLocale} from '../store/settings/settings.actions';
import {SettingsActionTypes} from '../store/settings/settings.actions';
import {Actions, ofType} from '@ngrx/effects';
import {LocaleCode} from './locale.model';
import {
  DEFAULT_LOCALE, isLocaleCode, LANGUAGE_MENU_LOCALES, normalizeLocale, SUPPORTED_LOCALES
} from './locale.registry';

@Injectable({providedIn: 'root'})
export class LocaleService {
  private readonly transloco = inject(TranslocoService);
  private readonly userService = inject(UserService);
  private readonly store = inject(Store<AppState>);
  private readonly actions = inject(Actions, {optional: true});
  private readonly locale = signal<LocaleCode>(DEFAULT_LOCALE);
  private pendingPrevious: LocaleCode | null = null;

  readonly activeLocale = this.locale.asReadonly();
  readonly supportedLocales = LANGUAGE_MENU_LOCALES;
  readonly direction = computed(() =>
    SUPPORTED_LOCALES.find(item => item.code === this.locale())?.direction ?? 'ltr'
  );
  readonly isRtl = computed(() => this.direction() === 'rtl');

  constructor() {
    this.apply(this.locale());
    effect(() => {
      const user = this.userService.userSignal();
      if (!user) {
        return;
      }
      const locale = isLocaleCode(user.locale) ? user.locale : DEFAULT_LOCALE;
      this.apply(locale);
    });
    this.actions?.pipe(ofType(SettingsActionTypes.STORE_UPDATE_WEB_USER_LOCALE)).subscribe(() => {
        this.pendingPrevious = null;
      });
    this.actions?.pipe(ofType(SettingsActionTypes.StoreGotWebError)).subscribe(() => {
        if (this.pendingPrevious) {
          this.apply(this.pendingPrevious);
          this.pendingPrevious = null;
        }
      });
  }

  setLocale(value: unknown, persist = true): void {
    const locale = normalizeLocale(value);
    const previous = this.locale();
    this.apply(locale);
    const user = this.userService.userSignal();
    if (persist && user?.id) {
      this.pendingPrevious = previous;
      this.store.dispatch(new UpdateWebUserLocale({id: user.id, value: locale}));
    }
  }

  formatDate(value: Date | number | string, options?: Intl.DateTimeFormatOptions): string {
    return new Intl.DateTimeFormat(this.locale(), options).format(new Date(value));
  }

  formatNumber(value: number, options?: Intl.NumberFormatOptions): string {
    return new Intl.NumberFormat(this.locale(), options).format(value);
  }

  currentLocale() {
    return SUPPORTED_LOCALES.find(item => item.code === this.locale()) ?? SUPPORTED_LOCALES[0];
  }

  private apply(locale: LocaleCode): void {
    this.locale.set(locale);
    this.transloco.setActiveLang(locale);
    document.documentElement.lang = locale;
    document.documentElement.dir = this.direction();
  }
}
