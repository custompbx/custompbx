import {ComponentFixture, TestBed} from '@angular/core/testing';
import {provideTransloco, TranslocoLoader} from '@jsverse/transloco';
import {of} from 'rxjs';

import {DialplanApplicationRowComponent} from './dialplan-application-row.component';

class TestLoader implements TranslocoLoader {
  getTranslation() {
    return of({common: {moveItem: 'Move item', actions: {save: 'Save', cancel: 'Cancel', delete: 'Delete', disable: 'Disable', enable: 'Enable'}}});
  }
}

describe('DialplanApplicationRowComponent', () => {
  let fixture: ComponentFixture<DialplanApplicationRowComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DialplanApplicationRowComponent],
      providers: [provideTransloco({config: {availableLangs: ['en'], defaultLang: 'en'}, loader: TestLoader})],
    }).compileComponents();

    fixture = TestBed.createComponent(DialplanApplicationRowComponent);
    fixture.componentRef.setInput('item', {id: 1, position: 1, application: 'answer', data: '', inline: false, enabled: true});
    fixture.componentRef.setInput('rowKey', 'action1');
    fixture.detectChanges();
  });

  it('does not save an unchanged persisted row', () => {
    const save = fixture.nativeElement.querySelector('.cpbx-button--primary') as HTMLButtonElement;
    expect(save.disabled).toBeTrue();
  });

  it('uses cancel and no toggle for a draft row', () => {
    fixture.componentRef.setInput('draft', true);
    fixture.detectChanges();
    expect(fixture.nativeElement.textContent).toContain('Cancel');
    expect(fixture.nativeElement.textContent).not.toContain('Disable');
  });

  it('saves a changed valid row on Enter', () => {
    const emitted: unknown[] = [];
    fixture.componentInstance.save.subscribe(value => emitted.push(value));
    const inputs = fixture.nativeElement.querySelectorAll('input') as NodeListOf<HTMLInputElement>;
    inputs[1].value = 'sofia/internal/1000';
    inputs[1].dispatchEvent(new Event('input'));
    fixture.detectChanges();
    inputs[1].dispatchEvent(new KeyboardEvent('keydown', {key: 'Enter', bubbles: true, cancelable: true}));

    expect(emitted).toEqual([fixture.componentInstance.item()]);
  });

  it('does not save an unchanged persisted row on Enter', () => {
    fixture.componentRef.setInput('item', {id: 1, position: 1, application: 'answer', data: 'ok', inline: false, enabled: true});
    fixture.detectChanges();
    const emitted: unknown[] = [];
    fixture.componentInstance.save.subscribe(value => emitted.push(value));
    const input = fixture.nativeElement.querySelector('input') as HTMLInputElement;
    input.dispatchEvent(new KeyboardEvent('keydown', {key: 'Enter', bubbles: true, cancelable: true}));

    expect(emitted).toEqual([]);
  });

  it('emits toggle and remove actions', () => {
    const toggled: unknown[] = [];
    let removed = 0;
    fixture.componentInstance.toggle.subscribe(value => toggled.push(value));
    fixture.componentInstance.remove.subscribe(() => removed++);
    const buttons = fixture.nativeElement.querySelectorAll('button') as NodeListOf<HTMLButtonElement>;
    buttons[2].click();
    buttons[3].click();

    expect(toggled).toEqual([fixture.componentInstance.item()]);
    expect(removed).toBe(1);
  });
});
