import {Component, computed, effect, inject, OnDestroy, OnInit, signal} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {Observable, Subscription} from 'rxjs';
import {Idetails, Iusers, State as DirectoryState} from '../../../store/directory/directory.reducers';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState, selectDirectoryState, selectPhoneState} from '../../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {StoreCommand} from '../../../store/phone/phone.actions';
import {
  GetCallcenterQueues,
  SubscribeCallcenterAgents,
  SubscribeCallcenterTiers
} from '../../../store/config/callcenter/config.actions.callcenter';
import {FormsModule} from "@angular/forms";
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {FormatTimerPipe} from "../../../pipes/format-timer.pipe";
import {toSignal} from "@angular/core/rxjs-interop";
import {State as PhoneState} from "../../../store/phone/phone.reducers";
import {State as ConfigState} from "../../../store/config/config.state.struct";


@Component({
standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, FormatTimerPipe],
    selector: 'app-users-panel',
    templateUrl: './users-panel.component.html',
    styleUrls: ['./users-panel.component.css']
})
export class UsersPanelComponent implements OnInit, OnDestroy {

  public users: Observable<any>;
  public users$: Subscription;
  public webUsers: Observable<any>;
  public webUsers$: Subscription;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public timersIntervalUpdater: any;
  public timersIntervalUpdaterTier: any;
  public domainIds: Array<string> = [];
  public userStatuses: Array<string> = [
    'Calling',
    'Registered',
    'Enabled'
  ];
  public chosenUserStatuses: Array<string> = [];
  public phone: Observable<any>;
  public phone$: Subscription;

  public agentsListEnabled: boolean;
  public agentsListOnly: boolean;
  public config: Observable<any>;
  public config$: Subscription;

