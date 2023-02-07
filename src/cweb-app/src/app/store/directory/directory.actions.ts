import { Action } from '@ngrx/store';
import {StoreClearWebUserAvatar} from '../settings/settings.actions';

export enum DirectoryActionTypes {
  GetDirectoryDomains = 'GetDirectoryDomains',
  StoreGetDirectoryDomains = 'StoreGetDirectoryDomains',
  GetDirectoryDomainDetails = 'GetDirectoryDomainDetails',
  StoreGetDirectoryDomainDetails = 'StoreGetDirectoryDomainDetails',
  UpdateFailure = 'UpdateFailure',
  ClearDetails = 'ClearDetails',
  AddDirectoryDomain = 'AddDirectoryDomain',
  StoreAddDirectoryDomain = 'StoreAddDirectoryDomain',
  StoreAddNewDirectoryDomainParameter = 'StoreAddNewDirectoryDomainParameter',
  StoreAddNewDirectoryDomainVariable = 'StoreAddNewDirectoryDomainVariable',
  DeleteDirectoryDomain = 'DeleteDirectoryDomain',
  StoreDeleteDirectoryDomain = 'StoreDeleteDirectoryDomain',
  SwitchDirectoryDomain = 'SwitchDirectoryDomain',
  StoreSwitchDirectoryDomain = 'StoreSwitchDirectoryDomain',
  RenameDirectoryDomain = 'RenameDirectoryDomain',
  StoreRenameDirectoryDomain = 'StoreRenameDirectoryDomain',
  AddDirectoryDomainParameter = 'AddDirectoryDomainParameter',
  StoreAddDirectoryDomainParameter = 'StoreAddDirectoryDomainParameter',
  AddDirectoryDomainVariable = 'AddDirectoryDomainVariable',
  StoreAddDirectoryDomainVariable = 'StoreAddDirectoryDomainVariable',
  StoreDeleteNewDirectoryDomainParameter = 'StoreDeleteNewDirectoryDomainParameter',
  StoreDeleteNewDirectoryDomainVariable = 'StoreDeleteNewDirectoryDomainVariable',
  DropDirectoryDomainParameter = 'DropDirectoryDomainParameter',
  DropDirectoryDomainVariable = 'DropDirectoryDomainVariable',
  DeleteDirectoryDomainParameter = 'DeleteDirectoryDomainParameter',
  DeleteDirectoryDomainVariable = 'DeleteDirectoryDomainVariable',
  UpdateDirectoryDomainParameter = 'UpdateDirectoryDomainParameter',
  StoreUpdateDirectoryDomainParameter = 'StoreUpdateDirectoryDomainParameter',
  UpdateDirectoryDomainVariable = 'UpdateDirectoryDomainVariable',
  StoreUpdateDirectoryDomainVariable = 'StoreUpdateDirectoryDomainVariable',
  StorePasteDirectoryDomainParameters = 'StorePasteDirectoryDomainParameters',
  StorePasteDirectoryDomainVariables = 'StorePasteDirectoryDomainVariables',

  SwitchDirectoryDomainParameter = 'SwitchDirectoryDomainParameter',
  StoreSwitchDirectoryDomainParameter = 'StoreSwitchDirectoryDomainParameter',
  SwitchDirectoryDomainVariable = 'SwitchDirectoryDomainVariable',
  StoreSwitchDirectoryDomainVariable = 'StoreSwitchDirectoryDomainVariable',

