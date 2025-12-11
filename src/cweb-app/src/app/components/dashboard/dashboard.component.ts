import {Component, computed, effect, inject, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {filter, map} from 'rxjs/operators';
import {Breakpoints, BreakpointObserver} from '@angular/cdk/layout';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState, selectDataFlowState} from '../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {Idashboard} from '../../store/dataFlow/dataFlow.reducers';
import {ChartOptions, ChartDataset} from 'chart.js';
import * as pluginDataLabels from 'chartjs-plugin-datalabels';
import {ActivatedRoute} from '@angular/router';
import {State} from '../../store/config/config.state.struct';
import {FormsModule} from "@angular/forms";
import {InnerHeaderComponent} from "../inner-header/inner-header.component";
import {BaseChartDirective, provideCharts, withDefaultRegisterables} from "ng2-charts";
import {toSignal} from "@angular/core/rxjs-interop";

// --- Type Definitions for Signals (for better structure and initial values) ---
interface DashboardDataState {
  loadCounter: number;
  dashboardData: Idashboard; // The raw data (list)
  errorMessage: string | null;
}

interface ConfigsState {
  loadCounter: number;
  errorMessage: string | null;
  sofia: { id: any, loaded: boolean }; // Added sofia properties for showCard logic
}

interface ChartDataMap {
  [key: string]: ChartDataset[];
}

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, BaseChartDirective],
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css'],
  providers: [provideCharts(withDefaultRegisterables())],
  // ChangeDetectionStrategy.OnPush is often used with Signals for maximum performance
})
export class DashboardComponent implements OnDestroy {

  public dashboard: Observable<any>;
  public dashboard$: Subscription;
  public selectedIndex: number;
  private lastErrorMessage: string;
  private newContextId: number;
  private contextId: number;
  private newContextName: string;
  private newExtensionName: string;
  private expanded = [];

  public configs: Observable<any>;
  public configs$: Subscription;
  // pie
  public pieChartOptions: ChartOptions<'pie'> = {
    responsive: true,
    maintainAspectRatio: false,
    animation: {duration: 0},
    plugins: {
      legend: {
        position: 'bottom',
      },
      datalabels: {
        anchor: 'center',
        backgroundColor: null,
        color: 'white',
        display: function (context) {
          const label = context.chart.data.labels[context.dataIndex];
          // return <string>label;
          return 'HI hi';
        },
        font: {
          weight: 'bold'
        },
        padding: 6,
        // formatter: Math.round
        formatter: (value, ctx) => {
          const label = ctx.chart.data.labels[ctx.dataIndex];
          return label;
        },
      },
    }
  };
  public pieChartLegend = true;
  public pieChartColors = [
    {
      backgroundColor: ['rgba(255,0,0,0.3)', 'rgba(0,255,0,0.3)', 'rgba(0,0,255,0.3)'],
    },
  ];

  // bar
  public barChartOptions: ChartOptions = {
    animation: {
      duration: 0,
    },
    responsive: true,
    maintainAspectRatio: false,
    scales: {
      y: {
        suggestedMin: 0,
        suggestedMax: 100
      }
    },
    plugins: {
      legend: {position: 'bottom'},
    }
  };
  public sipRegbarChartOptions: ChartOptions = {
    animation: {
      duration: 0
    },
    responsive: true,
    maintainAspectRatio: false,
    scales: {
      y: {
        suggestedMin: 0,
        suggestedMax: 20
      }
    },
    plugins: {
      datalabels: {
        anchor: 'end',
        align: 'end',
      }
    }
  };
  public barChartLegend = true;
  public barChartPlugins = [pluginDataLabels];
  public barChartColors = [
    {
      backgroundColor: ['rgba(0,0,255,0.3)'],
    },
  ];

  private breakpointObserver = inject(BreakpointObserver);
  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);

