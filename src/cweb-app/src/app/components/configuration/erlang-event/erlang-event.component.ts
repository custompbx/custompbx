import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Iitem, Ilcr, IsimpleModule} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  DelErlangEventParameter,
  AddErlangEventParameter,
  StoreNewErlangEventParameter,
  StoreDropNewErlangEventParameter,
  SwitchErlangEventParameter,
  UpdateErlangEventParameter
} from '../../../store/config/erlang_event/config.actions.erlang_event';

@Component({
  selector: 'app-erlang-event',
  templateUrl: './erlang-event.component.html',
  styleUrls: ['./erlang-event.component.css']
})
export class ErlangEventComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: IsimpleModule;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public globalSettingsDispatchers: object;

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
      this.list = configs.erlang_event;
      this.lastErrorMessage = configs.erlang_event && configs.erlang_event.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewErlangEventParam.bind(this),
      switchItem: this.switchErlangEventParam.bind(this),
      addItem: this.newErlangEventParam.bind(this),
      dropNewItem: this.dropNewErlangEventParam.bind(this),
      deleteItem: this.deleteErlangEventParam.bind(this),
      updateItem: this.updateErlangEventParam.bind(this),
      pasteItems: null,
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateErlangEventParam(param: Iitem) {
    this.store.dispatch(new UpdateErlangEventParameter({param: param}));
  }

  switchErlangEventParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchErlangEventParameter({param: newParam}));
  }

  newErlangEventParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddErlangEventParameter({index: index, param: param}));
  }

  deleteErlangEventParam(param: Iitem) {
    this.store.dispatch(new DelErlangEventParameter({param: param}));
  }

  addNewErlangEventParam() {
    this.store.dispatch(new StoreNewErlangEventParameter(null));
  }

  dropNewErlangEventParam(index: number) {
    this.store.dispatch(new StoreDropNewErlangEventParameter({index: index}));
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

