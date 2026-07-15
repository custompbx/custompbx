import {inject, Injectable} from '@angular/core';
import {Actions, createEffect} from '@ngrx/effects';
import {filter, tap} from 'rxjs/operators';
import {ToastService} from '../services/toast.service';

const persistedOperations = /Add|Create|Update|Delete|Remove|Del|Switch|Paste|Import|Truncate|Rename|Save|Move|Load|Unload|Reload|Autoload|Clear|From scratch/i;

export function hasOperationError(response: any): boolean {
  return Boolean(response?.error || response?.Error || response?.errorMessage || response?.errors?.length);
}

export function isConfirmedOperation(action: any): boolean {
  const type = action?.type ?? '';
  const hasResponse = Object.prototype.hasOwnProperty.call(action?.payload ?? {}, 'response');

  return type.includes('Store')
    && persistedOperations.test(type)
    && !/Reset|Delete\s*new|DeleteNew|DelNew|DropNew/i.test(type)
    && hasResponse
    && !hasOperationError(action.payload.response);
}

export function operationMessage(type: string): string {
  if (/Add|Create/.test(type)) return 'Item added successfully.';
  if (/Delete|Remove|Del/.test(type)) return 'Item removed successfully.';
  if (/Switch/.test(type)) return 'Status updated successfully.';
  if (/Import/.test(type)) return 'Import completed successfully.';
  if (/Paste/.test(type)) return 'Items pasted successfully.';
  if (/Rename/.test(type)) return 'Item renamed successfully.';
  if (/Move/.test(type)) return 'Order updated successfully.';
  if (/Reload/.test(type)) return 'Reload completed successfully.';
  if (/Unload/.test(type)) return 'Module unloaded successfully.';
  if (/Load/.test(type)) return 'Module loaded successfully.';
  if (/Autoload/.test(type)) return 'Autoload setting updated successfully.';
  if (/Clear/.test(type)) return 'Item cleared successfully.';
  if (/Truncate/.test(type)) return 'Configuration cleared successfully.';
  if (/From scratch/i.test(type)) return 'Configuration created successfully.';
  return 'Changes saved successfully.';
}

@Injectable()
export class OperationFeedbackEffects {
  private readonly actions$ = inject(Actions);
  private readonly toast = inject(ToastService);

  readonly completed$ = createEffect(() => this.actions$.pipe(
    filter(isConfirmedOperation),
    tap((action: any) => this.toast.success(operationMessage(action.type)))
  ), {dispatch: false});
}
