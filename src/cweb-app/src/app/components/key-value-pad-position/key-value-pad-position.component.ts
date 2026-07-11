import {Component, Input, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {AbstractControl, FormsModule} from '@angular/forms';
import {CdkDragDrop} from '@angular/cdk/drag-drop';
import {ResizeInputDirective} from "../../directives/resize-input.directive";
import {MatBottomSheet} from "@angular/material/bottom-sheet";
import {ConfirmBottomSheetComponent} from "../confirm-bottom-sheet/confirm-bottom-sheet.component";

@Component({
standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, ResizeInputDirective],
    selector: 'app-key-value-pad-position',
    templateUrl: './key-value-pad-position.component.html',
    styleUrls: ['./key-value-pad-position.component.css']
})
export class KeyValuePadPositionComponent implements OnInit {

  @Input() items: object;
  @Input() newItems: Array<any>;
  @Input() exist: boolean;
  @Input() id: number;
  @Input() toCopy: number;
  @Input() dispatchersCallbacks: any;
  @Input() fieldsMask: {
    name: {name: string, style?: object, depend?: string, value?: string},
    value?: {name: string, style?: object, depend?: string, value?: string},
    extraField1?: {name: string, style?: object, depend?: string, value?: string},
    extraField2?: {name: string, style?: object, depend?: string, value?: string},
  };
  @Input() grandParentId: number;
  @Input() pending = false;
  public filterText = '';

  constructor(private bottomSheet: MatBottomSheet) { }

  ngOnInit() {
    if (!this.fieldsMask) {
      this.fieldsMask = {name: {name: 'name'}, value: {name: 'value'}};
    }
  }

  isReadyToSend(nameObject: AbstractControl,
                valueObject: AbstractControl, valueObject1: AbstractControl, valueObject2: AbstractControl): boolean {
    let obj1 = false;
    let obj2 = false;
    let value = true;
    if (valueObject1) {
      obj1 = valueObject1.dirty && valueObject1.valid;
    }
    if (valueObject2) {
      obj2 = valueObject2.dirty && valueObject2.valid;
    }
    if (valueObject) {
      value = valueObject.dirty && valueObject.valid;
    }
    return nameObject && nameObject.valid && (nameObject.dirty || obj1 || obj2 || value);
  }

  canSaveExisting(item: any, nameObject: AbstractControl,
                valueObject: AbstractControl, valueObject1: AbstractControl, valueObject2: AbstractControl): boolean {
    return this.isReadyToSend(nameObject, valueObject, valueObject1, valueObject2)
      && !this.nameValidationMessage(nameObject?.value, item?.id);
  }

  isNewReadyToSend(nameObject: AbstractControl,
                   valueObject: AbstractControl, valueObject1: AbstractControl, valueObject2: AbstractControl): boolean {
    let obj1 = true;
    let obj2 = true;
    let value = true;
    if (valueObject1 && !valueObject1.disabled) {
      obj1 = valueObject1.valid;
    }
    if (valueObject2 && !valueObject2.disabled) {
      obj2 = valueObject2.valid;
    }
    if (valueObject && !valueObject.disabled) {
      value = valueObject.valid;
    }
    return nameObject && nameObject.valid && value && obj1 && obj2;
  }

  canAddNew(index: number, nameObject: AbstractControl,
                   valueObject: AbstractControl, valueObject1: AbstractControl, valueObject2: AbstractControl): boolean {
    return this.isNewReadyToSend(nameObject, valueObject, valueObject1, valueObject2)
      && !this.nameValidationMessage(nameObject?.value, null, index);
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFnId(index, item) {
    if (item.id) {
      return item.id;
    }
  }

  clearFilter(): void {
    this.filterText = '';
  }

  addNewItemField() {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['addNewItemField']) {
      return;
    }
    if (this.id === null) {
      this.dispatchersCallbacks['addNewItemField']();
      return;
    }
    if (this.grandParentId) {
      this.dispatchersCallbacks['addNewItemField'](this.grandParentId, this.id);
      return;
    }
    this.dispatchersCallbacks['addNewItemField'](this.id);
  }

