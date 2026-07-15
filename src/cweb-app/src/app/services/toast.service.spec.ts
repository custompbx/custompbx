import {fakeAsync, TestBed, tick} from '@angular/core/testing';
import {ToastService} from './toast.service';

describe('ToastService', () => {
  let service: ToastService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
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
});
