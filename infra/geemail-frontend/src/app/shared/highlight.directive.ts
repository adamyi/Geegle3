import { Directive, Renderer2, ElementRef, HostListener, } from '@angular/core';

@Directive({
  selector: '[appHighlight]'
})
export class HighlightDirective  {

  constructor(private elementRef: ElementRef,
              private renderer: Renderer2) { }

  @HostListener('mouseenter') mouseover() {
    this.renderer.setStyle(this.elementRef.nativeElement, 'box-shadow', '#DADCE0 1px 0px 0px 0px inset, #DADCE0 -1px 0px 0px 0px inset, rgba(60,64,67,0.3) 0px 1px 2px 0px, rgba(60,64,67,0.15) 0px 1px 3px 1px');
  }

  @HostListener('mouseleave') mouseleave() {
    this.renderer.setStyle(this.elementRef.nativeElement, 'box-shadow', 'none');
  }
}
