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
  DelPerlParameter,
  AddPerlParameter,
  StoreNewPerlParameter,
  StoreDropNewPerlParameter,
  SwitchPerlParameter,
  UpdatePerlParameter
} from '../../../store/config/perl/config.actions.perl'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-perl', // Changed selector
  templateUrl: './perl.component.html', // Kept original template reference
  styleUrls: ['./perl.component.css']
})
export class PerlComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Perl';

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
        perl: {} as IsimpleModule, // Initial state set to perl
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().perl); // Accessing perl state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().perl?.errorMessage || null); // Accessing perl error message

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
    // Initialize dispatchers here, updated for Perl
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewPerlParam.bind(this),
      switchItem: this.switchPerlParam.bind(this),
      addItem: this.newPerlParam.bind(this),
      dropNewItem: this.dropNewPerlParam.bind(this),
      deleteItem: this.deletePerlParam.bind(this),
      updateItem: this.updatePerlParam.bind(this),
      pasteItems: null,
    };
  }

  updatePerlParam(param: Iitem) {
    this.store.dispatch(new UpdatePerlParameter({param: param}));
  }

  switchPerlParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchPerlParameter({param: newParam}));
  }

  newPerlParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddPerlParameter({index: index, param: param}));
  }

  deletePerlParam(param: Iitem) {
    this.store.dispatch(new DelPerlParameter({param: param}));
  }

  addNewPerlParam() {
    this.store.dispatch(new StoreNewPerlParameter(null));
  }

  dropNewPerlParam(index: number) {
    this.store.dispatch(new StoreDropNewPerlParameter({index: index}));
  }

}
