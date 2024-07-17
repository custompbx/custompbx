import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Iconference, Iitem} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {ConfirmBottomSheetComponent} from '../../confirm-bottom-sheet/confirm-bottom-sheet.component';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  AddConferenceProfile,
  AddConferenceProfileParameter,
  AddConferenceRoom,
  DelConferenceProfile,
  DelConferenceProfileParameter,
  DelConferenceRoom,
  GetConferenceProfileParameters,
  UpdateConferenceProfile,
  StoreDropNewConferenceProfileParameter,
  StoreDropNewConferenceRoom,
  StoreNewConferenceProfileParameter,
  StoreNewConferenceRoom,
  StorePasteConferenceProfileParameters,
  SwitchConferenceProfileParameter,
  SwitchConferenceRoom,
  UpdateConferenceProfileParameter,
  UpdateConferenceRoom,
  UpdateConferenceCallerControl,
  AddConferenceCallerControl,
  DelConferenceCallerControl,
  StoreDropNewConferenceCallerControl,
  SwitchConferenceCallerControl,
  AddConferenceCallerControlGroup,
  GetConferenceCallerControls,
  StoreNewConferenceCallerControl,
  StorePasteConferenceCallerControls,
  DelConferenceCallerControlGroup,
  UpdateConferenceCallerControlGroup,
  AddConferenceChatPermissionUser,
  AddConferenceChatPermission,
  StoreDropNewConferenceChatPermissionUser,
  GetConferenceChatPermissionUsers,
  StoreNewConferenceChatPermissionUser,
  SwitchConferenceChatPermissionUser,
  DelConferenceChatPermissionUser,
  UpdateConferenceChatPermissionUser,
  StorePasteConferenceChatPermissionUsers,
  DelConferenceChatPermission,
  UpdateConferenceChatPermission,
  GetConferenceLayouts,
  GetConferenceLayoutImages,
  GetConferenceLayoutGroupLayouts,
  StoreNewConferenceLayoutImage,
  StoreDropConferenceLayoutImage,
  SwitchConferenceLayoutImage,
  AddConferenceLayoutImage,
  StorePasteConferenceLayoutImage,
  DelConferenceLayoutImage,
  UpdateConferenceLayoutImage,
  StoreNewConferenceLayoutGroupLayout,
  SwitchConferenceLayoutGroupLayout,
  AddConferenceLayoutGroupLayout,
  StoreDropConferenceLayoutGroupLayout,
  DelConferenceLayoutGroupLayout,
  UpdateConferenceLayoutGroupLayout,
  StorePasteConferenceLayoutGroupLayout,
  AddConferenceLayout,
  AddConferenceLayoutGroup,
  DelConferenceLayout,
  UpdateConferenceLayout, DelConferenceLayoutGroup, UpdateConferenceLayoutGroup, UpdateConferenceLayout3D,
} from '../../../store/config/conference/config.actions.conference';

