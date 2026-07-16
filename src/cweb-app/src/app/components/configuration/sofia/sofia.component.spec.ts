import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { signal } from '@angular/core';

import { SofiaComponent } from './sofia.component';

describe('SofiaComponent', () => {
  let component: SofiaComponent;
  let fixture: ComponentFixture<SofiaComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [ SofiaComponent ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SofiaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('groups settings and profiles into labeled cards', () => {
    fixture.destroy();
    fixture = TestBed.createComponent(SofiaComponent);
    component = fixture.componentInstance;
    (component as any).list = signal({
      exists: true,
      global_settings: {},
      profiles: {},
    });

    fixture.detectChanges();

    const cards = fixture.nativeElement.querySelectorAll('.cpbx-tab-panel > .cpbx-card');
    expect(cards.length).toBe(2);
    expect(cards[0].querySelector(':scope > .cpbx-card__header').textContent).toContain('Settings');
    expect(cards[0].querySelector('app-disclosure')).toBeTruthy();
    expect(cards[1].querySelector(':scope > .cpbx-card__header').textContent).toContain('Profiles');
  });
});
