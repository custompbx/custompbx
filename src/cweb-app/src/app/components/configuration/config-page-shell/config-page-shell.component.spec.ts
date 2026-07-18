import {ComponentFixture, TestBed} from '@angular/core/testing';
import {ConfigPageShellComponent} from './config-page-shell.component';
import {customPbxTestProviders} from '../../../testing/test-providers';

describe('ConfigPageShellComponent', () => {
  let fixture: ComponentFixture<ConfigPageShellComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ConfigPageShellComponent],
      providers: customPbxTestProviders(),
    }).compileComponents();

    fixture = TestBed.createComponent(ConfigPageShellComponent);
    fixture.componentRef.setInput('name', 'Sofia');
    fixture.componentRef.setInput('module', {exists: true});
    fixture.componentRef.setInput('tabs', ['List', 'Add']);
    fixture.detectChanges();
  });

  it('renders the shared page header and tabs', () => {
    expect(fixture.nativeElement.textContent).toContain('Sofia');
    expect(fixture.nativeElement.querySelectorAll('[role="tab"]').length).toBe(2);
  });

  it('hides tabs when the module does not exist', () => {
    fixture.componentRef.setInput('module', {exists: false});
    fixture.detectChanges();
    expect(fixture.nativeElement.querySelector('[role="tablist"]')).toBeNull();
  });
});
