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
  DelNibblebillParameter,
  AddNibblebillParameter,
  StoreNewNibblebillParameter,
  StoreDropNewNibblebillParameter,
  SwitchNibblebillParameter,
  UpdateNibblebillParameter
} from '../../../store/config/nibblebill/config.actions.nibblebill'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-nibblebill', // Changed selector
  templateUrl: './nibblebill.component.html', // Kept original template reference
  styleUrls: ['./nibblebill.component.css']
})
export class NibblebillComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Nibblebill';

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
        nibblebill: {} as IsimpleModule, // Initial state set to nibblebill
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().nibblebill); // Accessing nibblebill state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().nibblebill?.errorMessage || null); // Accessing nibblebill error message

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
    // Initialize dispatchers here, updated for Nibblebill
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewNibblebillParam.bind(this),
      switchItem: this.switchNibblebillParam.bind(this),
      addItem: this.newNibblebillParam.bind(this),
      dropNewItem: this.dropNewNibblebillParam.bind(this),
      deleteItem: this.deleteNibblebillParam.bind(this),
      updateItem: this.updateNibblebillParam.bind(this),
      pasteItems: null,
    };
  }

  updateNibblebillParam(param: Iitem) {
    this.store.dispatch(new UpdateNibblebillParameter({param: param}));
  }

  switchNibblebillParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchNibblebillParameter({param: newParam}));
  }

  newNibblebillParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddNibblebillParameter({index: index, param: param}));
  }

  deleteNibblebillParam(param: Iitem) {
    this.store.dispatch(new DelNibblebillParameter({param: param}));
  }

  addNewNibblebillParam() {
    this.store.dispatch(new StoreNewNibblebillParameter(null));
  }

  dropNewNibblebillParam(index: number) {
    this.store.dispatch(new StoreDropNewNibblebillParameter({index: index}));
  }

}