  GetDirectoryUsers = 'GetDirectoryUsers',
  StoreGetDirectoryUsers = 'StoreGetDirectoryUsers',
  GetDirectoryUserDetails = 'GetDirectoryUserDetails',
  StoreGetDirectoryUserDetails = 'StoreGetDirectoryUserDetails',
  AddDirectoryUserParameter = 'AddDirectoryUserParameter',
  StoreAddDirectoryUserParameter = 'StoreAddDirectoryUserParameter',
  AddDirectoryUserVariable = 'AddDirectoryUserVariable',
  StoreAddDirectoryUserVariable = 'StoreAddDirectoryUserVariable',
  StoreAddNewDirectoryUserParameter = 'StoreAddNewDirectoryUserParameter',
  StoreAddNewDirectoryUserVariable = 'StoreAddNewDirectoryUserVariable',
  StoreDeleteNewDirectoryUserParameter = 'StoreDeleteNewDirectoryUserParameter',
  StoreDeleteNewDirectoryUserVariable = 'StoreDeleteNewDirectoryUserVariable',
  DeleteDirectoryUserParameter = 'DeleteDirectoryUserParameter',
  StoreDeleteDirectoryUserParameter = 'StoreDeleteDirectoryUserParameter',
  DeleteDirectoryUserVariable = 'DeleteDirectoryUserVariable',
  StoreDeleteDirectoryUserVariable = 'StoreDeleteDirectoryUserVariable',
  UpdateDirectoryUserParameter = 'UpdateDirectoryUserParameter',
  StoreUpdateDirectoryUserParameter = 'StoreUpdateDirectoryUserParameter',
  UpdateDirectoryUserVariable = 'UpdateDirectoryUserVariable',
  StoreUpdateDirectoryUserVariable = 'StoreUpdateDirectoryUserVariable',
  UpdateDirectoryUserCache = 'UpdateDirectoryUserCache',
  StoreUpdateDirectoryUserCache = 'StoreUpdateDirectoryUserCache',
  UpdateDirectoryUserNumberAlias = 'UpdateDirectoryUserNumberAlias',
  StoreUpdateDirectoryUserNumberAlias = 'StoreUpdateDirectoryUserNumberAlias',
  UpdateDirectoryUserCidr = 'UpdateDirectoryUserCidr',
  StoreUpdateDirectoryUserCidr = 'StoreUpdateDirectoryUserCidr',
  AddDirectoryUser = 'AddDirectoryUser',
  StoreAddDirectoryUser = 'StoreAddDirectoryUser',
  DeleteDirectoryUser = 'DeleteDirectoryUser',
  StoreDeleteDirectoryUser = 'StoreDeleteDirectoryUser',
  UpdateDirectoryUserName = 'UpdateDirectoryUserName',
  StoreUpdateDirectoryUserName = 'StoreUpdateDirectoryUserName',
  StorePasteDirectoryUserParameters = 'StorePasteDirectoryUserParameters',
  StorePasteDirectoryUserVariables = 'StorePasteDirectoryUserVariables',

  SwitchDirectoryUser = 'SwitchDirectoryUser',
  StoreSwitchDirectoryUser = 'StoreSwitchDirectoryUser',
  SwitchDirectoryUserParameter = 'SwitchDirectoryUserParameter',
  StoreSwitchDirectoryUserParameter = 'StoreSwitchDirectoryUserParameter',
  SwitchDirectoryUserVariable = 'SwitchDirectoryUserVariable',
  StoreSwitchDirectoryUserVariable = 'StoreSwitchDirectoryUserVariable',

  GetDirectoryGroups = 'GetDirectoryGroups',
  StoreGetDirectoryGroups = 'StoreGetDirectoryGroups',
  GetDirectoryGroupUsers = 'GetDirectoryGroupUsers',
  StoreGetDirectoryGroupUsers = 'StoreGetDirectoryGroupUsers',
  AddNewDirectoryGroup = 'AddNewDirectoryGroup',
  StoreAddNewDirectoryGroup = 'StoreAddNewDirectoryGroup',
  DeleteDirectoryGroup = 'DeleteDirectoryGroup',
  StoreDeleteDirectoryGroup = 'StoreDeleteDirectoryGroup',
  UpdateDirectoryGroupName = 'UpdateDirectoryGroupName',
  StoreUpdateDirectoryGroupName = 'StoreUpdateDirectoryGroupName',
  AddDirectoryGroupUser = 'AddDirectoryGroupUser',
  StoreAddDirectoryGroupUser = 'StoreAddDirectoryGroupUser',
  DeleteDirectoryGroupUser = 'DeleteDirectoryGroupUser',
  StoreDeleteDirectoryGroupUser = 'StoreDeleteDirectoryGroupUser',

