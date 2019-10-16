import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { MatDialogModule } from '@angular/material/dialog';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { MailItemsComponent } from './mail-items/mail-items.component';
import { HeaderComponent } from './header/header.component';
import { MailFoldersComponent } from './mail-folders/mail-folders.component';
import { AppRoutingModule } from './app-routing.module';
import { OpenMailComponent } from './mail-folders/open-mail/open-mail.component';
import { MailService } from './shared/mail.service';
import { InboxComponent } from './mail-folders/inbox/inbox.component';
import { SentComponent } from './mail-folders/sent/sent.component';
import { NewMailComponent } from './mail-items/new-mail/new-mail.component';
import { DataService } from './shared/data.service';
import { FormsModule } from '@angular/forms';
import { HighlightDirective } from './shared/highlight.directive';
import { HttpClientModule } from '@angular/common/http';
import { Base64Pipe } from './b64decode.pipe';
import { ReplaceLineBreaksPipe } from './replaceLineBreaks.pipe';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { WithCredentialsInterceptor } from './credentials.interceptor';

@NgModule({
  declarations: [
    AppComponent,
    MailItemsComponent,
    HeaderComponent,
    MailFoldersComponent,
    OpenMailComponent,
    InboxComponent,
    SentComponent,
    NewMailComponent,
    HighlightDirective,
    Base64Pipe,
    ReplaceLineBreaksPipe,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MatDialogModule,
    BrowserAnimationsModule,
    FormsModule,
    HttpClientModule,
  ],
  providers: [MailService, DataService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: WithCredentialsInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent],
  entryComponents: [NewMailComponent]
})
export class AppModule { }
