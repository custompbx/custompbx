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
  DelCepstralParameter,
  AddCepstralParameter,
  StoreNewCepstralParameter,
  StoreDropNewCepstralParameter,
  SwitchCepstralParameter,
  UpdateCepstralParameter
} from '../../../store/config/cepstral/config.actions.cepstral'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-cepstral', // Changed selector
  templateUrl: './cepstral.component.html', // Kept original template reference
  styleUrls: ['./cepstral.component.css']
})
export class CepstralComponent implements OnInit { // Removed OnDestroy

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
        cepstral: {} as IsimpleModule, // Initial state set to cepstral
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().cepstral); // Accessing cepstral state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().cepstral?.errorMessage || null); // Accessing cepstral error message

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
    // Initialize dispatchers here, updated for Cepstral
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewCepstralParam.bind(this),
      switchItem: this.switchCepstralParam.bind(this),
      addItem: this.newCepstralParam.bind(this),
      dropNewItem: this.dropNewCepstralParam.bind(this),
      deleteItem: this.deleteCepstralParam.bind(this),
      updateItem: this.updateCepstralParam.bind(this),
      pasteItems: null,
    };
  }

  updateCepstralParam(param: Iitem) {
    this.store.dispatch(new UpdateCepstralParameter({param: param}));
  }

  switchCepstralParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchCepstralParameter({param: newParam}));
  }

  newCepstralParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddCepstralParameter({index: index, param: param}));
  }

  deleteCepstralParam(param: Iitem) {
    this.store.dispatch(new DelCepstralParameter({param: param}));
  }

  addNewCepstralParam() {
    this.store.dispatch(new StoreNewCepstralParameter(null));
  }

  dropNewCepstralParam(index: number) {
    this.store.dispatch(new StoreDropNewCepstralParameter({index: index}));
  }

}
