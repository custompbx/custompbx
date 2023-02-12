import {ChangeDetectionStrategy, Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {IglobalVariable, IglobalVariables} from '../../store/global-variables/global-variables.reducer';
import {select, Store} from '@ngrx/store';
import {AppState, selectGlobalVariablesState} from '../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
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
import {CdkDragDrop} from '@angular/cdk/drag-drop';


@Component({
  selector: 'app-global-variables',
  templateUrl: './global-variables.component.html',
  styleUrls: ['./global-variables.component.css'],
  // changeDetection: ChangeDetectionStrategy.OnPush,
})
export class GlobalVariablesComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: IglobalVariables;
  public newList: Array<IglobalVariable>;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public globalSettingsDispatchers: object;
  public variableTypes = ['set', 'exec-set', 'env-set', 'stun-set'];

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.configs = this.store.pipe(select(selectGlobalVariablesState));
  }

  ngOnInit() {
    this.configs$ = this.configs.subscribe((globalVars) => {
      this.loadCounter = globalVars.loadCounter;
      this.list = globalVars.globalVariables;
      this.newList = globalVars.newGlobalVariables;
      this.lastErrorMessage = globalVars.globalVariables && globalVars.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
/*    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewGlobalVariable.bind(this),
      switchItem: this.switchGlobalVariable.bind(this),
      addItem: this.newGlobalVariable.bind(this),
      dropNewItem: this.dropNewGlobalVariable.bind(this),
      deleteItem: this.deleteGlobalVariable.bind(this),
      updateItem: this.updateGlobalVariable.bind(this),
      pasteItems: null,
    };*/
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
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

  trackByFn(index, item) {
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

  trackByIdFn(index, item) {
    return item.id;
  }

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }
    this.store.dispatch(new MoveGlobalVariable({
      previous_index: parent[event.previousIndex].position,
      current_index: parent[event.currentIndex].position,
      id: parent[event.previousIndex].id
    }));
  }

}

