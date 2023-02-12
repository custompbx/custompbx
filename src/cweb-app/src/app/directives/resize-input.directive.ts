import {Directive, ElementRef, EventEmitter, HostListener, Input, Output} from '@angular/core';

@Directive({
  selector: '[appResizeInput]'
})
export class ResizeInputDirective {

  private mainWidth = 180;
  private textareaNow = 1000;

  @Input() resizeOnString: string;
  @Output() inputType = new EventEmitter();

  @HostListener('focus') onFocus() {
    this.resizeInput();
  }

  @HostListener('paste') onPaste() {
    setTimeout(this.resizeInput.bind(this), 10);
  }

  @HostListener('cut') onCut() {
    setTimeout(this.resizeInput.bind(this), 10);
  }

  @HostListener('input') onInput() {
    this.resizeInput();
  }

  @HostListener('blur') onBlur() {
    this.el['focused'] = false;
    setTimeout(this.noResize.bind(this), 500);
  }

  constructor(private el: ElementRef) { }

  private resizeInput() {
    const length = this.resizeOnString ? this.resizeOnString.length : 0;
    this.el['focused'] = true;
    innerWidth = (length * 8);
    /*if (innerWidth >= this.textareaNow ) {
      this.inputType.next(true);
      this.el.nativeElement.parentElement.style.setProperty('width', this.textareaNow + 'ch');
    } else */if (innerWidth > this.mainWidth) {
      this.inputType.next(false);
      this.el.nativeElement.parentElement.style.setProperty('width', (length + 5) + 'ch');
    } else {
      this.inputType.next(false);
      this.el.nativeElement.parentElement.style.setProperty('width', this.mainWidth + 'px');
    }
  }

  private noResize() {
    if (this.el['focused'] === true) {
      return;
    }
    this.el.nativeElement.parentElement.style.setProperty('width', this.mainWidth + 'px');
  }
}
