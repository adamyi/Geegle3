import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { MailService } from '../../shared/mail.service';
import { DataService } from '../../shared/data.service';
import { Mail } from '../../shared/mail.model';

@Component({
  selector: 'app-new-mail',
  templateUrl: './new-mail.component.html',
  styleUrls: ['./new-mail.component.css']
})
export class NewMailComponent implements OnInit {

  checkIfopen: boolean;

  mail = new Mail('', '', '', '', 0);
  monthNames = ["January", "February", "March", "April", "May", "June",
  "July", "August", "September", "October", "November", "December"];
  d = new Date();

  onSubmit() {
    this.mail.time = new Date().getTime();
    this.mail.sender = this.mailService.username;
    this.mail.body = btoa(this.mail.body);
    this.mailService.sendMail(this.mail).subscribe(mail => {
      this.mailService.sentMail.unshift(this.mail);
      if (this.mail.receiver == this.mail.sender) {
        this.mailService.inboxMail.unshift(this.mail);
      }
      // this.mailService.retrieveMails();
    });
    this.data.changeButton(false);
    this.dialogRef.close();
  }

  constructor(public dialogRef: MatDialogRef<NewMailComponent>,
              private mailService: MailService,
              private data: DataService) { }

  ngOnInit() {
    this.data.currentCheck.subscribe(check => this.checkIfopen = check);
  }

  onCancel() {
    this.data.changeButton(false);
    this.dialogRef.close();
  }
}
