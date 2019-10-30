#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct user {
    char *name;
    int auth;
};

struct user *user = NULL;

void get_flag() {
    char flag[48];
    FILE *f = fopen("flag", "r");
    if (f == NULL) {
        printf("Flag file not found\n");
        exit(-1);
    }

    fgets(flag, 48, f);
    puts(flag);
    fclose(f);
}

int logged_in() {
    if (user == NULL) {
        puts("Not logged in.");
        return 0;
    } else {
        return 1;
    }
}


void menu() {
    puts("Available commands:");
    puts("\tuser - shows current user info");
    puts("\tlogin user - Login to user");
    puts("\tsudo level - Set authorization level (must be < 9)");
    puts("\tgetflag - Prints flag (requires auth level 9)");
    puts("\tlogout - log out and reset auth");
    puts("\tquit");
}

void user_info() {
    if (logged_in()) {
        printf("Logged in %s [%u]\n", user->name, user->auth);
    }
}

void login(char* buf) {
    if (user != NULL) {
        puts("Already logged in. logout first");
        return;
    }

    char* arg = strtok(&buf[6], "\n");
    if (arg == NULL) {
        puts("Invalid command");
        return;
    }

    user = (struct user*) malloc(sizeof(struct user));
    if (user == NULL) {
        puts("Malloc failed");
        exit(-1);
    }

    user->name = strdup(arg);
    printf("Logged in as \"%s\"\n", arg);
}

void sudo(char* buf) {
    if (!logged_in()) {
       return;
    }

    char* arg = strtok(&buf[5], "\n");
    if (arg == NULL) {
        puts("Invalid command");
        return;
    }

    int level = atoi(arg);

    if (level >= 9) {
        puts("Can only set below 9");
        return;
    }
    user->auth = level;
    printf("Set auth level to %u\n", level);
}

void print_flag() {
    if (!logged_in()) {
        return;
    }

    if (user->auth != 9) {
        puts("Unauthorized");
        return;
    }
    get_flag();
}

void logout() {
    if (!logged_in()) {
        return;
    }

    free(user->name);
    user = NULL;
    puts("Logged out");
}

int main(int argc, char* argv[], char* envp[]) {
    setbuf(stdout,  NULL);

    menu();

    while(1) {
        char buf[512];

        puts("Enter cmd: ");
        putchar('>');
        putchar(' ');

        if (fgets(buf, 512, stdin) == NULL)
            break;

        if (!strncmp(buf, "user", 4)) {
            user_info();
        } else if(!strncmp(buf, "login", 5)) {
            login(buf);
        } else if (!strncmp(buf, "sudo", 4)) {
            sudo(buf);
        } else if (!strncmp(buf, "getflag", 7)) {
            print_flag();
        } else if(!strncmp(buf, "logout", 6)) {
            logout();
        } else {
            puts("What?");
            return 0;
        }
    }
}

