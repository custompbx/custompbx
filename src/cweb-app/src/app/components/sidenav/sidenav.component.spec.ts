import {Component, signal} from '@angular/core';
import {ComponentFixture, TestBed} from '@angular/core/testing';
import {provideRouter, Router} from '@angular/router';

import {UserService} from '../../services/user.service';
import {SidenavComponent} from './sidenav.component';

@Component({standalone: true, template: ''})
class RouteStubComponent {}

describe('SidenavComponent', () => {
  let component: SidenavComponent;
  let fixture: ComponentFixture<SidenavComponent>;
  let router: Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SidenavComponent],
      providers: [
        provideRouter([
          {path: 'dashboard', component: RouteStubComponent},
          {path: 'monitoring/users-panel', component: RouteStubComponent},
          {path: 'directory/users', component: RouteStubComponent},
          {path: 'configuration/verto', component: RouteStubComponent},
          {path: 'cdr', component: RouteStubComponent},
        ]),
        {provide: UserService, useValue: {userSignal: signal({group_id: 1})}},
      ],
    }).compileComponents();

    router = TestBed.inject(Router);
    fixture = TestBed.createComponent(SidenavComponent);
    component = fixture.componentInstance;
  });

  it('opens only the section containing the active child route', async () => {
    await router.navigateByUrl('/directory/users');
    fixture.detectChanges();

    const panels = Array.from(fixture.nativeElement.querySelectorAll('details')) as HTMLDetailsElement[];
    const openPanels = panels.filter(panel => panel.open);

    expect(openPanels.length).toBe(1);
    expect(openPanels[0].textContent).toContain('Directory');
    expect(fixture.nativeElement.querySelector('[aria-current="page"]')?.textContent).toContain('Users');
  });

  it('updates the expanded section when navigation changes', async () => {
    await router.navigateByUrl('/directory/users');
    fixture.detectChanges();
    await router.navigateByUrl('/configuration/verto');
    fixture.detectChanges();

    const openPanel = fixture.nativeElement.querySelector('details[open]') as HTMLDetailsElement;
    expect(openPanel.textContent).toContain('Configuration');
    expect(openPanel.textContent).not.toContain('Domains');
  });

  it('keeps native summaries keyboard accessible', async () => {
    await router.navigateByUrl('/dashboard');
    fixture.detectChanges();

    const summaries = Array.from(fixture.nativeElement.querySelectorAll('summary')) as HTMLElement[];
    expect(summaries.length).toBeGreaterThan(0);
    expect(summaries.every(summary => Boolean(summary.getAttribute('aria-label')))).toBeTrue();
  });

  it('uses labelled icon links when collapsed', async () => {
    await router.navigateByUrl('/cdr');
    component.collapsed = true;
    fixture.detectChanges();

    const links = Array.from(fixture.nativeElement.querySelectorAll('.menu-rail-item')) as HTMLAnchorElement[];
    expect(links.length).toBe(component.menuItems.length);
    expect(links.every(link => Boolean(link.getAttribute('aria-label')))).toBeTrue();
    expect(fixture.nativeElement.querySelector('.menu-rail-item.menu-item-active')?.getAttribute('aria-label')).toBe('CDR');
  });
});
