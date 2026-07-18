import {ChangeDetectionStrategy, Component, input, output} from '@angular/core';
import {DragDropModule, CdkDragDrop} from '@angular/cdk/drag-drop';
import {FormsModule} from '@angular/forms';
import {TranslocoPipe} from '@jsverse/transloco';

import {CpbxSelectDirective} from '../../../directives/cpbx-select.directive';
import {Iaction, Iantiaction, Icondition, Iregex} from '../../../store/dialplan/dialplan.reducers';
import {DialplanApplicationRowComponent} from '../application-row/dialplan-application-row.component';

export interface IndexedDialplanItem<T> {
  index: number;
  item: T;
}

type ConditionTextField = 'dst' | 'hour' | 'mday' | 'minday' | 'minute' | 'mon' |
  'mweek' | 'date_time' | 'time_of_day' | 'tz_offset' | 'wday' | 'week' | 'yday' | 'year';

@Component({
  selector: 'app-dialplan-condition-editor',
  standalone: true,
  imports: [FormsModule, DragDropModule, TranslocoPipe, CpbxSelectDirective, DialplanApplicationRowComponent],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './dialplan-condition-editor.component.html',
  styleUrl: './dialplan-condition-editor.component.css',
})
export class DialplanConditionEditorComponent {
  readonly condition = input.required<Icondition>();
  readonly inlineActions = input<Record<string, boolean>>({});

  readonly conditionSave = output<Icondition>();
  readonly conditionToggle = output<Icondition>();
  readonly conditionRemove = output<void>();

  readonly regexSave = output<Iregex>();
  readonly regexToggle = output<Iregex>();
  readonly regexRemove = output<Iregex>();
  readonly regexAdd = output<IndexedDialplanItem<Iregex>>();
  readonly regexDraftAdd = output<void>();
  readonly regexDraftRemove = output<number>();

  readonly actionSave = output<Iaction>();
  readonly actionToggle = output<Iaction>();
  readonly actionRemove = output<Iaction>();
  readonly actionAdd = output<IndexedDialplanItem<Iaction>>();
  readonly actionDraftAdd = output<void>();
  readonly actionDraftRemove = output<number>();
  readonly actionDropped = output<CdkDragDrop<string[]>>();

  readonly antiactionSave = output<Iantiaction>();
  readonly antiactionToggle = output<Iantiaction>();
  readonly antiactionRemove = output<Iantiaction>();
  readonly antiactionAdd = output<IndexedDialplanItem<Iantiaction>>();
  readonly antiactionDraftAdd = output<void>();
  readonly antiactionDraftRemove = output<number>();
  readonly antiactionDropped = output<CdkDragDrop<string[]>>();

  readonly scheduleFieldPairs: ReadonlyArray<readonly [ConditionTextField, ConditionTextField]> = [
    ['dst', 'hour'],
    ['mday', 'minday'],
    ['minute', 'mon'],
    ['mweek', 'date_time'],
    ['time_of_day', 'tz_offset'],
    ['wday', 'week'],
    ['yday', 'year'],
  ];

  conditionValue(field: ConditionTextField): string {
    return this.condition()[field] ?? '';
  }

  setConditionValue(field: ConditionTextField, value: string): void {
    this.condition()[field] = value;
  }

  supportsInline(application: string): boolean {
    return !!this.inlineActions()[application];
  }
}
