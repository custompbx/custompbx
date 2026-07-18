import {fakeAsync, TestBed, tick} from '@angular/core/testing';
import {ToastService} from './toast.service';
import {customPbxTestProviders} from '../testing/test-providers';
import {TranslocoService} from '@jsverse/transloco';

describe('ToastService', () => {
  let service: ToastService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        ...customPbxTestProviders(),
        {
          provide: TranslocoService,
          useValue: {
            translate: (key: string, params?: {message?: string}) => {
              if (key === 'common.copied') return 'Copied!';
              if (key === 'feedback.errorWithDetail') return `Error: ${params?.message}`;
              return key;
            },
          },
        },
      ],
    });
    service = TestBed.inject(ToastService);
  });

  it('publishes and automatically dismisses a message', fakeAsync(() => {
    service.open('Saved', null, {duration: 100});
    expect(service.messages().length).toBe(1);
    expect(service.messages()[0].tone).toBe('success');

    tick(100);
    expect(service.messages()).toEqual([]);
  }));

  it('marks error messages as errors', () => {
    service.open('Error: failed');
    expect(service.messages()[0].tone).toBe('error');
  });

  it('localizes shared copied and backend error feedback', () => {
    service.copied();
    service.backendError('failed');

    expect(service.messages()[0]).toEqual(jasmine.objectContaining({
      message: 'Copied!',
      tone: 'success',
    }));
    expect(service.messages()[1]).toEqual(jasmine.objectContaining({
      message: 'Error: failed',
      tone: 'error',
    }));
  });
});
