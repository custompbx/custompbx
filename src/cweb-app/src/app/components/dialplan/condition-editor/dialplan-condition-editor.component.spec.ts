import {ComponentFixture, TestBed} from '@angular/core/testing';
import {provideTransloco, TranslocoLoader} from '@jsverse/transloco';
import {of} from 'rxjs';

import {DialplanConditionEditorComponent} from './dialplan-condition-editor.component';
import {Icondition} from '../../../store/dialplan/dialplan.reducers';

class TestLoader implements TranslocoLoader {
  getTranslation() {
    return of({
      common: {moveItem: 'Move item', actions: {save: 'Save', cancel: 'Cancel', delete: 'Delete', disable: 'Disable', enable: 'Enable', add: 'Add'}},
      dialplan: {conditionFields: 'Condition fields', field: 'Field', expression: 'Expression', break: 'Break', defaultOnFalse: 'Default (on-false)', regexes: 'Regexes', actions: 'Actions', antiactions: 'Antiactions'},
    });
  }
}

const condition = (): Icondition => ({
  id: 1,
  position: 1,
  enabled: true,
  field: 'destination_number',
  expression: '^1000$',
  break: undefined,
  regexes: [],
  actions: [],
  antiactions: [],
  newRegexes: [],
  newActions: [],
  newAntiactions: [],
  new: [],
});

describe('DialplanConditionEditorComponent', () => {
  let fixture: ComponentFixture<DialplanConditionEditorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DialplanConditionEditorComponent],
      providers: [provideTransloco({config: {availableLangs: ['en'], defaultLang: 'en'}, loader: TestLoader})],
    }).compileComponents();

    fixture = TestBed.createComponent(DialplanConditionEditorComponent);
    fixture.componentRef.setInput('condition', condition());
    fixture.detectChanges();
  });

  it('keeps the default break value empty for serialization', () => {
    const select = fixture.nativeElement.querySelector('.condition-break-field select') as HTMLSelectElement;
    expect(select.value).toBe('');
    expect(fixture.componentInstance.condition().break).toBeUndefined();
  });

  it('does not enable save before a condition changes', () => {
    const save = fixture.nativeElement.querySelector('.dialplan-editor-row .cpbx-button--primary') as HTMLButtonElement;
    expect(save.disabled).toBeTrue();
  });

  it('emits condition changes only after the form becomes dirty', () => {
    const emitted: Icondition[] = [];
    fixture.componentInstance.conditionSave.subscribe(value => emitted.push(value));
    const field = fixture.nativeElement.querySelector('.dialplan-editor-row input.cpbx-input') as HTMLInputElement;
    field.value = 'destination_number_v2';
    field.dispatchEvent(new Event('input'));
    fixture.detectChanges();

    const save = fixture.nativeElement.querySelector('.dialplan-editor-row .cpbx-button--primary') as HTMLButtonElement;
    expect(save.disabled).toBeFalse();
    save.click();
    expect(emitted.length).toBe(1);
  });

  it('emits enable/disable and delete independently from editing', () => {
    const toggled: Icondition[] = [];
    let removed = 0;
    fixture.componentInstance.conditionToggle.subscribe(value => toggled.push(value));
    fixture.componentInstance.conditionRemove.subscribe(() => removed++);
    const toggle = fixture.nativeElement.querySelector('.dialplan-editor-row .switch-button') as HTMLButtonElement;
    const remove = fixture.nativeElement.querySelector('.dialplan-editor-row .cpbx-button--danger-outline') as HTMLButtonElement;

    toggle.click();
    remove.click();

    expect(toggled).toEqual([fixture.componentInstance.condition()]);
    expect(removed).toBe(1);
  });

  it('preserves explicit break values for serialization', async () => {
    const explicitFixture = TestBed.createComponent(DialplanConditionEditorComponent);
    explicitFixture.componentRef.setInput('condition', {...condition(), break: 'always'});
    explicitFixture.detectChanges();
    await explicitFixture.whenStable();
    explicitFixture.detectChanges();
    const select = explicitFixture.nativeElement.querySelector('.condition-break-field select') as HTMLSelectElement;
    expect(select.value).toBe('always');
    expect(explicitFixture.componentInstance.condition().break).toBe('always');
    explicitFixture.destroy();
  });

  it('emits a draft action without coercing its string values', () => {
    const draft = {id: 0, position: 0, application: 'set', data: 'x=1', inline: false, enabled: true};
    fixture.componentRef.setInput('condition', {...condition(), newActions: [draft]});
    const emitted: unknown[] = [];
    fixture.componentInstance.actionAdd.subscribe(value => emitted.push(value));
    fixture.detectChanges();

    fixture.componentInstance.actionAdd.emit({index: 0, item: draft});
    expect(emitted).toEqual([{index: 0, item: draft}]);
    expect(typeof draft.data).toBe('string');
  });
});
