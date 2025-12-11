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
  DelMongoParameter,
  AddMongoParameter,
  StoreNewMongoParameter,
  StoreDropNewMongoParameter,
  SwitchMongoParameter,
  UpdateMongoParameter
} from '../../../store/config/mongo/config.actions.mongo'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-mongo', // Changed selector
  templateUrl: './mongo.component.html', // Kept original template reference
  styleUrls: ['./mongo.component.css']
})
export class MongoComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Mongo';

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
        mongo: {} as IsimpleModule, // Initial state set to mongo
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().mongo); // Accessing mongo state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().mongo?.errorMessage || null); // Accessing mongo error message

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
    // Initialize dispatchers here, updated for Mongo
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewMongoParam.bind(this),
      switchItem: this.switchMongoParam.bind(this),
      addItem: this.newMongoParam.bind(this),
      dropNewItem: this.dropNewMongoParam.bind(this),
      deleteItem: this.deleteMongoParam.bind(this),
      updateItem: this.updateMongoParam.bind(this),
      pasteItems: null,
    };
  }

  updateMongoParam(param: Iitem) {
    this.store.dispatch(new UpdateMongoParameter({param: param}));
  }

  switchMongoParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchMongoParameter({param: newParam}));
  }

  newMongoParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddMongoParameter({index: index, param: param}));
  }

  deleteMongoParam(param: Iitem) {
    this.store.dispatch(new DelMongoParameter({param: param}));
  }

  addNewMongoParam() {
    this.store.dispatch(new StoreNewMongoParameter(null));
  }

  dropNewMongoParam(index: number) {
    this.store.dispatch(new StoreDropNewMongoParameter({index: index}));
  }

}
