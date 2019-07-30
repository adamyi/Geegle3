import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Mail } from '../shared/mail.model';
import { MailService } from '../shared/mail.service';

@Component({
  selector: 'app-mail-items',
  templateUrl: './mail-items.component.html',
  styleUrls: ['./mail-items.component.css']
})
export class MailItemsComponent implements OnInit {

  

  constructor(private mailService: MailService,
              private router: Router) { }

  ngOnInit() {

  }

}
