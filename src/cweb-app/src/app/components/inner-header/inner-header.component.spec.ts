import {ComponentFixture, TestBed} from '@angular/core/testing';
import {InnerHeaderComponent} from './inner-header.component';

describe('InnerHeaderComponent', () => {
  let fixture: ComponentFixture<InnerHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({imports: [InnerHeaderComponent]}).compileComponents();
    fixture = TestBed.createComponent(InnerHeaderComponent);
  });

  it('renders title, subtitle, status, and loading state', () => {
    Object.assign(fixture.componentInstance, {name: 'Dashboard', subtitle: 'System health', status: 'Online', loadCounter: 1});
    fixture.detectChanges();
    expect(fixture.nativeElement.querySelector('h1').textContent).toContain('Dashboard');
    expect(fixture.nativeElement.querySelector('p').textContent).toContain('System health');
    expect(fixture.nativeElement.querySelector('.status-badge').textContent).toContain('Online');
    expect(fixture.nativeElement.querySelector('.cpbx-loading-bar')).toBeTruthy();
    expect(fixture.nativeElement.querySelector('.cpbx-loading-bar').getAttribute('role')).toBe('progressbar');
  });
});
