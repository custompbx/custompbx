import {
  Component,
  ComponentFactory,
  ComponentFactoryResolver, ComponentRef,
  Input,
  OnDestroy,
  OnInit,
  ViewChild,
  ViewContainerRef
} from '@angular/core';
import {UserService} from '../../services/user.service';
import {Iuser} from '../../store/auth/auth.reducers';
import {Subscription} from 'rxjs';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit, OnDestroy {

  public user: Iuser;
  public showPhone: boolean;
  private hidePhone = true;
  public getState$: Subscription;
  public clone = true;

  @Input() currentComponent;
  @ViewChild('componentContainer', { read: ViewContainerRef }) container: ViewContainerRef;
  componentRef: ComponentRef<any>;

  constructor(
    private userService: UserService,
    private resolver: ComponentFactoryResolver
  ) {
    this.user = this.userService.user;
  }

  ngOnInit() {
    this.getState$ = this.userService.getState.subscribe((state) => {
      this.user = state.user;
    });
  }

  ngOnDestroy() {
    this.getState$.unsubscribe();
    if (this.componentRef) {
      this.componentRef.destroy();
    }
  }

  logOut(): void {
    this.userService.logOut();
  }

  showHidePhone() {
    if (!this.showPhone) {
      this.showPhone = true;
    }
    this.hidePhone = !this.hidePhone;
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
    let factory: ComponentFactory<any>;
    if (typeof this.currentComponent.getChildComponentFactory === 'function') {
      factory = this.currentComponent.getChildComponentFactory();
    } else {
      factory = this.resolver.resolveComponentFactory(this.currentComponent.constructor);
    }

    this.componentRef = this.container.createComponent(factory);
    // this.componentRef.instance.type = type;
    // this.componentRef.instance.output.subscribe(event => console.log(event));
  }

}
