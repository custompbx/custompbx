import {Component, Input, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {AbstractControl, FormsModule} from '@angular/forms';
import {ResizeInputDirective} from "../../directives/resize-input.directive";
import {MatBottomSheet} from "@angular/material/bottom-sheet";
import {ConfirmBottomSheetComponent} from "../confirm-bottom-sheet/confirm-bottom-sheet.component";

@Component({
standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, ResizeInputDirective],
    selector: 'app-key-value-pad-2',
    templateUrl: './key-value-pad-2.component.html',
    styleUrls: ['./key-value-pad-2.component.css']
})
export class KeyValuePad2Component implements OnInit {

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
    extraField3?: {name: string, style?: object, depend?: string, value?: string},
    extraField4?: {name: string, style?: object, depend?: string, value?: string},
    extraField5?: {name: string, style?: object, depend?: string, value?: string},
    extraField6?: {name: string, style?: object, depend?: string, value?: string},
    extraField7?: {name: string, style?: object, depend?: string, value?: string},
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

  isReadyToSend(
    nameObject: AbstractControl,
    valueObject: AbstractControl,
    valueObject1: AbstractControl,
    valueObject2: AbstractControl,
    valueObject3: AbstractControl,
    valueObject4: AbstractControl,
    valueObject5: AbstractControl,
    valueObject6: AbstractControl,
    valueObject7: AbstractControl,
  ): boolean {
    const optionalControls = this.presentControls(
      valueObject,
      valueObject1,
      valueObject2,
      valueObject3,
      valueObject4,
      valueObject5,
      valueObject6,
      valueObject7,
    );
    return !!nameObject && nameObject.valid && this.hasDirtyValidControl(nameObject, ...optionalControls);
  }

  canSaveExisting(
    item: any,
    nameObject: AbstractControl,
    valueObject: AbstractControl,
    valueObject1: AbstractControl,
    valueObject2: AbstractControl,
    valueObject3: AbstractControl,
    valueObject4: AbstractControl,
    valueObject5: AbstractControl,
    valueObject6: AbstractControl,
    valueObject7: AbstractControl,
  ): boolean {
    return this.isReadyToSend(
      nameObject,
      valueObject,
      valueObject1,
      valueObject2,
      valueObject3,
      valueObject4,
      valueObject5,
      valueObject6,
      valueObject7,
    ) && !this.nameValidationMessage(nameObject?.value, item?.id);
  }

  isNewReadyToSend(
    nameObject: AbstractControl,
    valueObject: AbstractControl,
    valueObject1: AbstractControl,
    valueObject2: AbstractControl,
    valueObject3: AbstractControl,
    valueObject4: AbstractControl,
    valueObject5: AbstractControl,
    valueObject6: AbstractControl,
    valueObject7: AbstractControl,
  ): boolean {
    return !!nameObject && nameObject.valid && this.presentControls(
      valueObject,
      valueObject1,
      valueObject2,
      valueObject3,
      valueObject4,
      valueObject5,
      valueObject6,
      valueObject7,
    ).every((control) => control.disabled || control.valid);
  }

  canAddNew(
    index: number,
    nameObject: AbstractControl,
    valueObject: AbstractControl,
    valueObject1: AbstractControl,
    valueObject2: AbstractControl,
    valueObject3: AbstractControl,
    valueObject4: AbstractControl,
    valueObject5: AbstractControl,
    valueObject6: AbstractControl,
    valueObject7: AbstractControl,
  ): boolean {
    return this.isNewReadyToSend(
      nameObject,
      valueObject,
      valueObject1,
      valueObject2,
      valueObject3,
      valueObject4,
      valueObject5,
      valueObject6,
      valueObject7,
    ) && !this.nameValidationMessage(nameObject?.value, null, index);
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

  updateItem(
    id: number,
    name: string,
    value?: string,
    exta1?: string,
    exta2?: string,
    exta3?: string,
    exta4?: string,
    exta5?: string,
    exta6?: string,
    exta7?: string
    ) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['updateItem']) {
      return;
    }
    const obj = {
      id: id,
      [this.fieldsMask.name.name]: this.trimValue(name),
      ...this.maskedValues([value, exta1, exta2, exta3, exta4, exta5, exta6, exta7]),
    };
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

