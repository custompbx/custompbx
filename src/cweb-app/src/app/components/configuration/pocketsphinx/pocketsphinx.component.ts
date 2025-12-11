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
  DelPocketsphinxParameter,
  AddPocketsphinxParameter,
  StoreNewPocketsphinxParameter,
  StoreDropNewPocketsphinxParameter,
  SwitchPocketsphinxParameter,
  UpdatePocketsphinxParameter
} from '../../../store/config/pocketsphinx/config.actions.pocketsphinx'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-pocketsphinx', // Changed selector
  templateUrl: './pocketsphinx.component.html', // Kept original template reference
  styleUrls: ['./pocketsphinx.component.css']
})
export class PocketsphinxComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Pocketsphinx';

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
        pocketsphinx: {} as IsimpleModule, // Initial state set to pocketsphinx
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().pocketsphinx); // Accessing pocketsphinx state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().pocketsphinx?.errorMessage || null); // Accessing pocketsphinx error message

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
    // Initialize dispatchers here, updated for Pocketsphinx
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewPocketsphinxParam.bind(this),
      switchItem: this.switchPocketsphinxParam.bind(this),
      addItem: this.newPocketsphinxParam.bind(this),
      dropNewItem: this.dropNewPocketsphinxParam.bind(this),
      deleteItem: this.deletePocketsphinxParam.bind(this),
      updateItem: this.updatePocketsphinxParam.bind(this),
      pasteItems: null,
    };
  }

  updatePocketsphinxParam(param: Iitem) {
    this.store.dispatch(new UpdatePocketsphinxParameter({param: param}));
  }

  switchPocketsphinxParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchPocketsphinxParameter({param: newParam}));
  }

  newPocketsphinxParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddPocketsphinxParameter({index: index, param: param}));
  }

  deletePocketsphinxParam(param: Iitem) {
    this.store.dispatch(new DelPocketsphinxParameter({param: param}));
  }

  addNewPocketsphinxParam() {
    this.store.dispatch(new StoreNewPocketsphinxParameter(null));
  }

  dropNewPocketsphinxParam(index: number) {
    this.store.dispatch(new StoreDropNewPocketsphinxParameter({index: index}));
  }

}
