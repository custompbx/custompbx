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
  DelSndfileParameter,
  AddSndfileParameter,
  StoreNewSndfileParameter,
  StoreDropNewSndfileParameter,
  SwitchSndfileParameter,
  UpdateSndfileParameter
} from '../../../store/config/sndfile/config.actions.sndfile'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-sndfile', // Changed selector
  templateUrl: './sndfile.component.html', // Kept original template reference
  styleUrls: ['./sndfile.component.css']
})
export class SndfileComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Sndfile';

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
        sndfile: {} as IsimpleModule, // Initial state set to sndfile
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().sndfile); // Accessing sndfile state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().sndfile?.errorMessage || null); // Accessing sndfile error message

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
    // Initialize dispatchers here, updated for Sndfile
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewSndfileParam.bind(this),
      switchItem: this.switchSndfileParam.bind(this),
      addItem: this.newSndfileParam.bind(this),
      dropNewItem: this.dropNewSndfileParam.bind(this),
      deleteItem: this.deleteSndfileParam.bind(this),
      updateItem: this.updateSndfileParam.bind(this),
      pasteItems: null,
    };
  }

  updateSndfileParam(param: Iitem) {
    this.store.dispatch(new UpdateSndfileParameter({param: param}));
  }

  switchSndfileParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchSndfileParameter({param: newParam}));
  }

  newSndfileParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddSndfileParameter({index: index, param: param}));
  }

  deleteSndfileParam(param: Iitem) {
    this.store.dispatch(new DelSndfileParameter({param: param}));
  }

  addNewSndfileParam() {
    this.store.dispatch(new StoreNewSndfileParameter(null));
  }

  dropNewSndfileParam(index: number) {
    this.store.dispatch(new StoreDropNewSndfileParameter({index: index}));
  }

}
