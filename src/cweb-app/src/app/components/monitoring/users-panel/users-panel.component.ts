import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Idetails} from '../../../store/directory/directory.reducers';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState, selectDirectoryState, selectPhoneState} from '../../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {StoreMakePhoneCall} from '../../../store/phone/phone.actions';
import {
  GetCallcenterQueues,
  SubscribeCallcenterAgents,
  SubscribeCallcenterTiers
} from '../../../store/config/callcenter/config.actions.callcenter';

@Component({
  selector: 'app-users-panel',
  templateUrl: './users-panel.component.html',
  styleUrls: ['./users-panel.component.css']
})
export class UsersPanelComponent implements OnInit, OnDestroy {

  public users: Observable<any>;
  public users$: Subscription;
  public webUsers: Observable<any>;
  public webUsers$: Subscription;
  public list: any;
  private userList: any;
  private additionalData: any;
  public listDetails: Idetails;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public timersIntervalUpdater: any;
  public domainIds: Array<string> = [];
  public userStatuses: Array<string> = [
    'Calling',
    'Registered',
    'Enabled'
  ];
  public chosenUserStatuses: Array<string> = [];
  public phone: Observable<any>;
  public phone$: Subscription;
  public phoneStatus: boolean;
  public phoneUser: string;

  public agentsListEnabled: boolean;
  public agentsListOnly: boolean;
  public config: Observable<any>;
  public config$: Subscription;
  public agentsList: object;
  public queuesList: object;
  public tiersList: {[name: string]: Array<object>};
  public agentsListByName: {[name: string]: object};

