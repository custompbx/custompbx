import {Component, Input} from '@angular/core';

import {RouterLink} from "@angular/router";

@Component({
standalone: true,
  imports: [RouterLink],
    selector: 'app-module-not-exists-banner',
    templateUrl: './module-not-exists-banner.component.html'
})
export class ModuleNotExistsBannerComponent {

  @Input() list: {exists: boolean};

}