  GetDirectoryUserGateways = 'GetDirectoryUserGateways',
  StoreGetDirectoryUserGateways = 'StoreGetDirectoryUserGateways',
  GetDirectoryUserGatewayDetails = 'GetDirectoryUserGatewayDetails',
  StoreGetDirectoryUserGatewayDetails = 'StoreGetDirectoryUserGatewayDetails',
  StoreNewDirectoryUserGatewayParameter = 'StoreNewDirectoryUserGatewayParameter',
  DropNewDirectoryUserGatewayParameter = 'DropNewDirectoryUserGatewayParameter',
  AddDirectoryUserGatewayParameter = 'AddDirectoryUserGatewayParameter',
  StoreDirectoryUserGatewayParameter = 'StoreDirectoryUserGatewayParameter',
  DeleteDirectoryUserGatewayParameter = 'DeleteDirectoryUserGatewayParameter',
  StoreDeleteDirectoryUserGatewayParameter = 'StoreDeleteDirectoryUserGatewayParameter',
  UpdateDirectoryUserGatewayParameter = 'UpdateDirectoryUserGatewayParameter',
  StoreUpdateDirectoryUserGatewayParameter = 'StoreUpdateDirectoryUserGatewayParameter',
  SwitchDirectoryUserGatewayParameter = 'SwitchDirectoryUserGatewayParameter',
  StoreSwitchDirectoryUserGatewayParameter = 'StoreSwitchDirectoryUserGatewayParameter',
  AddDirectoryUserGateway = 'AddDirectoryUserGateway',
  StoreAddDirectoryUserGateway = 'StoreAddDirectoryUserGateway',
  DeleteDirectoryUserGateway = 'DeleteDirectoryUserGateway',
  StoreDeleteDirectoryUserGateway = 'StoreDeleteDirectoryUserGateway',
  UpdateDirectoryUserGatewayName = 'UpdateDirectoryUserGatewayName',
  StoreUpdateDirectoryUserGatewayName = 'StoreUpdateDirectoryUserGatewayName',
  StorePasteDirectoryUserGatewayParameters = 'StorePasteDirectoryUserGatewayParameters',
  StorePasteDirectoryUserGatewayVariables = 'StorePasteDirectoryUserGatewayVariables',

  UpdateDirectoryUserGatewayVariable = 'UpdateDirectoryUserGatewayVariable',
  StoreUpdateDirectoryUserGatewayVariable = 'StoreUpdateDirectoryUserGatewayVariable',
  SwitchDirectoryUserGatewayVariable = 'SwitchDirectoryUserGatewayVariable',
  StoreSwitchDirectoryUserGatewayVariable = 'StoreSwitchDirectoryUserGatewayVariable',
  AddDirectoryUserGatewayVariable = 'AddDirectoryUserGatewayVariable',
  StoreAddDirectoryUserGatewayVariable = 'StoreAddDirectoryUserGatewayVariable',
  DeleteDirectoryUserGatewayVariable = 'DeleteDirectoryUserGatewayVariable',
  StoreDeleteDirectoryUserGatewayVariable = 'StoreDeleteDirectoryUserGatewayVariable',
  StoreNewDirectoryUserGatewayVariable = 'StoreNewDirectoryUserGatewayVariable',
  DropNewDirectoryUserGatewayVariable = 'DropNewDirectoryUserGatewayVariable',

