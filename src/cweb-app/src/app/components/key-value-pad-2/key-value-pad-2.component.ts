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
    let obj1 = false;
    let obj2 = false;
    let obj3 = false;
    let obj4 = false;
    let obj5 = false;
    let obj7 = false;
    let obj6 = false;
    let value = true;
    if (valueObject1) {
      obj1 = valueObject1.dirty && valueObject1.valid;
    }
    if (valueObject2) {
      obj2 = valueObject2.dirty && valueObject2.valid;
    }
    if (valueObject3) {
      obj3 = valueObject3.dirty && valueObject3.valid;
    }
    if (valueObject4) {
      obj4 = valueObject4.dirty && valueObject4.valid;
    }
    if (valueObject5) {
      obj5 = valueObject5.dirty && valueObject5.valid;
    }
    if (valueObject6) {
      obj6 = valueObject6.dirty && valueObject6.valid;
    }
    if (valueObject7) {
      obj7 = valueObject7.dirty && valueObject7.valid;
    }
    if (valueObject) {
      value = valueObject.dirty && valueObject.valid;
    }
    return nameObject && nameObject.valid && (nameObject.dirty || obj1 || obj2 || obj3 || obj4 || obj5 || obj6 || obj7 || value);
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
    let obj1 = true;
    let obj2 = true;
    let obj3 = true;
    let obj4 = true;
    let obj5 = true;
    let obj6 = true;
    let obj7 = true;
    let value = true;
    if (valueObject1 && !valueObject1.disabled) {
      obj1 = valueObject1.valid;
    }
    if (valueObject2 && !valueObject2.disabled) {
      obj2 = valueObject2.valid;
    }
    if (valueObject3 && !valueObject3.disabled) {
      obj3 = valueObject3.valid;
    }
    if (valueObject4 && !valueObject4.disabled) {
      obj4 = valueObject4.valid;
    }
    if (valueObject5 && !valueObject5.disabled) {
      obj5 = valueObject5.valid;
    }
    if (valueObject6 && !valueObject6.disabled) {
      obj6 = valueObject6.valid;
    }
    if (valueObject7 && !valueObject7.disabled) {
      obj7 = valueObject7.valid;
    }
    if (valueObject && !valueObject.disabled) {
      value = valueObject.valid;
    }
    return nameObject && nameObject.valid && value && obj1 && obj2 && obj3 && obj4 && obj5 && obj6 && obj7;
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
    const obj = {id: id, [this.fieldsMask.name.name]: name};

    if (this.fieldsMask.value) {
      obj[this.fieldsMask.value.name] = value;
    }
    if (this.fieldsMask.extraField1) {
      obj[this.fieldsMask.extraField1.name] = exta1;
    }
    if (this.fieldsMask.extraField2) {
      obj[this.fieldsMask.extraField2.name] = exta2;
    }
    if (this.fieldsMask.extraField3) {
      obj[this.fieldsMask.extraField3.name] = exta3;
    }
    if (this.fieldsMask.extraField4) {
      obj[this.fieldsMask.extraField4.name] = exta4;
    }
    if (this.fieldsMask.extraField5) {
      obj[this.fieldsMask.extraField5.name] = exta5;
    }
    if (this.fieldsMask.extraField6) {
      obj[this.fieldsMask.extraField6.name] = exta6;
    }
    if (this.fieldsMask.extraField7) {
      obj[this.fieldsMask.extraField7.name] = exta7;
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
    if (this.fieldsMask.extraField1) {
      args.push(extraField1);
    }
    if (this.fieldsMask.extraField2) {
      args.push(extraField2);
    }
    if (this.fieldsMask.extraField3) {
      args.push(extraField3);
    }
    if (this.fieldsMask.extraField4) {
      args.push(extraField4);
    }
    if (this.fieldsMask.extraField5) {
      args.push(extraField5);
    }
    if (this.fieldsMask.extraField6) {
      args.push(extraField6);
    }
    if (this.fieldsMask.extraField7) {
      args.push(extraField7);
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

}
