import { TestBed } from '@angular/core/testing';

import { WsDataService } from './ws-data.service';

describe('WsDataService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: WsDataService = TestBed.get(WsDataService);
    expect(service).toBeTruthy();
  });
});