  public queueIds: Array<string> = [];
  public agentStatuses: Array<string> = [
    'Available',
    'Logged Out',
    'On Break'
  ];
  public chosenAgentStatuses: Array<string> = [];
  public usersListByAgentId: {[id: number]: object};

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.users = this.store.pipe(select(selectDirectoryState));
    this.phone = this.store.pipe(select(selectPhoneState));
    this.config = this.store.pipe(select(selectConfigurationState));
  }

  ngOnInit() {
    this.users$ = this.users.subscribe((users) => {
      this.loadCounter = users.loadCounter;
      this.list = users.domains;
      this.userList = users.users;
      this.listDetails = users.userDetails;
      this.additionalData = users.additionalData;
      this.lastErrorMessage = users.errorMessage;
      if (!this.lastErrorMessage) {
        this.selectedIndex = this.selectedIndex === 1 ? 0 : this.selectedIndex;
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
      this.usersListByAgentId = {};
      if (this.userList) {
          Object.values(this.userList).forEach(user => {
            if (!user['cc_agent']) {
              return;
            }
            this.usersListByAgentId[user['cc_agent']] = {[user['id']]: user};
          });
      }
    });
    this.phone$ = this.phone.subscribe((phone) => {
      this.phoneStatus = phone.phoneStatus.isRunning;
      if (phone.phoneCreds) {
        this.phoneUser = phone.phoneCreds.user_name || '';
      }
    });
    this.config$ = this.config.subscribe((config) => {
      if (config.callcenter && config.callcenter.queues) {
        this.queuesList = config.callcenter.queues;
      }
      if (config.callcenter && config.callcenter.agents && config.callcenter.agents.list) {
        this.agentsList = config.callcenter.agents.list;
        this.agentsListByName = <{[name: string]: object}>{};
        if (this.agentsList) {
          Object.keys(this.agentsList).forEach(key => {
            this.agentsListByName[this.agentsList[key].name] = this.agentsList[key];
          });
        }
      }
      this.tiersList = <{[name: string]: Array<object>}>{};
      if (config.callcenter && config.callcenter.tiers && config.callcenter.tiers.list) {
        Object.keys(config.callcenter.tiers.list).forEach(key => {
          if (!this.tiersList[config.callcenter.tiers.list[key].queue]) {
            this.tiersList[config.callcenter.tiers.list[key].queue] = [];
          }
          this.tiersList[config.callcenter.tiers.list[key].queue] =
            [...this.tiersList[config.callcenter.tiers.list[key].queue], config.callcenter.tiers.list[key]];
        });
      }
    });
    this.updateTimers();
    this.timersIntervalUpdater = setInterval(this.updateTimers.bind(this), 1000);
    this.updateAgentTimers();
    this.timersIntervalUpdater = setInterval(this.updateAgentTimers.bind(this), 1000);
  }

  ngOnDestroy() {
    this.timersIntervalUpdater = null;
    this.users$.unsubscribe();
    this.phone$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  trackByFnId(index, item) {
    return item.value.id;
  }

  getUserCardColor(user: object): string {
    switch (true) {
      case user['talking']:
        return 'warn';
      case user['in_call']:
        return 'accent';
      case !user['enabled']:
        return 'basic';
      case user['sip_register']:
      case user['verto_register']:
        return 'green';
      default:
        return 'basic';
    }
  }

  getUserStatus(user: object): string {
    switch (true) {
      case user['talking']:
        return 'talking';
      case user['in_call']:
        return 'ringing';
      case !user['enabled']:
        return 'user is disabled';
      case user['sip_register']:
      case user['verto_register']:
        return 'Idle';
      default:
        return 'away';
    }
  }

  updateTimers() {
    const now = Math.floor(Date.now() / 1000);
    if (!this.userList) {
      return now;
    }
    Object.values(this.userList).forEach(
      (user) => {
        user['actionTimer'] =
          now - Number(user['call_date'] || now);
      }
    );
  }

  domainFilter (): object {
    if (this.domainIds.length === 0) {
      return this.list;
    }

    const res: object = {};
    if (!this.list) {
      return res;
    }
    Object.keys(this.list).forEach(
      key => {
        if (this.domainIds.includes(key)) {
          res[key] = this.list[key];
        }
      }
    );

    return res;
  }

  usersFilter (users: []): Array<any> {
    const res = [];
    if (!users && Array.isArray(users)) {
      return res;
    }
    users.forEach(
      user => {
        if (this.chosenUserStatuses.length === 0 && !this.agentsListOnly) {
          res.push(user);
          return;
        }

        if (this.agentsListOnly) {
          if (this.chosenUserStatuses.length === 0 && user['cc_agent']) {
            res.push(user);
            return;
          } else if (!user['cc_agent']) {
            return;
          }
        }
        switch (true) {
          case this.chosenUserStatuses.includes('Calling'):
            if (user['in_call']) {
              res.push(user);
            }
            break;
          case this.chosenUserStatuses.includes('Registered'):
            if (user['sip_register'] || user['verto_register']) {
              res.push(user);
            }
            break;
          case this.chosenUserStatuses.includes('Enabled'):
            if (user['enabled']) {
              res.push(user);
            }
            break;
        }
      }
    );

    return res.sort((a, b) => a.name.localeCompare(b.name));
  }

  connectToUser(user: object, domainName: string) {
    if (!this.phoneStatus || this.phoneUser === user['name'] || user['name'] === '') {
      return;
    }
      const fullName = user['name'] + '@' + domainName;

      if (user['in_call'] && user['last_uuid']) {
        this.store.dispatch(new StoreMakePhoneCall({user: 'eavesdrop::' + user['last_uuid']}));
      } else if (user['sip_register'] || user['verto_register']) {
        this.store.dispatch(new StoreMakePhoneCall({user: user['name']}));
      }
  }

  agentsListChange() {
    if (!this.agentsListEnabled || this.agentsList) {
      return;
    }
    this.store.dispatch(new SubscribeCallcenterAgents({keep_subscription: true}));
  }

  getAgentState(user: object) {
    if (!this.agentsListEnabled || !this.agentsList || !user['cc_agent']) {
      return '';
    }

    return this.agentsList[user['cc_agent']];
  }

  getAgentCardColor(agent: object): string {
    switch (true) {
      case agent['state'] === 'Idle':
        return 'basic';
      case agent['state'] === 'Receiving':
        return 'accent';
      case agent['state'] === 'In a queue call':
        return 'warn';
      case agent['status'] === 'Logged Out':
        return 'basic';
      case agent['status'] === 'On Break':
        return 'yellow';
      case agent['status'] === 'Available':
      case agent['status'] === 'Available (On Demand)':
        return 'green';
      default:
        return 'basic';
    }
  }

  getAgentStatus(agent: object): string {
    switch (true) {
      case agent['state'] === 'Idle':
        return 'Idle';
      case agent['state'] === 'Receiving':
        return 'Receiving';
      case agent['state'] === 'In a queue call':
        return 'In a queue call';
      case agent['status'] === 'Logged Out':
        return 'Logged Out';
      case agent['status'] === 'On Break':
        return 'On Break';
      case agent['status'] === 'Available':
      case agent['status'] === 'Available (On Demand)':
        return 'Available';
      default:
        return 'Unknown';
    }
  }

  mainTabChanged(event) {
    if (event === 1) {
      this.store.dispatch(new SubscribeCallcenterAgents({keep_subscription: true}));
      this.store.dispatch(new GetCallcenterQueues(null));
      this.store.dispatch(new SubscribeCallcenterTiers({keep_subscription: true}));
    }
  }

  queueFilter (): object {
    if (this.queueIds.length === 0) {
      return this.queuesList;
    }

    const res: object = {};
    if (!this.queuesList) {
      return res;
    }
    Object.keys(this.queuesList).forEach(
      key => {
        if (this.queueIds.includes(key)) {
          res[key] = this.queuesList[key];
        }
      }
    );

    return res;
  }

  agentsFilter(queueName: string): Array<any> {
    let res = [];
    if (!this.tiersList || !this.tiersList[queueName]) {
      return res;
    }
    this.tiersList[queueName].forEach(
      tier => {
        if (!this.agentsListByName || !this.agentsListByName[tier['agent']]) {
          return;
        }
        if (this.agentsListByName[tier['agent']]) {
          res = [...res, this.agentsListByName[tier['agent']]];
        }
      }
    );

    if (this.chosenAgentStatuses.length) {
      res = res.filter(
        agent => {
          switch (true) {
            case this.chosenAgentStatuses.includes('Available'):
              return (agent.status === 'Available');
            case this.chosenAgentStatuses.includes('Logged Out'):
              return (agent.status === 'Logged Out');
            case this.chosenAgentStatuses.includes('On Break'):
              return (agent.status === 'On Break');
            default:
              return false;
          }
        }
      );

    }
    return res.sort((a, b) => a.name.localeCompare(b.name));
  }

  getUserForAgent(id: number): object {
    if (!this.usersListByAgentId || !this.usersListByAgentId[id]) {
      return null;
    }
    return this.usersListByAgentId[id];
  }

  updateAgentTimers() {
    if (!this.agentsListByName) {
      return;
    }
    const now = Math.floor(Date.now() / 1000);
    Object.keys(this.agentsListByName).forEach(
      (agentName) => {
        this.agentsListByName[agentName]['actionTimer'] =
          now - Number(this.agentsListByName[agentName]['last_status_change'] || now);
      }
    );
  }

  cutNameAndDomain(fullName: string): Array<string> {
    if (!fullName) {
      return ['', ''];
    }

    const posAt = fullName.indexOf('@');
    switch (posAt) {
      case -1:
      return [fullName, 'Agent'];
      case 0:
        return [fullName, 'Agent'];
      default:
        return [fullName.slice(0, posAt), fullName.slice(posAt + 1, fullName.length)];
    }
  }

  onlyValuesByParent(obj: object, parentId: number): Array<any> {
    if (!obj) {
      return [];
    }
    return Object.values(obj).filter(u => u.parent.id === Number(parentId));
  }

}
