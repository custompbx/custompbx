import {Component, Input} from '@angular/core';

import {MaterialModule} from "../../../../material-module";
import {RouterLink} from "@angular/router";

@Component({
standalone: true,
  imports: [MaterialModule, RouterLink],
    selector: 'app-module-not-exists-banner',
    templateUrl: './module-not-exists-banner.component.html'
})
export class ModuleNotExistsBannerComponent {

  @Input() list: {exists: boolean};

}

