import { Pipe, PipeTransform } from '@angular/core';
@Pipe({name: 'b64decode'})
export class Base64Pipe implements PipeTransform {
  transform(value: string): string {
    return atob(value);
  }
}
