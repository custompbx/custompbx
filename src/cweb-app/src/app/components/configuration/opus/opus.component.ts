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
  DelOpusParameter,
  AddOpusParameter,
  StoreNewOpusParameter,
  StoreDropNewOpusParameter,
  SwitchOpusParameter,
  UpdateOpusParameter
} from '../../../store/config/opus/config.actions.opus'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-opus', // Changed selector
  templateUrl: './opus.component.html', // Kept original template reference
  styleUrls: ['./opus.component.css']
})
export class OpusComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Opus';

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
        opus: {} as IsimpleModule, // Initial state set to opus
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().opus); // Accessing opus state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().opus?.errorMessage || null); // Accessing opus error message

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
    // Initialize dispatchers here, updated for Opus
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewOpusParam.bind(this),
      switchItem: this.switchOpusParam.bind(this),
      addItem: this.newOpusParam.bind(this),
      dropNewItem: this.dropNewOpusParam.bind(this),
      deleteItem: this.deleteOpusParam.bind(this),
      updateItem: this.updateOpusParam.bind(this),
      pasteItems: null,
    };
  }

  updateOpusParam(param: Iitem) {
    this.store.dispatch(new UpdateOpusParameter({param: param}));
  }

  switchOpusParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchOpusParameter({param: newParam}));
  }

  newOpusParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddOpusParameter({index: index, param: param}));
  }

  deleteOpusParam(param: Iitem) {
    this.store.dispatch(new DelOpusParameter({param: param}));
  }

  addNewOpusParam() {
    this.store.dispatch(new StoreNewOpusParameter(null));
  }

  dropNewOpusParam(index: number) {
    this.store.dispatch(new StoreDropNewOpusParameter({index: index}));
  }

}
