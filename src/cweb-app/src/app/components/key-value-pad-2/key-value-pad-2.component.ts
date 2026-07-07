import {Component, Input, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {AbstractControl, FormsModule} from '@angular/forms';
import {ResizeInputDirective} from "../../directives/resize-input.directive";

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

  constructor() { }

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

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFnId(index, item) {
    if (item.id) {
      return item.id;
    }
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
      [this.fieldsMask.name.name]: name,
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
    args.push(name);

    if (value !== undefined) {
      args.push(value);
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

  private presentControls(...controls: AbstractControl[]): AbstractControl[] {
    return controls.filter(Boolean);
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
      .map((entry) => entry.value);
  }

  private fieldValueEntries(values: string[]): Array<{name: string, value: string}> {
    return [
      {name: this.fieldsMask.value?.name, value: values[0]},
      ...this.extraFieldValueEntries(values.slice(1)),
    ].filter((entry): entry is {name: string, value: string} => !!entry.name);
  }

  private extraFieldValueEntries(values: string[]): Array<{name?: string, value: string}> {
    return [
      {name: this.fieldsMask.extraField1?.name, value: values[0]},
      {name: this.fieldsMask.extraField2?.name, value: values[1]},
      {name: this.fieldsMask.extraField3?.name, value: values[2]},
      {name: this.fieldsMask.extraField4?.name, value: values[3]},
      {name: this.fieldsMask.extraField5?.name, value: values[4]},
      {name: this.fieldsMask.extraField6?.name, value: values[5]},
      {name: this.fieldsMask.extraField7?.name, value: values[6]},
    ].filter((entry) => !!entry.name);
  }

}
