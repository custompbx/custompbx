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
  DelSangomaCodecParameter,
  AddSangomaCodecParameter,
  StoreNewSangomaCodecParameter,
  StoreDropNewSangomaCodecParameter,
  SwitchSangomaCodecParameter,
  UpdateSangomaCodecParameter
} from '../../../store/config/sangoma_codec/config.actions.sangoma_codec';

@Component({
  selector: 'app-sangoma-codec',
  templateUrl: './sangoma-codec.component.html',
  styleUrls: ['./sangoma-codec.component.css']
})
export class SangomaCodecComponent implements OnInit, OnDestroy {

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
      this.list = configs.sangoma_codec;
      this.lastErrorMessage = configs.sangoma_codec && configs.sangoma_codec.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewSangomaCodecParam.bind(this),
      switchItem: this.switchSangomaCodecParam.bind(this),
      addItem: this.newSangomaCodecParam.bind(this),
      dropNewItem: this.dropNewSangomaCodecParam.bind(this),
      deleteItem: this.deleteSangomaCodecParam.bind(this),
      updateItem: this.updateSangomaCodecParam.bind(this),
      pasteItems: null,
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateSangomaCodecParam(param: Iitem) {
    this.store.dispatch(new UpdateSangomaCodecParameter({param: param}));
  }

  switchSangomaCodecParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchSangomaCodecParameter({param: newParam}));
  }

  newSangomaCodecParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddSangomaCodecParameter({index: index, param: param}));
  }

  deleteSangomaCodecParam(param: Iitem) {
    this.store.dispatch(new DelSangomaCodecParameter({param: param}));
  }

  addNewSangomaCodecParam() {
    this.store.dispatch(new StoreNewSangomaCodecParameter(null));
  }

  dropNewSangomaCodecParam(index: number) {
    this.store.dispatch(new StoreDropNewSangomaCodecParameter({index: index}));
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