@Component({
  selector: 'app-conference',
  templateUrl: './conference.component.html',
  styleUrls: ['./conference.component.css']
})
export class ConferenceComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: Iconference;
  public newProfileName: string;
  public newControlGroupName: string;
  public newChatPermissionName: string;
  private newGroupName: string;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public profileId: number;
  public controlGroupId: number;
  public chatPermissionId: number;
  public loadCounter: number;
  public toCopyProfile: number;
  public toCopyUser: number;
  public toCopyGroup: number;
  public globalSettingsDispatchers: object;
  public groupSettingsDispatchers: object;
  public profileSettingsDispatchers: object;
  public chatPermissionSettingsDispatchers: object;
  public advertiseMask: object;
  public chatPermissionMask: object;
  public controlMask: object;
  public layoutImageMask: object;
  public layoutGroupMask: object;
  public layoutImageDispatchers: object;
  public layoutGroupDispatchers: object;
  public toCopylayoutImage: number;
  public toCopylayoutGroup: number;
  public newLayoutName: string;
  public newLayoutGroupName: string;

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
      this.list = configs.conference;
      this.lastErrorMessage = configs.conference && configs.conference.errorMessage || null;
      if (!this.lastErrorMessage) {
        this.newProfileName = '';
        this.newControlGroupName = '';
        this.newChatPermissionName = '';
        this.newLayoutName = '';
        this.newLayoutGroupName = '';
        this.profileId = (this.list && this.list.profiles && this.list.profiles[this.profileId]) ? this.profileId : 0;
        this.controlGroupId = (this.list && this.list.caller_controls && this.list.caller_controls[this.controlGroupId]) ?
          this.controlGroupId : 0;
        this.chatPermissionId = (this.list && this.list.chat_profiles && this.list.chat_profiles[this.chatPermissionId]) ?
          this.chatPermissionId : 0;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewConferenceRoom.bind(this),
      switchItem: this.switchConferenceRoom.bind(this),
      addItem: this.newConferenceRoom.bind(this),
      dropNewItem: this.dropNewConferenceRoom.bind(this),
      deleteItem: this.deleteConferenceRoom.bind(this),
      updateItem: this.updateConferenceRoom.bind(this),
      pasteItems: null,
    };
    this.groupSettingsDispatchers = {
      addNewItemField: this.addNewCallerControl.bind(this),
      switchItem: this.switchCallerControl.bind(this),
      addItem: this.newCallerControl.bind(this),
      dropNewItem: this.dropNewCallerControl.bind(this),
      deleteItem: this.deleteCallerControl.bind(this),
      updateItem: this.updateCallerControl.bind(this),
      pasteItems: this.pasteCallerControls.bind(this),
    };
    this.profileSettingsDispatchers = {
      addNewItemField: this.addNewProfileParam.bind(this),
      switchItem: this.switchProfileParam.bind(this),
      addItem: this.newProfileParam.bind(this),
      dropNewItem: this.dropNewProfileParam.bind(this),
      deleteItem: this.deleteProfileParam.bind(this),
      updateItem: this.updateProfileParam.bind(this),
      pasteItems: this.pasteProfileParams.bind(this),
    };
    this.chatPermissionSettingsDispatchers = {
      addNewItemField: this.addNewChatPermissionUser.bind(this),
      switchItem: this.switchChatPermissionUser.bind(this),
      addItem: this.newChatPermissionUser.bind(this),
      dropNewItem: this.dropNewChatPermissionUser.bind(this),
      deleteItem: this.deleteChatPermissionUser.bind(this),
      updateItem: this.updateChatPermissionUser.bind(this),
      pasteItems: this.pasteChatPermissionUsers.bind(this),
    };
    this.layoutImageDispatchers = {
      addNewItemField: this.storeNewConferenceLayoutImage.bind(this),
      switchItem: this.switchConferenceLayoutImage.bind(this),
      addItem: this.addConferenceLayoutImage.bind(this),
      dropNewItem: this.storeDropConferenceLayoutImage.bind(this),
      deleteItem: this.delConferenceLayoutImage.bind(this),
      updateItem: this.updateConferenceLayoutImage.bind(this),
      pasteItems: this.storePasteConferenceLayoutImage.bind(this),
    };
    this.layoutGroupDispatchers = {
      addNewItemField: this.storeNewConferenceLayoutGroupLayout.bind(this),
      switchItem: this.switchConferenceLayoutGroupLayout.bind(this),
      addItem: this.addConferenceLayoutGroupLayout.bind(this),
      dropNewItem: this.storeDropConferenceLayoutGroupLayout.bind(this),
      deleteItem: this.delConferenceLayoutGroupLayout.bind(this),
      updateItem: this.updateConferenceLayoutGroupLayout.bind(this),
      pasteItems: this.storePasteConferenceLayoutGroupLayout.bind(this),
    };
    this.advertiseMask = {name: {name: 'name'}, value: {name: 'status'}};
    this.chatPermissionMask = {name: {name: 'name'}, value: {name: 'commands'}};
    this.controlMask = {name: {name: 'action'}, value: {name: 'digits'}};
    this.layoutImageMask = {
      name: {name: 'x'},
      value: {name: 'y'},
      extraField1: {name: 'scale'},
      extraField2: {name: 'hscale'},
      extraField3: {name: 'zoom'},
      extraField4: {name: 'floor'},
      extraField5: {name: 'floor_only'},
      extraField6: {name: 'overlap'},
      extraField7: {name: 'reservation_id'},
    };
    this.layoutGroupMask = {name: {name: 'body'}};
  }
