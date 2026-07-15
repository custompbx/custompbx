import {Component, Input} from '@angular/core';
import {IconComponent} from '../icon/icon.component';


@Component({
  standalone: true,
    imports: [IconComponent],
    selector: 'app-inner-header',
    templateUrl: './inner-header.component.html',
    styleUrls: ['./inner-header.component.css']
})
export class InnerHeaderComponent {

  @Input() name: string;
  @Input() subtitle: string;
  @Input() status: string;
  @Input() statusTone: 'default' | 'success' | 'warning' | 'danger' = 'default';
  @Input() loadCounter: number;
  @Input() errorMessage: string | null;

}
