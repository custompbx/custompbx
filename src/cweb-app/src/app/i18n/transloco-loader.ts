import {HttpClient} from '@angular/common/http';
import {inject, Injectable} from '@angular/core';
import {Translation, TranslocoLoader} from '@jsverse/transloco';
import {Observable} from 'rxjs';
import {map} from 'rxjs/operators';
import {normalizeLocale} from './locale.registry';
import {featureTranslation} from './feature-translations';

function mergeTranslations(base: Translation, additions: Translation): Translation {
  const result: Translation = {...base};
  for (const [key, value] of Object.entries(additions)) {
    const current = result[key];
    result[key] = current && value && typeof current === 'object' && typeof value === 'object'
      ? mergeTranslations(current as Translation, value as Translation)
      : value;
  }
  return result;
}

@Injectable({providedIn: 'root'})
export class CpbxTranslocoLoader implements TranslocoLoader {
  private readonly http = inject(HttpClient);

  getTranslation(lang: string): Observable<Translation> {
    const locale = normalizeLocale(lang);
    return this.http.get<Translation>(`./assets/i18n/${locale}.json`).pipe(
      map(base => mergeTranslations(base, featureTranslation(locale))),
    );
  }
}
