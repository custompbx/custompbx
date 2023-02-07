import {Component, Input, OnInit} from '@angular/core';
import {AbstractControl} from '@angular/forms';
import {Store} from '@ngrx/store';

@Component({
  selector: 'app-key-value-pad',
  templateUrl: './key-value-pad.component.html',
  styleUrls: ['./key-value-pad.component.css']
})
export class KeyValuePadComponent implements OnInit {

  @Input() items: object;
  @Input() newItems: Array<any>;
  @Input() exist: boolean;
  @Input() id: number;
  @Input() toCopy: number;
  @Input() store: Store<any>;
  @Input() dispatchers: any;

  constructor() { }

  ngOnInit() {
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFnId(index, item) {
    if (item.id) {
      return item.id;
    }
  }

  addItemField() {
    if (!this.dispatchers || !this.dispatchers['addItemField']) {
      return;
    }
    this.store.dispatch(new this.dispatchers['addItemField']({id: this.id}));
  }

  switchItem(object) {
    if (!this.dispatchers || !this.dispatchers['switchItem']) {
      return;
    }
    this.store.dispatch(new this.dispatchers['switchItem']({id: object.id, enabled: !object.enabled}));
  }

  newItem(index: number, name: string, value: string) {
    if (!this.dispatchers || !this.dispatchers['newItem']) {
      return;
    }
    this.store.dispatch(new this.dispatchers['newItem']({id: this.id, index: index, name: name, value: value}));
  }

  dropNewItem(index: number) {
    if (!this.dispatchers || !this.dispatchers['dropNewItem']) {
      return;
    }
    this.store.dispatch(new this.dispatchers['dropNewItem']({id: this.id, index: index}));
  }

  deleteItem(index: number) {
    if (!this.dispatchers || !this.dispatchers['deleteItem']) {
      return;
    }
    this.store.dispatch(new this.dispatchers['deleteItem']({id: this.id, index: index}));
  }

  updateItem(id: number, name: string, value: string) {
    if (!this.dispatchers || !this.dispatchers['updateItem']) {
      return;
    }
    this.store.dispatch(new this.dispatchers['updateItem']({id: id, name: name, value: value}));
  }

  pasteItems() {
    if (!this.dispatchers || !this.dispatchers['pasteItems']) {
      return;
    }
    this.store.dispatch(new this.dispatchers['pasteItems']({from_id: this.toCopy, to_id: this.id}));
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}