  ImportDirectory = 'ImportDirectory',
  ReduceLoadCounter = 'ReduceLoadCounter',
  GetWebUsersByDirectory = 'GetWebUsersByDirectory',
  StoreGetWebUsersByDirectory = 'StoreGetWebUsersByDirectory',
  ImportXMLDomain = 'ImportXMLDomain',
  StoreImportXMLDomain = 'StoreImportXMLDomain',
  ImportXMLDomainUser = 'ImportXMLDomainUser',
  StoreImportXMLDomainUser = 'StoreImportXMLDomainUser',

  GetWebDirectoryUsersTemplatesList = 'GetWebDirectoryUsersTemplatesList',
  StoreGetWebDirectoryUsersTemplatesList = 'StoreGetWebDirectoryUsersTemplatesList',
  GetWebDirectoryUsersTemplateForm = 'GetWebDirectoryUsersTemplateForm',
  StoreGetWebDirectoryUsersTemplateForm = 'StoreGetWebDirectoryUsersTemplateForm',
  CreateWebDirectoryUsersByTemplate = 'CreateWebDirectoryUsersByTemplate',
  StoreCreateWebDirectoryUsersByTemplate = 'StoreCreateWebDirectoryUsersByTemplate',

  StoreDirectoryError = 'StoreDirectoryError',
}

export class ImportXMLDomainUser implements Action {
  readonly type = DirectoryActionTypes.ImportXMLDomainUser;
  constructor(public payload: any) {}
}

export class StoreImportXMLDomainUser implements Action {
  readonly type = DirectoryActionTypes.StoreImportXMLDomainUser;
  constructor(public payload: any) {}
}

export class ImportXMLDomain implements Action {
  readonly type = DirectoryActionTypes.ImportXMLDomain;
  constructor(public payload: any) {}
}

export class StoreImportXMLDomain implements Action {
  readonly type = DirectoryActionTypes.StoreImportXMLDomain;
  constructor(public payload: any) {}
}

export class ReduceLoadCounter implements Action {
  readonly type = DirectoryActionTypes.ReduceLoadCounter;
  // constructor(public payload: any) {}
}

export class ImportDirectory implements Action {
  readonly type = DirectoryActionTypes.ImportDirectory;
  constructor(public payload: any) {}
}

// domains
export class SwitchDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.SwitchDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.StoreSwitchDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class SwitchDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.SwitchDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.StoreSwitchDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryDomains implements Action {
  readonly type = DirectoryActionTypes.StoreGetDirectoryDomains;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryDomainDetails implements Action {
  readonly type = DirectoryActionTypes.StoreGetDirectoryDomainDetails;
  constructor(public payload: any) {}
}

export class UpdateFailure implements Action {
  readonly type = DirectoryActionTypes.UpdateFailure;
  constructor(public payload: any) {}
}

export class GetDirectoryDomains implements Action {
  readonly type = DirectoryActionTypes.GetDirectoryDomains;
  constructor(public payload: any) {}
}

export class StoreAddNewDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.StoreAddNewDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class StoreAddNewDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.StoreAddNewDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class GetDirectoryDomainDetails implements Action {
  readonly type = DirectoryActionTypes.GetDirectoryDomainDetails;
  constructor(public payload: any) {}
}

export class ClearDetails implements Action {
  readonly type = DirectoryActionTypes.ClearDetails;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryDomain implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryDomain;
  constructor(public payload: any) {}
}

export class AddDirectoryDomain implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryDomain;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryDomain implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryDomain;
  constructor(public payload: any) {}
}

export class DeleteDirectoryDomain implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryDomain;
  constructor(public payload: any) {}
}

export class RenameDirectoryDomain implements Action {
  readonly type = DirectoryActionTypes.RenameDirectoryDomain;
  constructor(public payload: any) {}
}

export class SwitchDirectoryDomain implements Action {
  readonly type = DirectoryActionTypes.SwitchDirectoryDomain;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryDomain implements Action {
  readonly type = DirectoryActionTypes.StoreSwitchDirectoryDomain;
  constructor(public payload: any) {}
}

export class AddDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class AddDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class StoreDeleteNewDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteNewDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class StoreDeleteNewDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteNewDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class DropDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.DropDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class DropDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.DropDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class DeleteDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class DeleteDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class UpdateDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryDomainParameter implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryDomainParameter;
  constructor(public payload: any) {}
}

