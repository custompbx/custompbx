import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { HeaderComponent } from './header.component';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [ HeaderComponent ],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('renders key action buttons with accessible labels', () => {
    const labels = Array.from<HTMLElement>(fixture.nativeElement.querySelectorAll('[aria-label]'))
      .map((element) => element.getAttribute('aria-label'));

    expect(labels).toContain('Toggle navigation');
    expect(labels).toContain('CustomPBX dashboard');
    expect(labels).toContain('Open conversations');
    expect(labels).toContain('Toggle phone');
    expect(labels).toContain('Open user menu');
  });

  it('does not require a loaded user before first render', () => {
    component.user = null;

    expect(() => fixture.detectChanges()).not.toThrow();
  });
});
