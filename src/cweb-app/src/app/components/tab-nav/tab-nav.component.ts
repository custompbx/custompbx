import {ChangeDetectionStrategy, Component, input, model} from '@angular/core';

@Component({
  selector: 'app-tab-nav',
  standalone: true,
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <div class="cpbx-tab-nav" role="tablist" [attr.aria-label]="ariaLabel()">
      @for (tab of tabs(); track tab; let index = $index) {
        <button
          type="button"
          class="cpbx-tab-nav__item"
          role="tab"
          [class.is-active]="selectedIndex() === index"
          [attr.aria-selected]="selectedIndex() === index"
          [attr.tabindex]="selectedIndex() === index ? 0 : -1"
          [disabled]="disabled()[index]"
          (click)="select(index)"
          (keydown)="handleKeydown($event, index)">
          {{ tab }}
        </button>
      }
    </div>
  `,
})
export class TabNavComponent {
  readonly tabs = input.required<readonly string[]>();
  readonly disabled = input<readonly boolean[]>([]);
  readonly ariaLabel = input('Sections');
  readonly selectedIndex = model(0);

  select(index: number): void {
    if (!this.disabled()[index]) {
      this.selectedIndex.set(index);
    }
  }

  handleKeydown(event: KeyboardEvent, currentIndex: number): void {
    const keys = ['ArrowLeft', 'ArrowRight', 'Home', 'End'];
    if (!keys.includes(event.key)) {
      return;
    }

    event.preventDefault();
    const tabs = this.tabs();
    const enabled = tabs.map((_, index) => index).filter(index => !this.disabled()[index]);
    if (!enabled.length) {
      return;
    }

    const position = enabled.indexOf(currentIndex);
    let nextIndex: number;
    if (event.key === 'Home') {
      nextIndex = enabled[0];
    } else if (event.key === 'End') {
      nextIndex = enabled[enabled.length - 1];
    } else {
      const direction = event.key === 'ArrowRight' ? 1 : -1;
      nextIndex = enabled[(position + direction + enabled.length) % enabled.length];
    }

    this.select(nextIndex);
    const tabList = (event.currentTarget as HTMLElement).parentElement;
    (tabList?.querySelectorAll<HTMLButtonElement>('[role="tab"]')[nextIndex])?.focus();
  }
}
