import {
  AfterViewInit,
  Directive,
  DoCheck,
  ElementRef,
  HostListener,
  NgZone,
  OnDestroy,
  Renderer2,
} from '@angular/core';

@Directive({
  selector: 'select.cpbx-select',
  standalone: true,
})
export class CpbxSelectDirective implements AfterViewInit, DoCheck, OnDestroy {
  private wrapper?: HTMLElement;
  private button?: HTMLButtonElement;
  private menu?: HTMLElement;
  private observer?: MutationObserver;
  private removeDocumentClick?: () => void;
  private removeSelectChange?: () => void;
  private open = false;
  private renderedState = '';

  constructor(
    private readonly elementRef: ElementRef<HTMLSelectElement>,
    private readonly renderer: Renderer2,
    private readonly zone: NgZone,
  ) {}

  ngAfterViewInit(): void {
    const select = this.select;

    this.wrapper = this.renderer.createElement('div');
    this.renderer.addClass(this.wrapper, 'cpbx-select-menu');

    this.button = this.renderer.createElement('button');
    this.button.type = 'button';
    this.renderer.addClass(this.button, 'cpbx-select-menu__button');

    this.menu = this.renderer.createElement('div');
    this.renderer.addClass(this.menu, 'cpbx-select-menu__panel');
    this.renderer.setAttribute(this.menu, 'role', select.multiple ? 'listbox' : 'menu');
    if (select.multiple) this.renderer.setAttribute(this.menu, 'aria-multiselectable', 'true');

    const parent = select.parentElement;
    if (!parent) return;
    parent.insertBefore(this.wrapper, select);
    this.wrapper.appendChild(select);
    this.wrapper.appendChild(this.button);
    this.wrapper.appendChild(this.menu);

    this.renderer.addClass(select, 'cpbx-select--native-hidden');

    this.removeSelectChange = this.renderer.listen(select, 'change', () => this.refresh());
    this.renderer.listen(this.button, 'click', (event: MouseEvent) => {
      event.preventDefault();
      event.stopPropagation();
      if (select.disabled) return;
      this.setOpen(!this.open);
    });

    this.zone.runOutsideAngular(() => {
      this.removeDocumentClick = this.renderer.listen('document', 'click', (event: MouseEvent) => {
        if (!this.wrapper?.contains(event.target as Node)) this.setOpen(false);
      });
      this.observer = new MutationObserver(() => {
        // Angular can register async options before its value accessor updates
        // their selected state. Refresh after the current render turn.
        queueMicrotask(() => this.refresh());
      });
      this.observer.observe(select, { childList: true, subtree: true, attributes: true });
    });

    this.refresh();
  }

  ngDoCheck(): void {
    if (!this.button || !this.menu) return;
    const state = this.currentState();
    if (state !== this.renderedState) this.refresh();
  }

  ngOnDestroy(): void {
    this.observer?.disconnect();
    this.removeDocumentClick?.();
    this.removeSelectChange?.();
  }

  @HostListener('disabled')
  refresh(): void {
    if (!this.button || !this.menu) return;
    const select = this.select;
    this.button.disabled = select.disabled;
    this.button.textContent = this.buttonText();
    this.button.setAttribute('aria-expanded', String(this.open));
    this.button.setAttribute('aria-haspopup', 'listbox');
    this.renderOptions();
    this.renderedState = this.currentState();
  }

  private currentState(): string {
    const select = this.select;
    return JSON.stringify({
      disabled: select.disabled,
      optionCount: select.options.length,
      selected: Array.from(select.options)
        .filter((option) => option.selected)
        .map((option) => option.value),
    });
  }

  private get select(): HTMLSelectElement {
    return this.elementRef.nativeElement;
  }

  private buttonText(): string {
    const selected = Array.from(this.select.selectedOptions)
      .filter((option) => !option.disabled || option.value)
      .map((option) => option.textContent?.trim())
      .filter(Boolean) as string[];

    if (selected.length) return selected.join(', ');

    const placeholder = Array.from(this.select.options).find((option) => option.disabled && option.hidden);
    return placeholder?.textContent?.trim() || 'Select';
  }

  private renderOptions(): void {
    if (!this.menu) return;
    this.menu.replaceChildren();

    Array.from(this.select.children).forEach((child) => {
      if (child instanceof HTMLOptGroupElement) {
        const group = this.renderer.createElement('div') as HTMLElement;
        this.renderer.addClass(group, 'cpbx-select-menu__group');
        const label = this.renderer.createElement('div') as HTMLElement;
        this.renderer.addClass(label, 'cpbx-select-menu__group-label');
        label.textContent = child.label;
        group.appendChild(label);
        Array.from(child.children).forEach((option) => {
          if (option instanceof HTMLOptionElement) group.appendChild(this.optionButton(option));
        });
        this.menu?.appendChild(group);
      } else if (child instanceof HTMLOptionElement) {
        if (child.hidden) return;
        this.menu?.appendChild(this.optionButton(child));
      }
    });
  }

  private optionButton(option: HTMLOptionElement): HTMLElement {
    const row = this.renderer.createElement('button') as HTMLButtonElement;
    row.type = 'button';
    this.renderer.addClass(row, 'cpbx-select-menu__option');
    if (option.selected) this.renderer.addClass(row, 'cpbx-select-menu__option--selected');
    if (option.disabled) row.disabled = true;
    row.setAttribute('role', 'option');
    row.setAttribute('aria-selected', String(option.selected));

    const mark = this.renderer.createElement('span') as HTMLElement;
    this.renderer.addClass(mark, 'cpbx-select-menu__mark');
    mark.textContent = option.selected ? '✓' : '';

    const text = this.renderer.createElement('span') as HTMLElement;
    text.textContent = option.textContent?.trim() || '';

    row.appendChild(mark);
    row.appendChild(text);

    this.renderer.listen(row, 'click', (event: MouseEvent) => {
      event.preventDefault();
      event.stopPropagation();
      if (option.disabled) return;
      this.selectOption(option);
    });

    return row;
  }

  private selectOption(option: HTMLOptionElement): void {
    const select = this.select;
    if (select.multiple) {
      option.selected = !option.selected;
    } else {
      Array.from(select.options).forEach((candidate) => (candidate.selected = candidate === option));
      this.setOpen(false);
    }

    select.dispatchEvent(new Event('input', { bubbles: true }));
    select.dispatchEvent(new Event('change', { bubbles: true }));
    this.refresh();
  }

  private setOpen(next: boolean): void {
    this.open = next;
    this.wrapper?.classList.toggle('cpbx-select-menu--open', this.open);
    this.button?.setAttribute('aria-expanded', String(this.open));
  }
}
