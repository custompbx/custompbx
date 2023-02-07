import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {State} from '../../store/fscli/fscli.reducers';
import {select, Store} from '@ngrx/store';
import {AppState, selectFSCLIState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {SendFSCLICommand} from '../../store/fscli/fscli.actions';

@Component({
  selector: 'app-fs-cli',
  templateUrl: './fs-cli.component.html',
  styleUrls: ['./fs-cli.component.css']
})
export class FsCliComponent implements OnInit, OnDestroy {

  public fscli: Observable<any>;
  public fscli$: Subscription;
  public list: State;
  public loadCounter: number;
  public command: string;

  constructor(
    private store: Store<AppState>,
    private _snackBar: MatSnackBar,
  ) {
    this.fscli = this.store.pipe(select(selectFSCLIState));
  }

  ngOnInit() {
    this.fscli$ = this.fscli.subscribe((fscli) => {
      this.loadCounter = fscli.loadCounter;
      this.list = fscli;
      if (!fscli.errorMessage) {

      } else {
        this._snackBar.open('Error: ' + fscli.errorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.fscli$.unsubscribe();
  }

  runCommand() {
    this.store.dispatch(new SendFSCLICommand({name: this.command}));
  }

}
