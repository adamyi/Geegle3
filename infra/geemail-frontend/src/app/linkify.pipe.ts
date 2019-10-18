import { Pipe, PipeTransform } from '@angular/core';
@Pipe({name: 'linkify'})
export class LinkifyPipe implements PipeTransform {
  transform(value: string): string {
    return this.linkify(value);
  }
  private linkify(plainText): string{
    let replacedText;
    let replacePattern1;
    let replacePattern2;

    //URLs starting with http://, https://, or ftp://
    replacePattern1 = /(\b(https?|ftp):\/\/[-A-Z0-9+&@#\/%?=~_|!:,.;]*[-A-Z0-9+&@#\/%=~_|])/gim;
    replacedText = plainText.replace(replacePattern1, '<a href="$1" target="_blank">$1</a>');

    //URLs starting with "www." (without // before it, or it'd re-link the ones done above).
    replacePattern2 = /(^|[^\/])(www\.[\S]+(\b|$))/gim;
    replacedText = replacedText.replace(replacePattern2, '$1<a href="http://$2" target="_blank">$2</a>');

    return replacedText;
  }
}
