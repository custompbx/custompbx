import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {AbstractControl} from '@angular/forms';
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
import {Iacl, Inode} from '../../../store/config/config.state.struct';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {CdkDragDrop} from '@angular/cdk/drag-drop';

@Component({
  selector: 'app-acl',
  templateUrl: './acl.component.html',
  styleUrls: ['./acl.component.css']
})
export class AclComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Iacl;
  public newItemName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public aclBehavior = ['allow', 'deny'];
  public defaultBehavior = 'deny';
  public loadCounter: number;

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
      this.list = configs.acl;
      this.lastErrorMessage = configs.acl && configs.acl.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newItemName = '';
        this.selectedIndex = 0;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
      if (this.route.snapshot?.data?.reconnectUpdater) {
        this.route.snapshot.data.reconnectUpdater.unsubscribe();
      }
    }
  }

  getDetails(id) {
    this.store.dispatch(GetAclNodes({id}));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isvalueReadyToSend(valueObject: AbstractControl): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isReadyToSendThree(mainObject: AbstractControl, object2: AbstractControl, object3: AbstractControl): boolean {
    return mainObject && mainObject.valid && mainObject.dirty
      && ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty));
  }

  updateNode(node: Inode) {
    this.store.dispatch(UpdateAclNode({node: node}));
  }

  updateDefault(id: number, valueObject: AbstractControl) {
    const value = valueObject.value;
    valueObject.reset();
    this.store.dispatch(UpdateAclListDefault({value: value, id: id}));
  }

  switchNode(node: Inode) {
    const newNode = <Inode>{...node};
    newNode.enabled = !node.enabled;
    this.store.dispatch(SwitchAclNode({node: newNode}));
  }

  deleteNode(node: Inode) {
    this.store.dispatch(DelAclNode({id: node.id}));
  }

  clearDetails(id) {
    //  this.store.dispatch(ClearDetails(id));
  }

  addAclNodeField(id) {
    this.store.dispatch(StoreNewAclNode({id}));
  }

  dropNewNode(id: number, index: number) {
    this.store.dispatch(DropNewAclNode({id: id, index: index}));
  }

  newNode(id: number, index: number, type: string, cidr: AbstractControl, domain: AbstractControl) {
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
    this.store.dispatch(AddAclList({name: this.newItemName, default: this.defaultBehavior}));
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  trackByPositionFn(index, item) {
    return item.position;
  }

  trackByIdFn(index, item) {
    return item.id;
  }

  openBottomSheet(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete list "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename list "' + oldName + '" to "' + newName + '"?',
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

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
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

  dropAction(event: CdkDragDrop<string[]>, parent: Array<any>) {
    if (parent[event.previousIndex].position === parent[event.currentIndex].position) {
      return;
    }
    this.store.dispatch(MoveAclListNode({
      previous_index: parent[event.previousIndex].position,
      current_index: parent[event.currentIndex].position,
      id: parent[event.previousIndex].id
    }));
  }

}
