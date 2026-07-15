import {TestBed} from '@angular/core/testing';
import {ViewportService} from './viewport.service';

describe('ViewportService', () => {
  let matches = false;
  let changeListener: (() => void) | undefined;

  beforeEach(() => {
    matches = false;
    changeListener = undefined;

    const mediaQuery = {
      get matches() { return matches; },
      media: '(max-width: 1023px)',
      onchange: null,
      addEventListener: (_type: string, listener: EventListenerOrEventListenerObject) => {
        changeListener = typeof listener === 'function'
          ? () => listener(new Event('change'))
          : () => listener.handleEvent(new Event('change'));
      },
      removeEventListener: () => undefined,
      addListener: () => undefined,
      removeListener: () => undefined,
      dispatchEvent: () => true,
    } as MediaQueryList;

    spyOn(window, 'matchMedia').and.returnValue(mediaQuery);
    TestBed.configureTestingModule({});
  });

  it('uses the current compact-navigation media-query value', () => {
    matches = true;
    const service = TestBed.inject(ViewportService);

    expect(service.compactNavigation()).toBeTrue();
    expect(window.matchMedia).toHaveBeenCalledWith('(max-width: 1023px)');
  });

  it('updates the signal when the media query changes', () => {
    const service = TestBed.inject(ViewportService);
    expect(service.compactNavigation()).toBeFalse();

    matches = true;
    changeListener?.();

    expect(service.compactNavigation()).toBeTrue();
  });
});
