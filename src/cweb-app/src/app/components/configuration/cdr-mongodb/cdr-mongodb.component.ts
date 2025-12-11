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
  DelCdrMongodbParameter,
  AddCdrMongodbParameter,
  StoreNewCdrMongodbParameter,
  StoreDropNewCdrMongodbParameter,
  SwitchCdrMongodbParameter,
  UpdateCdrMongodbParameter
} from '../../../store/config/cdr_mongodb/config.actions.cdr_mongodb'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-cdr-mongodb', // Changed selector
  templateUrl: './cdr-mongodb.component.html', // Kept original template reference
  styleUrls: ['./cdr-mongodb.component.css']
})
export class CdrMongodbComponent implements OnInit { // Removed OnDestroy

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
        cdr_mongodb: {} as IsimpleModule, // Initial state set to cdr_mongodb
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().cdr_mongodb); // Accessing cdr_mongodb state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().cdr_mongodb?.errorMessage || null); // Accessing cdr_mongodb error message

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
    // Initialize dispatchers here, updated for CdrMongodb
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewCdrMongodbParam.bind(this),
      switchItem: this.switchCdrMongodbParam.bind(this),
      addItem: this.newCdrMongodbParam.bind(this),
      dropNewItem: this.dropNewCdrMongodbParam.bind(this),
      deleteItem: this.deleteCdrMongodbParam.bind(this),
      updateItem: this.updateCdrMongodbParam.bind(this),
      pasteItems: null,
    };
  }

  updateCdrMongodbParam(param: Iitem) {
    this.store.dispatch(new UpdateCdrMongodbParameter({param: param}));
  }

  switchCdrMongodbParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchCdrMongodbParameter({param: newParam}));
  }

  newCdrMongodbParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddCdrMongodbParameter({index: index, param: param}));
  }

  deleteCdrMongodbParam(param: Iitem) {
    this.store.dispatch(new DelCdrMongodbParameter({param: param}));
  }

  addNewCdrMongodbParam() {
    this.store.dispatch(new StoreNewCdrMongodbParameter(null));
  }

  dropNewCdrMongodbParam(index: number) {
    this.store.dispatch(new StoreDropNewCdrMongodbParameter({index: index}));
  }

}
