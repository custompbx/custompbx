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
  DelErlangEventParameter,
  AddErlangEventParameter,
  StoreNewErlangEventParameter,
  StoreDropNewErlangEventParameter,
  SwitchErlangEventParameter,
  UpdateErlangEventParameter
} from '../../../store/config/erlang_event/config.actions.erlang_event'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-erlang-event', // Changed selector
  templateUrl: './erlang-event.component.html', // Kept original template reference
  styleUrls: ['./erlang-event.component.css']
})
export class ErlangEventComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'ErlangEvent';

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
        erlang_event: {} as IsimpleModule, // Initial state set to erlang_event
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().erlang_event); // Accessing erlang_event state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().erlang_event?.errorMessage || null); // Accessing erlang_event error message

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
    // Initialize dispatchers here, updated for ErlangEvent
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewErlangEventParam.bind(this),
      switchItem: this.switchErlangEventParam.bind(this),
      addItem: this.newErlangEventParam.bind(this),
      dropNewItem: this.dropNewErlangEventParam.bind(this),
      deleteItem: this.deleteErlangEventParam.bind(this),
      updateItem: this.updateErlangEventParam.bind(this),
      pasteItems: null,
    };
  }

  updateErlangEventParam(param: Iitem) {
    this.store.dispatch(new UpdateErlangEventParameter({param: param}));
  }

  switchErlangEventParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchErlangEventParameter({param: newParam}));
  }

  newErlangEventParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddErlangEventParameter({index: index, param: param}));
  }

  deleteErlangEventParam(param: Iitem) {
    this.store.dispatch(new DelErlangEventParameter({param: param}));
  }

  addNewErlangEventParam() {
    this.store.dispatch(new StoreNewErlangEventParameter(null));
  }

  dropNewErlangEventParam(index: number) {
    this.store.dispatch(new StoreDropNewErlangEventParameter({index: index}));
  }

}
