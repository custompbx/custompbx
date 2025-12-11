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
  DelXmlRpcParameter,
  AddXmlRpcParameter,
  StoreNewXmlRpcParameter,
  StoreDropNewXmlRpcParameter,
  SwitchXmlRpcParameter,
  UpdateXmlRpcParameter
} from '../../../store/config/xml_rpc/config.actions.xml_rpc'; // Changed path
import {InnerHeaderComponent} from "../../inner-header/inner-header.component";
import {ModuleNotExistsBannerComponent} from "../module-not-exists-banner/module-not-exists-banner.component";
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {KeyValuePad2Component} from "../../key-value-pad-2/key-value-pad-2.component";

@Component({
  standalone: true,
  imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent, ModuleNotExistsBannerComponent, KeyValuePad2Component],
  selector: 'app-xml-rpc', // Changed selector
  templateUrl: './xml-rpc.component.html', // Kept original template reference
  styleUrls: ['./xml-rpc.component.css']
})
export class XmlRpcComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'XmlRpc';

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
        xml_rpc: {} as IsimpleModule, // Initial state set to xml_rpc
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().xml_rpc); // Accessing xml_rpc state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().xml_rpc?.errorMessage || null); // Accessing xml_rpc error message

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
    // Initialize dispatchers here, updated for XmlRpc
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewXmlRpcParam.bind(this),
      switchItem: this.switchXmlRpcParam.bind(this),
      addItem: this.newXmlRpcParam.bind(this),
      dropNewItem: this.dropNewXmlRpcParam.bind(this),
      deleteItem: this.deleteXmlRpcParam.bind(this),
      updateItem: this.updateXmlRpcParam.bind(this),
      pasteItems: null,
    };
  }

  updateXmlRpcParam(param: Iitem) {
    this.store.dispatch(new UpdateXmlRpcParameter({param: param}));
  }

  switchXmlRpcParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchXmlRpcParameter({param: newParam}));
  }

  newXmlRpcParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddXmlRpcParameter({index: index, param: param}));
  }

  deleteXmlRpcParam(param: Iitem) {
    this.store.dispatch(new DelXmlRpcParameter({param: param}));
  }

  addNewXmlRpcParam() {
    this.store.dispatch(new StoreNewXmlRpcParameter(null));
  }

  dropNewXmlRpcParam(index: number) {
    this.store.dispatch(new StoreDropNewXmlRpcParameter({index: index}));
  }

}
