import { Pipe, PipeTransform } from '@angular/core';
@Pipe({name: 'htmlToPlain'})
export class HtmlToPlainPipe implements PipeTransform {
  transform(value: string): string {
    return value.replace(/<[^>]*>/g, '');
  }
}