  addItem(
    index: number,
    name: string,
    value: string,
    extraField1: string,
    extraField2: string,
    extraField3: string,
    extraField4: string,
    extraField5: string,
    extraField6: string,
    extraField7: string
    ) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['addItem']) {
      return;
    }
    const args = [];
    if (this.id !== null) {
      args.push(this.id);
    }

    args.push(index);
    args.push(this.trimValue(name));

    if (value !== undefined) {
      args.push(this.trimValue(value));
    }
    this.enabledExtraFieldValues([extraField1, extraField2, extraField3, extraField4, extraField5, extraField6, extraField7])
      .forEach((extraValue) => args.push(extraValue));
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

  itemValues(obj: object): Array<any> {
    return this.onlyValues(obj).filter((item) => item?.id && !this.isArray(item));
  }

  filteredValues(obj: object): Array<any> {
    const values = this.itemValues(obj);
    const query = this.filterText.trim().toLowerCase();
    if (!query) {
      return values;
    }
    return values.filter((item) => this.searchableText(item).includes(query));
  }

  hasItems(obj: object): boolean {
    return this.itemValues(obj).length > 0;
  }

  totalItemCount(obj: object): number {
    return this.itemValues(obj).length;
  }

  showFilter(obj: object): boolean {
    return this.totalItemCount(obj) > 1;
  }

  visibleItemCount(obj: object): number {
    return this.filteredValues(obj).length;
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
    this.setControlValue(form, 'itemExtraField3' + item.id, item?.[this.fieldsMask.extraField3?.name]);
    this.setControlValue(form, 'itemExtraField4' + item.id, item?.[this.fieldsMask.extraField4?.name]);
    this.setControlValue(form, 'itemExtraField5' + item.id, item?.[this.fieldsMask.extraField5?.name]);
    this.setControlValue(form, 'itemExtraField6' + item.id, item?.[this.fieldsMask.extraField6?.name]);
    this.setControlValue(form, 'itemExtraField7' + item.id, item?.[this.fieldsMask.extraField7?.name]);
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

  private presentControls(...controls: AbstractControl[]): AbstractControl[] {
    return controls.filter(Boolean);
  }

  private controlsForItem(form: any, id: number): AbstractControl[] {
    return [
      form?.controls?.['itemName' + id],
      form?.controls?.['itemValue' + id],
      form?.controls?.['itemExtraField1' + id],
      form?.controls?.['itemExtraField2' + id],
      form?.controls?.['itemExtraField3' + id],
      form?.controls?.['itemExtraField4' + id],
      form?.controls?.['itemExtraField5' + id],
      form?.controls?.['itemExtraField6' + id],
      form?.controls?.['itemExtraField7' + id],
    ].filter(Boolean);
  }

  private setControlValue(form: any, name: string, value: any): void {
    const control = form?.controls?.[name];
    if (control) {
      control.setValue(value ?? '');
    }
  }

  private searchableText(item: any): string {
    return this.fieldValueEntries([
      item?.[this.fieldsMask.value?.name],
      item?.[this.fieldsMask.extraField1?.name],
      item?.[this.fieldsMask.extraField2?.name],
      item?.[this.fieldsMask.extraField3?.name],
      item?.[this.fieldsMask.extraField4?.name],
      item?.[this.fieldsMask.extraField5?.name],
      item?.[this.fieldsMask.extraField6?.name],
      item?.[this.fieldsMask.extraField7?.name],
    ])
      .map((entry) => entry.value)
      .concat(item?.[this.fieldsMask.name.name], item?.id?.toString())
      .filter((value) => value !== undefined && value !== null)
      .join(' ')
      .toLowerCase();
  }

  private hasDirtyValidControl(...controls: AbstractControl[]): boolean {
    return controls.some((control) => control.dirty && control.valid);
  }

  private maskedValues(values: string[]): object {
    return this.fieldValueEntries(values)
      .reduce((result, entry) => ({...result, [entry.name]: entry.value}), {});
  }

  private enabledExtraFieldValues(values: string[]): string[] {
    return this.extraFieldValueEntries(values)
      .map((entry) => this.trimValue(entry.value));
  }

  private fieldValueEntries(values: string[]): Array<{name: string, value: string}> {
    return [
      {name: this.fieldsMask.value?.name, value: this.trimValue(values[0])},
      ...this.extraFieldValueEntries(values.slice(1)),
    ].filter((entry): entry is {name: string, value: string} => !!entry.name);
  }

  private extraFieldValueEntries(values: string[]): Array<{name?: string, value: string}> {
    return [
      {name: this.fieldsMask.extraField1?.name, value: this.trimValue(values[0])},
      {name: this.fieldsMask.extraField2?.name, value: this.trimValue(values[1])},
      {name: this.fieldsMask.extraField3?.name, value: this.trimValue(values[2])},
      {name: this.fieldsMask.extraField4?.name, value: this.trimValue(values[3])},
      {name: this.fieldsMask.extraField5?.name, value: this.trimValue(values[4])},
      {name: this.fieldsMask.extraField6?.name, value: this.trimValue(values[5])},
      {name: this.fieldsMask.extraField7?.name, value: this.trimValue(values[6])},
    ].filter((entry) => !!entry.name);
  }

  private hasDuplicateName(name: string, id?: number, newIndex?: number): boolean {
    const existingDuplicate = this.onlyValues(this.items).some((item) => {
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
