import {ChangeDetectionStrategy, Component, HostListener, inject} from '@angular/core';
import {ConfirmationService} from '../../services/confirmation.service';

@Component({
  selector: 'app-confirmation-dialog',
  standalone: true,
  templateUrl: './confirmation-dialog.component.html',
  styleUrl: './confirmation-dialog.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ConfirmationDialogComponent {
  readonly confirmation = inject(ConfirmationService);

  @HostListener('document:keydown.escape')
  closeOnEscape(): void {
    this.confirmation.close(false);
  }

  title(): string {
    const data = this.confirmation.request()?.data;
    if (data?.title) return data.title;
    if (data?.action === 'delete') return 'Confirm deletion';
    if (data?.action === 'rename') return 'Confirm rename';
    return 'Confirm action';
  }

  message(): string {
    const data = this.confirmation.request()?.data;
    if (data?.message) return data.message;
    if (data?.action === 'delete' && data.case1Text) return data.case1Text;
    if (data?.action === 'rename' && data.case2Text) return data.case2Text;
    return 'This action may change the current configuration.';
  }
}
