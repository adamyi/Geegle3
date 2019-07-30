import { environment } from '../../environments/environment';
import { HttpClient } from '@angular/common/http';

import { Mail } from "./mail.model";
import { BehaviorSubject } from "rxjs";

const API_URL = environment.apiUrl;

export class MailService {

  constructor(
    private http: HttpClient
  ) {
    this.retrieveMails();
  }

  inboxMail: Mail[] = [];

  sentMail: Mail[] = [];

  check: boolean = false;

  private inboxMails = new BehaviorSubject<Mail[]>(this.inboxMail);
  currentInbox = this.inboxMails.asObservable();

  private sentMails = new BehaviorSubject<Mail[]>(this.sentMail);
  currentSent = this.sentMails.asObservable();

  getCheck() {
    return this.check;
  }

  getInboxMail(index: number) {
    console.log(this.inboxMail);
    return this.inboxMail[index];
  }

  getSentMail(index: number) {
    return this.sentMail[index];
  }

  retrieveMails() {
    this.http.get(API_URL + '/api/userinfo').subscribe(
      response => {
        this.inboxMail = response['inbox'];
        this.sentMail = response['sent'];
        this.inboxMails.next(this.inboxMail);
        this.sentMails.next(this.sentMail);
      },
    );
  }

  mailCount(path: string) {
    if (path === "inbox") {
      if (this.inboxMail.length > 0) {
        return this.inboxMail.length;
      }
      return "";
    } else if (path === "sent") {
      if (this.sentMail.length > 0) {
        return this.sentMail.length;
      }
      return "";
    }
  }
}