  public queueIds: Array<string> = [];
  public agentStatuses: Array<string> = [
    'Available',
    'Logged Out',
    'On Break'
  ];
  public chosenAgentStatuses: Array<string> = [];

  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);

  private directoryState = toSignal(
    this.store.pipe(select(selectDirectoryState)),
    { initialValue: {} as DirectoryState }
  );

  private phoneState = toSignal(
    this.store.pipe(select(selectPhoneState)),
    { initialValue: { phoneStatus: { isRunning: false }, phoneCreds: null } as PhoneState }
  );

  private configState = toSignal(
    this.store.pipe(select(selectConfigurationState)),
    { initialValue: { callcenter: null, errorMessage: null } as ConfigState }
  );

  // Directly exposed state data
  public list = computed(() => this.directoryState().domains);
  public userList = computed(() => this.directoryState().users);
  public listDetails = computed(() => this.directoryState().userDetails);
  public additionalData = computed(() => this.directoryState().additionalData);
  public loadCounter = computed(() => this.directoryState().loadCounter);

  // Phone state
  public phoneStatus = computed(() => this.phoneState().phoneStatus.isRunning);
  public phoneUser = computed(() => this.phoneState().phoneCreds?.user_name || '');

  // Config/Callcenter state
  public queuesList = computed(() => this.configState().callcenter?.queues || {});
  public agentsList = computed(() => this.configState().callcenter?.agents?.list || {});

  public agentsListByName = computed<{[name: string]: object}>(() => {
    const agents = this.agentsList();
    const result: {[name: string]: object} = {};
    if (agents) {
      Object.keys(agents).forEach(key => {
        result[agents[key].name] = agents[key];
      });
    }
    return result;
  });
  // 1. Dedicated Signal to act as a change trigger
  private refreshTrigger = signal(0);
  // 2. Computed Signal: Calculates timers based on the trigger
  public userListWithTimers = computed(() => {
    this.refreshTrigger();

    const directoryData = this.directoryState().users;
    if (!directoryData) {
      return [];
    }
    const now = Math.floor(Date.now() / 1000);

    // IMPORTANT: Clone the objects to create a NEW REFERENCE
    // for Angular's change detection.
    return Object.values(directoryData).map((user) => {
      // Calculate the timer and return a new object reference
      let actionTimer = 0;
      if (user.call_date) {
        const callTimestamp = Number(user.call_date);
        actionTimer = now - callTimestamp;
      }

      return {
        ...user, // Copy all existing properties
        actionTimer: actionTimer, // Add the calculated dynamic timer
      };
    });
  });


  public tiersList = computed<{[name: string]: Array<object>}>(() => {
    const tiers = this.configState().callcenter?.tiers?.list;
    const result: {[name: string]: Array<object>} = {};
    if (tiers) {
      Object.keys(tiers).forEach(key => {
        const queueName = tiers[key].queue;
        if (!result[queueName]) {
          result[queueName] = [];
        }
        result[queueName] = [...result[queueName], tiers[key]];
      });
    }
    return result;
  });

  // Derived Directory State (Users List by Agent ID)
  public usersListByAgentId = computed<{[id: number]: object}>(() => {
    const users = this.userList();
    const result: {[id: number]: object} = {};
    if (users) {
      Object.values(users).forEach(user => {
        if (user['cc_agent']) {
          result[user['cc_agent']] = {[user['id']]: user};
        }
      });
    }
    return result;
  });

  // --- Side Effect (Snackbar Logic) ---
  // Use effect() to handle side effects (snackbars) whenever an error changes.
  private directoryErrorEffect = effect(() => {
    const directoryError = this.directoryState().errorMessage;
    if (directoryError) {
      this._snackBar.open('Error: ' + directoryError + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  private configErrorEffect = effect(() => {
    const configError = this.configState().errorMessage;
    if (configError) {
      this._snackBar.open('Error: ' + configError + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });


  ngOnInit() {
    this.timersIntervalUpdater = setInterval(() => this.updateTimers(), 1000);
    this.updateAgentTimers();
    this.timersIntervalUpdaterTier = setInterval(this.updateAgentTimers.bind(this), 1000);
  }

  ngOnDestroy() {
    // Clear the interval set in ngOnInit
    if (this.timersIntervalUpdater) {
      clearInterval(this.timersIntervalUpdater);
      this.timersIntervalUpdater = null;
    }
    if (this.timersIntervalUpdaterTier) {
      clearInterval(this.timersIntervalUpdaterTier);
      this.timersIntervalUpdaterTier = null;
    }
    // Subscriptions (users$, phone$, config$) removed by toSignal
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


  // 3. Simple Interval Function: Only updates the trigger signal
  updateTimers() {
    // This single line causes userListWithTimers to re-run,
    // generating a new list with updated timers.
    this.refreshTrigger.update(v => v + 1);
  }

  domainFilter (): Record<number, object>  {
    if (this.domainIds.length === 0) {
      return this.list(); // Read signal
    }

    const res: Record<number, object> = {};
    const list = this.list(); // Read signal
    if (!list) {
      return res;
    }
    Object.keys(list).forEach(
      key => {
        if (this.domainIds.includes(key)) {
          res[key] = list[key];
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
    // Read signal values
    if (!this.phoneStatus() || this.phoneUser() === user['name'] || user['name'] === '') {
      return;
    }
    const fullName = user['name'] + '@' + domainName;

    if (user['in_call'] && user['last_uuid']) {
      this.store.dispatch(StoreCommand({callTo: 'eavesdrop::' + user['last_uuid']}));
    } else if (user['sip_register'] || user['verto_register']) {
      this.store.dispatch(StoreCommand({callTo: fullName}));
    }
  }

  agentsListChange() {
    // Read signal value
    if (!this.agentsListEnabled || Object.keys(this.agentsList()).length > 0) {
      return;
    }
    this.store.dispatch(new SubscribeCallcenterAgents({keep_subscription: true}));
  }

  getAgentState(user: object) {
    const agents = this.agentsList(); // Read signal value
    if (!this.agentsListEnabled || !agents || !user['cc_agent']) {
      return '';
    }

    return agents[user['cc_agent']];
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
      // Dispatch actions to load callcenter data when the tab is switched
      this.store.dispatch(new SubscribeCallcenterAgents({keep_subscription: true}));
      this.store.dispatch(new GetCallcenterQueues(null));
      this.store.dispatch(new SubscribeCallcenterTiers({keep_subscription: true}));
    }
  }

  queueFilter (): Record<number, object> {
    if (this.queueIds.length === 0) {
      return this.queuesList(); // Read signal
    }

    const res: Record<number, object> = {};
    const queues = this.queuesList(); // Read signal
    if (!queues) {
      return res;
    }
    Object.keys(queues).forEach(
      key => {
        if (this.queueIds.includes(key)) {
          res[key] = queues[key];
        }
      }
    );

    return res;
  }

  agentsFilter(queueName: string): Array<any> {
    let res = [];
    const tiers = this.tiersList(); // Read signal
    const agentsByName = this.agentsListByName(); // Read signal

    if (!tiers || !tiers[queueName]) {
      return res;
    }
    tiers[queueName].forEach(
      tier => {
        if (!agentsByName || !agentsByName[tier['agent']]) {
          return;
        }
        if (agentsByName[tier['agent']]) {
          res = [...res, agentsByName[tier['agent']]];
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
    const usersByAgent = this.usersListByAgentId(); // Read signal
    if (!usersByAgent || !usersByAgent[id]) {
      return null;
    }
    return usersByAgent[id];
  }

  updateAgentTimers() {
    const agentsByName = this.agentsListByName(); // Read signal
    if (!agentsByName) {
      return;
    }
    const now = Math.floor(Date.now() / 1000);
    Object.keys(agentsByName).forEach(
      (agentName) => {
        // Since agentsByName is a computed signal, this mutation only affects the local copy
        agentsByName[agentName]['actionTimer'] =
          now - Number(agentsByName[agentName]['last_status_change'] || now);
      }
    );
  }

  cutNameAndDomain(fullName: string): Array<string> { /* ... kept as is ... */
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

  protected readonly Number = Number;
}
