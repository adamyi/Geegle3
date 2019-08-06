import { Component, OnInit } from '@angular/core';
import { MailService } from '../shared/mail.service';
import { MatDialog, MatDialogConfig } from '@angular/material/dialog';
import { NewMailComponent } from '../mail-items/new-mail/new-mail.component';
import { DataService } from '../shared/data.service';

@Component({
  selector: 'app-mail-folders',
  templateUrl: './mail-folders.component.html',
  styleUrls: ['./mail-folders.component.css']
})
export class MailFoldersComponent implements OnInit {

  checkIfopen: boolean = false;

  constructor(private mailService: MailService,
              private dialog: MatDialog,
              private data: DataService) { }

  ngOnInit() {
    this.data.currentCheck.subscribe(check => this.checkIfopen = check);
  }

  mailCounter(path: string) {
    return this.mailService.mailCount(path);
  }

  newMail() {
    const dialogConfig = new MatDialogConfig();
    dialogConfig.disableClose = true;
    dialogConfig.autoFocus = true;
    this.dialog.open(NewMailComponent, dialogConfig);
    this.checkIfopen = true;
  }
}
