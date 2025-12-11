import {Component, DestroyRef, inject, computed, effect} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {Iitem, Iunicall, State} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {
  AddUnicallSpan,
  AddUnicallSpanParameter,
  AddUnicallParameter,
  DelUnicallSpan, DelUnicallSpanParameter, DelUnicallParameter, GetUnicallSpanParameters,
  UpdateUnicallSpan, StoreDropNewUnicallSpanParameter, StoreDropNewUnicallParameter, StoreNewUnicallSpanParameter, StoreNewUnicallParameter,
  StorePasteUnicallSpanParameters, SwitchUnicallSpanParameter,
  SwitchUnicallParameter, UpdateUnicallSpanParameter,
  UpdateUnicallParameter
} from '../../../store/config/unicall/config.actions.unicall';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-unicall',
  templateUrl: './unicall.component.html',
  styleUrls: ['./unicall.component.css']
})
export class UnicallComponent {

  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);

  private configsObservable = this.store.pipe(select(selectConfigurationState));
  private configsSignal = toSignal(this.configsObservable, { initialValue: {} as State });

  public list = computed(() => this.configsSignal().unicall || {} as Iunicall);
  public loadCounter = computed(() => this.configsSignal().loadCounter || 0);
  private lastErrorMessage = computed(() => this.list().errorMessage || null);

  public newSpanId: string = '';
  public selectedIndex: number = 0;
  public spanId: number = 0;
  public panelCloser: any = {};
  public toCopySpan: number = 0;

  public globalSettingsDispatchers: object;
  public spanSettingsDispatchers: object;

  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    } else {
      // Logic from ngOnInit subscription: Reset newSpanId and ensure spanId is valid
      this.newSpanId = '';
      const spans = this.list().spans;
      if (spans && spans[this.spanId]) {
        // If spanId is still valid, keep it. Otherwise, default to 0.
      } else {
        this.spanId = 0;
      }
    }
  });

  constructor() {
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewUnicallParam.bind(this),
      switchItem: this.switchUnicallParam.bind(this),
      addItem: this.newUnicallParam.bind(this),
      dropNewItem: this.dropNewUnicallParam.bind(this),
      deleteItem: this.deleteUnicallParam.bind(this),
      updateItem: this.updateUnicallParam.bind(this),
      pasteItems: null,
    };
    this.spanSettingsDispatchers = {
      addNewItemField: this.addNewSpanParam.bind(this),
      switchItem: this.switchSpanParam.bind(this),
      addItem: this.newSpanParam.bind(this),
      dropNewItem: this.dropNewSpanParam.bind(this),
      deleteItem: this.deleteSpanParam.bind(this),
      updateItem: this.updateSpanParam.bind(this),
      pasteItems: this.pasteSpanParams.bind(this),
    };
  }

  updateUnicallParam(param: Iitem) {
    this.store.dispatch(new UpdateUnicallParameter({param: param}));
  }

  switchUnicallParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchUnicallParameter({param: newParam}));
  }

  newUnicallParam(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddUnicallParameter({index: index, param: param}));
  }

  deleteUnicallParam(param: Iitem) {
    this.store.dispatch(new DelUnicallParameter({param: param}));
  }

  addNewUnicallParam() {
    this.store.dispatch(new StoreNewUnicallParameter(null));
  }

  dropNewUnicallParam(index: number) {
    this.store.dispatch(new StoreDropNewUnicallParameter({index: index}));
  }

  getUnicallSpansParams(id: number) {
    this.panelCloser['span' + id] = true;
    this.store.dispatch(new GetUnicallSpanParameters({id: id}));
  }

  updateSpanParam(param: Iitem) {
    this.store.dispatch(new UpdateUnicallSpanParameter({param: param}));
  }

  switchSpanParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchUnicallSpanParameter({param: newParam}));
  }

  newSpanParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddUnicallSpanParameter({id: parentId, index: index, param: param}));
  }

  deleteSpanParam(param: Iitem) {
    this.store.dispatch(new DelUnicallSpanParameter({param: param}));
  }

  addNewSpanParam(parentId: number) {
    this.store.dispatch(new StoreNewUnicallSpanParameter({id: parentId}));
  }

  dropNewSpanParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewUnicallSpanParameter({id: parentId, index: index}));
  }

  onSpanSubmit() {
    this.store.dispatch(new AddUnicallSpan({name: this.newSpanId}));
  }

  checkDirty(condition: AbstractControl | null): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isReadyToSendThree(mainObject: AbstractControl | null, object2: AbstractControl | null, object3: AbstractControl | null): boolean {
    return (mainObject && mainObject.valid && mainObject.dirty)
      || ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty));
  }

  isvalueReadyToSend(valueObject: AbstractControl | null): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  isReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index: number, item: any) {
    return index; // or item.id
  }

  isNewReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  copySpan(key: number) {
    const spans = this.list().spans;
    if (!spans || !spans[key]) {
      this.toCopySpan = 0;
      return;
    }
    this.toCopySpan = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteSpanParams(to: number) {
    this.store.dispatch(new StorePasteUnicallSpanParameters({from_id: this.toCopySpan, to_id: to}));
  }

  openBottomSheetSpan(id: number, newName: string, oldName: string, action: 'delete' | 'rename'): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete span "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename span "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelUnicallSpan({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateUnicallSpan({id: id, name: newName}));
      }
    });
  }

  onlyValues(obj: object | null): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}
