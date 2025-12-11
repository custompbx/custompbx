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
  DelFaxParameter,
  AddFaxParameter,
  StoreNewFaxParameter,
  StoreDropNewFaxParameter,
  SwitchFaxParameter,
  UpdateFaxParameter
} from '../../../store/config/fax/config.actions.fax'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-fax', // Changed selector
  templateUrl: './fax.component.html', // Kept original template reference
  styleUrls: ['./fax.component.css']
})
export class FaxComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'Fax';

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
        fax: {} as IsimpleModule, // Initial state set to fax
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().fax); // Accessing fax state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().fax?.errorMessage || null); // Accessing fax error message

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
    // Initialize dispatchers here, updated for Fax
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewFaxParam.bind(this),
      switchItem: this.switchFaxParam.bind(this),
      addItem: this.newFaxParam.bind(this),
      dropNewItem: this.dropNewFaxParam.bind(this),
      deleteItem: this.deleteFaxParam.bind(this),
      updateItem: this.updateFaxParam.bind(this),
      pasteItems: null,
    };
  }

  updateFaxParam(param: Iitem) {
    this.store.dispatch(new UpdateFaxParameter({param: param}));
  }

  switchFaxParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchFaxParameter({param: newParam}));
  }

  newFaxParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddFaxParameter({index: index, param: param}));
  }

  deleteFaxParam(param: Iitem) {
    this.store.dispatch(new DelFaxParameter({param: param}));
  }

  addNewFaxParam() {
    this.store.dispatch(new StoreNewFaxParameter(null));
  }

  dropNewFaxParam(index: number) {
    this.store.dispatch(new StoreDropNewFaxParameter({index: index}));
  }

}
