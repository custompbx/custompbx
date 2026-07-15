import {Component, OnDestroy, inject, signal, computed, effect} from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';

import {CommonModule} from "@angular/common";
import {select, Store} from '@ngrx/store';
import {AppState, selectDirectoryState} from '../../../store/app.states';
import {
  GetDirectoryGroupUsers,
  AddNewDirectoryGroup,
  DeleteDirectoryGroup,
  UpdateDirectoryGroupName,
  AddDirectoryGroupUser,
  DeleteDirectoryGroupUser,
  ImportDirectory
} from '../../../store/directory/directory.actions';
import {AbstractControl, FormsModule} from '@angular/forms';
import {ActivatedRoute, RouterLink} from '@angular/router';
import {ToastService} from '../../../services/toast.service';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {TabNavComponent} from '../../tab-nav/tab-nav.component';
import {DisclosureComponent} from '../../disclosure/disclosure.component';
import {State} from "../../../store/directory/directory.reducers";
import {CpbxSelectDirective} from '../../../directives/cpbx-select.directive';
import {ConfirmationService} from '../../../services/confirmation.service';
import {IconComponent} from '../../icon/icon.component';

@Component({
  standalone: true,
  imports: [CommonModule, FormsModule, InnerHeaderComponent, RouterLink, CpbxSelectDirective, TabNavComponent, DisclosureComponent, IconComponent],
  selector: 'app-groups',
  templateUrl: './groups.component.html',
  styleUrls: ['./groups.component.css']
})
export class GroupsComponent implements OnDestroy {

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private confirmation = inject(ConfirmationService);
  private route = inject(ActivatedRoute);
  private _snackBar = inject(ToastService);

  // --- Reactive State from NgRx using toSignal ---
  private directoryState = toSignal(
    this.store.pipe(select(selectDirectoryState)),
    {
      initialValue: {
        domains: {},
        groupNames: {},
        groupUsers: {},
        users: {},
        errorMessage: null,
        loadCounter: 0
      } as State
    }
  );

  // --- Computed State for Template Access ---
  public list = computed(() => this.directoryState().domains || {}); // domains
  protected groupList = computed(() => this.directoryState().groupNames || {});
  public listGroupUsers = computed(() => this.directoryState().groupUsers || {});
  public listDomainUsers = computed(() => this.directoryState().users || {});
  public loadCounter = computed(() => this.directoryState().loadCounter || 0);

  // --- Local State as Signals/Properties ---
  public newGroupName = signal('');
  public domainId: number;
  public selectedIndex: number = 0;
  public opened: boolean = false; // Initialize boolean properties

  // Note: `bound` and `manageGroup` are objects that seem to be dynamically updated
  // based on reactive data. We will make `bound` a computed signal.

  public manageGroup: object = {};

  // Computed Signal for the `bound` map
  public bound = computed(() => {
    const groupNames = this.groupList();
    const groupUsers = this.listGroupUsers();

    // 1. Initialize bound map for all groups
    const newBound: { [groupId: number]: { [userId: number]: boolean } } = {};
    Object.values(groupNames).forEach((g: any) => newBound[g?.id] = {});

    // 2. Map group users to the bound structure
    Object.values(groupUsers).forEach((u: any) => {
        const groupId = u.parent?.id;
        const userId = u.user?.id;
        if (groupId && userId) {
          if (!newBound[groupId]) {
            newBound[groupId] = {};
          }
          newBound[groupId][userId] = true;
        }
      }
    );
    return newBound;
  });

  private groupEffect =  effect(() => {
    const users = this.directoryState();
    const lastErrorMessage = users.errorMessage;

    if (lastErrorMessage) {
      this._snackBar.open('Error: ' + lastErrorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    } else {
      // If the error message clears, reset the new group name input and selected index
      this.newGroupName.set('');
      // This is simplified, as complex index/domain resets are often better handled by routing/user input
      this.selectedIndex = this.selectedIndex === 1 ? 0 : this.selectedIndex;
    }
  });

  ngOnDestroy() {
    // toSignal handles users$.unsubscribe() automatically.
    if (this.route.snapshot?.data?.reconnectUpdater) {
      this.route.snapshot.data.reconnectUpdater.unsubscribe();
    }
  }

  importDirectory() {
    this.store.dispatch(new ImportDirectory(null));
  }

  getDetails(id) {
    this.store.dispatch(new GetDirectoryGroupUsers({id: Number(id)}));
  }

  clearDetails(id) {
    //  this.store.dispatch(new ClearDetails(id));
  }

  addUser(userId, groupId) {
    this.store.dispatch(new AddDirectoryGroupUser({id_int: Number(userId), id: Number(groupId)}));
  }

  delUser(id) {
    this.store.dispatch(new DeleteDirectoryGroupUser({id: Number(id)}));
  }

  onGroupSubmit() {
    this.store.dispatch(new AddNewDirectoryGroup({userName: this.newGroupName(), domainId: Number(this.domainId)}));
  }

  checkDirty(condition: AbstractControl<any, any>): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  onlyValuesByParent(obj: object, parentId: number): Array<any> {
    if (!obj) {
      return [];
    }
    // Access the computed signal listDomainUsers()
    return Object.values(obj).filter((u: any) => u.parent?.id === Number(parentId));
  }

  openBottomSheet(id, newName, oldName, action): void {
    const previousName = typeof oldName === 'string' ? oldName : oldName?.name;
    const message = action === 'delete'
      ? `Delete group "${previousName}"?`
      : `Rename group "${previousName}" to "${newName}"?`;
    const sheet = this.confirmation.open({data: {action, message}});
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DeleteDirectoryGroup({id: Number(id)}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateDirectoryGroupName({id: Number(id), name: newName}));
      }
    });
  }

  // bingToGroup is no longer needed as the `bound` computed signal replaces it.
  // It is also explicitly marked as deprecated in the old code so we can remove it.
  /*
  bingToGroup(): void {
    // ... logic moved to computed signal `bound` ...
  }
  */
  protected readonly Number = Number;
}
