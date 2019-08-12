import { Component, OnInit } from "@angular/core";
import { MailService } from "../../shared/mail.service";
import { Mail } from "../../shared/mail.model";
import { ActivatedRoute, Router, Params } from "@angular/router";

@Component({
  selector: "app-open-mail",
  templateUrl: "./open-mail.component.html",
  styleUrls: ["./open-mail.component.css"]
})
export class OpenMailComponent implements OnInit {
  openMail;
  id: number;
  currentFolder: string;

  constructor(
    private mailService: MailService,
    private route: ActivatedRoute,
    private router: Router
  ) {}

  ngOnInit() {
    this.route.params.subscribe((params: Params) => {
      this.id = +params["id"];
      if (this.router.url === "/inbox/" + this.id) {
        this.openMail = this.mailService.getInboxMail(this.id);
      } else if (this.router.url === "/sent/" + this.id) {
        this.openMail = this.mailService.getSentMail(this.id);
      }
      if (this.openMail == null) {
        this.router.navigate(['../']);
      }
    });
    this.currentFolder = this.route.snapshot.url[0].path;
  }

  goBack() {
    this.router.navigate([this.route.snapshot.url[0].path]);
  }
}
