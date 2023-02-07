import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
  name: 'formatTimer'
})
export class FormatTimerPipe implements PipeTransform {

  transform(time: number): string {
    if (isNaN(time) || time < 0) {
      time = 0;
    }
    const hours = Math.floor(Number(time) / 3600);
    const minutes = Math.floor((Number(time) - (Number(hours) * 3600)) / 60);
    const seconds = Number(time) - (Number(hours) * 3600) - (Number(minutes) * 60);

    return this.padTime(hours) + ':' + this.padTime(minutes) + ':' + this.padTime(seconds);
  }

  padTime(time) {
    return (Number(time) < 10 ? '0' : '') + String(time);
  }

}
