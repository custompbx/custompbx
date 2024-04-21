import {Injectable} from '@angular/core';
import {Actions} from '@ngrx/effects';
import {Observable} from 'rxjs';
import {WsDataService} from '../../../services/ws-data.service';
import {createEffectForActions} from '../../../services/rxjs-helper/effects-helper';

import {
  AddAclList, AddAclNode,
  DelAclList, DelAclNode,
  DropAclList,
  GetAclLists,
  GetAclNodes,
  MoveAclListNode, StoreAclList, StoreAclLists,
  StoreAclNode,
  StoreAclNodes,
  StoreDelAclNode,
  StoreGotAclError,
  StoreMoveAclListNode,
  StoreSwitchAclNode,
  StoreUpdatedAclList,
  StoreUpdatedAclListDefault,
  StoreUpdatedAclNode,
  SwitchAclNode,
  UpdateAclList,
  UpdateAclListDefault, UpdateAclNode
} from './config.actions.acl';

@Injectable()
export class ConfigEffectsAcl {

  constructor(
    private actions: Actions,
    private ws: WsDataService,
  ) {
  }

  GetAclLists: Observable<any> = createEffectForActions(this.actions, this.ws, GetAclLists, StoreAclLists, StoreGotAclError);
  AddAclList: Observable<any> = createEffectForActions(this.actions, this.ws, AddAclList, StoreAclList, StoreGotAclError);
  UpdateAclList: Observable<any> = createEffectForActions(this.actions, this.ws, UpdateAclList, StoreUpdatedAclList, StoreGotAclError);
  UpdateAclListDefault: Observable<any> = createEffectForActions(this.actions, this.ws, UpdateAclListDefault, StoreUpdatedAclListDefault, StoreGotAclError);
  DelAclList: Observable<any> = createEffectForActions(this.actions, this.ws, DelAclList, DropAclList, StoreGotAclError);
  AddAclNode: Observable<any> = createEffectForActions(this.actions, this.ws, AddAclNode, StoreAclNode, StoreGotAclError, 'index');
  UpdateAclNode: Observable<any> = createEffectForActions(this.actions, this.ws, UpdateAclNode, StoreUpdatedAclNode, StoreGotAclError);
  DelAclNode: Observable<any> = createEffectForActions(this.actions, this.ws, DelAclNode, StoreDelAclNode, StoreGotAclError);
  SwitchAclNode: Observable<any> = createEffectForActions(this.actions, this.ws, SwitchAclNode, StoreSwitchAclNode, StoreGotAclError);
  MoveAclListNode: Observable<any> = createEffectForActions(this.actions, this.ws, MoveAclListNode, StoreMoveAclListNode, StoreGotAclError);
  GetAclNodes: Observable<any> = createEffectForActions(this.actions, this.ws, GetAclNodes, StoreAclNodes, StoreGotAclError, 'id');
}