/*
x
y
scale
hscale

zoom
floor
floor_only
overlap
reservation_id
*/
  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  mainTabChanged(event) {
    if (event === 1) {
      this.store.dispatch(GetConferenceLayouts(null));
    }
  }

  updateConferenceRoom(param: Iitem) {
    this.store.dispatch(new UpdateConferenceRoom({param: param}));
  }

  switchConferenceRoom(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchConferenceRoom({param: newParam}));
  }

  newConferenceRoom(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddConferenceRoom({index: index, param: param}));
  }

  deleteConferenceRoom(param: Iitem) {
    this.store.dispatch(new DelConferenceRoom({param: param}));
  }

  addNewConferenceRoom() {
    this.store.dispatch(new StoreNewConferenceRoom(null));
  }

  dropNewConferenceRoom(index: number) {
    this.store.dispatch(new StoreDropNewConferenceRoom({index: index}));
  }

  getConferenceCallerControls(id) {
    this.store.dispatch(new GetConferenceCallerControls({id: id}));
  }

  updateCallerControl(param) {
    const para = <Iitem>{};
    para.id = param.id;
    para.name = param.action;
    para.value = param.digits;
    this.store.dispatch(new UpdateConferenceCallerControl({param: para}));
  }

  switchCallerControl(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchConferenceCallerControl({param: newParam}));
  }

  newCallerControl(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddConferenceCallerControl({id: parentId, index: index, param: param}));
  }

  deleteCallerControl(param: Iitem) {
    this.store.dispatch(new DelConferenceCallerControl({param: param}));
  }

  addNewCallerControl(parentId: number) {
    this.store.dispatch(new StoreNewConferenceCallerControl({id: parentId}));
  }

  dropNewCallerControl(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewConferenceCallerControl({id: parentId, index: index}));
  }

  getConferenceProfilesParams(id) {
    this.store.dispatch(new GetConferenceProfileParameters({id: id}));
  }

  updateProfileParam(param: Iitem) {
    this.store.dispatch(new UpdateConferenceProfileParameter({param: param}));
  }

  switchProfileParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchConferenceProfileParameter({param: newParam}));
  }

  newProfileParam(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddConferenceProfileParameter({id: parentId, index: index, param: param}));
  }

  deleteProfileParam(param: Iitem) {
    this.store.dispatch(new DelConferenceProfileParameter({param: param}));
  }

  addNewProfileParam(parentId: number) {
    this.store.dispatch(new StoreNewConferenceProfileParameter({id: parentId}));
  }

  dropNewProfileParam(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewConferenceProfileParameter({id: parentId, index: index}));
  }

  onProfileSubmit() {
    this.store.dispatch(new AddConferenceProfile({name: this.newProfileName}));
  }

  onControlGroupSubmit() {
    this.store.dispatch(new AddConferenceCallerControlGroup({name: this.newControlGroupName}));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isReadyToSendThree(mainObject: AbstractControl, object2: AbstractControl, object3: AbstractControl): boolean {
    return (mainObject && mainObject.valid && mainObject.dirty)
      || ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty));
  }

  isvalueReadyToSend(valueObject: AbstractControl): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

  copyCallerControlGroup(key) {
    if (!this.list.caller_controls[key]) {
      this.toCopyGroup = 0;
      return;
    }
    this.toCopyGroup = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteCallerControls(to: number) {
    this.store.dispatch(new StorePasteConferenceCallerControls({from_id: this.toCopyGroup, to_id: to}));
  }

  openBottomSheetControlGroup(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete group "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename group "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelConferenceCallerControlGroup({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateConferenceCallerControlGroup({id: id, name: newName}));
      }
    });
  }

  copyProfile(key) {
    if (!this.list.profiles[key]) {
      this.toCopyProfile = 0;
      return;
    }
    this.toCopyProfile = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  pasteProfileParams(to: number) {
    this.store.dispatch(new StorePasteConferenceProfileParameters({from_id: this.toCopyProfile, to_id: to}));
  }

  openBottomSheetProfile(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete profile "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename profile "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelConferenceProfile({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateConferenceProfile({id: id, name: newName}));
      }
    });
  }

  getConferenceChatPermissionsUsers(id) {
    this.store.dispatch(new GetConferenceChatPermissionUsers({id: id}));
  }

  updateChatPermissionUser(param: Iitem) {
    this.store.dispatch(new UpdateConferenceChatPermissionUser({param: param}));
  }

  switchChatPermissionUser(param: Iitem) {
    const newUser = <Iitem>{...param};
    newUser.enabled = !newUser.enabled;
    this.store.dispatch(new SwitchConferenceChatPermissionUser({param: newUser}));
  }

  newChatPermissionUser(parentId: number, index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddConferenceChatPermissionUser({id: parentId, index: index, param: param}));
  }

  deleteChatPermissionUser(param: Iitem) {
    this.store.dispatch(new DelConferenceChatPermissionUser({param: param}));
  }

  addNewChatPermissionUser(parentId: number) {
    this.store.dispatch(new StoreNewConferenceChatPermissionUser({id: parentId}));
  }

  dropNewChatPermissionUser(parentId: number, index: number) {
    this.store.dispatch(new StoreDropNewConferenceChatPermissionUser({id: parentId, index: index}));
  }

  onChatPermissionSubmit() {
    this.store.dispatch(new AddConferenceChatPermission({name: this.newChatPermissionName}));
  }

  pasteChatPermissionUsers(to: number) {
    this.store.dispatch(new StorePasteConferenceChatPermissionUsers({from_id: this.toCopyUser, to_id: to}));
  }

  openBottomSheetChatPermissionProfile(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete profile "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename profile "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(new DelConferenceChatPermission({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(new UpdateConferenceChatPermission({id: id, name: newName}));
      }
    });
  }

  onlyValues(obj: object): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj);
  }

  getConferenceLayoutImages(id) {
    this.store.dispatch(GetConferenceLayoutImages({id: id}));
  }

  getConferenceLayoutGroupLayouts(id) {
    this.store.dispatch(GetConferenceLayoutGroupLayouts({id: id}));
  }

  updateConferenceLayout3D(id, value) {
    this.store.dispatch(UpdateConferenceLayout3D({id, value}));
  }

  copyLayoutImage(key) {
    if (!this.list.layouts.conference_layouts[key]) {
      this.toCopylayoutImage = 0;
      return;
    }
    this.toCopylayoutImage = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  copyLayoutGroup(key) {
    if (!this.list.layouts.conference_layouts_groups[key]) {
      this.toCopylayoutGroup = 0;
      return;
    }
    this.toCopylayoutGroup = key;
    this._snackBar.open('Copied!', null, {
      duration: 700,
    });
  }

  storeNewConferenceLayoutImage(parentId: number) {
    this.store.dispatch(StoreNewConferenceLayoutImage({id: parentId}));
  }
  switchConferenceLayoutImage(param: Iitem) {
    const newUser = <Iitem>{...param};
    newUser.enabled = !newUser.enabled;
    this.store.dispatch(SwitchConferenceLayoutImage({param: newUser}));
  }
  addConferenceLayoutImage(parentId: number, index: number, x: string, y: string, scale: string, hscale: string, zoom: string, floor: string, floor_only: string, overlap: string, reservation_id: string) {
    this.store.dispatch(AddConferenceLayoutImage({id: parentId, index: index, layout_images: {x, y, scale, hscale, zoom, floor, floor_only, overlap, reservation_id}}));
  }
  storeDropConferenceLayoutImage(parentId: number, index: number) {
    this.store.dispatch(StoreDropConferenceLayoutImage({id: parentId, index: index}));
  }
  delConferenceLayoutImage(param: Iitem) {
    this.store.dispatch(DelConferenceLayoutImage({param: param}));
  }
  updateConferenceLayoutImage(layout_images) {
    this.store.dispatch(UpdateConferenceLayoutImage({layout_images}));
  }
  storePasteConferenceLayoutImage(to: number) {
    this.store.dispatch(StorePasteConferenceLayoutImage({from_id: this.toCopyUser, to_id: to}));
  }

  storeNewConferenceLayoutGroupLayout(parentId: number) {
    this.store.dispatch(StoreNewConferenceLayoutGroupLayout({id: parentId}));
  }
  switchConferenceLayoutGroupLayout(param: Iitem) {
    const newUser = <Iitem>{...param};
    newUser.enabled = !newUser.enabled;
    this.store.dispatch(SwitchConferenceLayoutGroupLayout({param: newUser}));
  }
  addConferenceLayoutGroupLayout(parentId: number, index: number, body: string) {
    this.store.dispatch(AddConferenceLayoutGroupLayout({id: parentId, index: index, enabled: true, value: body}));
  }
  storeDropConferenceLayoutGroupLayout(parentId: number, index: number) {
    this.store.dispatch(StoreDropConferenceLayoutGroupLayout({id: parentId, index: index}));
  }
  delConferenceLayoutGroupLayout(param: Iitem) {
    this.store.dispatch(DelConferenceLayoutGroupLayout({param: param}));
  }
  updateConferenceLayoutGroupLayout(param: Iitem) {
    this.store.dispatch(UpdateConferenceLayoutGroupLayout({param: param}));
  }
  storePasteConferenceLayoutGroupLayout(to: number) {
    this.store.dispatch(StorePasteConferenceLayoutGroupLayout({from_id: this.toCopyUser, to_id: to}));
  }

  onLayoutSubmit() {
    this.store.dispatch(AddConferenceLayout({name: this.newLayoutName}));
  }

  onLayoutGroupSubmit() {
    this.store.dispatch(AddConferenceLayoutGroup({name: this.newLayoutGroupName}));
  }

  openBottomSheetLayout(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete layout "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename layout "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(DelConferenceLayout({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(UpdateConferenceLayout({id: id, name: newName}));
      }
    });
  }
  openBottomSheetLayoutGroup(id, newName, oldName, action): void {
    const config = {
      data:
        {
          newName: newName,
          oldName: oldName,
          action: action,
          case1Text: 'Are you sure you want to delete layout "' + oldName + '"?',
          case2Text: 'Are you sure you want to rename layout "' + oldName + '" to "' + newName + '"?',
        }
    };
    const sheet = this.bottomSheet.open(ConfirmBottomSheetComponent, config);
    sheet.afterDismissed().subscribe(result => {
      if (!result) {
        return;
      }
      if (action === 'delete') {
        this.store.dispatch(DelConferenceLayoutGroup({id: id}));
      } else if (action === 'rename') {
        this.store.dispatch(UpdateConferenceLayoutGroup({id: id, name: newName}));
      }
    });
  }
}
