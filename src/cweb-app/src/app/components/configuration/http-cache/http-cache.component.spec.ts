import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { HttpCacheComponent } from './http-cache.component';

describe('SofiaComponent', () => {
  let component: HttpCacheComponent;
  let fixture: ComponentFixture<HttpCacheComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [ HttpCacheComponent ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HttpCacheComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
