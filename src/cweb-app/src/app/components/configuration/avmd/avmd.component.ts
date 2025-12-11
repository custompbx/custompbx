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
  DelAvmdParameter,
  AddAvmdParameter,
  StoreNewAvmdParameter,
  StoreDropNewAvmdParameter,
  SwitchAvmdParameter,
  UpdateAvmdParameter
} from '../../../store/config/avmd/config.actions.avmd'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-avmd', // Changed selector
  templateUrl: './avmd.component.html', // Kept original template reference
  styleUrls: ['./avmd.component.css']
})
export class AvmdComponent implements OnInit { // Removed OnDestroy

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
        avmd: {} as IsimpleModule, // Initial state set to avmd
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().avmd); // Accessing avmd state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().avmd?.errorMessage || null); // Accessing avmd error message

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
    // Initialize dispatchers here, updated for Avmd
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewAvmdParam.bind(this),
      switchItem: this.switchAvmdParam.bind(this),
      addItem: this.newAvmdParam.bind(this),
      dropNewItem: this.dropNewAvmdParam.bind(this),
      deleteItem: this.deleteAvmdParam.bind(this),
      updateItem: this.updateAvmdParam.bind(this),
      pasteItems: null,
    };
  }

  updateAvmdParam(param: Iitem) {
    this.store.dispatch(new UpdateAvmdParameter({param: param}));
  }

  switchAvmdParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchAvmdParameter({param: newParam}));
  }

  newAvmdParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddAvmdParameter({index: index, param: param}));
  }

  deleteAvmdParam(param: Iitem) {
    this.store.dispatch(new DelAvmdParameter({param: param}));
  }

  addNewAvmdParam() {
    this.store.dispatch(new StoreNewAvmdParameter(null));
  }

  dropNewAvmdParam(index: number) {
    this.store.dispatch(new StoreDropNewAvmdParameter({index: index}));
  }

}
