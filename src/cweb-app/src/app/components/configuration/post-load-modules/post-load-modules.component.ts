import {Component, OnDestroy, OnInit} from '@angular/core';
import {Observable, Subscription} from 'rxjs';
import {Iitem, IpostLoadModules} from '../../../store/config/config.state.struct';
import {select, Store} from '@ngrx/store';
import {AppState, selectConfigurationState} from '../../../store/app.states';
import {AbstractControl} from '@angular/forms';
import {MatBottomSheet} from '@angular/material/bottom-sheet';
import {MatSnackBar} from '@angular/material/snack-bar';
import {ActivatedRoute} from '@angular/router';
import {
  DelPostLoadModule,
  AddPostLoadModule,
  StoreNewPostLoadModule,
  StoreDropNewPostLoadModule,
  SwitchPostLoadModule,
  UpdatePostLoadModule
} from '../../../store/config/post_load_modules/config.actions.PostLoadModules';

@Component({
  selector: 'app-post-load-modules',
  templateUrl: './post-load-modules.component.html',
  styleUrls: ['./post-load-modules.component.css']
})
export class PostLoadModulesComponent implements OnInit, OnDestroy {

  public configs: Observable<any>;
  public configs$: Subscription;
  public list: IpostLoadModules;
  public selectedIndex: number;
  private lastErrorMessage: string;
  public loadCounter: number;
  public globalSettingsDispatchers: object;

  constructor(
    private store: Store<AppState>,
    private bottomSheet: MatBottomSheet,
    private _snackBar: MatSnackBar,
    private route: ActivatedRoute,
  ) {
    this.selectedIndex = 0;
    this.configs = this.store.pipe(select(selectConfigurationState));
  }

  ngOnInit() {
    this.configs$ = this.configs.subscribe((configs) => {
      this.loadCounter = configs.loadCounter;
      this.list = configs.post_load_modules;
      this.lastErrorMessage = configs.post_load_modules && configs.post_load_modules.errorMessage || null;
      if (!this.lastErrorMessage) {
      } else {
        this._snackBar.open('Error: ' + this.lastErrorMessage + '!', null, {
          duration: 3000,
          panelClass: ['error-snack'],
        });
      }
    });
    this.globalSettingsDispatchers = {
      addNewItemField: this.addNewPostLoadModule.bind(this),
      switchItem: this.switchPostLoadModule.bind(this),
      addItem: this.newPostLoadModule.bind(this),
      dropNewItem: this.dropNewPostLoadModule.bind(this),
      deleteItem: this.deletePostLoadModule.bind(this),
      updateItem: this.updatePostLoadModule.bind(this),
      pasteItems: null,
    };
  }

  ngOnDestroy() {
    this.configs$.unsubscribe();
    if (this.route.snapshot?.data?.reconnectUpdater) {
       this.route.snapshot.data.reconnectUpdater.unsubscribe();
     }
  }

  updatePostLoadModule(param: Iitem) {
    this.store.dispatch(new UpdatePostLoadModule({param: param}));
  }

  switchPostLoadModule(param: Iitem) {
    const newParam = <Iitem>{...param};
    newParam.enabled = !newParam.enabled;
    this.store.dispatch(new SwitchPostLoadModule({param: newParam}));
  }

  newPostLoadModule(index: number, name: string, value: string) {
    const param = <Iitem>{};
    param.enabled = true;
    param.name = name;
    param.value = value;

    this.store.dispatch(new AddPostLoadModule({index: index, param: param}));
  }

  deletePostLoadModule(param: Iitem) {
    this.store.dispatch(new DelPostLoadModule({param: param}));
  }

  addNewPostLoadModule() {
    this.store.dispatch(new StoreNewPostLoadModule(null));
  }

  dropNewPostLoadModule(index: number) {
    this.store.dispatch(new StoreDropNewPostLoadModule({index: index}));
  }

  checkDirty(condition: AbstractControl): boolean {
    if (condition) {
      return !condition.dirty;
    } else {
      return true;
    }
  }

  isReadyToSendThree(mainObject: AbstractControl, object2: AbstractControl, object3: AbstractControl): boolean {
    return (mainObject && mainObject.valid && mainObject.dirty)
      || ((object2 && object2.valid && object2.dirty) || (object3 && object3.valid && object3.dirty));
  }

  isvalueReadyToSend(valueObject: AbstractControl): boolean {
    return valueObject && valueObject.dirty && valueObject.valid;
  }

  isReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && (nameObject.dirty || valueObject.dirty) && nameObject.valid && valueObject.valid;
  }

  isArray(obj: any): boolean {
    return Array.isArray(obj);
  }

  trackByFn(index, item) {
    return index; // or item.id
  }

  isNewReadyToSend(nameObject: AbstractControl, valueObject: AbstractControl): boolean {
    return nameObject && valueObject && nameObject.valid && valueObject.valid;
  }

}

