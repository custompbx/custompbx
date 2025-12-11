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
  DelCurlParameter,
  AddCurlParameter,
  StoreNewCurlParameter,
  StoreDropNewCurlParameter,
  SwitchCurlParameter,
  UpdateCurlParameter
} from '../../../store/config/curl/config.actions.curl'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-curl', // Changed selector
  templateUrl: './curl.component.html', // Kept original template reference
  styleUrls: ['./curl.component.css']
})
export class CurlComponent implements OnInit { // Removed OnDestroy

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
        curl: {} as IsimpleModule, // Initial state set to curl
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().curl); // Accessing curl state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().curl?.errorMessage || null); // Accessing curl error message

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
    // Initialize dispatchers here, updated for Curl
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewCurlParam.bind(this),
      switchItem: this.switchCurlParam.bind(this),
      addItem: this.newCurlParam.bind(this),
      dropNewItem: this.dropNewCurlParam.bind(this),
      deleteItem: this.deleteCurlParam.bind(this),
      updateItem: this.updateCurlParam.bind(this),
      pasteItems: null,
    };
  }

  updateCurlParam(param: Iitem) {
    this.store.dispatch(new UpdateCurlParameter({param: param}));
  }

  switchCurlParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchCurlParameter({param: newParam}));
  }

  newCurlParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddCurlParameter({index: index, param: param}));
  }

  deleteCurlParam(param: Iitem) {
    this.store.dispatch(new DelCurlParameter({param: param}));
  }

  addNewCurlParam() {
    this.store.dispatch(new StoreNewCurlParameter(null));
  }

  dropNewCurlParam(index: number) {
    this.store.dispatch(new StoreDropNewCurlParameter({index: index}));
  }

}
