import {Component, Input, OnInit} from '@angular/core';
import {AbstractControl} from '@angular/forms';
import {CdkDragDrop} from '@angular/cdk/drag-drop';

@Component({
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

  constructor() { }

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

  updateItem(id: number, name: string, value: string, exta1?: string, exta2?: string) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['updateItem']) {
      return;
    }
    const obj = {id: id, name: name, value: value};

    if (this.fieldsMask.extraField1) {
      obj[this.fieldsMask.extraField1.name] = exta1;
    }
    if (this.fieldsMask.extraField2) {
      obj[this.fieldsMask.extraField2.name] = exta2;
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

    const args = [this.id, index, name];
    if (value !== undefined) {
      args.push(value);
    }
    if (this.fieldsMask.extraField1) {
      args.push(extraField1);
    }
    if (this.fieldsMask.extraField2) {
      args.push(extraField2);
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
    const arr = Object.values(obj).sort(
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

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (!this.dispatchersCallbacks || !this.dispatchersCallbacks['dropActionItem']) {
      return;
    }
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }

    this.dispatchersCallbacks['dropActionItem'](event, parent);
  }
}