export class UpdateDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryDomainVariable implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryDomainVariable;
  constructor(public payload: any) {}
}

export class StorePasteDirectoryDomainParameters implements Action {
  readonly type = DirectoryActionTypes.StorePasteDirectoryDomainParameters;
  constructor(public payload: any) {}
}

export class StorePasteDirectoryDomainVariables implements Action {
  readonly type = DirectoryActionTypes.StorePasteDirectoryDomainVariables;
  constructor(public payload: any) {}
}

export class StoreRenameDirectoryDomain implements Action {
  readonly type = DirectoryActionTypes.StoreRenameDirectoryDomain;
  constructor(public payload: any) {}
}

// users
export class StorePasteDirectoryUserParameters implements Action {
  readonly type = DirectoryActionTypes.StorePasteDirectoryUserParameters;
  constructor(public payload: any) {}
}

export class StorePasteDirectoryUserVariables implements Action {
  readonly type = DirectoryActionTypes.StorePasteDirectoryUserVariables;
  constructor(public payload: any) {}
}

export class SwitchDirectoryUser implements Action {
  readonly type = DirectoryActionTypes.SwitchDirectoryUser;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryUser implements Action {
  readonly type = DirectoryActionTypes.StoreSwitchDirectoryUser;
  constructor(public payload: any) {}
}

export class SwitchDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.SwitchDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.StoreSwitchDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class SwitchDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.SwitchDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.StoreSwitchDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class GetDirectoryUsers implements Action {
  readonly type = DirectoryActionTypes.GetDirectoryUsers;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryUsers implements Action {
  readonly type = DirectoryActionTypes.StoreGetDirectoryUsers;
  constructor(public payload: any) {}
}

export class GetDirectoryUserDetails implements Action {
  readonly type = DirectoryActionTypes.GetDirectoryUserDetails;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryUserDetails implements Action {
  readonly type = DirectoryActionTypes.StoreGetDirectoryUserDetails;
  constructor(public payload: any) {}
}

export class UpdateDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class UpdateDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class UpdateDirectoryUserName implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserName;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserName implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserName;
  constructor(public payload: any) {}
}

export class UpdateDirectoryUserCache implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserCache;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserCache implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserCache;
  constructor(public payload: any) {}
}

export class UpdateDirectoryUserNumberAlias implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserNumberAlias;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserNumberAlias implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserNumberAlias;
  constructor(public payload: any) {}
}

export class UpdateDirectoryUserCidr implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserCidr;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserCidr implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserCidr;
  constructor(public payload: any) {}
}

export class AddDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class AddDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class StoreDeleteNewDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteNewDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class StoreDeleteNewDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteNewDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class StoreAddNewDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.StoreAddNewDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class StoreAddNewDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.StoreAddNewDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class DeleteDirectoryUserParameter implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryUserParameter;
  constructor(public payload: any) {}
}

export class DeleteDirectoryUserVariable implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryUserVariable;
  constructor(public payload: any) {}
}

export class AddDirectoryUser implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryUser;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryUser implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryUser;
  constructor(public payload: any) {}
}

export class DeleteDirectoryUser implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryUser;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryUser implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryUser;
  constructor(public payload: any) {}
}

