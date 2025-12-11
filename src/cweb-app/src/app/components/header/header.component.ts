import {
  Component,
  ComponentRef, effect, EventEmitter, inject,
  Input,
  OnDestroy,
  OnInit, Output,
  ViewChild,
  ViewContainerRef
} from '@angular/core';
import {UserService} from '../../services/user.service';
import {Iuser} from '../../store/auth/auth.reducers';
import {Observable, Subscription} from 'rxjs';
import {select, Store} from '@ngrx/store';
import {AppState, selectHeader} from '../../store/app.states';
import {StartPhone, ToggleShowPhone} from '../../store/header/header.actions';
import {MaterialModule} from "../../../material-module";
import {PhoneComponent} from "../phone/phone.component";
import {RouterLink} from "@angular/router";


@Component({
  standalone: true,
  imports: [MaterialModule, PhoneComponent, RouterLink],
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit, OnDestroy {

  public user: Iuser;
  public startPhone: boolean;
  public hidePhone = true;
  public clone = true;
  public phoneState: Observable<any>;
  public phoneState$: Subscription;

  @Input() currentComponent;
  @Output() showRightSideNav = new EventEmitter<boolean>();
  @ViewChild('componentContainer', {read: ViewContainerRef}) container: ViewContainerRef;
  componentRef: ComponentRef<any>;

  private userService = inject(UserService);
  private resolver = inject(ViewContainerRef);
  private store = inject(Store<AppState>);

  private menuUpdateEffect = effect(() => {
    this.user = this.userService.userSignal();
  });

  constructor(
  ) {
    this.phoneState = this.store.pipe(select(selectHeader));
  }

  ngOnInit() {
    this.phoneState$ = this.phoneState.subscribe((phone) => {
      this.startPhone = phone.phone.started;
      this.hidePhone = !phone.phone.shown;
    });
  }

  ngOnDestroy() {
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
      this.store.dispatch(StartPhone(null));
    }
    this.store.dispatch(ToggleShowPhone(null));
  }

  showHideConversations() {
    this.showRightSideNav.emit(true);
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
