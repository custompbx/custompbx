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
  DelOrekaParameter,
  AddOrekaParameter,
  StoreNewOrekaParameter,
  StoreDropNewOrekaParameter,
  SwitchOrekaParameter,
  UpdateOrekaParameter
} from '../../../store/config/oreka/config.actions.oreka'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-oreka', // Changed selector
  templateUrl: './oreka.component.html', // Kept original template reference
  styleUrls: ['./oreka.component.css']
})
export class OrekaComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Oreka';

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
        oreka: {} as IsimpleModule, // Initial state set to oreka
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().oreka); // Accessing oreka state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().oreka?.errorMessage || null); // Accessing oreka error message

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
    // Initialize dispatchers here, updated for Oreka
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewOrekaParam.bind(this),
      switchItem: this.switchOrekaParam.bind(this),
      addItem: this.newOrekaParam.bind(this),
      dropNewItem: this.dropNewOrekaParam.bind(this),
      deleteItem: this.deleteOrekaParam.bind(this),
      updateItem: this.updateOrekaParam.bind(this),
      pasteItems: null,
    };
  }

  updateOrekaParam(param: Iitem) {
    this.store.dispatch(new UpdateOrekaParameter({param: param}));
  }

  switchOrekaParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchOrekaParameter({param: newParam}));
  }

  newOrekaParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddOrekaParameter({index: index, param: param}));
  }

  deleteOrekaParam(param: Iitem) {
    this.store.dispatch(new DelOrekaParameter({param: param}));
  }

  addNewOrekaParam() {
    this.store.dispatch(new StoreNewOrekaParameter(null));
  }

  dropNewOrekaParam(index: number) {
    this.store.dispatch(new StoreDropNewOrekaParameter({index: index}));
  }

}
