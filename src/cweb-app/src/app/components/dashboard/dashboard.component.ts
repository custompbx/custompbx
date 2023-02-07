import {Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {filter, map} from 'rxjs/operators';
import {Breakpoints, BreakpointObserver} from '@angular/cdk/layout';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState, selectDataFlowState} from '../../store/app.states';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {Idashboard} from '../../store/dataFlow/dataFlow.reducers';
import {ChartOptions, ChartDataset} from 'chart.js';
// import {Label} from 'ng2-charts';
import * as pluginDataLabels from 'chartjs-plugin-datalabels';
import {ActivatedRoute} from '@angular/router';
import {State} from '../../store/config/config.state.struct';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css'],
})
export class DashboardComponent implements OnInit, OnDestroy {

  public dashboard: Observable<any>;
  public dashboard$: Subscription;
  public list: Idashboard;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  private newContextId: number;
  private contextId: number;
  private newContextName: string;
  private newExtensionName: string;
  private expanded = [];
  public cards: Observable<any>;
  private pieData = {};
  private barData = {};

  public configs: Observable<any>;
  public configs$: Subscription;
  public confList: State;
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
        display: function(context) {
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
    Object.keys(data.domain_sip_regs).forEach(function (key, index) {
      result.push({data: [data.domain_sip_regs[key]], label: key});
    });
    return result;
  }

  constructor(
    private breakpointObserver: BreakpointObserver,
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.dashboard = this.store.pipe(select(selectDataFlowState));
    this.configs = this.store.pipe(select(selectConfigurationState));
  }

  ngOnInit() {
    this.dashboard$ = this.dashboard.subscribe((dashboard) => {
      this.loadCounter = dashboard.loadCounter;
      this.list = dashboard.dashboardData;

      this.pieData = {
        'percentage_used_memory':
          [{data: this.list.dynamic_metrics ? this.getPercentage(this.list.dynamic_metrics.percentage_used_memory) : <ChartDataset><any>[]}],
        'percentage_disk_usage':
          [{data: this.list.dynamic_metrics ? this.getPercentage(this.list.dynamic_metrics.percentage_disk_usage) : <ChartDataset><any>[]}],
      };
      this.barData = {
        'core_utilization': this.list.dynamic_metrics ? this.convertCPUCoreData(this.list) : <ChartDataset><any>[],
        'domain_sip_regs': this.list.domain_sip_regs ? this.convertSIPRegsData(this.list) : <ChartDataset><any>[],
      };

      this.lastErrorMessage = dashboard && dashboard.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.configs$ = this.configs.subscribe((configs) => {
      this.loadCounter = configs.loadCounter;
      this.confList = configs;
      this.lastErrorMessage = configs.errorMessage;
      if (!this.lastErrorMessage) {

      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    /** Based on the screen size, switch from standard to one column per row */
    this.cards = this.breakpointObserver.observe(Breakpoints.Handset).pipe(
      map(({matches}) => {
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
      }),
    );

  }

  ngOnDestroy() {
    this.dashboard$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
      this.route.snapshot.data.reconnectUpdater.unsubscribe();
    }
  }

  showCard(moduleName: string): boolean {
    switch (moduleName) {
      case 'sofia':
        return this.confList.sofia && this.confList.sofia.id && this.confList.sofia.loaded;
      default:
        return true;
    }
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

}
