import {ChangeDetectionStrategy, Component, inject, signal, computed, effect, OnInit} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';

import {IglobalVariable, IglobalVariables} from '../../store/global-variables/global-variables.reducer';
import {select, Store} from '@ngrx/store';
import {AppState, selectGlobalVariablesState} from '../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ToastService} from '../../services/toast.service';
import {ActivatedRoute} from '@angular/router';
import {
  DelGlobalVariable,
  AddGlobalVariable,
  StoreNewGlobalVariable,
  StoreDropNewGlobalVariable,
  SwitchGlobalVariable,
  UpdateGlobalVariable,
  ImportGlobalVariables,
  MoveGlobalVariable
} from '../../store/global-variables/global-variables.actions';
import {CdkDragDrop, DragDropModule} from '@angular/cdk/drag-drop';
import {CommonModule} from "@angular/common";
import {InnerHeaderComponent} from "../inner-header/inner-header.component";
import {ResizeInputDirective} from "../../directives/resize-input.directive";
import {CpbxSelectDirective} from '../../directives/cpbx-select.directive';
import {resolvePositionedReorder} from '../../utils/reorder';


@Component({
  standalone: true,
  imports: [CommonModule, DragDropModule, FormsModule, InnerHeaderComponent, ResizeInputDirective, CpbxSelectDirective],
  selector: 'app-global-variables',
  templateUrl: './global-variables.component.html',
  styleUrls: ['./global-variables.component.css'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class GlobalVariablesComponent implements OnInit { // Removed OnDestroy

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private _snackBar = inject(ToastService);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private globalVarsState = toSignal(
    this.store.pipe(select(selectGlobalVariablesState)),
    {
      initialValue: {
        globalVariables: {} as IglobalVariables,
        newGlobalVariables: [],
        errorMessage: null,
        loadCounter: 0,
      }
    }
  );

  // --- Computed State from NgRx State ---
  public list = computed(() => this.globalVarsState().globalVariables);
  public newList = computed(() => this.globalVarsState().newGlobalVariables);
  public loadCounter = computed(() => this.globalVarsState().loadCounter);
  public lastErrorMessage = computed(() => this.globalVarsState().errorMessage);

  // --- Local State as Signal ---
  public selectedIndex = signal<number>(0);

  // --- Constant Properties ---
  public globalSettingsDispatchers: object;
  public variableTypes = ['set', 'exec-set', 'env-set', 'stun-set'];

  // --- Effect for Side Effects (Replaces subscription error handling) ---
  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  ngOnInit() {
    // The subscription logic is handled by toSignal and the effect.

    // The following commented out block is preserved if needed later, but the subscription
    // cleanup in ngOnDestroy is no longer necessary.

    /*
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewGlobalVariable.bind(this),
      switchItem: this.switchGlobalVariable.bind(this),
      addItem: this.newGlobalVariable.bind(this),
      dropNewItem: this.dropNewGlobalVariable.bind(this),
      deleteItem: this.deleteGlobalVariable.bind(this),
      updateItem: this.updateGlobalVariable.bind(this),
      pasteItems: null,
    };
    */
  }

  updateGlobalVariable(variable: IglobalVariable) {
    this.store.dispatch(new UpdateGlobalVariable({variable: variable}));
  }

  switchGlobalVariable(variable: IglobalVariable) {
    const newParam = <IglobalVariable>{...variable};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchGlobalVariable({variable: newParam}));
  }

  newGlobalVariable(index: number, name: string, value: string, type: string) {
    const variable = <IglobalVariable>{};
    variable.enabled = true;
    variable.name = name;
    variable.value = value;
    variable.type = type;

    this.store.dispatch(new AddGlobalVariable({index: index, variable: variable}));
  }

  deleteGlobalVariable(variable: IglobalVariable) {
    this.store.dispatch(new DelGlobalVariable({variable: variable}));
  }

  addNewGlobalVariable() {
    this.store.dispatch(new StoreNewGlobalVariable(null));
  }

  dropNewGlobalVariable(index: number) {
    this.store.dispatch(new StoreDropNewGlobalVariable({index: index}));
  }

  ImportGlobalVariables() {
    this.store.dispatch(new ImportGlobalVariables(null));
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

  trackByFn(index: number, item: any): number {
    return index; // or item.id
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  ObjLength(item: object): number {
    return Object.keys(item).length;
  }

  onlySortedValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    const arr = Object.values(obj).sort(
      function (a, b) {
        if (a.position > b.position) {
          return 1;
        }
        if (a.position < b.position) {
          return -1;
        }
        return 0;
      }
    );
    return arr;
  }

  trackByIdFn(index: number, item: any) {
    return item.id;
  }

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    const change = resolvePositionedReorder(parent, event.previousIndex, event.currentIndex);
    if (change) this.store.dispatch(new MoveGlobalVariable(change.move));
  }

}
