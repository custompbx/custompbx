import {ChangeDetectionStrategy, Component, input, output, signal} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {AppAutoFocusDirective} from '../../directives/auto-focus.directive';
import {IconComponent} from '../icon/icon.component';

export interface EditableTableCellEvent {
  row: Record<string, any>;
  column: string;
  value: any;
  changed: boolean;
}

@Component({
  selector: 'app-editable-table',
  standalone: true,
  imports: [FormsModule, AppAutoFocusDirective, IconComponent],
  templateUrl: './editable-table.component.html',
  styleUrl: './editable-table.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class EditableTableComponent {
  readonly rows = input<Record<string, any>[]>([]);
  readonly columns = input<string[]>([]);
  readonly rowKey = input.required<string>();
  readonly editable = input(false);
  readonly readonlyColumns = input<string[]>([]);
  readonly deleteColumn = input<string | null>(null);
  readonly warningColumn = input<string | null>(null);
  readonly changedCells = input<Record<string, Record<string, boolean>> | Record<number, Record<string, boolean>>>({});

  readonly columnSelected = output<string>();
  readonly cellChanged = output<EditableTableCellEvent>();
  readonly cellSaved = output<EditableTableCellEvent>();
  readonly rowDeleted = output<Record<string, any>>();

  readonly editing = signal<{key: string; column: string} | null>(null);
  private originalValue: any;

  rowIdentity(row: Record<string, any>): string {
    return String(row[this.rowKey()]);
  }

  isEditing(row: Record<string, any>, column: string): boolean {
    const editing = this.editing();
    return editing?.key === this.rowIdentity(row) && editing.column === column;
  }

  isChanged(row: Record<string, any>, column: string): boolean {
    const changed = this.changedCells() as Record<string, Record<string, boolean>>;
    return !!changed[this.rowIdentity(row)]?.[column];
  }

  isWarning(row: Record<string, any>, column: string): boolean {
    return this.warningColumn() === column && (row[column] === '' || row[column] == null);
  }

  isCellEditable(column: string): boolean {
    return this.editable()
      && this.deleteColumn() !== column
      && !this.readonlyColumns().includes(column);
  }

  beginEdit(row: Record<string, any>, column: string): void {
    if (!this.isCellEditable(column) || this.isEditing(row, column)) {
      return;
    }
    this.originalValue = row[column];
    this.editing.set({key: this.rowIdentity(row), column});
  }

  valueChanged(row: Record<string, any>, column: string, value: any): void {
    row[column] = value;
    this.cellChanged.emit({
      row,
      column,
      value,
      changed: !this.valuesEqual(value, this.originalValue),
    });
  }

  save(row: Record<string, any>, column: string): void {
    if (!this.isEditing(row, column)) {
      return;
    }
    const changed = this.isChanged(row, column) || !this.valuesEqual(row[column], this.originalValue);
    this.editing.set(null);
    if (changed) {
      this.cellSaved.emit({row, column, value: row[column], changed: true});
    }
  }

  finishEdit(row: Record<string, any>, column: string): void {
    if (this.isEditing(row, column)) {
      this.editing.set(null);
    }
  }

  cancel(row: Record<string, any>, column: string): void {
    if (!this.isEditing(row, column)) {
      return;
    }
    row[column] = this.originalValue;
    this.editing.set(null);
  }

  deleteRow(event: Event, row: Record<string, any>): void {
    event.stopPropagation();
    this.rowDeleted.emit(row);
  }

  private valuesEqual(left: any, right: any): boolean {
    return String(left ?? '') === String(right ?? '');
  }
}
