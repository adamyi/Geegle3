import { Component, OnInit } from '@angular/core';
import { MailService } from '../../shared/mail.service';
import { Mail } from '../../shared/mail.model';
import { Params, ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-inbox',
  templateUrl: './inbox.component.html',
  styleUrls: ['./inbox.component.css']
})
export class InboxComponent implements OnInit {

  inboxMails: Mail[];
  id: number;
  currentFolder: string;

  constructor(private mailService: MailService,
              private route: ActivatedRoute) { }

  ngOnInit() {
    this.mailService.currentInbox.subscribe(mail => this.inboxMails = mail);
    this.route.params.subscribe((params: Params) => {
      this.id = +params["id"];
    });
  }
}
