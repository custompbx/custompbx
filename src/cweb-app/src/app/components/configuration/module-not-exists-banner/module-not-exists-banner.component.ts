import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-module-not-exists-banner',
  templateUrl: './module-not-exists-banner.component.html',
  styleUrls: ['./module-not-exists-banner.component.css']
})
export class ModuleNotExistsBannerComponent implements OnInit {

  @Input() list: {exists: boolean};

  constructor() { }

  ngOnInit() {}

}

