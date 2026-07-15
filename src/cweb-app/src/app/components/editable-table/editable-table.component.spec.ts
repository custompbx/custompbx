import {ComponentFixture, TestBed} from '@angular/core/testing';
import {EditableTableCellEvent, EditableTableComponent} from './editable-table.component';

describe('EditableTableComponent', () => {
  let fixture: ComponentFixture<EditableTableComponent>;
  let component: EditableTableComponent;
  let rows: Record<string, any>[];

  beforeEach(async () => {
    await TestBed.configureTestingModule({imports: [EditableTableComponent]}).compileComponents();
    fixture = TestBed.createComponent(EditableTableComponent);
    component = fixture.componentInstance;
    rows = [{id: 7, name: 'agent-7', contact: ''}];
    fixture.componentRef.setInput('rows', rows);
    fixture.componentRef.setInput('columns', ['id', 'name', 'contact']);
    fixture.componentRef.setInput('rowKey', 'id');
    fixture.componentRef.setInput('editable', true);
    fixture.componentRef.setInput('deleteColumn', 'id');
    fixture.componentRef.setInput('warningColumn', 'contact');
    fixture.detectChanges();
  });

  it('makes an editable cell visibly enter edit mode', () => {
    const nameCell = fixture.nativeElement.querySelectorAll('tbody td')[1] as HTMLElement;
    nameCell.click();
    fixture.detectChanges();

    expect(nameCell.classList).toContain('cpbx-data-grid__cell--editing');
    expect(nameCell.querySelector('input')).not.toBeNull();
    expect(nameCell.textContent).toContain('Enter to save');
  });

  it('saves on Enter and restores the original value on Escape', () => {
    const saved: EditableTableCellEvent[] = [];
    component.cellSaved.subscribe(event => saved.push(event));
    const nameCell = fixture.nativeElement.querySelectorAll('tbody td')[1] as HTMLElement;

    nameCell.click();
    fixture.detectChanges();
    let input = nameCell.querySelector('input') as HTMLInputElement;
    input.value = 'renamed';
    input.dispatchEvent(new Event('input'));
    input.dispatchEvent(new KeyboardEvent('keydown', {key: 'Enter'}));
    fixture.detectChanges();

    expect(rows[0]['name']).toBe('renamed');
    expect(saved[0].value).toBe('renamed');

    nameCell.click();
    fixture.detectChanges();
    input = nameCell.querySelector('input') as HTMLInputElement;
    input.value = 'discarded';
    input.dispatchEvent(new Event('input'));
    input.dispatchEvent(new KeyboardEvent('keydown', {key: 'Escape'}));
    fixture.detectChanges();

    expect(rows[0]['name']).toBe('renamed');
    expect(nameCell.querySelector('input')).toBeNull();
  });

  it('does not save an unchanged cell when Enter is pressed', () => {
    const saved: EditableTableCellEvent[] = [];
    component.cellSaved.subscribe(event => saved.push(event));
    const nameCell = fixture.nativeElement.querySelectorAll('tbody td')[1] as HTMLElement;

    nameCell.click();
    fixture.detectChanges();
    const input = nameCell.querySelector('input') as HTMLInputElement;
    input.dispatchEvent(new KeyboardEvent('keydown', {key: 'Enter'}));
    fixture.detectChanges();

    expect(saved).toEqual([]);
    expect(nameCell.querySelector('input')).toBeNull();
  });

  it('marks a changed cell without saving it on blur', () => {
    const changed: EditableTableCellEvent[] = [];
    const saved: EditableTableCellEvent[] = [];
    component.cellChanged.subscribe(event => changed.push(event));
    component.cellSaved.subscribe(event => saved.push(event));
    const nameCell = fixture.nativeElement.querySelectorAll('tbody td')[1] as HTMLElement;

    nameCell.click();
    fixture.detectChanges();
    const input = nameCell.querySelector('input') as HTMLInputElement;
    input.value = 'pending-name';
    input.dispatchEvent(new Event('input'));
    input.dispatchEvent(new Event('blur'));
    fixture.detectChanges();

    expect(rows[0]['name']).toBe('pending-name');
    expect(changed[changed.length - 1]?.changed).toBeTrue();
    expect(saved).toEqual([]);
    expect(nameCell.querySelector('input')).toBeNull();
  });

  it('clears the changed state when the original value is restored', () => {
    const changed: EditableTableCellEvent[] = [];
    component.cellChanged.subscribe(event => changed.push(event));
    const nameCell = fixture.nativeElement.querySelectorAll('tbody td')[1] as HTMLElement;

    nameCell.click();
    fixture.detectChanges();
    const input = nameCell.querySelector('input') as HTMLInputElement;
    input.value = 'pending-name';
    input.dispatchEvent(new Event('input'));
    input.value = 'agent-7';
    input.dispatchEvent(new Event('input'));

    expect(changed[changed.length - 1]?.changed).toBeFalse();
  });

  it('saves an already dirty cell only when Enter is pressed', () => {
    const saved: EditableTableCellEvent[] = [];
    component.cellSaved.subscribe(event => saved.push(event));
    fixture.componentRef.setInput('changedCells', {'7': {name: true}});
    fixture.detectChanges();
    const nameCell = fixture.nativeElement.querySelectorAll('tbody td')[1] as HTMLElement;

    expect(nameCell.textContent).toContain('Unsaved');
    nameCell.click();
    fixture.detectChanges();
    const input = nameCell.querySelector('input') as HTMLInputElement;
    input.dispatchEvent(new KeyboardEvent('keydown', {key: 'Enter'}));
    fixture.detectChanges();

    expect(saved.length).toBe(1);
    expect(saved[0].value).toBe('agent-7');
  });

  it('keeps read-only tables non-editable', () => {
    fixture.componentRef.setInput('editable', false);
    fixture.detectChanges();
    const nameCell = fixture.nativeElement.querySelectorAll('tbody td')[1] as HTMLElement;
    nameCell.click();
    fixture.detectChanges();

    expect(nameCell.querySelector('input')).toBeNull();
  });

  it('keeps identifier and configured read-only columns non-editable', () => {
    fixture.componentRef.setInput('readonlyColumns', ['contact']);
    fixture.detectChanges();
    const cells = fixture.nativeElement.querySelectorAll('tbody td') as NodeListOf<HTMLElement>;

    cells[0].click();
    cells[2].click();
    fixture.detectChanges();

    expect(cells[0].querySelector('input')).toBeNull();
    expect(cells[2].querySelector('input')).toBeNull();
    expect(cells[0].classList).not.toContain('cpbx-data-grid__cell--editable');
    expect(cells[2].classList).not.toContain('cpbx-data-grid__cell--editable');
  });
});
