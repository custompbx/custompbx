import {computed, Component, effect, OnDestroy, OnInit} from '@angular/core';
import {CommonModule} from "@angular/common";
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {ConfirmationService} from '../../../services/confirmation.service';
import {AbstractControl, FormsModule} from '@angular/forms';
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
import {ToastService} from '../../../services/toast.service';
import {ActivatedRoute} from '@angular/router';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {TabNavComponent} from '../../tab-nav/tab-nav.component';
import {DisclosureComponent} from '../../disclosure/disclosure.component';
import {toSignal} from "@angular/core/rxjs-interop";

@Component({
standalone: true,
  imports: [CommonModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, TabNavComponent, DisclosureComponent],
  selector: 'app-distributor',
  templateUrl: './distributor.component.html',
  styleUrls: ['./distributor.component.css']
})
export class DistributorComponent implements OnInit, OnDestroy {

  private configState = toSignal(this.store.pipe(select(selectConfigurationState)), {initialValue: {} as any});
  public list = computed(() => this.configState().distributor as Idistributor);
  public loadCounter = computed(() => this.configState().loadCounter || 0);
  public errorMessage = computed(() => this.configState().distributor?.errorMessage || null);
  public statusText = computed(() => this.loadCounter() > 0 ? 'Saving…' : null);
  public statusTone = computed(() => this.errorMessage() ? 'danger' : this.loadCounter() > 0 ? 'warning' : 'default');
  private newItemName: string;
  public selectedIndex: number;
  public expandedLists = [];

  constructor(
    private store: Store<AppState>,
    private bottomSheet: ConfirmationService,
    private _snackBar: ToastService,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    effect(() => {
      const errorMessage = this.errorMessage();
      if (!errorMessage) {
        this.newItemName = '';
        this.selectedIndex = 0;
      } else {
        this._snackBar.open('Error: ' + errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnInit() {
  }

  ngOnDestroy() {
    if (this.route.snapshot?.data?.reconnectUpdater) {
      if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
    }
  }

  getDetails(id) {
    this.expandedLists['list' + id] = true;
    this.store.dispatch(new GetDistributorNodes({id: id}));
  }

  expandAllPanels() {
    this.onlyValues(this.list()?.lists).forEach((list) => {
      if (list?.id) {
        this.expandedLists['list' + list.id] = true;
        this.store.dispatch(new GetDistributorNodes({id: list.id}));
      }
    });
  }

  collapseAllPanels() {
    this.onlyValues(this.list()?.lists).forEach((list) => {
      if (list?.id) {
        this.expandedLists['list' + list.id] = false;
      }
    });
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
    this.expandedLists['list' + id] = false;
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
    const sheet = this.bottomSheet.open(config);
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
