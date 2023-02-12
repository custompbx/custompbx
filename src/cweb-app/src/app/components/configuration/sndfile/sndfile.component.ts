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
  DelSndfileParameter,
  AddSndfileParameter,
  StoreNewSndfileParameter,
  StoreDropNewSndfileParameter,
  SwitchSndfileParameter,
  UpdateSndfileParameter
} from '../../../store/config/sndfile/config.actions.sndfile';

@Component({
  selector: 'app-sndfile',
  templateUrl: './sndfile.component.html',
  styleUrls: ['./sndfile.component.css']
})
export class SndfileComponent implements OnInit, OnDestroy {

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
      this.list = configs.sndfile;
      this.lastErrorMessage = configs.sndfile && configs.sndfile.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewSndfileParam.bind(this),
      switchItem: this.switchSndfileParam.bind(this),
      addItem: this.newSndfileParam.bind(this),
      dropNewItem: this.dropNewSndfileParam.bind(this),
      deleteItem: this.deleteSndfileParam.bind(this),
      updateItem: this.updateSndfileParam.bind(this),
      pasteItems: null,
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateSndfileParam(param: Iitem) {
    this.store.dispatch(new UpdateSndfileParameter({param: param}));
  }

  switchSndfileParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchSndfileParameter({param: newParam}));
  }

  newSndfileParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddSndfileParameter({index: index, param: param}));
  }

  deleteSndfileParam(param: Iitem) {
    this.store.dispatch(new DelSndfileParameter({param: param}));
  }

  addNewSndfileParam() {
    this.store.dispatch(new StoreNewSndfileParameter(null));
  }

  dropNewSndfileParam(index: number) {
    this.store.dispatch(new StoreDropNewSndfileParameter({index: index}));
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

