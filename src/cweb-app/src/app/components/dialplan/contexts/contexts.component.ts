import {Component, inject, signal, computed, effect, OnInit, Pipe, PipeTransform} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {select, Store} from '@ngrx/store';
import {AppState, selectDialplanState} from '../../../store/app.states';
import {Iaction, Iantiaction, Icondition, Icontexts, Idebug, Iextension, Iregex} from '../../../store/dialplan/dialplan.reducers';
import {
  AddContext,
  AddExtension,
  DeleteAction,
  DeleteAntiaction,
  DeleteCondition,
  DeleteContext,
  DeleteExtension,
  GetExtensionDetails,
  GetConditions,
  GetExtensions,
  MoveAction,
  MoveAntiaction,
  MoveCondition,
  MoveExtension,
  RenameContext,
  RenameExtension,
  SwitchAction,
  SwitchAntiaction,
  SwitchCondition,
  UpdateAction,
  UpdateAntiaction,
  UpdateCondition,
  UpdateRegex,
  SwitchRegex,
  DeleteRegex,
  AddNewAction,
  AddNewAntiaction,
  AddNewRegex,
  DeleteNewAction,
  DeleteNewAntiaction,
  DeleteNewRegex,
  AddRegex,
  AddAction,
  AddAntiaction,
  AddCondition,
  SwitchExtensionContinue,
  ImportDialplan,
  DialplanDebug, SwitchDialplanDebug, StoreClearDialplanDebug, DialplanSettings, SwitchDialplanStatic
} from '../../../store/dialplan/dialplan.actions';
import {CdkDragDrop} from '@angular/cdk/drag-drop';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {ActivatedRoute} from '@angular/router';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ResizeInputDirective} from "../../../directives/resize-input.directive";

@Pipe({
  name: 'objectDataToName',
  standalone: true // Making the pipe standalone
})
export class ObjectToNamePipe implements PipeTransform {
  transform(value: object): string {
    const keys = Object.keys(value);
    let result = '';
    keys.forEach(
      (key) => {
        // Exclude internal fields and empty strings
        if (typeof value[key] !== 'string' || value[key] === '' || key === 'id' || key === 'position' || key === 'enabled') {
          return;
        }
        result += key + '=' + value[key] + ' ';
      }
    );

    return result.trim() || 'no conditions';
  }
}

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ObjectToNamePipe, ResizeInputDirective], // Include the new standalone pipe
  selector: 'app-contexts',
  templateUrl: './contexts.component.html',
  styleUrls: ['./contexts.component.css']
})
export class ContextsComponent { // Removed OnDestroy

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private dialplanState = toSignal(
    this.store.pipe(select(selectDialplanState)),
    {
      initialValue: {
        contexts: {} as Icontexts,
        debug: {} as Idebug,
        staticDialplan: false,
        errorMessage: null,
        loadCounter: 0,
      }
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.dialplanState().contexts);
  public debug = computed(() => this.dialplanState().debug);
  public staticDialplan = computed(() => this.dialplanState().staticDialplan);
  public loadCounter = computed(() => this.dialplanState().loadCounter);
  public lastErrorMessage = computed(() => this.dialplanState().errorMessage);

  // --- Local Component State as Signals/Properties ---
  public selectedIndex = signal<number>(0);

  // Properties used for forms (kept as properties for two-way binding with ngModel)
  public newContextName: string = '';
  public newExtensionName: string = '';
  public newContextId: number | null = null;
  public contextId: number | null = null;

  private expanded = []; // This should probably be a signal if it controls UI state

  protected inlineActions = {
    'check_acl': true,
    'eval': true,
    'event': true,
    'export': true,
    'log': true,
    'presence': true,
    'set': true,
    'set_global': true,
    'set_profile_var': true,
    'set_user': true,
    'unset': true,
    'verbose_events': true,
    'cidlookup': true,
    'curl': true,
    'easyroute': true,
    'enum': true,
    'lcr': true,
    'nibblebill': true,
    'odbc_query': true,
  };

  // --- Effect for Side Effects (Error handling) ---
  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  dropExtension(event: CdkDragDrop<string[]>, parent: Array<any>) {
    const previousItem = parent[event.previousIndex];
    const currentItem = parent[event.currentIndex];

    if (!previousItem || !currentItem || previousItem.position === currentItem.position) {
      return;
    }

    this.store.dispatch(new MoveExtension({
      previous_index: previousItem.position,
      current_index: currentItem.position,
      id: previousItem.id
    }));
  }

  dropCondition(event: CdkDragDrop<string[]>, parent: Array<any>) {
    const previousItem = parent[event.previousIndex];
    const currentItem = parent[event.currentIndex];

    if (!previousItem || !currentItem || previousItem.position === currentItem.position) {
      return;
    }
    this.store.dispatch(new MoveCondition({
      previous_index: previousItem.position,
      current_index: currentItem.position,
      id: previousItem.id
    }));
  }

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    const previousItem = parent[event.previousIndex];
    const currentItem = parent[event.currentIndex];

    if (!previousItem || !currentItem || previousItem.position === currentItem.position) {
      return;
    }
    this.store.dispatch(new MoveAction({
      previous_index: previousItem.position,
      current_index: currentItem.position,
      id: previousItem.id
    }));
  }

