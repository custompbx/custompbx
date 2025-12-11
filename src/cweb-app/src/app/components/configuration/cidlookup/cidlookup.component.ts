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
  DelCidlookupParameter,
  AddCidlookupParameter,
  StoreNewCidlookupParameter,
  StoreDropNewCidlookupParameter,
  SwitchCidlookupParameter,
  UpdateCidlookupParameter
} from '../../../store/config/cidlookup/config.actions.cidlookup'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-cidlookup', // Changed selector
  templateUrl: './cidlookup.component.html', // Kept original template reference
  styleUrls: ['./cidlookup.component.css']
})
export class CidlookupComponent implements OnInit { // Removed OnDestroy

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
        cidlookup: {} as IsimpleModule, // Initial state set to cidlookup
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().cidlookup); // Accessing cidlookup state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().cidlookup?.errorMessage || null); // Accessing cidlookup error message

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
    // Initialize dispatchers here, updated for Cidlookup
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewCidlookupParam.bind(this),
      switchItem: this.switchCidlookupParam.bind(this),
      addItem: this.newCidlookupParam.bind(this),
      dropNewItem: this.dropNewCidlookupParam.bind(this),
      deleteItem: this.deleteCidlookupParam.bind(this),
      updateItem: this.updateCidlookupParam.bind(this),
      pasteItems: null,
    };
  }

  updateCidlookupParam(param: Iitem) {
    this.store.dispatch(new UpdateCidlookupParameter({param: param}));
  }

  switchCidlookupParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchCidlookupParameter({param: newParam}));
  }

  newCidlookupParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddCidlookupParameter({index: index, param: param}));
  }

  deleteCidlookupParam(param: Iitem) {
    this.store.dispatch(new DelCidlookupParameter({param: param}));
  }

  addNewCidlookupParam() {
    this.store.dispatch(new StoreNewCidlookupParameter(null));
  }

  dropNewCidlookupParam(index: number) {
    this.store.dispatch(new StoreDropNewCidlookupParameter({index: index}));
  }

}
