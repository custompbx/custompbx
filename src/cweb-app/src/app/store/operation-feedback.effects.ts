import {inject, Injectable} from '@angular/core';
import {Actions, createEffect} from '@ngrx/effects';
import {filter, tap} from 'rxjs/operators';
import {ToastService} from '../services/toast.service';

@Injectable()
export class OperationFeedbackEffects {
  private readonly actions$ = inject(Actions);
  private readonly toast = inject(ToastService);

  readonly completed$ = createEffect(() => this.actions$.pipe(
    filter((action: any) => /^Store(?:Add|Update|Delete|Switch|Paste|Import|Truncate|Rename|Save)/.test(action.type ?? '')),
    filter((action: any) => !this.hasError(action.payload?.response)),
    tap((action: any) => this.toast.success(this.messageFor(action.type)))
  ), {dispatch: false});

  private hasError(response: any): boolean {
    return Boolean(response?.error || response?.Error || response?.errorMessage || response?.errors?.length);
  }

  private messageFor(type: string): string {
    if (/Add/.test(type)) return 'Item added successfully.';
    if (/Delete/.test(type)) return 'Item removed successfully.';
    if (/Switch/.test(type)) return 'Status updated successfully.';
    if (/Import/.test(type)) return 'Import completed successfully.';
    if (/Paste/.test(type)) return 'Items pasted successfully.';
    if (/Rename/.test(type)) return 'Item renamed successfully.';
    return 'Changes saved successfully.';
  }
}
