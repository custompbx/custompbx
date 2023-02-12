import { TestBed } from '@angular/core/testing';

import { CookiesStorageService } from './cookies-storage.service';

describe('CookiesStorageService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: CookiesStorageService = TestBed.get(CookiesStorageService);
    expect(service).toBeTruthy();
  });
});
