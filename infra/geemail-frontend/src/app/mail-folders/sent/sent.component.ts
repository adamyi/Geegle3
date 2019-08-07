import { Component, OnInit } from '@angular/core';
import { MailService } from 'src/app/shared/mail.service';
import { Mail } from 'src/app/shared/mail.model';
import { ActivatedRoute, Router, Params } from '@angular/router';

@Component({
  selector: 'app-sent',
  templateUrl: './sent.component.html',
  styleUrls: ['./sent.component.css']
})
export class SentComponent implements OnInit {

  sentMails: Mail[];
  id: number;
  currentFolder: string;

  constructor(private mailService: MailService,
              private route: ActivatedRoute,
              private router: Router) { }

  ngOnInit() {
    this.mailService.currentSent.subscribe(mail => this.sentMails = mail);
    this.route.params.subscribe((params: Params) => {
      this.id = +params["id"];
    });
  }
}
