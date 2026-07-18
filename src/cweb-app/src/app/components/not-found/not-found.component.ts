import {Component} from '@angular/core';
import {TranslocoPipe} from '@jsverse/transloco';

@Component({
standalone: true,
  imports: [TranslocoPipe],
  selector: 'app-not-found',
  templateUrl: './not-found.component.html',
  styleUrls: ['./not-found.component.css']
})
export class NotFoundComponent {}
