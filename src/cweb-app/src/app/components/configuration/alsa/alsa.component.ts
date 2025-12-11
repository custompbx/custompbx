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
  DelAlsaParameter,
  AddAlsaParameter,
  StoreNewAlsaParameter,
  StoreDropNewAlsaParameter,
  SwitchAlsaParameter,
  UpdateAlsaParameter
} from '../../../store/config/alsa/config.actions.alsa';
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-alsa',
  templateUrl: './alsa.component.html',
  styleUrls: ['./alsa.component.css']
})
export class AlsaComponent implements OnInit { // Removed OnDestroy

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
        alsa: {} as IsimpleModule,
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().alsa);
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().alsa?.errorMessage || null);

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
    // Initialize dispatchers here, as the component logic uses methods defined on `this`.
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewAlsaParam.bind(this),
      switchItem: this.switchAlsaParam.bind(this),
      addItem: this.newAlsaParam.bind(this),
      dropNewItem: this.dropNewAlsaParam.bind(this),
      deleteItem: this.deleteAlsaParam.bind(this),
      updateItem: this.updateAlsaParam.bind(this),
      pasteItems: null,
    };
  }

  updateAlsaParam(param: Iitem) {
    this.store.dispatch(new UpdateAlsaParameter({param: param}));
  }

  switchAlsaParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchAlsaParameter({param: newParam}));
  }

  newAlsaParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddAlsaParameter({index: index, param: param}));
  }

  deleteAlsaParam(param: Iitem) {
    this.store.dispatch(new DelAlsaParameter({param: param}));
  }

  addNewAlsaParam() {
    this.store.dispatch(new StoreNewAlsaParameter(null));
  }

  dropNewAlsaParam(index: number) {
    this.store.dispatch(new StoreDropNewAlsaParameter({index: index}));
  }

}
