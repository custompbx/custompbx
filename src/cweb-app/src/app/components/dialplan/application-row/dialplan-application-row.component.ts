import {ChangeDetectionStrategy, Component, input, output, signal} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {DragDropModule} from '@angular/cdk/drag-drop';
import {TranslocoPipe} from '@jsverse/transloco';

import {Iaction, Iantiaction} from '../../../store/dialplan/dialplan.reducers';
import {ResizeInputDirective} from '../../../directives/resize-input.directive';
import {CpbxSelectDirective} from '../../../directives/cpbx-select.directive';

export type DialplanApplication = Iaction | Iantiaction;

@Component({
  selector: 'app-dialplan-application-row',
  standalone: true,
  imports: [FormsModule, DragDropModule, TranslocoPipe, ResizeInputDirective, CpbxSelectDirective],
  changeDetection: ChangeDetectionStrategy.OnPush,
  templateUrl: './dialplan-application-row.component.html',
  styleUrl: './dialplan-application-row.component.css',
})
export class DialplanApplicationRowComponent {
  readonly item = input.required<DialplanApplication>();
  readonly draft = input(false);
  readonly rowKey = input.required<string>();
  readonly supportsInline = input(false);

  readonly save = output<DialplanApplication>();
  readonly toggle = output<DialplanApplication>();
  readonly remove = output<void>();
  readonly inlineDirty = signal(false);

  isAction(value: DialplanApplication): value is Iaction {
    return 'inline' in value;
  }

  saveIfReady(event: Event, valid: boolean, changed: boolean): void {
    event.preventDefault();
    if (valid && (this.draft() || changed)) {
      this.save.emit(this.item());
    }
  }
}
