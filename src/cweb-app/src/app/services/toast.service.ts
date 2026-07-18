import {inject, Injectable, signal} from '@angular/core';
import {TranslocoService} from '@jsverse/transloco';

export type ToastTone = 'info' | 'success' | 'warning' | 'error';

export interface ToastMessage {
  id: number;
  message: string;
  tone: ToastTone;
}

export interface ToastOptions {
  duration?: number;
  panelClass?: string | string[];
  tone?: ToastTone;
}

@Injectable({providedIn: 'root'})
export class ToastService {
  readonly messages = signal<readonly ToastMessage[]>([]);

  private readonly i18n = inject(TranslocoService);
  private nextId = 0;

  open(message: string, _action?: string | null, options: ToastOptions = {}): void {
    const id = ++this.nextId;
    const tone = options.tone ?? this.inferTone(message, options.panelClass);

    this.messages.update(messages => [...messages, {id, message, tone}]);
    window.setTimeout(() => this.dismiss(id), options.duration ?? 7500);
  }

  success(message: string, duration = 6000): void {
    this.open(message, null, {duration, tone: 'success'});
  }

  error(message: string, duration = 9000): void {
    this.open(message, null, {duration, tone: 'error'});
  }

  copied(duration = 5000): void {
    this.success(this.i18n.translate('common.copied'), duration);
  }

  backendError(message: string, duration = 9000): void {
    this.error(this.i18n.translate('feedback.errorWithDetail', {message}), duration);
  }

  dismiss(id: number): void {
    this.messages.update(messages => messages.filter(message => message.id !== id));
  }

  private inferTone(message: string, panelClass?: string | string[]): ToastTone {
    const classes = Array.isArray(panelClass) ? panelClass.join(' ') : panelClass ?? '';
    if (/error|danger|warn/i.test(classes) || /^error\b/i.test(message)) return 'error';
    if (/success/i.test(classes) || /^(copied|saved|created|updated)\b/i.test(message)) return 'success';
    return 'info';
  }
}
