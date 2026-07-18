import {signal} from '@angular/core';
import {TestBed} from '@angular/core/testing';
import {TranslocoService} from '@jsverse/transloco';
import {Store} from '@ngrx/store';
import {UserService} from '../services/user.service';
import {UpdateWebUserLocale} from '../store/settings/settings.actions';
import {LocaleService} from './locale.service';

describe('LocaleService', () => {
  let service: LocaleService;
  let store: jasmine.SpyObj<Store>;
  let transloco: jasmine.SpyObj<TranslocoService>;
  const userSignal = signal<any>({id: 7, locale: 'en'});

  beforeEach(() => {
    store = jasmine.createSpyObj<Store>('Store', ['dispatch']);
    transloco = jasmine.createSpyObj<TranslocoService>('TranslocoService', ['setActiveLang']);
    userSignal.set({id: 7, locale: 'en'});
    TestBed.configureTestingModule({
      providers: [
        LocaleService,
        {provide: Store, useValue: store},
        {provide: TranslocoService, useValue: transloco},
        {provide: UserService, useValue: {userSignal}},
      ],
    });
    service = TestBed.inject(LocaleService);
  });

  afterEach(() => {
    document.documentElement.lang = 'en';
    document.documentElement.dir = 'ltr';
  });

  it('applies RTL only to RTL locales', () => {
    service.setLocale('ar', false);
    expect(document.documentElement.lang).toBe('ar');
    expect(document.documentElement.dir).toBe('rtl');
    expect(service.isRtl()).toBeTrue();

    service.setLocale('de', false);
    expect(document.documentElement.lang).toBe('de');
    expect(document.documentElement.dir).toBe('ltr');
    expect(service.isRtl()).toBeFalse();
  });

  it('persists a selected locale on the current web user', () => {
    service.setLocale('ja');
    const action = store.dispatch.calls.mostRecent().args[0] as unknown as UpdateWebUserLocale;
    expect(action).toEqual(jasmine.any(UpdateWebUserLocale));
    expect(action.payload).toEqual({id: 7, value: 'ja'});
  });

  it('restores the persisted user locale when user state is received again', () => {
    userSignal.set({id: 7, locale: 'fa'});
    TestBed.tick();
    expect(service.activeLocale()).toBe('fa');
    expect(document.documentElement.dir).toBe('rtl');

    userSignal.set({id: 7, locale: 'it'});
    TestBed.tick();
    expect(service.activeLocale()).toBe('it');
    expect(document.documentElement.dir).toBe('ltr');
  });

  it('uses locale-aware date and number formatting', () => {
    service.setLocale('en', false);
    const englishNumber = service.formatNumber(1234.5);
    service.setLocale('de', false);
    const germanNumber = service.formatNumber(1234.5);

    expect(englishNumber).not.toBe(germanNumber);
    expect(service.formatDate(new Date(2026, 6, 17), {year: 'numeric', month: 'long', day: 'numeric'}))
      .toContain('2026');
  });

  it('falls back safely for unsupported stored locales', () => {
    service.setLocale('unsupported', false);
    expect(service.activeLocale()).toBe('en');
    expect(document.documentElement.dir).toBe('ltr');
  });
});
