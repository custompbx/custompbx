import {Injectable, signal} from '@angular/core';
import {Observable, Subject} from 'rxjs';

export interface ConfirmationData {
  action?: 'delete' | 'rename' | string;
  title?: string;
  message?: string;
  case1Text?: string | null;
  case2Text?: string | null;
  confirmText?: string;
  cancelText?: string;
  [key: string]: unknown;
}

export interface ConfirmationConfig {
  data?: ConfirmationData;
}

interface ConfirmationRequest {
  data: ConfirmationData;
  result: Subject<boolean>;
}

export interface ConfirmationRef {
  afterDismissed(): Observable<boolean>;
}

@Injectable({providedIn: 'root'})
export class ConfirmationService {
  readonly request = signal<ConfirmationRequest | null>(null);

  open(config: ConfirmationConfig = {}): ConfirmationRef {
    this.close(false);
    const result = new Subject<boolean>();
    this.request.set({data: config.data ?? {}, result});
    return {afterDismissed: () => result.asObservable()};
  }

  close(confirmed: boolean): void {
    const current = this.request();
    if (!current) return;
    this.request.set(null);
    current.result.next(confirmed);
    current.result.complete();
  }
}
