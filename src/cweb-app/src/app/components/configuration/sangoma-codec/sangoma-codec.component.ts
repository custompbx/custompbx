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
  DelSangomaCodecParameter,
  AddSangomaCodecParameter,
  StoreNewSangomaCodecParameter,
  StoreDropNewSangomaCodecParameter,
  SwitchSangomaCodecParameter,
  UpdateSangomaCodecParameter
} from '../../../store/config/sangoma_codec/config.actions.sangoma_codec'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-sangoma-codec', // Changed selector
  templateUrl: './sangoma-codec.component.html', // Kept original template reference
  styleUrls: ['./sangoma-codec.component.css']
})
export class SangomaCodecComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'SangomaCodec';

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
        sangoma_codec: {} as IsimpleModule, // Initial state set to sangoma_codec
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().sangoma_codec); // Accessing sangoma_codec state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().sangoma_codec?.errorMessage || null); // Accessing sangoma_codec error message

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
    // Initialize dispatchers here, updated for SangomaCodec
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewSangomaCodecParam.bind(this),
      switchItem: this.switchSangomaCodecParam.bind(this),
      addItem: this.newSangomaCodecParam.bind(this),
      dropNewItem: this.dropNewSangomaCodecParam.bind(this),
      deleteItem: this.deleteSangomaCodecParam.bind(this),
      updateItem: this.updateSangomaCodecParam.bind(this),
      pasteItems: null,
    };
  }

  updateSangomaCodecParam(param: Iitem) {
    this.store.dispatch(new UpdateSangomaCodecParameter({param: param}));
  }

  switchSangomaCodecParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchSangomaCodecParameter({param: newParam}));
  }

  newSangomaCodecParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddSangomaCodecParameter({index: index, param: param}));
  }

  deleteSangomaCodecParam(param: Iitem) {
    this.store.dispatch(new DelSangomaCodecParameter({param: param}));
  }

  addNewSangomaCodecParam() {
    this.store.dispatch(new StoreNewSangomaCodecParameter(null));
  }

  dropNewSangomaCodecParam(index: number) {
    this.store.dispatch(new StoreDropNewSangomaCodecParameter({index: index}));
  }

}
