import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {select, Store} from '@ngrx/store';
import {AppState, selectDirectoryState} from '../../../store/app.states';
import {Observable, Subscription} from 'rxjs';
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
import {AbstractControl} from '@angular/forms';
import { MatBottomSheet } from '@angular/material/bottom-sheet';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {Idetails} from '../../../store/directory/directory.reducers';
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-domains',
  templateUrl: './domains.component.html',
  styleUrls: ['./domains.component.css']
})
export class DomainsComponent implements OnInit, OnDestroy {

  public domains: Observable<any>;
  public domains$: Subscription;
  public list: any;
  public listDetails: Idetails;
  public newDomainName: string;
  public selectedIndex: number;
  public lastErrorMessage: string;
  public loadCounter: number;
  private toCopy: number;
  public domainParamDispatchers: object;
  public domainVarDispatchers: object;
  public XMLBody: string;
  public editorInited: boolean;

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
  ) {
    this.selectedIndex = 0;
    this.domains = this.store.pipe(select(selectDirectoryState));
  }

  ngOnInit() {
    this.domains$ = this.domains.subscribe((domains) => {
      this.loadCounter = domains.loadCounter;
      this.list = domains.domains;
      this.listDetails = domains.domainDetails;
      this.lastErrorMessage = domains.errorMessage;
      if (!this.lastErrorMessage) {
        this.newDomainName = '';
        // this.selectedIndex = 0;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
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

  ngOnDestroy() {
    this.domains$.unsubscribe();
  }

  importDirectory() {
    this.store.dispatch(new ImportDirectory(null));
  }

  getDetails(id) {
    this.store.dispatch(new GetDirectoryDomainDetails({id: id}));
  }

  clearDetails(id) {
    //  this.store.dispatch(new ClearDetails(id));
  }

  switchDomain(object) {
    this.store.dispatch(new SwitchDirectoryDomain({id: object.id, enabled: !object.enabled}));
  }
  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  onDomainSubmit() {
    this.store.dispatch(new AddDirectoryDomain({name: this.newDomainName}));
  }

  ImportXMLDomain() {
    this.store.dispatch(new ImportXMLDomain({file: this.XMLBody}));
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
    if (!this.listDetails[key]) {
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
        this.store.dispatch(new DeleteDirectoryDomain({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new RenameDirectoryDomain({id: id, name: newName}));
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
