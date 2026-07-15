import {ChangeDetectionStrategy, Component, input, model, output} from '@angular/core';

let nextDisclosureId = 0;

@Component({
  selector: 'app-disclosure',
  standalone: true,
  changeDetection: ChangeDetectionStrategy.OnPush,
  host: {
    class: 'cpbx-disclosure-host',
  },
  template: `
    <section class="cpbx-disclosure" [class.is-open]="expanded()">
      <button
        type="button"
        class="cpbx-disclosure__trigger"
        [attr.aria-controls]="contentId"
        [attr.aria-expanded]="expanded()"
        [disabled]="disabled()"
        (click)="toggle()">
        <span class="cpbx-disclosure__title">
          <ng-content select="[disclosureTitle]" />
        </span>
        <span class="cpbx-disclosure__description">
          <ng-content select="[disclosureDescription]" />
        </span>
        <span class="cpbx-disclosure__chevron" aria-hidden="true"></span>
      </button>

      @if (expanded()) {
        <div class="cpbx-disclosure__content" [id]="contentId">
          <ng-content />
        </div>
      }
    </section>
  `,
  styles: `
    :host {
      display: block;
      min-width: 0;
      width: 100%;
    }

    .cpbx-disclosure {
      background: var(--cpbx-surface, #fff);
      border: 1px solid var(--cpbx-border, #dbe3ef);
      min-width: 0;
      overflow: clip;
      width: 100%;
    }

    .cpbx-disclosure__trigger {
      align-items: center;
      background: var(--cpbx-surface-subtle, #f8fafc);
      border: 0;
      color: var(--cpbx-text, #0f172a);
      cursor: pointer;
      display: grid;
      font: inherit;
      grid-template-columns: minmax(10rem, 1fr) minmax(10rem, 1fr) 1.5rem;
      min-height: 4.25rem;
      padding: .875rem 1.5rem;
      text-align: left;
      transition: background-color 150ms ease, color 150ms ease;
      width: 100%;
    }

    .cpbx-disclosure__trigger:hover {
      background: var(--cpbx-hover, #f1f5f9);
    }

    .cpbx-disclosure__trigger:focus-visible {
      outline: 3px solid color-mix(in srgb, var(--cpbx-primary, #4355c5) 25%, transparent);
      outline-offset: -3px;
    }

    .cpbx-disclosure__trigger:disabled {
      cursor: not-allowed;
      opacity: .6;
    }

    .cpbx-disclosure__title {
      font-weight: 500;
      min-width: 0;
    }

    .cpbx-disclosure__description {
      color: var(--cpbx-text-muted, #64748b);
      min-width: 0;
    }

    .cpbx-disclosure__chevron {
      border-bottom: 2px solid currentColor;
      border-right: 2px solid currentColor;
      height: .55rem;
      justify-self: end;
      transform: rotate(45deg) translate(-.1rem, .1rem);
      transition: transform 150ms ease;
      width: .55rem;
    }

    .is-open .cpbx-disclosure__chevron {
      transform: rotate(225deg) translate(-.05rem, -.05rem);
    }

    .cpbx-disclosure__content {
      border-top: 1px solid var(--cpbx-border, #dbe3ef);
      min-width: 100%;
      padding: 1.5rem;
    }

    @media (max-width: 720px) {
      .cpbx-disclosure__trigger {
        grid-template-columns: minmax(0, 1fr) 1.5rem;
        padding-inline: 1rem;
      }

      .cpbx-disclosure__description {
        display: none;
      }

      .cpbx-disclosure__content {
        padding: 1rem;
      }
    }
  `,
})
export class DisclosureComponent {
  readonly expanded = model(false);
  readonly disabled = input(false);
  readonly opened = output<void>();
  readonly closed = output<void>();
  readonly contentId = `cpbx-disclosure-content-${nextDisclosureId++}`;

  toggle(): void {
    if (this.disabled()) {
      return;
    }

    const next = !this.expanded();
    this.expanded.set(next);
    if (next) {
      this.opened.emit();
    } else {
      this.closed.emit();
    }
  }
}
