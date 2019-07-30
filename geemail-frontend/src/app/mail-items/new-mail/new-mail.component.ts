import { Component, OnInit } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { MailService } from 'src/app/shared/mail.service';
import { DataService } from 'src/app/shared/data.service';
import { Mail } from 'src/app/shared/mail.model';

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
    // this.mail.time = (this.monthNames[this.d.getMonth()]).substring(0, 3) + " " + this.d.getDate() + ", " + this.d.getHours() + ":" + this.d.getMinutes() + "h";
    this.mail.time = new Date().getTime()
    this.mail.body = btoa(this.mail.body);
    this.mailService.sentMail.unshift(this.mail);
    alert("Backend for sending email is under development, but it should appear in the frontend now");
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
