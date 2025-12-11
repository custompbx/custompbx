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
  DelDbParameter,
  AddDbParameter,
  StoreNewDbParameter,
  StoreDropNewDbParameter,
  SwitchDbParameter,
  UpdateDbParameter
} from '../../../store/config/db/config.actions.db'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-db', // Changed selector
  templateUrl: './db.component.html', // Kept original template reference
  styleUrls: ['./db.component.css']
})
export class DbComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Db';

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
        db: {} as IsimpleModule, // Initial state set to db
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().db); // Accessing db state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().db?.errorMessage || null); // Accessing db error message

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
    // Initialize dispatchers here, updated for Db
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewDbParam.bind(this),
      switchItem: this.switchDbParam.bind(this),
      addItem: this.newDbParam.bind(this),
      dropNewItem: this.dropNewDbParam.bind(this),
      deleteItem: this.deleteDbParam.bind(this),
      updateItem: this.updateDbParam.bind(this),
      pasteItems: null,
    };
  }

  updateDbParam(param: Iitem) {
    this.store.dispatch(new UpdateDbParameter({param: param}));
  }

  switchDbParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchDbParameter({param: newParam}));
  }

  newDbParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddDbParameter({index: index, param: param}));
  }

  deleteDbParam(param: Iitem) {
    this.store.dispatch(new DelDbParameter({param: param}));
  }

  addNewDbParam() {
    this.store.dispatch(new StoreNewDbParameter(null));
  }

  dropNewDbParam(index: number) {
    this.store.dispatch(new StoreDropNewDbParameter({index: index}));
  }

}
