import {DestroyRef, Injectable, PLATFORM_ID, inject, signal} from '@angular/core';
import {isPlatformBrowser} from '@angular/common';

const COMPACT_NAVIGATION_QUERY = '(max-width: 1023px)';

@Injectable({providedIn: 'root'})
export class ViewportService {
  private readonly destroyRef = inject(DestroyRef);
  private readonly platformId = inject(PLATFORM_ID);
  private readonly compactNavigationState = signal(false);

  readonly compactNavigation = this.compactNavigationState.asReadonly();

  constructor() {
    if (!isPlatformBrowser(this.platformId)) return;

    const mediaQuery = window.matchMedia(COMPACT_NAVIGATION_QUERY);
    const update = (): void => this.compactNavigationState.set(mediaQuery.matches);

    update();
    mediaQuery.addEventListener('change', update);
    this.destroyRef.onDestroy(() => mediaQuery.removeEventListener('change', update));
  }
}
