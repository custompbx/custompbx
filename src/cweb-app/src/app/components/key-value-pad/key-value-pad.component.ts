import {Component, Input} from '@angular/core';
import {Store} from '@ngrx/store';
import {KeyValuePad2Component} from '../key-value-pad-2/key-value-pad-2.component';

@Component({
  standalone: true,
  imports: [KeyValuePad2Component],
  selector: 'app-key-value-pad',
  templateUrl: './key-value-pad.component.html',
  styleUrls: ['./key-value-pad.component.css'],
})
export class KeyValuePadComponent {
  @Input() items: object;
  @Input() newItems: Array<any>;
  @Input() exist: boolean;
  @Input() id: number;
  @Input() toCopy: number;
  @Input() store: Store<any>;
  @Input() dispatchers: any;

  readonly fieldsMask = {name: {name: 'name'}, value: {name: 'value'}};

  readonly dispatchersCallbacks = {
    addNewItemField: () => {
      if (this.dispatchers?.addItemField) {
        this.store.dispatch(new this.dispatchers.addItemField({id: this.id}));
      }
    },
    switchItem: (item: any) => {
      if (this.dispatchers?.switchItem) {
        this.store.dispatch(new this.dispatchers.switchItem({id: item.id, enabled: !item.enabled}));
      }
    },
    addItem: (...args: any[]) => {
      if (!this.dispatchers?.newItem) {
        return;
      }
      const index = args.length >= 4 ? args[1] : args[0];
      const name = args.length >= 4 ? args[2] : args[1];
      const value = args.length >= 4 ? args[3] : args[2];
      this.store.dispatch(new this.dispatchers.newItem({id: this.id, index, name, value}));
    },
    dropNewItem: (...args: any[]) => {
      if (!this.dispatchers?.dropNewItem) {
        return;
      }
      const index = args.length >= 2 ? args[1] : args[0];
      this.store.dispatch(new this.dispatchers.dropNewItem({id: this.id, index}));
    },
    deleteItem: (item: any) => {
      if (this.dispatchers?.deleteItem) {
        this.store.dispatch(new this.dispatchers.deleteItem({id: this.id, index: item.id}));
      }
    },
    updateItem: (item: any) => {
      if (this.dispatchers?.updateItem) {
        this.store.dispatch(new this.dispatchers.updateItem({id: item.id, name: item.name, value: item.value}));
      }
    },
    pasteItems: () => {
      if (this.dispatchers?.pasteItems) {
        this.store.dispatch(new this.dispatchers.pasteItems({from_id: this.toCopy, to_id: this.id}));
      }
    },
  };
}
