import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Iitem, IsimpleModule} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  DelPostSwitchParameter,
  AddPostSwitchParameter,
  StoreNewPostSwitchParameter,
  StoreDropNewPostSwitchParameter,
  SwitchPostSwitchParameter,
  UpdatePostSwitchParameter,
  DelPostSwitchDefaultPtime,
  DelPostSwitchCliKeybinding,
  StoreNewPostSwitchCliKeybinding,
  SwitchPostSwitchDefaultPtime,
  StoreNewPostSwitchDefaultPtime,
  UpdatePostSwitchCliKeybinding,
  UpdatePostSwitchDefaultPtime,
  AddPostSwitchDefaultPtime,
  StoreDropNewPostSwitchDefaultPtime,
  StoreDropNewPostSwitchCliKeybinding,
  SwitchPostSwitchCliKeybinding,
  AddPostSwitchCliKeybinding
} from '../../../store/config/post-switch/config.actions.post-switch';

@Component({
  selector: 'app-post-load-switch',
  templateUrl: './post-load-switch.component.html',
  styleUrls: ['./post-load-switch.component.css']
})
export class PostLoadSwitchComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: IsimpleModule;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public globalSettingsDispatchers: object;
  public cliKeybindingsDispatchers: object;
  public defaultPtimesDispatchers: object;
  public defaultPtimesMask: object;

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.configs = this.store.pipe(select(selectConfigurationState));
  }

  ngOnInit() {
    this.configs$ = this.configs.subscribe((configs) => {
      this.loadCounter = configs.loadCounter;
      this.list = configs.post_load_switch;
      this.lastErrorMessage = configs.post_load_switch && configs.post_load_switch.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewPostSwitchParam.bind(this),
      switchItem: this.switchPostSwitchParam.bind(this),
      addItem: this.newPostSwitchParam.bind(this),
      dropNewItem: this.dropNewPostSwitchParam.bind(this),
      deleteItem: this.deletePostSwitchParam.bind(this),
      updateItem: this.updatePostSwitchParam.bind(this),
      pasteItems: null,
    };
    this.cliKeybindingsDispatchers = {
      addNewItemField: this.addNewPostSwitchCliKeybinding.bind(this),
      switchItem: this.switchPostSwitchCliKeybinding.bind(this),
      addItem: this.newPostSwitchCliKeybinding.bind(this),
      dropNewItem: this.dropNewPostSwitchCliKeybinding.bind(this),
      deleteItem: this.deletePostSwitchCliKeybinding.bind(this),
      updateItem: this.updatePostSwitchCliKeybinding.bind(this),
      pasteItems: null,
    };
    this.defaultPtimesDispatchers = {
      addNewItemField: this.addNewPostSwitchDefaultPtime.bind(this),
      switchItem: this.switchPostSwitchDefaultPtime.bind(this),
      addItem: this.newPostSwitchDefaultPtime.bind(this),
      dropNewItem: this.dropNewPostSwitchDefaultPtime.bind(this),
      deleteItem: this.deletePostSwitchDefaultPtime.bind(this),
      updateItem: this.updatePostSwitchDefaultPtime.bind(this),
      pasteItems: null,
    };
    this.defaultPtimesMask = {
      name: {name: 'codec_name'},
      value: {name: 'codec_ptime'},
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updatePostSwitchParam(param: Iitem) {
    this.store.dispatch(new UpdatePostSwitchParameter({param: param}));
  }

  switchPostSwitchParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchPostSwitchParameter({param: newParam}));
  }

  newPostSwitchParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddPostSwitchParameter({index: index, param: param}));
  }

  deletePostSwitchParam(param: Iitem) {
    this.store.dispatch(new DelPostSwitchParameter({param: param}));
  }

  addNewPostSwitchParam() {
    this.store.dispatch(new StoreNewPostSwitchParameter(null));
  }

  dropNewPostSwitchParam(index: number) {
    this.store.dispatch(new StoreDropNewPostSwitchParameter({index: index}));
  }

  updatePostSwitchCliKeybinding(param: Iitem) {
    this.store.dispatch(new UpdatePostSwitchCliKeybinding({param: param}));
  }

  switchPostSwitchCliKeybinding(param: Iitem) {
    const newCliKeybinding = <Iitem>{...param};
    newCliKeybinding.enabled = !newCliKeybinding.enabled;
    this.store.dispatch(new SwitchPostSwitchCliKeybinding({param: newCliKeybinding}));
  }

  newPostSwitchCliKeybinding(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddPostSwitchCliKeybinding({index: index, param: param}));
  }

  deletePostSwitchCliKeybinding(param: Iitem) {
    this.store.dispatch(new DelPostSwitchCliKeybinding({param: param}));
  }

  addNewPostSwitchCliKeybinding() {
    this.store.dispatch(new StoreNewPostSwitchCliKeybinding(null));
  }

  dropNewPostSwitchCliKeybinding(index: number) {
    this.store.dispatch(new StoreDropNewPostSwitchCliKeybinding({index: index}));
  }

  updatePostSwitchDefaultPtime(param) {
    const para = <Iitem>{id: param.id, name: param.codec_name, value: param.codec_ptime};
    this.store.dispatch(new UpdatePostSwitchDefaultPtime({param: para}));
  }

  switchPostSwitchDefaultPtime(param: Iitem) {
    const newDefaultPtime = <Iitem>{...param};
    newDefaultPtime.enabled = !newDefaultPtime.enabled;
    this.store.dispatch(new SwitchPostSwitchDefaultPtime({param: newDefaultPtime}));
  }

  newPostSwitchDefaultPtime(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddPostSwitchDefaultPtime({index: index, param: param}));
  }

  deletePostSwitchDefaultPtime(param: Iitem) {
    this.store.dispatch(new DelPostSwitchDefaultPtime({param: param}));
  }

  addNewPostSwitchDefaultPtime() {
    this.store.dispatch(new StoreNewPostSwitchDefaultPtime(null));
  }

  dropNewPostSwitchDefaultPtime(index: number) {
    this.store.dispatch(new StoreDropNewPostSwitchDefaultPtime({index: index}));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isReadyToSendThree(mainObject: AbstractControl, object2: AbstractControl, object3: AbstractControl): boolean {
    return (mainObject && mainObject.valid && mainObject.dirty)
      || ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty));
  }

  isvalueReadyToSend(valueObject: AbstractControl): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

}

