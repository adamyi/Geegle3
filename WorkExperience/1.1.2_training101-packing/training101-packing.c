#include <stdio.h>
#include <stdlib.h>
#include <string.h>


char flag[64] = {0};

void load_flag();

int getFlag(void);

int actuallyDoThings(void);


int main(int argc, char *argv[]) {
    setbuf(stdout, NULL);
    return actuallyDoThings();
}


int actuallyDoThings(void) {

    char myTeam = 'A';
    char response[64];

    memset(response, 0, strlen(response));

    void *win = (void *)getFlag;

    printf("\n");
    printf(">> You approach the other teams base,\n");
    printf(">> and see the flag at %p.\n", win);
    printf("\n");
    printf("Guard: 'Halt! Who goes there?'\n");
    printf("\n");

    printf("You respond: ");
    gets(response);
    printf("\n");

    printf("Guard: 'Hi %s,'\n", response);
    printf("Guard: 'Nobody gets to see the flag,'\n");
    printf("Guard: 'New rules after the last break-in.\n\n");

    fflush(stdout);

    return EXIT_SUCCESS;
}

int getFlag(void) {

    load_flag();
    printf("%s\n", flag);

    return EXIT_SUCCESS;
}

void load_flag(){
    FILE *file = fopen("/flag", "r");

    if (file == NULL) {
        strcpy(flag, "FLAG{DEBUGGING_FLAG}");
        return;
    }

    fread(flag, sizeof(flag), 1, file);
    fclose(file);
}
