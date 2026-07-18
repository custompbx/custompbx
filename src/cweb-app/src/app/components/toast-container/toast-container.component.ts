import {ChangeDetectionStrategy, Component, inject} from '@angular/core';
import {ToastService} from '../../services/toast.service';
import {TranslocoPipe} from '@jsverse/transloco';

@Component({
  selector: 'app-toast-container',
  standalone: true,
  imports: [TranslocoPipe],
  templateUrl: './toast-container.component.html',
  styleUrl: './toast-container.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ToastContainerComponent {
  readonly toast = inject(ToastService);
}
