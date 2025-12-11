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
  DelTtsCommandlineParameter,
  AddTtsCommandlineParameter,
  StoreNewTtsCommandlineParameter,
  StoreDropNewTtsCommandlineParameter,
  SwitchTtsCommandlineParameter,
  UpdateTtsCommandlineParameter
} from '../../../store/config/tts_commandline/config.actions.tts_commandline'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-tts-commandline', // Changed selector
  templateUrl: './tts-commandline.component.html', // Kept original template reference
  styleUrls: ['./tts-commandline.component.css']
})
export class TtsCommandlineComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'TtsCommandline';

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
        tts_commandline: {} as IsimpleModule, // Initial state set to tts_commandline
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().tts_commandline); // Accessing tts_commandline state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().tts_commandline?.errorMessage || null); // Accessing tts_commandline error message

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
    // Initialize dispatchers here, updated for TtsCommandline
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewTtsCommandlineParam.bind(this),
      switchItem: this.switchTtsCommandlineParam.bind(this),
      addItem: this.newTtsCommandlineParam.bind(this),
      dropNewItem: this.dropNewTtsCommandlineParam.bind(this),
      deleteItem: this.deleteTtsCommandlineParam.bind(this),
      updateItem: this.updateTtsCommandlineParam.bind(this),
      pasteItems: null,
    };
  }

  updateTtsCommandlineParam(param: Iitem) {
    this.store.dispatch(new UpdateTtsCommandlineParameter({param: param}));
  }

  switchTtsCommandlineParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchTtsCommandlineParameter({param: newParam}));
  }

  newTtsCommandlineParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddTtsCommandlineParameter({index: index, param: param}));
  }

  deleteTtsCommandlineParam(param: Iitem) {
    this.store.dispatch(new DelTtsCommandlineParameter({param: param}));
  }

  addNewTtsCommandlineParam() {
    this.store.dispatch(new StoreNewTtsCommandlineParameter(null));
  }

  dropNewTtsCommandlineParam(index: number) {
    this.store.dispatch(new StoreDropNewTtsCommandlineParameter({index: index}));
  }

}
