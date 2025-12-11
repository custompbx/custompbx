import {Component, OnDestroy, OnInit, inject, signal, computed, effect} from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';

import {MaterialModule} from "../../../material-module";
import {State} from '../../store/fscli/fscli.reducers';
import {select, Store} from '@ngrx/store';
import {AppState, selectFSCLIState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {SendFSCLICommand} from '../../store/fscli/fscli.actions';
import {FormsModule} from "@angular/forms";
import {ResizeInputDirective} from "../../directives/resize-input.directive";

@Component({
  standalone: true,
  imports: [MaterialModule, FormsModule, ResizeInputDirective],
  selector: 'app-fs-cli',
  templateUrl: './fs-cli.component.html',
  styleUrls: ['./fs-cli.component.css']
})
export class FsCliComponent implements OnInit, OnDestroy {

  // --- Dependency Injection using inject() ---
  private store = inject(Store<AppState>);
  private _snackBar = inject(MatSnackBar);

  // --- Reactive State from NgRx using toSignal ---
  private fscliState = toSignal(
    this.store.pipe(select(selectFSCLIState)),
    {
      initialValue: {
        fsCliData: '',
        errorMessage: '',
        loadCounter: 0,
      } as State
    }
  );

  // --- Computed State ---
  public list = computed(() => this.fscliState());
  public loadCounter = computed(() => this.fscliState().loadCounter);

  // --- Local State as Signal (for two-way binding) ---
  public command = signal<string>('');

  private snackbarEffect = effect(() => {
    const fscli = this.fscliState();
    const errorMessage = fscli.errorMessage;

    if (errorMessage) {
      this._snackBar.open('Error: ' + errorMessage + '!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
    }
  });

  ngOnInit() {
    // The snackbarEffect runs automatically upon state changes.
  }

  ngOnDestroy() {
    // The subscription cleanup is handled automatically by toSignal.
  }

  runCommand() {
    // Access signal value using command()
    if (this.command().trim()) {
      this.store.dispatch(new SendFSCLICommand({name: this.command().trim()}));
    }
  }
}
