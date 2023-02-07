import {Component, Input, OnInit} from '@angular/core';
import {AbstractControl} from '@angular/forms';

@Component({
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
    value: {name: string, style?: object, depend?: string, value?: string},
    extraField1?: {name: string, style?: object, depend?: string, value?: string},
    extraField2?: {name: string, style?: object, depend?: string, value?: string},
  };
  @Input() grandParentId: number;

  constructor() { }

  ngOnInit() {
    if (!this.fieldsMask) {
      this.fieldsMask = {name: {name: 'name'}, value: {name: 'value'}};
    }
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl, valueObject1: AbstractControl, valueObject2: AbstractControl): boolean {
    let obj1 = false;
    let obj2 = false;
    if (valueObject1) {
      obj1 = valueObject1.dirty && valueObject1.valid;
    }
    if (valueObject2) {
      obj2 = valueObject2.dirty && valueObject2.valid;
    }
    return nameObject && valueObject && nameObject.valid && valueObject.valid && ((nameObject.dirty || valueObject.dirty) || obj1 || obj2 );
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl, valueObject1: AbstractControl, valueObject2: AbstractControl): boolean {
    let obj1 = true;
    let obj2 = true;
    if (valueObject1 && !valueObject1.disabled) {
      obj1 = valueObject1.valid;
    }
    if (valueObject2 && !valueObject1.disabled) {
      obj2 = valueObject2.valid;
    }
    return nameObject && valueObject && nameObject.valid && valueObject.valid && obj1 && obj2;
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

  updateItem(obj: object) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['updateItem']) {
      return;
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

  addItem(index: number, name: string, value: string, extraField1: string, extraField2: string) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['addItem']) {
      return;
    }
    if (this.id === null) {
      if (this.fieldsMask.extraField1 && this.fieldsMask.extraField2) {
        this.dispatchersCallbacks['addItem'](index, name, value, extraField1, extraField2);
        return;
      }
      if (this.fieldsMask.extraField1) {
        this.dispatchersCallbacks['addItem'](index, name, value, extraField1);
        return;
      }
      this.dispatchersCallbacks['addItem'](index, name, value);
      return;
    }

    if (this.fieldsMask.extraField1 && this.fieldsMask.extraField2) {
      this.dispatchersCallbacks['addItem'](this.id, index, name, value, extraField1, extraField2);
      return;
    }
    if (this.fieldsMask.extraField1) {
      this.dispatchersCallbacks['addItem'](this.id, index, name, value, extraField1);
      return;
    }
    this.dispatchersCallbacks['addItem'](this.id, index, name, value);
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
