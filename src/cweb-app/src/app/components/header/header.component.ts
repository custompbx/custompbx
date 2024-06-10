import {
  Component,
  ComponentRef, EventEmitter,
  Input,
  OnDestroy,
  OnInit, Output,
  ViewChild,
  ViewContainerRef
} from '@angular/core';
import {UserService} from '../../services/user.service';
import {Iuser} from '../../store/auth/auth.reducers';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from "@ngrx/store";
import {AppState, selectHeader, selectPhoneState} from "../../store/app.states";
import {StartPhone, ToggleShowPhone} from "../../store/header/header.actions";
import {ToggleShowConversations} from "../../store/app/app.actions";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit, OnDestroy {

  public user: Iuser;
  public startPhone: boolean;
  public hidePhone = true;
  public getState$: Subscription;
  public clone = true;
  public phoneState: Observable<any>;
  public phoneState$: Subscription;

  @Input() currentComponent;
  @ViewChild('componentContainer', {read: ViewContainerRef}) container: ViewContainerRef;
  componentRef: ComponentRef<any>;

  constructor(
    private userService: UserService,
    private resolver: ViewContainerRef,
    private store: Store<AppState>,
  ) {
    this.user = this.userService.user;
    this.phoneState = this.store.pipe(select(selectHeader));
  }

  ngOnInit() {
    this.getState$ = this.userService.getState.subscribe((state) => {
      this.user = state.user;
    });
    this.phoneState$ = this.phoneState.subscribe((phone) => {
      this.startPhone = phone.phone.started;
      this.hidePhone = !phone.phone.shown;
    });
  }

  ngOnDestroy() {
    this.getState$.unsubscribe();
    this.phoneState$.unsubscribe();
    if (this.componentRef) {
      this.componentRef.destroy();
    }
  }

  logOut(): void {
    this.userService.logOut();
  }

  showHidePhone() {
    if (!this.startPhone) {
      this.store.dispatch(StartPhone(null))
    }
    this.store.dispatch(ToggleShowPhone(null))
  }

  showHideConversations() {
    this.store.dispatch(ToggleShowConversations(null))
  }

  cloneComponent() {
    if (!this.currentComponent) {
      return;
    }
    this.clone = !this.clone;
    if (this.clone) {
      this.componentRef.destroy();
      this.container.clear();
      return;
    }
    this.container.clear();
    if (typeof this.currentComponent.getChildComponentFactory === 'function') {
      const componentClass = this.currentComponent.getChildComponentFactory().componentType;
      this.componentRef = this.resolver.createComponent(componentClass);
    } else {
      this.componentRef = this.resolver.createComponent(this.currentComponent.constructor);
    }
    const hostElement = this.componentRef.location.nativeElement;
    hostElement.classList.add('over-component');
  }

}
