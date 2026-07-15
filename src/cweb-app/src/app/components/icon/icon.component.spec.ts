import {ComponentFixture, TestBed} from '@angular/core/testing';
import {IconComponent} from './icon.component';

describe('IconComponent', () => {
  let fixture: ComponentFixture<IconComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({imports: [IconComponent]}).compileComponents();
    fixture = TestBed.createComponent(IconComponent);
    fixture.componentRef.setInput('name', 'phone');
    fixture.detectChanges();
  });

  it('points to the local SVG symbol', () => {
    expect(fixture.nativeElement.querySelector('use').getAttribute('href')).toBe('#phone');
  });

  it('hides decorative icons from assistive technology', () => {
    const svg = fixture.nativeElement.querySelector('svg');
    expect(svg.getAttribute('aria-hidden')).toBe('true');
    expect(svg.getAttribute('role')).toBeNull();
  });

  it('exposes an accessible name when supplied', () => {
    fixture.componentRef.setInput('label', 'Phone offline');
    fixture.detectChanges();
    const svg = fixture.nativeElement.querySelector('svg');
    expect(svg.getAttribute('aria-label')).toBe('Phone offline');
    expect(svg.getAttribute('role')).toBe('img');
    expect(svg.getAttribute('aria-hidden')).toBeNull();
  });
});
