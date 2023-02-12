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
  DelTtsCommandlineParameter,
  AddTtsCommandlineParameter,
  StoreNewTtsCommandlineParameter,
  StoreDropNewTtsCommandlineParameter,
  SwitchTtsCommandlineParameter,
  UpdateTtsCommandlineParameter
} from '../../../store/config/tts_commandline/config.actions.tts_commandline';

@Component({
  selector: 'app-tts-commandline',
  templateUrl: './tts-commandline.component.html',
  styleUrls: ['./tts-commandline.component.css']
})
export class TtsCommandlineComponent implements OnInit, OnDestroy {

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
      this.list = configs.tts_commandline;
      this.lastErrorMessage = configs.tts_commandline && configs.tts_commandline.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewTtsCommandlineParam.bind(this),
      switchItem: this.switchTtsCommandlineParam.bind(this),
      addItem: this.newTtsCommandlineParam.bind(this),
      dropNewItem: this.dropNewTtsCommandlineParam.bind(this),
      deleteItem: this.deleteTtsCommandlineParam.bind(this),
      updateItem: this.updateTtsCommandlineParam.bind(this),
      pasteItems: null,
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateTtsCommandlineParam(param: Iitem) {
    this.store.dispatch(new UpdateTtsCommandlineParameter({param: param}));
  }

  switchTtsCommandlineParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchTtsCommandlineParameter({param: newParam}));
  }

  newTtsCommandlineParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddTtsCommandlineParameter({index: index, param: param}));
  }

  deleteTtsCommandlineParam(param: Iitem) {
    this.store.dispatch(new DelTtsCommandlineParameter({param: param}));
  }

  addNewTtsCommandlineParam() {
    this.store.dispatch(new StoreNewTtsCommandlineParameter(null));
  }

  dropNewTtsCommandlineParam(index: number) {
    this.store.dispatch(new StoreDropNewTtsCommandlineParameter({index: index}));
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
