import {TestBed} from '@angular/core/testing';
import {ConfirmationService} from './confirmation.service';

describe('ConfirmationService', () => {
  let service: ConfirmationService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ConfirmationService);
  });

  it('resolves the active confirmation and clears it', () => {
    let result: boolean | undefined;
    service.open({data: {action: 'delete'}}).afterDismissed().subscribe(value => result = value);

    service.close(true);

    expect(result).toBeTrue();
    expect(service.request()).toBeNull();
  });

  it('cancels an existing request when a new one opens', () => {
    let firstResult: boolean | undefined;
    service.open().afterDismissed().subscribe(value => firstResult = value);

    service.open({data: {title: 'Second'}});

    expect(firstResult).toBeFalse();
    expect(service.request()?.data.title).toBe('Second');
  });
});
