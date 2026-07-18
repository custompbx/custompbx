import {Component, inject, computed, OnInit, effect} from '@angular/core';
import {toSignal} from '@angular/core/rxjs-interop';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {ToastService} from '../../../services/toast.service';
import {ActivatedRoute} from '@angular/router';
import {
  DelDialplanDirectoryParameter,
  AddDialplanDirectoryParameter,
  StoreNewDialplanDirectoryParameter,
  StoreDropNewDialplanDirectoryParameter,
  SwitchDialplanDirectoryParameter,
  UpdateDialplanDirectoryParameter
} from '../../../store/config/dialplan_directory/config.actions.dialplan_directory'; // Changed path
import {Iitem, IsimpleModule, State} from '../../../store/config/config.state.struct';
import {SimpleConfigPageComponent} from "../simple-config-page/simple-config-page.component";

@Component({
  standalone: true,
  imports: [SimpleConfigPageComponent],
  selector: 'app-dialplan-directory', // Changed selector
  templateUrl: './dialplan-directory.component.html', // Kept original template reference
  styleUrls: ['../generated-config-page.css']
})
export class DialplanDirectoryComponent implements OnInit { // Removed OnDestroy
  public moduleName: string = 'DialplanDirectory';

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private _snackBar = inject(ToastService);
  private route = inject(ActivatedRoute);

  // --- Reactive State from NgRx using toSignal ---
  private configState = toSignal(
    this.store.pipe(select(selectConfigurationState)),
    {
      initialValue: {
        dialplan_directory: {} as IsimpleModule, // Initial state set to dialplan_directory
        errorMessage: null,
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed/Derived State from NgRx State ---
  public list = computed(() => this.configState().dialplan_directory); // Accessing dialplan_directory state
  public loadCounter = computed(() => this.configState().loadCounter);
  private lastErrorMessage = computed(() => this.configState().dialplan_directory?.errorMessage || null); // Accessing dialplan_directory error message

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
    // Initialize dispatchers here, updated for DialplanDirectory
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewDialplanDirectoryParam.bind(this),
      switchItem: this.switchDialplanDirectoryParam.bind(this),
      addItem: this.newDialplanDirectoryParam.bind(this),
      dropNewItem: this.dropNewDialplanDirectoryParam.bind(this),
      deleteItem: this.deleteDialplanDirectoryParam.bind(this),
      updateItem: this.updateDialplanDirectoryParam.bind(this),
      pasteItems: null,
    };
  }

  updateDialplanDirectoryParam(param: Iitem) {
    this.store.dispatch(new UpdateDialplanDirectoryParameter({param: param}));
  }

  switchDialplanDirectoryParam(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchDialplanDirectoryParameter({param: newParam}));
  }

  newDialplanDirectoryParam(index: number, name: string, value: string) {
    const param = <Iitem>{
      enabled: true,
      name: name,
      value: value
    };

    this.store.dispatch(new AddDialplanDirectoryParameter({index: index, param: param}));
  }

  deleteDialplanDirectoryParam(param: Iitem) {
    this.store.dispatch(new DelDialplanDirectoryParameter({param: param}));
  }

  addNewDialplanDirectoryParam() {
    this.store.dispatch(new StoreNewDialplanDirectoryParameter(null));
  }

  dropNewDialplanDirectoryParam(index: number) {
    this.store.dispatch(new StoreDropNewDialplanDirectoryParameter({index: index}));
  }

}
