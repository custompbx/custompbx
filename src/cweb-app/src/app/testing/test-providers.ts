import {EnvironmentProviders, Provider} from '@angular/core';
import {provideHttpClient} from '@angular/common/http';
import {provideRouter} from '@angular/router';
import {combineReducers, provideStore} from '@ngrx/store';
import {config as websocketConfig} from '../services/websocket/websocket.config';
import {reducers} from '../store/app.states';
import {provideTransloco, Translation, TranslocoLoader} from '@jsverse/transloco';
import {Observable, of} from 'rxjs';
import {localeCodes} from '../i18n/locale.model';

class TestTranslocoLoader implements TranslocoLoader {
  getTranslation(): Observable<Translation> {
    return of({
      common: {
        language: 'Language', loading: 'Loading', pageActions: 'Page actions',
        lastOperationFailed: 'Last operation failed', notifications: 'Notifications',
        dismissNotification: 'Dismiss notification', default: 'Default', filterItems: 'Filter items',
        searchItems: 'Search name, value, or id', itemsShown: '{{visible}} / {{total}} shown',
        moveItem: 'Move item', draftItems: 'Draft items', noMatchingItems: 'No matching items',
        noItemsYet: 'No items yet', nameRequired: 'Name is required.', duplicateName: 'Duplicate name',
        deleteItemConfirm: 'Delete item?', deleteItemDescription: 'Delete item description',
        copied: 'Copied!', clickDetails: 'Click to get details', parameters: 'Parameters',
        variables: 'Variables', xmlEditor: 'XML editor',
        tabs: {list: 'List', add: 'Add', deleteRename: 'Delete/Rename'},
        fields: {
          name: 'Name', domain: 'Domain', user: 'User', cache: 'Cache', numberAlias: 'Number Alias',
          cidr: 'CIDR', bulk: 'Bulk', userName: 'User name', domainName: 'Domain name',
          groupName: 'Group name', gatewayName: 'Gateway name',
        },
        actions: {
          save: 'Save', cancel: 'Cancel', add: 'Add', delete: 'Delete', update: 'Update',
          reset: 'Reset', disable: 'Disable', enable: 'Enable', addItem: '+ Add item', paste: 'Paste',
          import: 'Import', copy: 'Copy', submit: 'Submit', goToDomains: 'Go to domains',
          manage: 'Manage', done: 'Done', remove: 'Remove',
        },
      },
      feedback: {
        changesSaved: 'Changes saved successfully.',
        itemAdded: 'Item added successfully.',
        itemRemoved: 'Item removed successfully.',
        statusUpdated: 'Status updated successfully.',
        importCompleted: 'Import completed successfully.',
        itemsPasted: 'Items pasted successfully.',
        itemRenamed: 'Item renamed successfully.',
        orderUpdated: 'Order updated successfully.',
        reloadCompleted: 'Reload completed successfully.',
        moduleUnloaded: 'Module unloaded successfully.',
        moduleLoaded: 'Module loaded successfully.',
        autoloadUpdated: 'Autoload setting updated successfully.',
        itemCleared: 'Item cleared successfully.',
        configurationCleared: 'Configuration cleared successfully.',
        configurationCreated: 'Configuration created successfully.',
        commandCompleted: 'Command completed successfully.',
        errorWithDetail: 'Error: {{message}}',
      },
      directory: {
        domainSections: 'Domain sections', userSections: 'User sections', groupSections: 'Group sections',
        gatewaySections: 'Gateway sections', noDomains: 'No domains yet',
        createOrImportDomain: 'Create or import a domain.', domainRequiredForUsers: 'A domain is required.',
        domainRequiredForGroups: 'A domain is required.', domainRequiredForGateways: 'A domain is required.',
        domainLabel: 'Domain: {{name}}', userLabel: 'User: {{name}}',
        copyParametersVariables: 'Copy parameters / variables', chooseDomain: 'Choose domain',
        newDomainName: 'New domain name', newUserName: 'New user name', newGroupName: 'New group name',
        newGatewayName: 'New gateway name', userTemplates: 'User templates',
        numberOrUserName: 'Number/User name', confirmDeleteDomain: 'Delete domain?',
        confirmRenameDomain: 'Rename domain?', confirmDeleteUser: 'Delete user?',
        confirmRenameUser: 'Rename user?', confirmDeleteGroup: 'Delete group?',
        confirmRenameGroup: 'Rename group?', confirmDeleteGateway: 'Delete gateway?',
        confirmRenameGateway: 'Rename gateway?',
      },
      configuration: {
        moduleSections: 'Module sections', aclSections: 'ACL sections',
        callcenterSections: 'Callcenter sections', sofiaSections: 'Sofia sections',
        vertoSections: 'Verto sections', importFromXml: 'Import from XML', truncate: 'Truncate',
        lists: 'Lists', nodes: 'Nodes', type: 'Type', newAclListName: 'New ACL list name',
        settings: 'Settings', profiles: 'Profiles', globalSettings: 'Global Settings',
        newProfileName: 'New profile name', profileName: 'Profile name',
        noAgents: 'There are no agents yet.',
      },
      ui: {
        field: 'Field', operand: 'Operand', value: 'Value', sortField: 'Sort field',
        descending: 'Descending', queues: 'Queues', queue: 'Queue', queueName: 'Queue name',
        newAgentName: 'New agent name', tierAgentName: 'Tier agent name', getList: 'Get list',
        addFilter: 'Add filter', editFilter: 'Edit filter', addSorting: 'Add sorting',
        gateways: 'Gateways', domains: 'Domains', aliases: 'Aliases', profileCommands: 'Profile commands',
        profile: 'Profile', gateway: 'Gateway', direction: 'Direction', alias: 'Alias', parse: 'Parse',
        advertise: 'Advertise', rooms: 'Rooms', callerControlGroups: 'Caller control groups',
        controls: 'Controls', chatPermissions: 'Chat permissions', layouts: 'Layouts',
        layoutGroups: 'Layout groups', images: 'Images', group: 'Group',
        permissionProfile: 'Permission profile',
        command: 'Command', enterSaveEscCancel: 'Enter to save · Esc to cancel',
        unsaved: 'Unsaved', noRecords: 'No records',
        exportPng: 'Export PNG', exportTxt: 'Export TXT',
      },
      dashboard: {subtitle: 'Live system capacity and service health'},
      header: {profile: 'Profile', settings: 'Settings', logOut: 'Log out'},
      navigation: {
        section: 'navigation section', monitoring: 'Monitoring', dashboard: 'Dashboard',
        usersPanel: 'Users Panel', directory: 'Directory', domains: 'Domains', users: 'Users',
        groups: 'Groups', gateways: 'Gateways', configuration: 'Configuration', modules: 'Modules',
        acl: 'ACL', callcenter: 'Callcenter', sofia: 'Sofia', verto: 'Verto', switch: 'Switch',
        dialplan: 'Dialplan', contexts: 'Contexts', cdr: 'CDR', logs: 'Logs', fsCli: 'FS_CLI',
        hep: 'HEP', instances: 'Instances', globalVariables: 'Global Variables',
      },
    });
  }
}

export function customPbxTestProviders(): Array<Provider | EnvironmentProviders> {
  return [
    provideHttpClient(),
    provideRouter([]),
    provideStore({app: combineReducers(reducers)}),
    provideTransloco({
      config: {availableLangs: [...localeCodes], defaultLang: 'en', fallbackLang: 'en', reRenderOnLangChange: true},
      loader: TestTranslocoLoader,
    }),
    {
      provide: websocketConfig,
      useValue: {url: 'ws://localhost/ws', reconnectAttempts: 0},
    },
  ];
}
