import { MailFoldersComponent } from "./mail-folders/mail-folders.component";
import { Routes, RouterModule } from "@angular/router";
import { NgModule } from "@angular/core";
import { EditorModule } from '@tinymce/tinymce-angular';

import { InboxComponent } from "./mail-folders/inbox/inbox.component";
import { SentComponent } from "./mail-folders/sent/sent.component";
import { OpenMailComponent } from "./mail-folders/open-mail/open-mail.component";

const appRoutes: Routes = [
  { path: '', redirectTo: "/inbox", pathMatch: 'full' },
  { path: '', component: MailFoldersComponent },
  { path: 'inbox', component: InboxComponent},
  { path: 'inbox/:id', component: OpenMailComponent },
  { path: 'sent', component: SentComponent },
  { path: 'sent/:id', component: OpenMailComponent },
  { path: '**', redirectTo: "/inbox", pathMatch: 'full'}
];

  @NgModule({
    imports: [RouterModule.forRoot(appRoutes), EditorModule],
    exports: [RouterModule]
  })

  export class AppRoutingModule {}
