import {Component, inject, signal, computed, effect, OnInit} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';

import {MaterialModule} from "../../../../material-module";
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {AbstractControl, FormsModule} from '@angular/forms';
import {
  AddAclList,
  AddAclNode,
  DelAclList,
  DelAclNode,
  DropNewAclNode,
  GetAclNodes,
  MoveAclListNode,
  StoreNewAclNode,
  SwitchAclNode,
  UpdateAclList,
  UpdateAclListDefault,
  UpdateAclNode
} from '../../../store/config/acl/config.actions.acl';
import {Iacl, Inode, State} from '../../../store/config/config.state.struct';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {CdkDragDrop} from '@angular/cdk/drag-drop';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {ResizeInputDirective} from "../../../directives/resize-input.directive";
import {JsonPipe} from "@angular/common";

@Component({
  standalone: true,
  imports: [MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, ResizeInputDirective, JsonPipe],
  selector: 'app-acl',
  templateUrl: './acl.component.html',
  styleUrls: ['./acl.component.css']
})
export class AclComponent { // Removed OnDestroy

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private configState = toSignal(
    this.store.pipe(select(selectConfigurationState)),
    {
      initialValue: {
        acl: {} as Iacl,
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().acl);
  public loadCounter = computed(() => this.configState().loadCounter);
  public lastErrorMessage = computed(() => this.configState().errorMessage);


  // --- Local Component State as Signals/Properties ---
  public newItemName: string = ''; // Kept as property for two-way binding
  public selectedIndex: number = 0; // Kept as property for binding
  public aclBehavior = ['allow', 'deny'];
  public defaultBehavior = 'deny';


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

  getDetails(id: number) {
    this.store.dispatch(GetAclNodes({id}));
  }

  checkDirty(condition: AbstractControl | null): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isvalueReadyToSend(valueObject: AbstractControl | null): boolean {
    return !!(valueObject && valueObject.dirty && valueObject.valid);
  }

  isReadyToSend(nameObject: AbstractControl | null, valueObject: AbstractControl | null): boolean {
    return !!(nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid);
  }

  isReadyToSendThree(mainObject: AbstractControl | null, object2: AbstractControl | null, object3: AbstractControl | null): boolean {
    return !!(mainObject && mainObject.valid && mainObject.dirty
      && ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty)));
  }

  updateNode(node: Inode) {
    this.store.dispatch(UpdateAclNode({node: node}));
  }

  updateDefault(id: number, valueObject: AbstractControl) {
    const value = valueObject.value;
    valueObject.reset(); // Reset control after dispatching action
    this.store.dispatch(UpdateAclListDefault({default: value, id: id}));
  }

  switchNode(node: Inode) {
    const newNode = <Inode>{...node};
    newNode.enabled = !node.enabled;
    this.store.dispatch(SwitchAclNode({node: newNode}));
  }

  deleteNode(node: Inode) {
    this.openBottomSheetNode(node.id, 'Are you sure you want to delete this ACL node?', 'delete');
  }

  // Refactored to use a dedicated bottom sheet for nodes
  openBottomSheetNode(id: number, message: string, action: 'delete'): void {
    const config = {
      data: {
        action: action,
        case1Text: message,
      }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(DelAclNode({id: id}));
      }
    });
  }


  clearDetails(id: number) {
    //  this.store.dispatch(ClearDetails(id));
  }

  addAclNodeField(id: number) {
    this.store.dispatch(StoreNewAclNode({id}));
  }

  dropNewNode(id: number, index: number) {
    this.store.dispatch(DropNewAclNode({id: id, index: index}));
  }

  newNode(id: number, index: number, type: string, cidr: AbstractControl | null, domain: AbstractControl | null) {
    const node = <Inode>{};
    node.enabled = true;
    node.type = type;

    if (cidr) {
      node.cidr = cidr.value;
    }
    if (domain) {
      node.domain = domain.value;
    }
    this.store.dispatch(AddAclNode({id: id, index: index, node: node}));
  }

  onAclListSubmit() {
    const name = this.newItemName.trim();
    if (name) {
      this.store.dispatch(AddAclList({name: name, default: this.defaultBehavior}));
      this.newItemName = ''; // Clear input field
    }
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  openBottomSheet(id: number, newName: string, oldName: string, action: 'delete' | 'rename'): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: action === 'delete' ? 'Are you sure you want to delete list "' + oldName + '"?' : null,
          case2Text: action === 'rename' ? 'Are you sure you want to rename list "' + oldName + '" to "' + newName + '"?' : null,
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(DelAclList({id}));
      } else if (action === 'rename') {
        this.store.dispatch(UpdateAclList({id: id, name: newName}));
      }
    });
  }

  onlyValues(obj: object | null): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  onlySortedValues(obj: object | null): Array<any> {
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

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    const previousItem = parent[event.previousIndex];
    const currentItem = parent[event.currentIndex];

    if (!previousItem || !currentItem || previousItem.position === currentItem.position) {
      return;
    }
    this.store.dispatch(MoveAclListNode({
      previous_index: previousItem.position,
      current_index: currentItem.position,
      id: previousItem.id
    }));
  }

}
