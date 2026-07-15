import {ChangeDetectionStrategy, Component, computed, input, output} from '@angular/core';

export interface CpbxPageEvent {
  pageIndex: number;
  pageSize: number;
  length?: number;
  previousPageIndex?: number;
}

@Component({
  selector: 'app-paginator',
  standalone: true,
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <nav class="cpbx-paginator" aria-label="Pagination">
      <span class="cpbx-paginator__summary">{{rangeLabel()}}</span>
      <label class="cpbx-paginator__size">
        <span>Rows</span>
        <select class="cpbx-select cpbx-select--compact" [value]="pageSize()" (change)="changeSize($event)">
          @for (size of pageSizeOptions(); track size) {
            <option [value]="size">{{size}}</option>
          }
        </select>
      </label>
      <div class="cpbx-paginator__actions">
        @if (showFirstLastButtons()) {
          <button type="button" class="cpbx-icon-button" aria-label="First page" [disabled]="pageIndex() === 0" (click)="goTo(0)">«</button>
        }
        <button type="button" class="cpbx-icon-button" aria-label="Previous page" [disabled]="pageIndex() === 0" (click)="goTo(pageIndex() - 1)">‹</button>
        <span class="cpbx-paginator__page">{{pageIndex() + 1}} / {{pageCount()}}</span>
        <button type="button" class="cpbx-icon-button" aria-label="Next page" [disabled]="pageIndex() >= pageCount() - 1" (click)="goTo(pageIndex() + 1)">›</button>
        @if (showFirstLastButtons()) {
          <button type="button" class="cpbx-icon-button" aria-label="Last page" [disabled]="pageIndex() >= pageCount() - 1" (click)="goTo(pageCount() - 1)">»</button>
        }
      </div>
    </nav>
  `,
  styleUrl: './paginator.component.css'
})
export class PaginatorComponent {
  readonly length = input(0);
  readonly pageIndex = input(0);
  readonly pageSize = input(25);
  readonly pageSizeOptions = input<number[]>([10, 25, 50, 100]);
  readonly showFirstLastButtons = input(false);
  readonly page = output<CpbxPageEvent>();

  readonly pageCount = computed(() => Math.max(1, Math.ceil(this.length() / Math.max(1, this.pageSize()))));
  readonly rangeLabel = computed(() => {
    if (!this.length()) {
      return '0 items';
    }
    const start = this.pageIndex() * this.pageSize() + 1;
    const end = Math.min(this.length(), start + this.pageSize() - 1);
    return `${start}–${end} of ${this.length()}`;
  });

  goTo(index: number): void {
    const next = Math.min(Math.max(0, index), this.pageCount() - 1);
    if (next === this.pageIndex()) {
      return;
    }
    this.page.emit({pageIndex: next, pageSize: this.pageSize(), length: this.length(), previousPageIndex: this.pageIndex()});
  }

  changeSize(event: Event): void {
    const pageSize = Number((event.target as HTMLSelectElement).value);
    this.page.emit({pageIndex: 0, pageSize, length: this.length(), previousPageIndex: this.pageIndex()});
  }
}
