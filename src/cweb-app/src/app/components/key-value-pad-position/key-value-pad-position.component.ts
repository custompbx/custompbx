import {Component, Input} from '@angular/core';
import {KeyValuePad2Component} from '../key-value-pad-2/key-value-pad-2.component';

@Component({
  standalone: true,
  imports: [KeyValuePad2Component],
  selector: 'app-key-value-pad-position',
  templateUrl: './key-value-pad-position.component.html',
  styleUrls: ['./key-value-pad-position.component.css'],
})
export class KeyValuePadPositionComponent {
  @Input() items: object;
  @Input() newItems: Array<any>;
  @Input() exist: boolean;
  @Input() id: number;
  @Input() toCopy: number;
  @Input() dispatchersCallbacks: any;
  @Input() fieldsMask: {
    name: {name: string, style?: object, depend?: string, value?: string, size?: 'sm' | 'md' | 'wide', required?: boolean},
    value?: {name: string, style?: object, depend?: string, value?: string, size?: 'sm' | 'md' | 'wide', required?: boolean},
    extraField1?: {name: string, style?: object, depend?: string, value?: string, size?: 'sm' | 'md' | 'wide', required?: boolean},
    extraField2?: {name: string, style?: object, depend?: string, value?: string, size?: 'sm' | 'md' | 'wide', required?: boolean},
  };
  @Input() grandParentId: number;
  @Input() pending = false;
}