// --- NgRx Observables converted to Signals (Primary State) ---
  // 1. Dashboard State
  private dashboardState = toSignal(
    this.store.pipe(select(selectDataFlowState)),
    {
      initialValue: {
        loadCounter: 0,
        dashboardData: {} as Idashboard,
        errorMessage: null
      } as DashboardDataState
    }
  );

  // 2. Configs State
  private configsState = toSignal(
    this.store.pipe(select(selectConfigurationState)),
    {
      initialValue: {
        loadCounter: 0,
        errorMessage: null,
        sofia: { id: null, loaded: false }
      } as State
    }
  );

  // Directly expose the raw data needed for the template (list)
  public list = computed(() => this.dashboardState().dashboardData);

  // Expose the configuration state (confList)
  public confList = computed(() => this.configsState());

  // Expose the combined loadCounter
  public loadCounter = computed(() =>
    Math.max(this.dashboardState().loadCounter, this.configsState().loadCounter)
  );

  // --- Computed Chart Data (Replacing pieData and barData assignment logic) ---

  // 3. Computed Pie Chart Data
  public pieData = computed<ChartDataMap>(() => {
    const list = this.list(); // Auto-tracks changes
    const dynamicMetrics = list.dynamic_metrics;

    if (!dynamicMetrics) return {};

    return {
      'percentage_used_memory': [{
        data: this.getPercentage(dynamicMetrics.percentage_used_memory)
      }],
      'percentage_disk_usage': [{
        data: this.getPercentage(dynamicMetrics.percentage_disk_usage)
      }],
    };
  });

  // 4. Computed Bar Chart Data
  public barData = computed<ChartDataMap>(() => {
    const list = this.list(); // Auto-tracks changes
    if (!list) return {};

    return {
      'core_utilization': this.convertCPUCoreData(list),
      'domain_sip_regs': this.convertSIPRegsData(list),
    };
  });

  // --- Screen Breakpoint Observable (Kept as Observable, bound with | async) ---
  public cards: Observable<any> = this.breakpointObserver.observe(Breakpoints.Handset).pipe(
    map(({ matches }) => {
      // ... (Your card definition logic remains the same) ...
      if (matches) {
        return [
          {title: 'Card 1', type: '', show: false},
          {title: 'Card 2', type: '', show: false},
          {title: 'Card 3', type: '', show: false},
          {title: 'Card 4', type: '', show: false},
          {title: 'Card 5', type: '', show: false},
          {title: 'Card 6', type: '', show: false},
          {title: 'Card 7', type: '', show: false},
        ];
      }

      return [
        {
          title: 'Sip Profiles',
          class: 'two-cols',
          type: 'table',
          field: 'sofia_profiles',
          tableFields: ['id', 'name', 'enabled', 'uri', 'state', 'started'],
          moduleName: 'sofia',
        },
        {
          title: 'Sip Gateways',
          type: 'table',
          field: 'sofia_gateways',
          tableFields: ['id', 'name', 'enabled', 'state', 'started'],
          moduleName: 'sofia',
        },
        {
          title: 'Sip Registrations',
          type: 'bar',
          field: 'domain_sip_regs',
          chartOption: this.sipRegbarChartOptions,
          chartLabels: /*<Label[]>*/['Domain'],
          moduleName: 'sofia',
        },
        {
          title: 'Channels Counter',
          type: 'number',
          field: 'calls_counter',
          show: true,
        },
        {
          title: 'Server Data',
          type: 'items',
          field: 'hostname',
          show: true,
        },
        {
          title: 'RAM',
          type: 'pie',
          field: 'percentage_used_memory',
          chartOption: this.pieChartOptions,
          chartLabels: /*<Label[]>*/['Utilized', 'Free'],
          show: true,
        },
        {
          title: 'HDD',
          type: 'pie',
          field: 'percentage_disk_usage',
          chartOption: this.pieChartOptions,
          chartLabels: /*<Label[]>*/['Utilized', 'Free'],
          show: true,
        },
        {
          title: 'CPU Cores',
          type: 'bar',
          field: 'core_utilization',
          chartOption: this.barChartOptions,
          chartLabels: /*<Label[]>*/['CPU Cores'],
          show: true,
        },
      ];
    })
  );
// --- Side Effect (Snackbar Logic) ---
  // 5. Use effect() to handle side effects (snackbars) whenever an error changes.
  private errorEffect = effect(() => {
    const dashboardError = this.dashboardState().errorMessage;
    const configError = this.configsState().errorMessage;

    // Prioritize the latest error
    const lastErrorMessage = dashboardError || configError;

    if (lastErrorMessage) {
      this._snackBar.open('Error: ' + lastErrorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  ngOnDestroy() {
    // Manual subscriptions (dashboard$, configs$) were removed by toSignal.
    // Keep this check for subscriptions possibly set up in resolvers/router data.
    if (this.route.snapshot?.data?.reconnectUpdater) {
      this.route.snapshot.data.reconnectUpdater.unsubscribe();
    }
  }

  showCard(moduleName: string): boolean {
    const conf = this.configsState(); // Read the latest configuration signal
    switch (moduleName) {
      case 'sofia':
        return conf.sofia && conf.sofia.id && conf.sofia.loaded;
      default:
        return true;
    }
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  convertCPUCoreData(data: Idashboard): ChartDataset[] {
    return data.dynamic_metrics.core_utilization.map((item, index) => {
      return {data: [Number(item.toFixed(1))], label: 'Core #' + index};
    });
  }

  getPercentage(digit: number): Array<any> {
    return [Number(digit).toFixed(1), (100 - Number(Number(digit).toFixed(1))).toFixed(1)];
  }

  convertSIPRegsData(data: Idashboard): ChartDataset[] {
    const result: ChartDataset[] = [];
    Object.keys(data?.domain_sip_regs || {}).forEach(function (key, index) {
      result.push({data: [data.domain_sip_regs[key]], label: key});
    });
    return result;
  }

}
