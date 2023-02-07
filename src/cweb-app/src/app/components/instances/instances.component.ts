import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectInstancesState} from '../../store/app.states';
import {MatSnackBar} from '@angular/material/snack-bar';
import {WsDataService} from '../../services/ws-data.service';
import {GetInstances, UpdateInstanceDescription} from '../../store/instances/instances.actions';
import {AbstractControl} from '@angular/forms';


@Component({
  selector: 'app-instances',
  templateUrl: './instances.component.html',
  styleUrls: ['./instances.component.css']
})
export class InstancesComponent implements OnInit, OnDestroy {

  public instances: Observable<any>;
  public instances$: Subscription;
  public list: any;
  private lastErrorMessage: string;
  public loadCounter: number;
  public currentInstanceId: number;

  constructor(
    private store: Store<AppState>,
    private _snackBar: MatSnackBar,
    private ws: WsDataService,
  ) {
    this.instances = this.store.pipe(select(selectInstancesState));
  }

  ngOnInit() {
    this.instances$ = this.instances.subscribe((instances) => {
      this.loadCounter = instances.loadCounter;
      this.lastErrorMessage = instances.errorMessage;
      this.list = instances.instances;
      this.currentInstanceId = instances.currentInstanceId;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  ngOnDestroy() {
    this.instances$.unsubscribe();
  }

  trackByFnId(index, item) {
    return item.value.id;
  }

  switchInstance(id) {
    if (!this.list[id]) {
      this._snackBar.open('Error:  Wrong instance id!', null, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
      return;
    }
    this.ws.connectToAnotherHost('wss://' + this.list[id].host + ':' + String(this.list[id].port) + '/ws').subscribe(
      connected => {
        if (connected) {
          this.store.dispatch(new GetInstances(null));
        }
      }
    );
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  updateDescription(instance, description) {
    const data = {id: instance.id, value: description};
    console.log(data);
    this.store.dispatch(new UpdateInstanceDescription(data));
  }
}
