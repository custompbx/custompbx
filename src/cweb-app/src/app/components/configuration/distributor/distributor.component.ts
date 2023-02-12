import {Component, Inject, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {AbstractControl} from '@angular/forms';
import {
  AddDistributorList,
  AddDistributorNode, DelDistributorList,
  DelDistributorNode,
  GetDistributorNodes, StoreDelNewDistributorNode,
  StoreNewDistributorNode,
  SwitchDistributorNode, UpdateDistributorList,
  UpdateDistributorNode
} from '../../../store/config/distributor/config.actions.distributor';
import {Idistributor, IdistributorNode, Inode} from '../../../store/config/config.state.struct';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-distributor',
  templateUrl: './distributor.component.html',
  styleUrls: ['./distributor.component.css']
})
export class DistributorComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Idistributor;
  private newItemName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
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
      this.list = configs.distributor;
      this.lastErrorMessage = configs.distributor && configs.distributor.errorMessage || null;
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
    this.store.dispatch(new GetDistributorNodes({id: id}));
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
    this.store.dispatch(new UpdateDistributorNode({distributor_node: node}));
  }

  switchNode(node: IdistributorNode) {
    const newNode = <IdistributorNode>{...node};
    newNode.enabled = !node.enabled;
    this.store.dispatch(new SwitchDistributorNode({distributor_node: newNode}));
  }

  deleteNode(node: Inode) {
    this.store.dispatch(new DelDistributorNode({distributor_node: node}));
  }
  clearDetails(id) {
    //  this.store.dispatch(new ClearDetails(id));
  }

  addDistributorNodeField(id) {
    this.store.dispatch(new StoreNewDistributorNode(id));
  }

  dropNewNode(id: number, index: number) {
    this.store.dispatch(new StoreDelNewDistributorNode({id: id, index: index}));
  }

  newNode(id: number, index: number, name: string, cidr: AbstractControl) {
    const node = {
      enabled: true,
      name: name,
      weight: cidr.value
    };

    this.store.dispatch(new AddDistributorNode({id: id, index: index, distributor_node: node}));
  }

  onDistributorListSubmit() {
    this.store.dispatch(new AddDistributorList({name: this.newItemName}));
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
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
        this.store.dispatch(new DelDistributorList({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateDistributorList({id: id, name: newName}));
      }
    });
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

}
