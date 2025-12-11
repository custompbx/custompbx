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
  DelEasyrouteParameter,
  AddEasyrouteParameter,
  StoreNewEasyrouteParameter,
  StoreDropNewEasyrouteParameter,
  SwitchEasyrouteParameter,
  UpdateEasyrouteParameter
} from '../../../store/config/easyroute/config.actions.easyroute'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-easyroute', // Changed selector
  templateUrl: './easyroute.component.html', // Kept original template reference
  styleUrls: ['./easyroute.component.css']
})
export class EasyrouteComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Easyroute';

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
        easyroute: {} as IsimpleModule, // Initial state set to easyroute
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().easyroute); // Accessing easyroute state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().easyroute?.errorMessage || null); // Accessing easyroute error message

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
    // Initialize dispatchers here, updated for Easyroute
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewEasyrouteParam.bind(this),
      switchItem: this.switchEasyrouteParam.bind(this),
      addItem: this.newEasyrouteParam.bind(this),
      dropNewItem: this.dropNewEasyrouteParam.bind(this),
      deleteItem: this.deleteEasyrouteParam.bind(this),
      updateItem: this.updateEasyrouteParam.bind(this),
      pasteItems: null,
    };
  }

  updateEasyrouteParam(param: Iitem) {
    this.store.dispatch(new UpdateEasyrouteParameter({param: param}));
  }

  switchEasyrouteParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchEasyrouteParameter({param: newParam}));
  }

  newEasyrouteParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddEasyrouteParameter({index: index, param: param}));
  }

  deleteEasyrouteParam(param: Iitem) {
    this.store.dispatch(new DelEasyrouteParameter({param: param}));
  }

  addNewEasyrouteParam() {
    this.store.dispatch(new StoreNewEasyrouteParameter(null));
  }

  dropNewEasyrouteParam(index: number) {
    this.store.dispatch(new StoreDropNewEasyrouteParameter({index: index}));
  }

}
