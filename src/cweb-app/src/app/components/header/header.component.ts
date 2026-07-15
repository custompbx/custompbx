import {
  Component,
  ComponentRef, computed, EventEmitter, inject,
  Input,
  OnDestroy, Output,
  ViewChild,
  ViewContainerRef
} from '@angular/core';
import {UserService} from '../../services/user.service';
import {Iuser} from '../../store/auth/auth.reducers';
import {Store} from '@ngrx/store';
import {AppState, selectHeader} from '../../store/app.states';
import {StartPhone, ToggleShowPhone} from '../../store/header/header.actions';
import {initialState as initialHeaderState} from '../../store/header/header.reducer';
import {PhoneComponent} from "../phone/phone.component";
import {RouterLink} from "@angular/router";
import {toSignal} from '@angular/core/rxjs-interop';
import {IconComponent} from '../icon/icon.component';


@Component({
  standalone: true,
  imports: [PhoneComponent, RouterLink, IconComponent],
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnDestroy {

  public readonly user = computed<Iuser | null>(() => this.userService.userSignal());
  public clone = true;

  @Input() currentComponent;
  @Output() showRightSideNav = new EventEmitter<boolean>();
  @Output() toggleMenu = new EventEmitter<void>();
  @ViewChild('componentContainer', {read: ViewContainerRef}) container: ViewContainerRef;
  componentRef: ComponentRef<any>;

  private userService = inject(UserService);
  private resolver = inject(ViewContainerRef);
  private store = inject(Store<AppState>);
  private readonly phoneState = toSignal(this.store.select(selectHeader), {initialValue: initialHeaderState});
  public readonly startPhone = computed(() => this.phoneState().phone.started);
  public readonly hidePhone = computed(() => !this.phoneState().phone.shown);

  ngOnDestroy() {
    this.componentRef?.destroy();
  }

  logOut(): void {
    this.userService.logOut();
  }

  showHidePhone() {
    if (!this.startPhone()) {
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
      this.componentRef?.destroy();
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