  switchItem(obj: object) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['switchItem']) {
      return;
    }
    this.dispatchersCallbacks['switchItem'](obj);
  }

  deleteItem(obj: object) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['deleteItem']) {
      return;
    }
    this.dispatchersCallbacks['deleteItem'](obj);
  }

  confirmDeleteItem(obj: any): void {
    const name = this.itemDisplayName(obj);
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, {
      data: {
        action: 'delete',
        case1Text: `Delete item "${name}"?`,
        message: 'This removes the item from the configuration after the server confirms the change.',
      },
    });

    sheet.afterDismissed().subscribe((confirmed) => {
      if (confirmed) {
        this.deleteItem(obj);
      }
    });
  }

  updateItem(id: number, name: string, value: string, exta1?: string, exta2?: string) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['updateItem']) {
      return;
    }
    const obj = {id: id, name: this.trimValue(name), value: this.trimValue(value)};

    if (this.fieldsMask.extraField1) {
      obj[this.fieldsMask.extraField1.name] = this.trimValue(exta1);
    }
    if (this.fieldsMask.extraField2) {
      obj[this.fieldsMask.extraField2.name] = this.trimValue(exta2);
    }
    this.dispatchersCallbacks['updateItem'](obj);
  }

  pasteItems() {
    if (!this.toCopy || !this.dispatchersCallbacks || !this.dispatchersCallbacks['pasteItems']) {
      return;
    }
    if (this.id === null) {
      return;
    }
    if (this.grandParentId) {
      this.dispatchersCallbacks['pasteItems'](this.grandParentId, this.id);
      return;
    }
    this.dispatchersCallbacks['pasteItems'](this.id);
  }

  addItem(index: number, name: string, value?: string, extraField1?: string, extraField2?: string) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['addItem']) {
      return;
    }

    const args = [this.id, index, this.trimValue(name)];
    if (value !== undefined) {
      args.push(this.trimValue(value));
    }
    if (this.fieldsMask.extraField1) {
      args.push(this.trimValue(extraField1));
    }
    if (this.fieldsMask.extraField2) {
      args.push(this.trimValue(extraField2));
    }

    this.dispatchersCallbacks['addItem'](...args);
  }

  dropNewItem(index: number) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['dropNewItem']) {
      return;
    }
    if (this.id === null) {
      this.dispatchersCallbacks['dropNewItem'](index);
      return;
    }
    if (this.grandParentId) {
      this.dispatchersCallbacks['dropNewItem'](this.grandParentId, this.id, index);
      return;
    }
    this.dispatchersCallbacks['dropNewItem'](this.id, index);
  }

  disIf(name: string, value: string): boolean {

    return name !== value;
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  onlySortedValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    const arr = Object.values(obj).filter((item) => item?.id && !this.isArray(item)).sort(
      function (a, b) {
        if (a.position > b.position) {
          return 1;
        }
        if (a.position < b.position) {
          return -1;
        }
        return 0;
      }
    );
    return arr;
  }

  filteredSortedValues(obj: object): Array<any> {
    const values = this.onlySortedValues(obj);
    const query = this.filterText.trim().toLowerCase();
    if (!query) {
      return values;
    }
    return values.filter((item) => this.searchableText(item).includes(query));
  }

  hasItems(obj: object): boolean {
    return this.onlySortedValues(obj).length > 0;
  }

  totalItemCount(obj: object): number {
    return this.onlySortedValues(obj).length;
  }

  showFilter(obj: object): boolean {
    return this.totalItemCount(obj) > 1;
  }

  visibleItemCount(obj: object): number {
    return this.filteredSortedValues(obj).length;
  }

  hasNewItems(): boolean {
    return (this.newItems || []).some((item) => !!item);
  }

  isRowDirty(form: any, id: number): boolean {
    return this.controlsForItem(form, id).some((control) => control?.dirty);
  }

  markItemPristine(form: any, id: number): void {
    this.controlsForItem(form, id).forEach((control) => control?.markAsPristine());
  }

  resetItemForm(form: any, item: any): void {
    this.setControlValue(form, 'itemName' + item.id, item?.[this.fieldsMask.name.name]);
    this.setControlValue(form, 'itemValue' + item.id, item?.[this.fieldsMask.value?.name]);
    this.setControlValue(form, 'itemExtraField1' + item.id, item?.[this.fieldsMask.extraField1?.name]);
    this.setControlValue(form, 'itemExtraField2' + item.id, item?.[this.fieldsMask.extraField2?.name]);
    this.markItemPristine(form, item.id);
  }

  nameValidationMessage(name: string, id?: number, newIndex?: number): string | null {
    const normalizedName = this.normalizeName(name);
    if (!normalizedName) {
      return 'Name is required.';
    }
    if (this.hasDuplicateName(normalizedName, id, newIndex)) {
      return 'Another item already uses this name.';
    }
    return null;
  }

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['dropActionItem']) {
      return;
    }
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }

    this.dispatchersCallbacks['dropActionItem'](event, parent);
  }

  private controlsForItem(form: any, id: number): AbstractControl[] {
    return [
      form?.controls?.['itemName' + id],
      form?.controls?.['itemValue' + id],
      form?.controls?.['itemExtraField1' + id],
      form?.controls?.['itemExtraField2' + id],
    ].filter(Boolean);
  }

  private setControlValue(form: any, name: string, value: any): void {
    const control = form?.controls?.[name];
    if (control) {
      control.setValue(value ?? '');
    }
  }

  private searchableText(item: any): string {
    return [
      item?.[this.fieldsMask.name.name],
      item?.[this.fieldsMask.value?.name],
      item?.[this.fieldsMask.extraField1?.name],
      item?.[this.fieldsMask.extraField2?.name],
      item?.id?.toString(),
    ]
      .filter((value) => value !== undefined && value !== null)
      .join(' ')
      .toLowerCase();
  }

  private hasDuplicateName(name: string, id?: number, newIndex?: number): boolean {
    const existingDuplicate = this.onlySortedValues(this.items).some((item) => {
      if (!item?.id || item.id === id || this.isArray(item)) {
        return false;
      }
      return this.normalizeName(item?.[this.fieldsMask.name.name]) === name;
    });

    if (existingDuplicate) {
      return true;
    }

    return (this.newItems || []).some((item, index) => {
      if (!item || index === newIndex) {
        return false;
      }
      return this.normalizeName(item.name) === name;
    });
  }

  private itemDisplayName(item: any): string {
    return this.trimValue(item?.[this.fieldsMask.name.name]) || `#${item?.id || ''}`.trim();
  }

  private normalizeName(value: any): string {
    return (this.trimValue(value) || '').toString().toLowerCase();
  }

  private trimValue(value: any): any {
    return typeof value === 'string' ? value.trim() : value;
  }
}
