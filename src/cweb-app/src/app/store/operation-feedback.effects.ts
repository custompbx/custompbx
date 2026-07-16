import {inject, Injectable} from '@angular/core';
import {Actions, createEffect} from '@ngrx/effects';
import {filter, tap} from 'rxjs/operators';
import {ToastService} from '../services/toast.service';
import {
  operationKindFromType,
  operationMetadata,
  operationSuccessMessage,
  OperationFeedbackAction,
} from '../services/operation-feedback';

export function hasOperationError(response: any): boolean {
  return Boolean(response?.error || response?.Error || response?.errorMessage || response?.errors?.length);
}

export function isConfirmedOperation(action: OperationFeedbackAction & {payload?: any}): boolean {
  const hasResponse = Object.prototype.hasOwnProperty.call(action?.payload ?? {}, 'response');

  return hasResponse
    && Boolean(operationMetadata(action))
    && !hasOperationError(action.payload.response);
}

export function operationMessage(type: string): string {
  const kind = operationKindFromType(type);
  return kind ? operationSuccessMessage(kind) : 'Changes saved successfully.';
}

@Injectable()
export class OperationFeedbackEffects {
  private readonly actions$ = inject(Actions);
  private readonly toast = inject(ToastService);

  readonly completed$ = createEffect(() => this.actions$.pipe(
    filter(isConfirmedOperation),
    tap((action: OperationFeedbackAction) => {
      const metadata = operationMetadata(action);
      if (metadata) this.toast.success(operationSuccessMessage(metadata.kind));
    })
  ), {dispatch: false});
}
