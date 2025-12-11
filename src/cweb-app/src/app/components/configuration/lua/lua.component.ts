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
  DelLuaParameter,
  AddLuaParameter,
  StoreNewLuaParameter,
  StoreDropNewLuaParameter,
  SwitchLuaParameter,
  UpdateLuaParameter
} from '../../../store/config/lua/config.actions.lua'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-lua', // Changed selector
  templateUrl: './lua.component.html', // Kept original template reference
  styleUrls: ['./lua.component.css']
})
export class LuaComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Lua';

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
        lua: {} as IsimpleModule, // Initial state set to lua
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().lua); // Accessing lua state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().lua?.errorMessage || null); // Accessing lua error message

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
    // Initialize dispatchers here, updated for Lua
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewLuaParam.bind(this),
      switchItem: this.switchLuaParam.bind(this),
      addItem: this.newLuaParam.bind(this),
      dropNewItem: this.dropNewLuaParam.bind(this),
      deleteItem: this.deleteLuaParam.bind(this),
      updateItem: this.updateLuaParam.bind(this),
      pasteItems: null,
    };
  }

  updateLuaParam(param: Iitem) {
    this.store.dispatch(new UpdateLuaParameter({param: param}));
  }

  switchLuaParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchLuaParameter({param: newParam}));
  }

  newLuaParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddLuaParameter({index: index, param: param}));
  }

  deleteLuaParam(param: Iitem) {
    this.store.dispatch(new DelLuaParameter({param: param}));
  }

  addNewLuaParam() {
    this.store.dispatch(new StoreNewLuaParameter(null));
  }

  dropNewLuaParam(index: number) {
    this.store.dispatch(new StoreDropNewLuaParameter({index: index}));
  }

}
