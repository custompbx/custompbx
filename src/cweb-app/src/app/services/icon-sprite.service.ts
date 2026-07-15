import {DOCUMENT} from '@angular/common';
import {inject, Injectable} from '@angular/core';

@Injectable({providedIn: 'root'})
export class IconSpriteService {
  private readonly document = inject(DOCUMENT);
  private loading?: Promise<void>;

  load(): void {
    if (this.document.getElementById('cpbx-icon-sprite') || this.loading) return;
    const url = new URL('assets/icons.svg', this.document.baseURI).toString();
    this.loading = fetch(url, {cache: 'force-cache'})
      .then(response => response.ok ? response.text() : Promise.reject(response.status))
      .then(source => {
        if (this.document.getElementById('cpbx-icon-sprite')) return;
        const parsed = new DOMParser().parseFromString(source, 'image/svg+xml').documentElement;
        parsed.id = 'cpbx-icon-sprite';
        parsed.setAttribute('aria-hidden', 'true');
        parsed.setAttribute('focusable', 'false');
        parsed.setAttribute('style', 'display:none');
        this.document.body.prepend(this.document.importNode(parsed, true));
      })
      .catch(() => undefined);
  }
}
