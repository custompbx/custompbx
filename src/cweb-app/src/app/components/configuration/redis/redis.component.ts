import {Component, inject, computed, OnInit, effect} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../../material-module";
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl, FormsModule} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  DelRedisParameter,
  AddRedisParameter,
  StoreNewRedisParameter,
  StoreDropNewRedisParameter,
  SwitchRedisParameter,
  UpdateRedisParameter
} from '../../../store/config/redis/config.actions.redis'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-redis', // Changed selector
  templateUrl: './redis.component.html', // Kept original template reference
  styleUrls: ['./redis.component.css']
})
export class RedisComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Redis';

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private bottomSheet = inject(MatBottomSheet);
  private _snackBar = inject(MatSnackBar);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private configState = toSignal(
    this.store.pipe(select(selectConfigurationState)),
    {
      initialValue: {
        redis: {} as IsimpleModule, // Initial state set to redis
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().redis); // Accessing redis state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().redis?.errorMessage || null); // Accessing redis error message

  // --- Local Component State ---
  public selectedIndex: number = 0;
  public globalSettingsDispatchers: object;

  // --- Effect for Side Effects (Error handling) ---
  private snackbarEffect = effect(() => {
    const errorMessage = this.lastErrorMessage();
    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  ngOnInit() {
    // Initialize dispatchers here, updated for Redis
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewRedisParam.bind(this),
      switchItem: this.switchRedisParam.bind(this),
      addItem: this.newRedisParam.bind(this),
      dropNewItem: this.dropNewRedisParam.bind(this),
      deleteItem: this.deleteRedisParam.bind(this),
      updateItem: this.updateRedisParam.bind(this),
      pasteItems: null,
    };
  }

  updateRedisParam(param: Iitem) {
    this.store.dispatch(new UpdateRedisParameter({param: param}));
  }

  switchRedisParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchRedisParameter({param: newParam}));
  }

  newRedisParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddRedisParameter({index: index, param: param}));
  }

  deleteRedisParam(param: Iitem) {
    this.store.dispatch(new DelRedisParameter({param: param}));
  }

  addNewRedisParam() {
    this.store.dispatch(new StoreNewRedisParameter(null));
  }

  dropNewRedisParam(index: number) {
    this.store.dispatch(new StoreDropNewRedisParameter({index: index}));
  }

}