// groups
export class GetDirectoryGroups implements Action {
  readonly type = DirectoryActionTypes.GetDirectoryGroups;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryGroups implements Action {
  readonly type = DirectoryActionTypes.StoreGetDirectoryGroups;
  constructor(public payload: any) {}
}

export class GetDirectoryGroupUsers implements Action {
  readonly type = DirectoryActionTypes.GetDirectoryGroupUsers;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryGroupUsers implements Action {
  readonly type = DirectoryActionTypes.StoreGetDirectoryGroupUsers;
  constructor(public payload: any) {}
}

export class AddDirectoryGroupUser implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryGroupUser;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryGroupUser implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryGroupUser;
  constructor(public payload: any) {}
}

export class DeleteDirectoryGroupUser implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryGroupUser;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryGroupUser implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryGroupUser;
  constructor(public payload: any) {}
}

export class UpdateDirectoryGroupName implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryGroupName;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryGroupName implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryGroupName;
  constructor(public payload: any) {}
}

export class AddNewDirectoryGroup implements Action {
  readonly type = DirectoryActionTypes.AddNewDirectoryGroup;
  constructor(public payload: any) {}
}

export class StoreAddNewDirectoryGroup implements Action {
  readonly type = DirectoryActionTypes.StoreAddNewDirectoryGroup;
  constructor(public payload: any) {}
}

export class DeleteDirectoryGroup implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryGroup;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryGroup implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryGroup;

  constructor(public payload: any) {
  }
}

// gateways
export class GetDirectoryUserGateways implements Action {
  readonly type = DirectoryActionTypes.GetDirectoryUserGateways;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryUserGateways implements Action {
  readonly type = DirectoryActionTypes.StoreGetDirectoryUserGateways;
  constructor(public payload: any) {}
}

export class GetDirectoryUserGatewayDetails implements Action {
  readonly type = DirectoryActionTypes.GetDirectoryUserGatewayDetails;
  constructor(public payload: any) {}
}

export class StoreGetDirectoryUserGatewayDetails implements Action {
  readonly type = DirectoryActionTypes.StoreGetDirectoryUserGatewayDetails;
  constructor(public payload: any) {}
}


export class UpdateDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class UpdateDirectoryUserGatewayName implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserGatewayName;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserGatewayName implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserGatewayName;
  constructor(public payload: any) {}
}

export class AddDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class StoreDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.StoreDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class StoreNewDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.StoreNewDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class DropNewDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.DropNewDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class DeleteDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class SwitchDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.SwitchDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryUserGatewayParameter implements Action {
  readonly type = DirectoryActionTypes.StoreSwitchDirectoryUserGatewayParameter;
  constructor(public payload: any) {}
}

export class AddDirectoryUserGateway implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryUserGateway;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryUserGateway implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryUserGateway;
  constructor(public payload: any) {}
}

export class DeleteDirectoryUserGateway implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryUserGateway;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryUserGateway implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryUserGateway;
  constructor(public payload: any) {}
}

export class UpdateDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.UpdateDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class StoreUpdateDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.StoreUpdateDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class SwitchDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.SwitchDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class StoreSwitchDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.StoreSwitchDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class AddDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.AddDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class StoreAddDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.StoreAddDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class DeleteDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.DeleteDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class StoreDeleteDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.StoreDeleteDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class StoreNewDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.StoreNewDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class DropNewDirectoryUserGatewayVariable implements Action {
  readonly type = DirectoryActionTypes.DropNewDirectoryUserGatewayVariable;
  constructor(public payload: any) {}
}

export class StorePasteDirectoryUserGatewayParameters implements Action {
  readonly type = DirectoryActionTypes.StorePasteDirectoryUserGatewayParameters;
  constructor(public payload: any) {}
}

export class StorePasteDirectoryUserGatewayVariables implements Action {
  readonly type = DirectoryActionTypes.StorePasteDirectoryUserGatewayVariables;
  constructor(public payload: any) {}
}

export class GetWebUsersByDirectory implements Action {
  readonly type = DirectoryActionTypes.GetWebUsersByDirectory;
  constructor(public payload: any) {}
}

