import {Component, computed, effect} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MaterialModule} from "../../../material-module";
import {Store} from '@ngrx/store';
import {AppState, selectInstancesState} from '../../store/app.states';
import {initialState, Iinstances} from '../../store/instances/instances.reducers';
import {MatSnackBar} from '@angular/material/snack-bar';
import {UpdateInstanceDescription} from '../../store/instances/instances.actions';
import {AbstractControl, FormsModule} from '@angular/forms';
import {InnerHeaderComponent} from "../inner-header/inner-header.component";
import {toSignal} from '@angular/core/rxjs-interop';


@Component({
standalone: true,
    imports: [CommonModule, MaterialModule, FormsModule, InnerHeaderComponent],
    selector: 'app-instances',
    templateUrl: './instances.component.html',
    styleUrls: ['./instances.component.css']
})
export class InstancesComponent {

  private readonly state = toSignal(this.store.select(selectInstancesState), {initialValue: initialState});
  public readonly list = computed(() => this.state().instances as Iinstances);
  public readonly loadCounter = computed(() => this.state().loadCounter);
  public readonly currentInstanceId = computed(() => this.state().currentInstanceId);
  private lastShownError: string | null = null;

  constructor(
    private store: Store<AppState>,
    private _snackBar: MatSnackBar,
  ) {
    effect(() => {
      const error = this.state().errorMessage;
      if (error && error !== this.lastShownError) {
        this.lastShownError = error;
        this._snackBar.open('Error: ' + error + '!', undefined, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
  }

  trackByFnId(_index: number, item: {value: {id: number}}): number {
    return item.value.id;
  }

  switchInstance(id: number): void {
    const instance = this.list()[id];
    if (!instance) {
      this._snackBar.open('Error: Wrong instance id!', undefined, {
        duration: 3000,
        panelClass: ['error-snack'],
      });
      return;
    }
    window.open(`https://${instance.host}:${instance.port}/cweb`, '_blank', 'noopener,noreferrer');
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  updateDescription(instance: {id: number}, description: string): void {
    const data = {id: instance.id, value: description};
    this.store.dispatch(new UpdateInstanceDescription(data));
  }
}
