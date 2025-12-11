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
  DelAmrParameter,
  AddAmrParameter,
  StoreNewAmrParameter,
  StoreDropNewAmrParameter,
  SwitchAmrParameter,
  UpdateAmrParameter
} from '../../../store/config/amr/config.actions.amr'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-amr', // Changed selector
  templateUrl: './amr.component.html', // Kept original template reference
  styleUrls: ['./amr.component.css']
})
export class AmrComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Amr';

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
        amr: {} as IsimpleModule, // Initial state set to amr
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().amr); // Accessing amr state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().amr?.errorMessage || null); // Accessing amr error message

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
    // Initialize dispatchers here, updated for Amr
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewAmrParam.bind(this),
      switchItem: this.switchAmrParam.bind(this),
      addItem: this.newAmrParam.bind(this),
      dropNewItem: this.dropNewAmrParam.bind(this),
      deleteItem: this.deleteAmrParam.bind(this),
      updateItem: this.updateAmrParam.bind(this),
      pasteItems: null,
    };
  }

  updateAmrParam(param: Iitem) {
    this.store.dispatch(new UpdateAmrParameter({param: param}));
  }

  switchAmrParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchAmrParameter({param: newParam}));
  }

  newAmrParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddAmrParameter({index: index, param: param}));
  }

  deleteAmrParam(param: Iitem) {
    this.store.dispatch(new DelAmrParameter({param: param}));
  }

  addNewAmrParam() {
    this.store.dispatch(new StoreNewAmrParameter(null));
  }

  dropNewAmrParam(index: number) {
    this.store.dispatch(new StoreDropNewAmrParameter({index: index}));
  }

}
