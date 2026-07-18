import {Component, Input, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {DragDropModule} from '@angular/cdk/drag-drop';
import {AbstractControl, FormsModule} from '@angular/forms';
import {CdkDragDrop} from '@angular/cdk/drag-drop';
import {ConfirmationService} from '../../services/confirmation.service';
import {resolvePositionedReorder} from '../../utils/reorder';
import {TranslocoPipe, TranslocoService} from '@jsverse/transloco';

type FieldSlot = 'name' | 'value' | 'extraField1' | 'extraField2' | 'extraField3' | 'extraField4' | 'extraField5' | 'extraField6' | 'extraField7';
type FieldSize = 'sm' | 'md' | 'wide';
type FieldConfig = {
  slot: FieldSlot,
  name: string,
  style?: Record<string, string>,
  depend?: string,
  value?: string,
  required?: boolean,
  options?: string[],
  size?: FieldSize,
};

@Component({
standalone: true,
  imports: [CommonModule, DragDropModule, FormsModule, TranslocoPipe],
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
    name: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
    value?: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
    extraField1?: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
    extraField2?: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
    extraField3?: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
    extraField4?: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
    extraField5?: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
    extraField6?: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
    extraField7?: {name: string, style?: Record<string, string>, depend?: string, value?: string, options?: string[], size?: FieldSize, required?: boolean},
  };
  @Input() grandParentId: number;
  @Input() pending = false;
  @Input() sortable = false;
  public filterText = '';
  protected readonly fieldSlots: FieldSlot[] = [
    'name',
    'value',
    'extraField1',
    'extraField2',
    'extraField3',
    'extraField4',
    'extraField5',
    'extraField6',
    'extraField7',
  ];

  constructor(
    private bottomSheet: ConfirmationService,
    private transloco: TranslocoService,
  ) { }

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

  fieldConfigs(): FieldConfig[] {
    return this.fieldSlots
      .map((slot) => ({slot, ...(this.fieldsMask?.[slot] || {})}))
      .filter((field): field is FieldConfig => !!field.name)
      .map((field) => ({...field, required: field.required ?? field.slot === 'name'}));
  }

  fieldClass(field: FieldConfig): string {
    const classes = [];
    if (field.size) {
      classes.push(`kv-field--${field.size}`);
      if (field.options?.length) {
        classes.push('kv-field--select');
      }
      return classes.join(' ');
    }
    const widthValue = field.style?.['max-width'] || field.style?.['width'];
    const width = typeof widthValue === 'string' ? Number.parseInt(widthValue, 10) : Number.NaN;
    if (Number.isFinite(width) && width <= 120) {
      classes.push('kv-field--sm');
    }
    if (Number.isFinite(width) && width >= 420) {
      classes.push('kv-field--wide');
    }
    if (field.options?.length) {
      classes.push('kv-field--select');
    }
    return classes.join(' ');
  }

  optionLabel(option: string): string {
    return option === '' ? this.transloco.translate('common.default') : option;
  }

  controlName(prefix: 'item' | 'newItem', field: FieldConfig, idOrIndex: number): string {
    const suffix = field.slot === 'name'
      ? 'Name'
      : field.slot === 'value'
        ? 'Value'
        : field.slot.charAt(0).toUpperCase() + field.slot.slice(1);
    return prefix + suffix + idOrIndex;
  }

  itemFieldValue(item: any, field: FieldConfig): any {
    return item?.[field.name];
  }

  setNewItemField(item: any, field: FieldConfig, value: any): void {
    if (!item) {
      return;
    }
    if (field.slot === 'name') {
      item.name = value;
      return;
    }
    if (field.slot === 'value') {
      item.value = value;
      return;
    }
    item[field.name] = value;
  }

  newItemFieldValue(item: any, field: FieldConfig): any {
    if (field.slot === 'name') {
      return item?.name;
    }
    if (field.slot === 'value') {
      return item?.value;
    }
    return item?.[field.name];
  }

  isFieldDisabled(item: any, field: FieldConfig): boolean {
    return !item?.enabled || this.disIf(field.depend, field.value);
  }

  isNewFieldDisabled(item: any, field: FieldConfig): boolean {
    return this.disIf(item?.[field.depend], field.value);
  }

  canSaveExistingForm(form: any, item: any): boolean {
    const controls = this.fieldConfigs().map((field) => form.controls[this.controlName('item', field, item.id)]);
    return this.canSaveExisting(
      item,
      controls[0],
      controls[1],
      controls[2],
      controls[3],
      controls[4],
      controls[5],
      controls[6],
      controls[7],
      controls[8],
    );
  }

  saveExistingFromForm(form: any, item: any): void {
    const values = this.fieldConfigs().map((field) => form.controls[this.controlName('item', field, item.id)]?.value);
    this.updateItem(item.id, values[0], values[1], values[2], values[3], values[4], values[5], values[6], values[7], values[8]);
  }

  canAddNewForm(form: any, index: number): boolean {
    const controls = this.fieldConfigs().map((field) => form.controls[this.controlName('newItem', field, index)]);
    return this.canAddNew(
      index,
      controls[0],
      controls[1],
      controls[2],
      controls[3],
      controls[4],
      controls[5],
      controls[6],
      controls[7],
      controls[8],
    );
  }

  addNewFromForm(form: any, index: number): void {
    const values = this.fieldConfigs().map((field) => form.controls[this.controlName('newItem', field, index)]?.value);
    this.addItem(index, values[0], values[1], values[2], values[3], values[4], values[5], values[6], values[7], values[8]);
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
    const sheet = this.bottomSheet.open({
      data: {
        action: 'delete',
        case1Text: this.transloco.translate('common.deleteItemConfirm', {name}),
        message: this.transloco.translate('common.deleteItemDescription'),
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
    const values = this.displayValues(obj);
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
    return this.displayValues(obj).length;
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
      return this.transloco.translate('common.nameRequired');
    }
    if (this.hasDuplicateName(normalizedName, id, newIndex)) {
      return this.transloco.translate('common.duplicateName');
    }
    return null;
  }

  private presentControls(...controls: AbstractControl[]): AbstractControl[] {
    return controls.filter(Boolean);
  }

  displayValues(obj: object): Array<any> {
    const values = this.itemValues(obj);
    if (!this.sortable) {
      return values;
    }
    return [...values].sort((a, b) => {
      if (a.position > b.position) {
        return 1;
      }
      if (a.position < b.position) {
        return -1;
      }
      return 0;
    });
  }

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (!this.sortable || !this.dispatchersCallbacks || !this.dispatchersCallbacks['dropActionItem']) {
      return;
    }
    if (resolvePositionedReorder(parent, event.previousIndex, event.currentIndex)) {
      this.dispatchersCallbacks['dropActionItem'](event, parent);
    }
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
      .filter((entry) => entry.value !== undefined)
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
