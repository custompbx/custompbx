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
  DelCepstralParameter,
  AddCepstralParameter,
  StoreNewCepstralParameter,
  StoreDropNewCepstralParameter,
  SwitchCepstralParameter,
  UpdateCepstralParameter
} from '../../../store/config/cepstral/config.actions.cepstral';

@Component({
  selector: 'app-cepstral',
  templateUrl: './cepstral.component.html',
  styleUrls: ['./cepstral.component.css']
})
export class CepstralComponent implements OnInit, OnDestroy {

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
      this.list = configs.cepstral;
      this.lastErrorMessage = configs.cepstral && configs.cepstral.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewCepstralParam.bind(this),
      switchItem: this.switchCepstralParam.bind(this),
      addItem: this.newCepstralParam.bind(this),
      dropNewItem: this.dropNewCepstralParam.bind(this),
      deleteItem: this.deleteCepstralParam.bind(this),
      updateItem: this.updateCepstralParam.bind(this),
      pasteItems: null,
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateCepstralParam(param: Iitem) {
    this.store.dispatch(new UpdateCepstralParameter({param: param}));
  }

  switchCepstralParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchCepstralParameter({param: newParam}));
  }

  newCepstralParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddCepstralParameter({index: index, param: param}));
  }

  deleteCepstralParam(param: Iitem) {
    this.store.dispatch(new DelCepstralParameter({param: param}));
  }

  addNewCepstralParam() {
    this.store.dispatch(new StoreNewCepstralParameter(null));
  }

  dropNewCepstralParam(index: number) {
    this.store.dispatch(new StoreDropNewCepstralParameter({index: index}));
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

