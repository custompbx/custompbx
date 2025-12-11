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
  DelMemcacheParameter,
  AddMemcacheParameter,
  StoreNewMemcacheParameter,
  StoreDropNewMemcacheParameter,
  SwitchMemcacheParameter,
  UpdateMemcacheParameter
} from '../../../store/config/memcache/config.actions.memcache'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-memcache', // Changed selector
  templateUrl: './memcache.component.html', // Kept original template reference
  styleUrls: ['./memcache.component.css']
})
export class MemcacheComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Memcache';

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
        memcache: {} as IsimpleModule, // Initial state set to memcache
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().memcache); // Accessing memcache state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().memcache?.errorMessage || null); // Accessing memcache error message

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
    // Initialize dispatchers here, updated for Memcache
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewMemcacheParam.bind(this),
      switchItem: this.switchMemcacheParam.bind(this),
      addItem: this.newMemcacheParam.bind(this),
      dropNewItem: this.dropNewMemcacheParam.bind(this),
      deleteItem: this.deleteMemcacheParam.bind(this),
      updateItem: this.updateMemcacheParam.bind(this),
      pasteItems: null,
    };
  }

  updateMemcacheParam(param: Iitem) {
    this.store.dispatch(new UpdateMemcacheParameter({param: param}));
  }

  switchMemcacheParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchMemcacheParameter({param: newParam}));
  }

  newMemcacheParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddMemcacheParameter({index: index, param: param}));
  }

  deleteMemcacheParam(param: Iitem) {
    this.store.dispatch(new DelMemcacheParameter({param: param}));
  }

  addNewMemcacheParam() {
    this.store.dispatch(new StoreNewMemcacheParameter(null));
  }

  dropNewMemcacheParam(index: number) {
    this.store.dispatch(new StoreDropNewMemcacheParameter({index: index}));
  }

}
