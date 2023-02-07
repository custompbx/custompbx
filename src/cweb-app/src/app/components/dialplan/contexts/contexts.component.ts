import {Component, OnDestroy, OnInit, Pipe, PipeTransform} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
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
import {AbstractControl} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-contexts',
  templateUrl: './contexts.component.html',
  styleUrls: ['./contexts.component.css']
})
export class ContextsComponent implements OnInit, OnDestroy, OnInit {

  public dialplan: Observable<any>;
  public dialplan$: Subscription;
  public list: Icontexts;
  public staticDialplan: boolean;
  private debug: Idebug;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  private newContextId: number;
  private contextId: number;
  private newContextName: string;
  private newExtensionName: string;
  private inlineActions = {
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
  private expanded = [];

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.dialplan = this.store.pipe(select(selectDialplanState));
  }

  ngOnInit() {
    this.dialplan$ = this.dialplan.subscribe((dialplan) => {
      this.loadCounter = dialplan.loadCounter;
      this.list = dialplan.contexts;
      this.debug = dialplan.debug;
      this.staticDialplan = dialplan.staticDialplan;
      this.lastErrorMessage = dialplan && dialplan.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.dialplan$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  dropExtension(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }
    this.store.dispatch(new MoveExtension({
      previous_index: parent[event.previousIndex].position,
      current_index: parent[event.currentIndex].position,
      id: parent[event.previousIndex].id
    }));
  }

  dropCondition(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }
    this.store.dispatch(new MoveCondition({
      previous_index: parent[event.previousIndex].position,
      current_index: parent[event.currentIndex].position,
      id: parent[event.previousIndex].id
    }));
  }

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }
    this.store.dispatch(new MoveAction({
      previous_index: parent[event.previousIndex].position,
      current_index: parent[event.currentIndex].position,
      id: parent[event.previousIndex].id
    }));
  }

  dropAntiaction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }
    this.store.dispatch(new MoveAntiaction({
      previous_index: parent[event.previousIndex].position,
      current_index: parent[event.currentIndex].position,
      id: parent[event.previousIndex].id
    }));
  }

  mainTabChanged(event) {
    if (event === 3) {
      this.store.dispatch(new DialplanDebug({keep_subscription: true}));
      this.store.dispatch(new DialplanSettings(null));
    }
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  trackById(index, item) {
    return item.id;
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isReadyToSendAction(nameObject: AbstractControl, valueObject: AbstractControl, inlineObject: AbstractControl): boolean {
    if (inlineObject) {
      return (nameObject && nameObject.valid && nameObject.dirty)
        || ((valueObject && valueObject.valid && valueObject.dirty) || (inlineObject && inlineObject.valid && inlineObject.dirty));
    }
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  getExtensions(id) {
    this.store.dispatch(new GetExtensions({id: id}));
  }

  selectContext(event) {
    this.store.dispatch(new GetExtensions({id: event.value}));
  }

  getConditions(id) {
    this.store.dispatch(new GetConditions({id: id}));
  }

  getActions(id) {
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
    this.store.dispatch(new DeleteCondition({condition: object}));
  }
  updateRegex(object: Iregex) {
    this.store.dispatch(new UpdateRegex({regex: object}));
  }
  switchRegex(object: Iregex) {
    this.store.dispatch(new SwitchRegex({regex: {...object, enabled: !object.enabled}}));
  }
  deleteRegex(object: Iregex) {
    this.store.dispatch(new DeleteRegex({regex: object}));
  }
  addRegex(contextId: number, extensionId: number, id: number, index: number, object: Iregex) {
    this.store.dispatch(new AddRegex({id: id, index: index, regex: object, contextId: contextId, extensionId: extensionId}));
  }
  addNewRegex(contextId: number, extensionId: number, conditionId: number) {
    this.store.dispatch(new AddNewRegex({contextId: contextId, extensionId: extensionId, conditionId: conditionId}));
  }
  delNewRegex(index: number, contextId: number, extensionId: number, conditionId: number) {
    this.store.dispatch(new DeleteNewRegex({contextId: contextId, extensionId: extensionId, conditionId: conditionId, index: index}));
  }
  updateAction(object: Iaction) {
    this.store.dispatch(new UpdateAction({action: {...object, inline: String(object.inline).toLowerCase() === 'true'}}));
  }
  switchAction(object: Iaction) {
    this.store.dispatch(new SwitchAction({action: {...object, enabled: !object.enabled}}));
  }
  deleteAction(object: Iaction) {
    this.store.dispatch(new DeleteAction({action: object}));
  }
  addAction(contextId: number, extensionId: number, id: number, index: number, object: Iaction) {
    this.store.dispatch(new AddAction(
      {id: id, index: index, action: {...object, inline: String(object.inline).toLowerCase() === 'true'},
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
    this.store.dispatch(new DeleteAntiaction({antiaction: object}));
  }
  addAntiaction(contextId: number, extensionId: number, id: number, index: number, object: Iantiaction) {
    this.store.dispatch(new AddAntiaction(
      {id: id, index: index, antiaction: object, contextId: contextId, extensionId: extensionId}
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
    this.store.dispatch(new AddContext({name: this.newContextName}));
  }

  onExtensionSubmit() {
    this.store.dispatch(new AddExtension({name: this.newExtensionName, id: this.newContextId}));
  }

  switchDebug() {
    this.store.dispatch(new SwitchDialplanDebug({enabled: !this.debug.enabled}));
  }

  switchNoProceed() {
    this.store.dispatch(new SwitchDialplanStatic({enabled: !this.staticDialplan}));
  }

  clearDebug() {
    this.store.dispatch(new StoreClearDialplanDebug(null));
  }

  openBottomSheetContext(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete context "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename context "' + oldName + '" to "' + newName + '"?',
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

  openBottomSheetExtension(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete extension "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename extension "' + oldName + '" to "' + newName + '"?',
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

  openBottomSheetCondition(object: Icondition): void {
    const config = {
      data:
        {
          action: 'delete',
          object: object,
          case1Text: 'Are you sure you want to delete this condition?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
        this.store.dispatch(new DeleteCondition({condition: object}));
    });
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}

@Pipe({
  name: 'objectDataToName'
})
export class ObjectToNamePipe implements PipeTransform {

  transform(value: object): string {
    const keys = Object.keys(value);
    let result = '';
    keys.forEach(
      (key) => {
        if (typeof value[key] !== 'string' || value[key] === '' || key === 'id' || key === 'position') {
          return;
        }
        result += key + '=' + value[key] + ' ';
      }
    );

    return result || 'no conditions';
  }

}
