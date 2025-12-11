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
  DelXmlCdrParameter,
  AddXmlCdrParameter,
  StoreNewXmlCdrParameter,
  StoreDropNewXmlCdrParameter,
  SwitchXmlCdrParameter,
  UpdateXmlCdrParameter
} from '../../../store/config/xml_cdr/config.actions.xml_cdr'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-xml-cdr', // Changed selector
  templateUrl: './xml-cdr.component.html', // Kept original template reference
  styleUrls: ['./xml-cdr.component.css']
})
export class XmlCdrComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'XmlCdr';

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
        xml_cdr: {} as IsimpleModule, // Initial state set to xml_cdr
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().xml_cdr); // Accessing xml_cdr state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().xml_cdr?.errorMessage || null); // Accessing xml_cdr error message

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
    // Initialize dispatchers here, updated for XmlCdr
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewXmlCdrParam.bind(this),
      switchItem: this.switchXmlCdrParam.bind(this),
      addItem: this.newXmlCdrParam.bind(this),
      dropNewItem: this.dropNewXmlCdrParam.bind(this),
      deleteItem: this.deleteXmlCdrParam.bind(this),
      updateItem: this.updateXmlCdrParam.bind(this),
      pasteItems: null,
    };
  }

  updateXmlCdrParam(param: Iitem) {
    this.store.dispatch(new UpdateXmlCdrParameter({param: param}));
  }

  switchXmlCdrParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchXmlCdrParameter({param: newParam}));
  }

  newXmlCdrParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddXmlCdrParameter({index: index, param: param}));
  }

  deleteXmlCdrParam(param: Iitem) {
    this.store.dispatch(new DelXmlCdrParameter({param: param}));
  }

  addNewXmlCdrParam() {
    this.store.dispatch(new StoreNewXmlCdrParameter(null));
  }

  dropNewXmlCdrParam(index: number) {
    this.store.dispatch(new StoreDropNewXmlCdrParameter({index: index}));
  }

}
