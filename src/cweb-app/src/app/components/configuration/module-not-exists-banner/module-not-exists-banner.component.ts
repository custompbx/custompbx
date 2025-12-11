import {Component, Input, OnInit} from '@angular/core';

import {MaterialModule} from "../../../../material-module";
import {RouterLink} from "@angular/router";

@Component({
standalone: true,
  imports: [MaterialModule, RouterLink],
    selector: 'app-module-not-exists-banner',
    templateUrl: './module-not-exists-banner.component.html',
    styleUrls: ['./module-not-exists-banner.component.css']
})
export class ModuleNotExistsBannerComponent implements OnInit {

  @Input() list: {exists: boolean};

  constructor() { }

  ngOnInit() {}

}

