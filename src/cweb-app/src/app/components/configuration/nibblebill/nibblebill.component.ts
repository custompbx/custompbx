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
  DelNibblebillParameter,
  AddNibblebillParameter,
  StoreNewNibblebillParameter,
  StoreDropNewNibblebillParameter,
  SwitchNibblebillParameter,
  UpdateNibblebillParameter
} from '../../../store/config/nibblebill/config.actions.nibblebill';

@Component({
  selector: 'app-nibblebill',
  templateUrl: './nibblebill.component.html',
  styleUrls: ['./nibblebill.component.css']
})
export class NibblebillComponent implements OnInit, OnDestroy {

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
      this.list = configs.nibblebill;
      this.lastErrorMessage = configs.nibblebill && configs.nibblebill.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewNibblebillParam.bind(this),
      switchItem: this.switchNibblebillParam.bind(this),
      addItem: this.newNibblebillParam.bind(this),
      dropNewItem: this.dropNewNibblebillParam.bind(this),
      deleteItem: this.deleteNibblebillParam.bind(this),
      updateItem: this.updateNibblebillParam.bind(this),
      pasteItems: null,
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updateNibblebillParam(param: Iitem) {
    this.store.dispatch(new UpdateNibblebillParameter({param: param}));
  }

  switchNibblebillParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchNibblebillParameter({param: newParam}));
  }

  newNibblebillParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddNibblebillParameter({index: index, param: param}));
  }

  deleteNibblebillParam(param: Iitem) {
    this.store.dispatch(new DelNibblebillParameter({param: param}));
  }

  addNewNibblebillParam() {
    this.store.dispatch(new StoreNewNibblebillParameter(null));
  }

  dropNewNibblebillParam(index: number) {
    this.store.dispatch(new StoreDropNewNibblebillParameter({index: index}));
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
