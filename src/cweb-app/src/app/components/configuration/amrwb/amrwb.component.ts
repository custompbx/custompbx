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
  DelAmrwbParameter,
  AddAmrwbParameter,
  StoreNewAmrwbParameter,
  StoreDropNewAmrwbParameter,
  SwitchAmrwbParameter,
  UpdateAmrwbParameter
} from '../../../store/config/amrwb/config.actions.amrwb'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-amrwb', // Changed selector
  templateUrl: './amrwb.component.html', // Kept original template reference
  styleUrls: ['./amrwb.component.css']
})
export class AmrwbComponent implements OnInit { // Removed OnDestroy

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
        amrwb: {} as IsimpleModule, // Initial state set to amrwb
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().amrwb); // Accessing amrwb state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().amrwb?.errorMessage || null); // Accessing amrwb error message

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
    // Initialize dispatchers here, updated for Amrwb
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewAmrwbParam.bind(this),
      switchItem: this.switchAmrwbParam.bind(this),
      addItem: this.newAmrwbParam.bind(this),
      dropNewItem: this.dropNewAmrwbParam.bind(this),
      deleteItem: this.deleteAmrwbParam.bind(this),
      updateItem: this.updateAmrwbParam.bind(this),
      pasteItems: null,
    };
  }

  updateAmrwbParam(param: Iitem) {
    this.store.dispatch(new UpdateAmrwbParameter({param: param}));
  }

  switchAmrwbParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchAmrwbParameter({param: newParam}));
  }

  newAmrwbParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddAmrwbParameter({index: index, param: param}));
  }

  deleteAmrwbParam(param: Iitem) {
    this.store.dispatch(new DelAmrwbParameter({param: param}));
  }

  addNewAmrwbParam() {
    this.store.dispatch(new StoreNewAmrwbParameter(null));
  }

  dropNewAmrwbParam(index: number) {
    this.store.dispatch(new StoreDropNewAmrwbParameter({index: index}));
  }

}
