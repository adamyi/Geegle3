#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <syslog.h>
#include <string.h>
#include <time.h>
#include <signal.h>

void bug_occured(int sig) {
    if (rand() % 100 < 10) {
         printf("An error has occurred... Saving logs and exiting\n"); 
        _exit(0);
    }

    alarm(10);
}

void flush(void) {
	char ch;
	while ((ch = getchar()) && ch != EOF && ch != '\n'){}
}


/**
 * Mailserver DEV Version.
 * All mails comes to this program.
 * Logs mail and replies if sent to debug@geegle.org
 */

void send_mail() {
    char to[64] = {0};
    char from[64] = {0};
    char mail[1000] = {0};

    printf("mailto: ");
    fgets(to, 64, stdin);
    printf("mailfrom: ");
    fgets(from, 64, stdin);
    printf("\nmail: \n");
    fgets(mail, 1000, stdin);

    if (strchr(to, '\n') != NULL) {
        char* newline = strchr(to, '\n');
        *newline = 0;
    }
    if (strchr(from, '\n') != NULL) {
        char* newline = strchr(from, '\n');
        *newline = 0;
    }

    if (strncmp(to, "debug@geegle.org", 64) == 0) {
        char sendmail[1200] = {0};
        char newmail[1000] = {0};
        
        snprintf(newmail, 1000, mail);

        snprintf(sendmail, 1200, "mailto: %s\nmailfrom: %s\n\n%s", to, from, newmail);

        // System sendmail or somethn
        // Reply with the fmt string

        printf("+ Debug message Logged\n"); 
        printf("+ Replying to %s with debug message\n", from);

        printf("%s", sendmail);

    } else {
        syslog(LOG_INFO, "Received Geegle email: from: %s\nto: %s\nmessage:%s\n", from, to, mail);
        printf("+ Message logged\n");
    }
}

void get_mail() {
    char email[64] = {0};

    printf("email: ");
    fgets(email, 64, stdin);

    char file[128];
    snprintf(file, 128, "/var/mail/%s", email);

    if (strchr(file, '\n') != NULL) {
        char* newline = strchr(file, '\n');
        *newline = 0;
    }


    FILE* mail_file = fopen(file, "r");

    if (mail_file == NULL) {
        printf("No mail found in /var/log/%s\n", email);
    } else {
        char mail[1024] = {0};
        fread(mail, 1024, 1, mail_file);

        printf("You have mail:\n\n%s", mail);
        fclose(mail_file);
    }

}

int main(int argc, char* argv[], char* envp[]) {
    srand(time(NULL));
    setbuf(stdout, NULL);
    
    struct sigaction sa;
    sa.sa_handler = bug_occured;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = SA_RESTART;

    sigaction(SIGALRM, &sa, NULL);

    alarm(10);



    while (1) {
        printf("Loading");
        sleep(1);
        printf(".");
        sleep(1);
        printf(".");
        sleep(1);
        printf(".");
        sleep(1);
        printf("\n");

        printf("> ");
        char option = getchar();    

        switch (option) {
            case 'g':
                flush();
                get_mail();
                printf("<--\n");
                break;
            case 's':
                flush();
                send_mail();
                printf("<--\n");
                break;
            default:
                bug_occured(0);
        }


        printf("\n\n");
    }
}

