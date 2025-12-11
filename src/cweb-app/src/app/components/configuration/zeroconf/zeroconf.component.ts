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
  DelZeroconfParameter,
  AddZeroconfParameter,
  StoreNewZeroconfParameter,
  StoreDropNewZeroconfParameter,
  SwitchZeroconfParameter,
  UpdateZeroconfParameter
} from '../../../store/config/zeroconf/config.actions.zeroconf'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-zeroconf', // Changed selector
  templateUrl: './zeroconf.component.html', // Kept original template reference
  styleUrls: ['./zeroconf.component.css']
})
export class ZeroconfComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Zeroconf';

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
        zeroconf: {} as IsimpleModule, // Initial state set to zeroconf
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().zeroconf); // Accessing zeroconf state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().zeroconf?.errorMessage || null); // Accessing zeroconf error message

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
    // Initialize dispatchers here, updated for Zeroconf
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewZeroconfParam.bind(this),
      switchItem: this.switchZeroconfParam.bind(this),
      addItem: this.newZeroconfParam.bind(this),
      dropNewItem: this.dropNewZeroconfParam.bind(this),
      deleteItem: this.deleteZeroconfParam.bind(this),
      updateItem: this.updateZeroconfParam.bind(this),
      pasteItems: null,
    };
  }

  updateZeroconfParam(param: Iitem) {
    this.store.dispatch(new UpdateZeroconfParameter({param: param}));
  }

  switchZeroconfParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchZeroconfParameter({param: newParam}));
  }

  newZeroconfParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddZeroconfParameter({index: index, param: param}));
  }

  deleteZeroconfParam(param: Iitem) {
    this.store.dispatch(new DelZeroconfParameter({param: param}));
  }

  addNewZeroconfParam() {
    this.store.dispatch(new StoreNewZeroconfParameter(null));
  }

  dropNewZeroconfParam(index: number) {
    this.store.dispatch(new StoreDropNewZeroconfParameter({index: index}));
  }

}
