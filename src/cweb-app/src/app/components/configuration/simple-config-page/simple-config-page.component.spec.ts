import {ComponentFixture, TestBed} from '@angular/core/testing';
import {SimpleConfigPageComponent} from './simple-config-page.component';
import {customPbxTestProviders} from '../../../testing/test-providers';

describe('SimpleConfigPageComponent', () => {
  let fixture: ComponentFixture<SimpleConfigPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SimpleConfigPageComponent],
      providers: customPbxTestProviders(),
    }).compileComponents();

    fixture = TestBed.createComponent(SimpleConfigPageComponent);
    fixture.componentRef.setInput('name', 'Alsa');
    fixture.componentRef.setInput('module', {exists: true, settings: {}});
    fixture.componentRef.setInput('dispatchersCallbacks', {});
    fixture.detectChanges();
  });

  it('renders the module name and localized shared section labels', () => {
    const text = fixture.nativeElement.textContent;
    expect(text).toContain('Alsa');
    expect(text).toContain('Settings');
    expect(text).toContain('Parameters');
  });

  it('does not render settings when the module does not exist', () => {
    fixture.componentRef.setInput('module', {exists: false});
    fixture.detectChanges();
    expect(fixture.nativeElement.querySelector('.simple-config-page__settings')).toBeNull();
  });
});
