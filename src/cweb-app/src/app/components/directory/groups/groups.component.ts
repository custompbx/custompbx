import {Component, Inject, OnDestroy, OnInit} from '@angular/core';
import {select, Store} from '@ngrx/store';
import {AppState, selectDirectoryState} from '../../../store/app.states';
import {Observable, Subscription} from 'rxjs';
import {
  GetDirectoryGroupUsers,
  AddNewDirectoryGroup,
  DeleteDirectoryGroup,
  UpdateDirectoryGroupName,
  AddDirectoryGroupUser,
  DeleteDirectoryGroupUser,
  ImportDirectory
} from '../../../store/directory/directory.actions';
import {UntypedFormControl} from '@angular/forms';
import {MAT_BOTTOM_SHEET_DATA, MatBottomSheet, MatBottomSheetRef} from '@angular/material/bottom-sheet';
import {ActivatedRoute} from '@angular/router';
import {MatSnackBar} from '@angular/material/snack-bar';

@Component({
  selector: 'app-groups',
  templateUrl: './groups.component.html',
  styleUrls: ['./groups.component.css']
})
export class GroupsComponent implements OnInit, OnDestroy {

  public users: Observable<any>;
  public users$: Subscription;
  private bound: object;
  private manageGroup: object;
  public list: any;
  private groupList: any;
  public listGroupUsers: any;
  public listDomainUsers: any;
  private newGroupName: string;
  private domainId: number;
  public selectedIndex: number;
  private lastErrorMessage: string;
  opened: boolean;
  public loadCounter: number;

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private route: ActivatedRoute,
    private _snackBar: MatSnackBar,
  ) {
    this.manageGroup = {};
    this.selectedIndex = 0;
    this.users = this.store.pipe(select(selectDirectoryState));
  }

  ngOnInit() {
    this.users$ = this.users.subscribe((users) => {
      this.loadCounter = users.loadCounter;
      this.list = users.domains;
      this.groupList = users.groupNames;
      this.listGroupUsers = users.groupUsers;
      this.listDomainUsers = users.users;
      this.lastErrorMessage = users.errorMessage;
      this.bound = {};
      this.bingToGroup();
      if (!this.lastErrorMessage) {
        this.newGroupName = '';
        // this.domainId = 0;
        this.selectedIndex = this.selectedIndex === 1 ? 0 : this.selectedIndex;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.users$.unsubscribe();
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
    this.store.dispatch(new AddNewDirectoryGroup({userName: this.newGroupName, domainId: Number(this.domainId)}));
  }

  checkDirty(condition: UntypedFormControl): boolean {
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
    return Object.values(obj).filter(u => u.parent?.id === Number(parentId));
  }

  openBottomSheet(id, newName, oldName, action): void {
    const sheet = this.bottomSheet.open(GroupsBottomSheetComponent, {data: {newName: newName, oldName: oldName, action: action}});
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

  bingToGroup(): void {
    Object.values(this.groupList).forEach((g: any) => this.bound[g?.id] = {});
    Object.values(this.listGroupUsers).forEach((u: any) => {
        if (!this.bound[u.parent?.id]) {
          this.bound[u.parent?.id] = {};
        }
        this.bound[u.parent?.id][u.user?.id] = true;
      }
    );
  }

}

@Component({
  selector: 'app-bottom-sheet-sheet',
  template: '<div [ngSwitch]="data.action">\n' +
    '  <h3 *ngSwitchCase="\'delete\'">Are you sure you want to delete group "{{data.oldName}}"?</h3>\n' +
    '  <h3 *ngSwitchCase="\'rename\'">Are you sure you want to rename group "{{data.oldName}}" to "{{data.newName}}"?</h3>\n' +
    '</div>' +
    '<mat-nav-list>\n' +
    '  <a mat-list-item><button mat-button mat-line color="warn" (click)="confirmAction(true)">Confirm</button></a>\n' +
    '  <a mat-list-item><button mat-button mat-line color="primary" (click)="confirmAction(false)">Cancel</button></a>\n' +
    '</mat-nav-list>',
})
export class GroupsBottomSheetComponent {
  constructor(
    private bottomSheetRef: MatBottomSheetRef<GroupsBottomSheetComponent>,
    @Inject(MAT_BOTTOM_SHEET_DATA) public data: any
  ) {
  }

  confirmAction(event: boolean): void {
    this.bottomSheetRef.dismiss(event);
  }
}
