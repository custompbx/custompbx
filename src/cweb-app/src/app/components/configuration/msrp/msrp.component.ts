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
  DelMsrpParameter,
  AddMsrpParameter,
  StoreNewMsrpParameter,
  StoreDropNewMsrpParameter,
  SwitchMsrpParameter,
  UpdateMsrpParameter
} from '../../../store/config/msrp/config.actions.msrp'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-msrp', // Changed selector
  templateUrl: './msrp.component.html', // Kept original template reference
  styleUrls: ['./msrp.component.css']
})
export class MsrpComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Msrp';

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
        msrp: {} as IsimpleModule, // Initial state set to msrp
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().msrp); // Accessing msrp state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().msrp?.errorMessage || null); // Accessing msrp error message

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
    // Initialize dispatchers here, updated for Msrp
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewMsrpParam.bind(this),
      switchItem: this.switchMsrpParam.bind(this),
      addItem: this.newMsrpParam.bind(this),
      dropNewItem: this.dropNewMsrpParam.bind(this),
      deleteItem: this.deleteMsrpParam.bind(this),
      updateItem: this.updateMsrpParam.bind(this),
      pasteItems: null,
    };
  }

  updateMsrpParam(param: Iitem) {
    this.store.dispatch(new UpdateMsrpParameter({param: param}));
  }

  switchMsrpParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchMsrpParameter({param: newParam}));
  }

  newMsrpParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddMsrpParameter({index: index, param: param}));
  }

  deleteMsrpParam(param: Iitem) {
    this.store.dispatch(new DelMsrpParameter({param: param}));
  }

  addNewMsrpParam() {
    this.store.dispatch(new StoreNewMsrpParameter(null));
  }

  dropNewMsrpParam(index: number) {
    this.store.dispatch(new StoreDropNewMsrpParameter({index: index}));
  }

}