export class StoreGetWebUsersByDirectory implements Action {
  readonly type = DirectoryActionTypes.StoreGetWebUsersByDirectory;
  constructor(public payload: any) {}
}

export class GetWebDirectoryUsersTemplatesList implements Action {
  readonly type = DirectoryActionTypes.GetWebDirectoryUsersTemplatesList;
  constructor(public payload: any) {}
}

export class StoreGetWebDirectoryUsersTemplatesList implements Action {
  readonly type = DirectoryActionTypes.StoreGetWebDirectoryUsersTemplatesList;
  constructor(public payload: any) {}
}

export class GetWebDirectoryUsersTemplateForm implements Action {
  readonly type = DirectoryActionTypes.GetWebDirectoryUsersTemplateForm;
  constructor(public payload: any) {}
}

export class StoreGetWebDirectoryUsersTemplateForm implements Action {
  readonly type = DirectoryActionTypes.StoreGetWebDirectoryUsersTemplateForm;
  constructor(public payload: any) {}
}

export class CreateWebDirectoryUsersByTemplate implements Action {
  readonly type = DirectoryActionTypes.CreateWebDirectoryUsersByTemplate;
  constructor(public payload: any) {}
}

export class StoreCreateWebDirectoryUsersByTemplate implements Action {
  readonly type = DirectoryActionTypes.StoreCreateWebDirectoryUsersByTemplate;
  constructor(public payload: any) {}
}

export class StoreDirectoryError implements Action {
  readonly type = DirectoryActionTypes.StoreDirectoryError;
  constructor(public payload: any) {}
}

