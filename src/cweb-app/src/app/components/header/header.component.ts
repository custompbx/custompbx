import {
  Component,
  ComponentRef, computed, ElementRef, EventEmitter, HostListener, inject,
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
import {TranslocoPipe} from '@jsverse/transloco';
import {LocaleService} from '../../i18n/locale.service';


@Component({
  standalone: true,
  imports: [PhoneComponent, RouterLink, IconComponent, TranslocoPipe],
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
  @ViewChild('localeMenu', {read: ElementRef}) localeMenu?: ElementRef<HTMLDetailsElement>;
  @ViewChild('userMenu', {read: ElementRef}) userMenu?: ElementRef<HTMLDetailsElement>;
  componentRef: ComponentRef<any>;

  private userService = inject(UserService);
  private resolver = inject(ViewContainerRef);
  private store = inject(Store<AppState>);
  public readonly localeService = inject(LocaleService);
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

  changeLocale(locale: string, menu?: HTMLDetailsElement): void {
    this.localeService.setLocale(locale);
    if (menu) {
      menu.open = false;
    }
  }

  @HostListener('document:click', ['$event'])
  closeMenusOnOutsideClick(event: Event): void {
    const target = event.target;
    if (!(target instanceof Node)) {
      return;
    }

    this.closeMenuWhenOutside(this.localeMenu?.nativeElement, target);
    this.closeMenuWhenOutside(this.userMenu?.nativeElement, target);
  }

  @HostListener('document:keydown.escape')
  closeMenusOnEscape(): void {
    this.closeMenu(this.localeMenu?.nativeElement);
    this.closeMenu(this.userMenu?.nativeElement);
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

  private closeMenuWhenOutside(menu: HTMLDetailsElement | undefined, target: Node): void {
    if (menu?.open && !menu.contains(target)) {
      this.closeMenu(menu);
    }
  }

  private closeMenu(menu: HTMLDetailsElement | undefined): void {
    if (menu) {
      menu.open = false;
    }
  }

}