  dropAntiaction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    const previousItem = parent[event.previousIndex];
    const currentItem = parent[event.currentIndex];

    if (!previousItem || !currentItem || previousItem.position === currentItem.position) {
      return;
    }
    this.store.dispatch(new MoveAntiaction({
      previous_index: previousItem.position,
      current_index: currentItem.position,
      id: previousItem.id
    }));
  }

  mainTabChanged(event: number) {
    this.selectedIndex.set(event);
    if (event === 3) { // Assuming index 3 is for Debug/Settings
      this.store.dispatch(new DialplanDebug({keep_subscription: true}));
      this.store.dispatch(new DialplanSettings(null));
    }
  }

  trackByFn(index: number, item: any): number {
    return index;
  }

  trackById(index: number, item: any) {
    return item.id;
  }

  isReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    // Check if both objects exist, and if either is dirty, and both are valid
    return !!(nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid);
  }

  isReadyToSendAction(nameObject: AbstractControl | null, valueObject: AbstractControl | null, inlineObject: AbstractControl | null): boolean {
    if (inlineObject) {
      // If inline is an input, check if any of the three are dirty and valid
      return (nameObject && nameObject.valid && nameObject.dirty)
        || ((valueObject && valueObject.valid && valueObject.dirty) || (inlineObject && inlineObject.valid && inlineObject.dirty));
    }
    // If inline is not a control, revert to standard name/value check
    return !!(nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid);
  }

  checkDirty(condition: AbstractControl | null): boolean {
    // Returns true if the control is NOT dirty, or if the control is null
    return !condition || !condition.dirty;
  }

  getExtensions(id: number) {
    this.store.dispatch(new GetExtensions({id: id}));
  }

  selectContext(event: any) {
    this.store.dispatch(new GetExtensions({id: event.value}));
  }

  getConditions(id: number) {
    this.store.dispatch(new GetConditions({id: id}));
  }

  getActions(id: number) {
    this.store.dispatch(new GetExtensionDetails({id: id}));
  }

  switchContinue(object: Iextension) {
    this.store.dispatch(new SwitchExtensionContinue({id: object.id, value: object.continue === 'true' ? '' : 'true'}));
  }

  addCondition(id: number) {
    this.store.dispatch(new AddCondition({id: id}));
  }

  updateCondition(object: Icondition) {
    this.store.dispatch(new UpdateCondition({condition: object}));
  }

  switchCondition(object: Icondition) {
    this.store.dispatch(new SwitchCondition({condition: {...object, enabled: !object.enabled}}));
  }

  deleteCondition(object: Icondition) {
    this.openBottomSheetCondition(object, 'condition');
  }

  updateRegex(object: Iregex) {
    this.store.dispatch(new UpdateRegex({regex: object}));
  }

  switchRegex(object: Iregex) {
    this.store.dispatch(new SwitchRegex({regex: {...object, enabled: !object.enabled}}));
  }

  deleteRegex(object: Iregex) {
    this.openBottomSheetCondition(object, 'regex');
  }

  addRegex(contextId: number, extensionId: number, conditionId: number, index: number, object: Iregex) {
    this.store.dispatch(new AddRegex({id: conditionId, index: index, regex: object, contextId: contextId, extensionId: extensionId}));
  }

  addNewRegex(contextId: number, extensionId: number, conditionId: number) {
    this.store.dispatch(new AddNewRegex({contextId: contextId, extensionId: extensionId, conditionId: conditionId}));
  }

  delNewRegex(index: number, contextId: number, extensionId: number, conditionId: number) {
    this.store.dispatch(new DeleteNewRegex({contextId: contextId, extensionId: extensionId, conditionId: conditionId, index: index}));
  }

  updateAction(object: Iaction) {
    // Convert 'inline' binding from string/boolean to true boolean
    const inlineValue = typeof object.inline === 'string' ? object.inline.toLowerCase() === 'true' : !!object.inline;
    this.store.dispatch(new UpdateAction({action: {...object, inline: inlineValue}}));
  }

  switchAction(object: Iaction) {
    this.store.dispatch(new SwitchAction({action: {...object, enabled: !object.enabled}}));
  }

  deleteAction(object: Iaction) {
    this.openBottomSheetCondition(object, 'action');
  }

  addAction(contextId: number, extensionId: number, conditionId: number, index: number, object: Iaction) {
    // Convert 'inline' binding from string/boolean to true boolean
    const inlineValue = typeof object.inline === 'string' ? object.inline.toLowerCase() === 'true' : !!object.inline;
    this.store.dispatch(new AddAction(
      {id: conditionId, index: index, action: {...object, inline: inlineValue},
        contextId: contextId, extensionId: extensionId}));
  }

  addNewAction(contextId: number, extensionId: number, conditionId: number) {
    this.store.dispatch(new AddNewAction({contextId: contextId, extensionId: extensionId, conditionId: conditionId}));
  }

  delNewAction(index: number, contextId: number, extensionId: number, conditionId: number) {
    this.store.dispatch(new DeleteNewAction({contextId: contextId, extensionId: extensionId, conditionId: conditionId, index: index}));
  }

  updateAntiaction(object: Iantiaction) {
    this.store.dispatch(new UpdateAntiaction({antiaction: object}));
  }

  switchAntiaction(object: Iantiaction) {
    this.store.dispatch(new SwitchAntiaction({antiaction: {...object, enabled: !object.enabled}}));
  }

  deleteAntiaction(object: Iantiaction) {
    this.openBottomSheetCondition(object, 'antiaction');
  }

  addAntiaction(contextId: number, extensionId: number, conditionId: number, index: number, object: Iantiaction) {
    this.store.dispatch(new AddAntiaction(
      {id: conditionId, index: index, antiaction: object, contextId: contextId, extensionId: extensionId}
    ));
  }

  addNewAntiaction(contextId: number, extensionId: number, conditionId: number) {
    this.store.dispatch(new AddNewAntiaction({contextId: contextId, extensionId: extensionId, conditionId: conditionId}));
  }

  delNewAntiaction(index: number, contextId: number, extensionId: number, conditionId: number) {
    this.store.dispatch(new DeleteNewAntiaction({contextId: contextId, extensionId: extensionId, conditionId: conditionId, index: index}));
  }

  importDialplan() {
    this.store.dispatch(new ImportDialplan(null));
  }

  onContextSubmit() {
    if (this.newContextName.trim()) {
      this.store.dispatch(new AddContext({name: this.newContextName.trim()}));
      this.newContextName = ''; // Clear input field
    }
  }

  onExtensionSubmit() {
    if (this.newExtensionName.trim() && this.newContextId !== null) {
      this.store.dispatch(new AddExtension({name: this.newExtensionName.trim(), id: this.newContextId}));
      this.newExtensionName = ''; // Clear input field
      this.newContextId = null;
    }
  }

  switchDebug() {
    // Read current state from the signal
    const currentDebug = this.debug();
    this.store.dispatch(new SwitchDialplanDebug({enabled: !currentDebug.enabled}));
  }

  switchNoProceed() {
    // Read current state from the signal
    const currentStatic = this.staticDialplan();
    this.store.dispatch(new SwitchDialplanStatic({enabled: !currentStatic}));
  }

  clearDebug() {
    this.store.dispatch(new StoreClearDialplanDebug(null));
  }

  openBottomSheetContext(id: number, newName: string, oldName: string, action: 'delete' | 'rename'): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: action === 'delete' ? 'Are you sure you want to delete context "' + oldName + '"?' : null,
          case2Text: action === 'rename' ? 'Are you sure you want to rename context "' + oldName + '" to "' + newName + '"?' : null,
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DeleteContext({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameContext({id: id, name: newName}));
      }
    });
  }

  openBottomSheetExtension(id: number, newName: string, oldName: string, action: 'delete' | 'rename'): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: action === 'delete' ? 'Are you sure you want to delete extension "' + oldName + '"?' : null,
          case2Text: action === 'rename' ? 'Are you sure you want to rename extension "' + oldName + '" to "' + newName + '"?' : null,
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DeleteExtension({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameExtension({id: id, name: newName}));
      }
    });
  }

  openBottomSheetCondition(object: Icondition | Iregex | Iaction | Iantiaction, type: string): void {
    const config = {
      data:
        {
          action: 'delete',
          object: object,
          case1Text: `Are you sure you want to delete this ${type}?`,
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      switch (type) {
        case 'condition':
          this.store.dispatch(new DeleteCondition({condition: object as Icondition}));
          break;
        case 'regex':
          this.store.dispatch(new DeleteRegex({regex: object as Iregex}));
          break;
        case 'action':
          this.store.dispatch(new DeleteAction({action: object as Iaction}));
          break;
        case 'antiaction':
          this.store.dispatch(new DeleteAntiaction({antiaction: object as Iantiaction}));
          break;
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