export type All =
  | ReduceLoadCounter
  | ImportDirectory
  | SwitchDirectoryDomainParameter
  | StoreSwitchDirectoryDomainParameter
  | SwitchDirectoryDomainVariable
  | StoreSwitchDirectoryDomainVariable
  | StoreGetDirectoryDomainDetails
  | UpdateFailure
  | GetDirectoryDomains
  | StoreGetDirectoryDomains
  | GetDirectoryDomainDetails
  | StoreAddNewDirectoryDomainParameter
  | StoreAddNewDirectoryDomainVariable
  | ClearDetails
  | SwitchDirectoryDomain
  | StoreSwitchDirectoryDomain
  | AddDirectoryDomain
  | StoreAddDirectoryDomain
  | DeleteDirectoryDomain
  | RenameDirectoryDomain
  | StoreDeleteDirectoryDomain
  | AddDirectoryDomainParameter
  | AddDirectoryDomainVariable
  | StoreDeleteNewDirectoryDomainParameter
  | StoreDeleteNewDirectoryDomainVariable
  | StoreAddDirectoryDomainParameter
  | StoreAddDirectoryDomainVariable
  | DropDirectoryDomainParameter
  | DropDirectoryDomainVariable
  | DeleteDirectoryDomainParameter
  | DeleteDirectoryDomainVariable
  | UpdateDirectoryDomainParameter
  | StoreUpdateDirectoryDomainParameter
  | UpdateDirectoryDomainVariable
  | StoreUpdateDirectoryDomainVariable
  | StorePasteDirectoryDomainParameters
  | StorePasteDirectoryDomainVariables
  | StoreRenameDirectoryDomain
  | StorePasteDirectoryUserParameters
  | StorePasteDirectoryUserVariables
  | SwitchDirectoryUser
  | StoreSwitchDirectoryUser
  | SwitchDirectoryUserParameter
  | StoreSwitchDirectoryUserParameter
  | SwitchDirectoryUserVariable
  | StoreSwitchDirectoryUserVariable
  | GetDirectoryUsers
  | StoreGetDirectoryUsers
  | GetDirectoryUserDetails
  | StoreGetDirectoryUserDetails
  | UpdateDirectoryUserParameter
  | StoreUpdateDirectoryUserParameter
  | UpdateDirectoryUserVariable
  | StoreUpdateDirectoryUserVariable
  | UpdateDirectoryUserName
  | StoreUpdateDirectoryUserName
  | UpdateDirectoryUserCache
  | StoreUpdateDirectoryUserCache
  | UpdateDirectoryUserNumberAlias
  | StoreUpdateDirectoryUserNumberAlias
  | UpdateDirectoryUserCidr
  | StoreUpdateDirectoryUserCidr
  | AddDirectoryUserParameter
  | StoreAddDirectoryUserParameter
  | AddDirectoryUserVariable
  | StoreAddDirectoryUserVariable
  | StoreDeleteNewDirectoryUserParameter
  | StoreDeleteNewDirectoryUserVariable
  | StoreAddNewDirectoryUserParameter
  | StoreAddNewDirectoryUserVariable
  | StoreDeleteDirectoryUserParameter
  | StoreDeleteDirectoryUserVariable
  | DeleteDirectoryUserParameter
  | DeleteDirectoryUserVariable
  | AddDirectoryUser
  | StoreAddDirectoryUser
  | DeleteDirectoryUser
  | StoreDeleteDirectoryUser
  | GetDirectoryGroups
  | StoreGetDirectoryGroups
  | GetDirectoryGroupUsers
  | StoreGetDirectoryGroupUsers
  | AddNewDirectoryGroup
  | StoreAddNewDirectoryGroup
  | DeleteDirectoryGroup
  | StoreDeleteDirectoryGroup
  | UpdateDirectoryGroupName
  | StoreUpdateDirectoryGroupName
  | AddDirectoryGroupUser
  | StoreAddDirectoryGroupUser
  | DeleteDirectoryGroupUser
  | StoreDeleteDirectoryGroupUser
  | GetDirectoryUserGateways
  | StoreGetDirectoryUserGateways
  | GetDirectoryUserGatewayDetails
  | StoreGetDirectoryUserGatewayDetails
  | StoreNewDirectoryUserGatewayParameter
  | DropNewDirectoryUserGatewayParameter
  | AddDirectoryUserGatewayParameter
  | StoreDirectoryUserGatewayParameter
  | DeleteDirectoryUserGatewayParameter
  | StoreDeleteDirectoryUserGatewayParameter
  | UpdateDirectoryUserGatewayParameter
  | StoreUpdateDirectoryUserGatewayParameter
  | AddDirectoryUserGateway
  | StoreAddDirectoryUserGateway
  | UpdateDirectoryUserGatewayName
  | StoreUpdateDirectoryUserGatewayName
  | DeleteDirectoryUserGateway
  | StoreDeleteDirectoryUserGateway
  | UpdateDirectoryUserGatewayVariable
  | StoreUpdateDirectoryUserGatewayVariable
  | SwitchDirectoryUserGatewayVariable
  | StoreSwitchDirectoryUserGatewayVariable
  | AddDirectoryUserGatewayVariable
  | StoreAddDirectoryUserGatewayVariable
  | DeleteDirectoryUserGatewayVariable
  | StoreDeleteDirectoryUserGatewayVariable
  | StoreNewDirectoryUserGatewayVariable
  | DropNewDirectoryUserGatewayVariable
  | StorePasteDirectoryUserGatewayParameters
  | StorePasteDirectoryUserGatewayVariables
  | SwitchDirectoryUserGatewayParameter
  | StoreSwitchDirectoryUserGatewayParameter
  | GetWebUsersByDirectory
  | StoreGetWebUsersByDirectory
  | ImportXMLDomain
  | StoreImportXMLDomain
  | ImportXMLDomainUser
  | StoreImportXMLDomainUser
  | GetWebDirectoryUsersTemplatesList
  | StoreGetWebDirectoryUsersTemplatesList
  | GetWebDirectoryUsersTemplateForm
  | StoreGetWebDirectoryUsersTemplateForm
  | CreateWebDirectoryUsersByTemplate
  | StoreCreateWebDirectoryUsersByTemplate
  | StoreDirectoryError
  ;
