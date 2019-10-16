#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef enum auth {
    ANON,
    GUESST,
    USER,
    ADMIN,
    ROOT
} auth_t;

struct user {
    char *name;
    auth_t auth;
};

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

void menu() {
    puts("Available commands:");
    puts("\tuser - shows current user info");
    puts("\tlogin <user> - Login to user");
    puts("\tsudo <level> - Set authorization level (must be < 5)");
    puts("\tget-flag - Prints flag (requires auth level 5)");
    puts("\tlogout - log out and reset auth");
    puts("\tquit");
}

int main(int argc, char* argv[], char* envp[]) {
    char buf[512];
    char *arg;
    int level;
    struct user *user = NULL;

    setbuf(stdout,  NULL);

    menu();

    while(1) {
        puts("Enter cmd: ");
        putchar('>');
        putchar(' ');

        if (fgets(buf, 512, stdin) == NULL)
            break;

        if (!strncmp(buf, "user", 4)) {
            if (user == NULL) {
                puts("Not logged in.");
            } else {
                printf("Logged in %s [%u]\n", user->name, user->auth);
            }
        } else if(!strncmp(buf, "login", 5)) {
            if (user != NULL) {
                puts("Already logged in. logout first");
                continue;
            }

            arg = strtok(&buf[6], "\n");
            if (arg == NULL) {
                puts("Invalid command");
                continue;
            }

            user = (struct user*) malloc(sizeof(struct user));
            if (user == NULL) {
                puts("Malloc failed");
                exit(-1);
            }

            user->name = strdup(arg);
            printf("Logged in as \"%s\"\n", arg);
        } else if (!strncmp(buf, "sudo", 4)) {
            if (user == NULL) {
                puts("Login first.");
                continue;
            }

            arg = strtok(&buf[5], "\n");
            if (arg == NULL) {
                puts("Invalid command");
                continue;
            }

            level = atoi(arg);

            if (level >= 5) {
                puts("Can only set below 5");
                continue;
            }
            user->auth = level;
            printf("Set auth level to %u\n", level);
        } else if (!strncmp(buf, "getflag", 7)) {
            if (user == NULL) {
                puts("Login first!");
                continue;
            }

            if (user->auth != 5) {
                puts("Unauthorized");
                continue;
            }
            get_flag();
        } else if(!strncmp(buf, "logout", 6)) {
            if (user == NULL) {
                puts("Not logged in!");
                continue;
            }

            free(user->name);
            user = NULL;
            puts("Logged out");
        } else {
            puts("What?");
            return 0;
        }
    }
}



















