#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void flush(void) {
    int ch;
    while ((ch = getchar()) && ch != EOF && ch != '\n'){}
    __asm__("int 0x80");
}

void _() {
    __asm__("xor ecx, ecx; xor edx, edx; ret");

    __asm("ret; pop eax; add al, 8; add ecx, ecx; ret");
}

char banner[] =
    "+=============================================+\n"
    "|Welcome Employee to the Geegle Payroll System|\n"
    "+=============================================+\n";

// HELPER FUNCTIONS

void print_banner() {
    puts(banner);
}

void waitForKey(void) {
    printf("Press any key to continue...\n");
    flush();
    printf("\n");
}

void stuff() {
    char buffer[58] = {0};


    do {
        char start[5] = {0};
        char end[5] = {0};

        printf("\n");
        printf("[A]dd hours to case\n");
        printf("[Q]uit\n");
        printf("\n");
        printf("Enter your choice, (or press enter to refresh): ");

        char ch = getchar();

        switch (ch) {
            case EOF:
            case 'q':
            case 'Q':
                return;

            case 'A':
            case 'a':
                flush();

                printf("Enter start hours (24H fmt HHMM): ");
                fgets(start, 5, stdin);
                flush();


                printf("Enter end hours (24H fmt HHMM): ");
                fgets(end, 5, stdin);
                flush();

                puts("Hours logged...\n");
                strncat(buffer, start, 4);
                strncat(buffer, "-", 1);
                strncat(buffer, end, 4);
                strncat(buffer, "\n", 1);
                printf("Current hours logged are: \n%s\n", buffer);
                break;

            default:
                flush();
                printf("\n");
                break;
        }
    } while (1);
}

int main(void) {
    setbuf(stdout, NULL);
    print_banner();
    printf("Please make sure you run this program under: /bin/sh");
    puts("");

    char name[40];
    printf("Enter employee name: ");
    fgets(name, 40, stdin);
    stuff();
}
