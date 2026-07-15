import {ComponentFixture, TestBed} from '@angular/core/testing';
import {LoadingBarComponent} from './loading-bar.component';

describe('LoadingBarComponent', () => {
  let fixture: ComponentFixture<LoadingBarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({imports: [LoadingBarComponent]}).compileComponents();
    fixture = TestBed.createComponent(LoadingBarComponent);
  });

  it('exposes an accessible indeterminate progress indicator', () => {
    fixture.componentRef.setInput('label', 'Loading conversations');
    fixture.detectChanges();

    const progress = fixture.nativeElement.querySelector('[role="progressbar"]') as HTMLElement;
    expect(progress.getAttribute('aria-label')).toBe('Loading conversations');
    expect(progress.querySelector('.cpbx-progress__bar')).not.toBeNull();
  });
});
