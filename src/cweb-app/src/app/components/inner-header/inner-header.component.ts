import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-inner-header',
  templateUrl: './inner-header.component.html',
  styleUrls: ['./inner-header.component.css']
})
export class InnerHeaderComponent implements OnInit {

  @Input() name: string;
  @Input() loadCounter: number;

  constructor() { }

  ngOnInit() {}

}
