import { Component, OnDestroy, OnInit, ViewChild, inject, signal, computed, effect } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';

import { MaterialModule } from "../../../../material-module";
import { select, Store } from '@ngrx/store';
import { AppState, selectDirectoryState } from '../../../store/app.states';
import {
  GetDirectoryDomainDetails,
  StoreAddNewDirectoryDomainParameter,
  StoreAddNewDirectoryDomainVariable,
  AddDirectoryDomain,
  RenameDirectoryDomain,
  DeleteDirectoryDomain,
  AddDirectoryDomainVariable,
  AddDirectoryDomainParameter,
  StoreDeleteNewDirectoryDomainVariable,
  StoreDeleteNewDirectoryDomainParameter,
  UpdateDirectoryDomainParameter,
  UpdateDirectoryDomainVariable,
  DeleteDirectoryDomainParameter,
  DeleteDirectoryDomainVariable,
  SwitchDirectoryDomainParameter,
  SwitchDirectoryDomainVariable,
  ImportDirectory,
  StorePasteDirectoryDomainVariables,
  StorePasteDirectoryDomainParameters, ImportXMLDomain, SwitchDirectoryDomain,
} from '../../../store/directory/directory.actions';
import { AbstractControl, FormsModule } from '@angular/forms';
import { MatBottomSheet } from '@angular/material/bottom-sheet';
import { ConfirmBottomSheetComponent } from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import { MatSnackBar } from '@angular/material/snack-bar';
import { InnerHeaderComponent } from "../../inner-header/inner-header.component";
import {KeyValuePadComponent} from "../../key-value-pad/key-value-pad.component";
import {CodeEditorComponent} from "../../code-editor/code-editor.component";
import {JsonPipe} from "@angular/common";
import {State} from "../../../store/directory/directory.reducers";

@Component({
  standalone: true,
  imports: [MaterialModule, FormsModule, InnerHeaderComponent, KeyValuePadComponent, CodeEditorComponent, JsonPipe],
  selector: 'app-domains',
  templateUrl: './domains.component.html',
  styleUrls: ['./domains.component.css']
})
export class DomainsComponent implements OnInit {

  // --- Dependency Injection using inject() ---
  public store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);

  private directoryState = toSignal(
    this.store.pipe(select(selectDirectoryState)),
    {
      initialValue: {
        domains: {},
        users: {},
        webUsersTemplates: {},
        templatesItems: {},
        errorMessage: null,
        loadCounter: 0
      } as State
    }
  );

  public domainsList = computed(() => this.directoryState().domains || []);
  public listDetails = computed(() => this.directoryState().domainDetails || {});
  public loadCounter = computed(() => this.directoryState().loadCounter || 0);

  // --- Local State as Signals/Properties ---
  public newDomainName = signal(''); // Local input state now a signal
  public selectedIndex: number = 0;
  public toCopy: number;
  public domainParamDispatchers: object;
  public domainVarDispatchers: object;
  public XMLBody: string;
  public editorInited: boolean;

  // NOTE: lastErrorMessage is now read directly from directoryState()
  // --- Effect for Side Effects (Replaces Subscription logic) ---
  private domainEffect = effect(() => {
    const errorMessage = this.directoryState().errorMessage;

    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    } else {
      // If the error message clears, reset the new domain name input
      this.newDomainName.set('');
    }
  });

  ngOnInit() {
    this.domainParamDispatchers = {
      addItemField: StoreAddNewDirectoryDomainParameter,
      switchItem: SwitchDirectoryDomainParameter,
      newItem: AddDirectoryDomainParameter,
      dropNewItem: StoreDeleteNewDirectoryDomainParameter,
      deleteItem: DeleteDirectoryDomainParameter,
      updateItem: UpdateDirectoryDomainParameter,
      pasteItems: StorePasteDirectoryDomainParameters,
    };
    this.domainVarDispatchers = {
      addItemField: StoreAddNewDirectoryDomainVariable,
      switchItem: SwitchDirectoryDomainVariable,
      newItem: AddDirectoryDomainVariable,
      dropNewItem: StoreDeleteNewDirectoryDomainVariable,
      deleteItem: DeleteDirectoryDomainVariable,
      updateItem: UpdateDirectoryDomainVariable,
      pasteItems: StorePasteDirectoryDomainVariables,
    };
  }

  importDirectory() {
    this.store.dispatch(new ImportDirectory(null));
  }

  getDetails(id) {
    this.store.dispatch(new GetDirectoryDomainDetails({ id: id }));
  }

  clearDetails(id) {
    //  this.store.dispatch(new ClearDetails(id));
  }

  switchDomain(object) {
    this.store.dispatch(new SwitchDirectoryDomain({ id: object.id, enabled: !object.enabled }));
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  onDomainSubmit() {
    // Read the signal value using newDomainName()
    this.store.dispatch(new AddDirectoryDomain({ name: this.newDomainName() }));
  }

  ImportXMLDomain() {
    this.store.dispatch(new ImportXMLDomain({ file: this.XMLBody }));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  trackByFnId(index, item) {
    return item.id;
  }

  trackByFnFields(index, item) {
    return item.id;
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  copy(key) {
    // Read computed signal value
    if (!this.listDetails()[key]) {
      this.toCopy = 0;
      return;
    }
    this.toCopy = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
      // horizontalPosition: 'right',
      // verticalPosition: 'top',
    });
  }

  openBottomSheet(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete domain "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename domain "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DeleteDirectoryDomain({ id: id }));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameDirectoryDomain({ id: id, name: newName }));
      }
    });
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  initEditor() {
    this.editorInited = true;
  }

}
