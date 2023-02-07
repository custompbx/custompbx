import {Pipe, PipeTransform} from '@angular/core';

@Pipe({name: 'log'})
export class LogPipe implements PipeTransform {
  public transform(value: object): void {
    console.log(value);
    return;
  }
}
