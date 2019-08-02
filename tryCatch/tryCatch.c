#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <string.h>
#include <unistd.h>

struct totalDataStore {
	int a;
	char b;
	char c[1024];
	void *d;
}; 

char flag[64] = {0};
void _printf(void);
void _puts(void);

void getShell(void);

void handler(int num);

void getFeedback(char *name);
void keyGen(char *cyphertext);
void createData(struct totalDataStore *block);

int main(int argc, char* argv[], char* envp[]) {

	// Set up exception handler for divby0
    struct sigaction sa;
    sa.sa_handler = handler;
    sigemptyset(&sa.sa_mask);
    sigaction(SIGFPE, &sa, NULL);

    printf("Welcome to the employee feedback portal!\nWe value your opinion.\n");
    printf("What's your name?\n");
    char name[1024];
    fgets(name, sizeof(name) - 1, stdin);
    for (int i = 0; i < strlen(name); i++) {
    	if (name[i] == '\n')
    		name[i] = '\0';
    }
    printf("\nHi %s, thanks for taking the time to help us, help you.\n", name);
    getFeedback(name);

	return EXIT_SUCCESS;
}

void getFeedback(char *name) {
label2:	setbuf(stdout, NULL);

label3:
	
	printf("\nHow long have you been working with us?\n");

	char timeWorking[64];
	fgets(timeWorking, sizeof(timeWorking) - 1, stdin);
	for (int i = 0; i < strlen(timeWorking); i++) {
		if (timeWorking[i] == '\n') timeWorking[i] = '\0';
	}

	char *p = timeWorking;
	int num = 0;
	while (*p) { // While there are more characters to process...
	    if ( isdigit(*p) || ( (*p=='-'||*p=='+') && isdigit(*(p+1)) )) {
	        // Found a number
	        long val = strtol(p, &p, 10); // Read number
	        num = (int)val;
        } else {
        // Otherwise, move on to the next character.
        label1: p++;
    	}
    }

    if (num >= 10) {
    	printf("Wow, %s?? That's so long! Thank you for your loyalty!\n\n", timeWorking);
    } else {
    	printf("Only %s? Well thank you for joining us! We hope you like it so far.\n\n", timeWorking);
    }

    int response;

    printf("On a scale of 1 to 10, how do you like the work so far? (10 being the best)\n");
    scanf("%2d", &response);
    printf("Response '%d' recorded.\n\n", response);
    int like = response;
    if (like < 0) goto label1;

    struct totalDataStore *block = malloc(sizeof( struct totalDataStore));
    createData(block);
    block->a = 'i';


    printf("On a scale of 1 to 10, how many pancakes do you want?\n");
    scanf("%2d", &response);
    printf("Response '%d' recorded.\n\n", response);
    int pancakes = response; 
    if (strncmp(block->c, name, strlen(name))) {
    	goto label3;
    }


    printf("On a scale of 1 to 10, how many friends have you made here?\n");
    scanf("%2d", &response);
    printf("Response '%d' recorded.\n\n", response);
    int friends = response;
    if (like == 0) goto label2;

    printf("How many years do you intend to work with us?\n");
    scanf("%2d", &response);
    printf("Response '%d' recorded.\n\n", response);
    int cont = response;

    printf("Calculating more data based on your responses");
    sleep(1);
    printf(".");
    sleep(1);
    printf(".");
    sleep(1);
    printf(".");
    sleep(1);  

    int a = like * pancakes;
    int b = like % friends;
label4: a -= 1;
    int c = cont + like;
    int d = pancakes / friends;
    long long e = 2;
    for (int i = 0; i < 8; i++) {
    	e *= 2;
    }
    if (e == 6) goto label4;
    long f = a + b + c + d + e;


    return;
}

void handler(int num) {
	char string[80] = "";
	strcat(string, ";{\x14Y\x15");
	strcat(string, "EX\x18J_C@J\x15\x15");
	strcat(string, "AR\x1fOU\x11ZZ@\x15WY\x18\\BC]A\x1a?f[]XCT\x12@Q[R\x17MJ\x10");
	strcat(string, "EZZG\x15SEJVB\x11Q\\PP\x0c\x17");
	strcat(string, "2p\x17\\\x12@[\x15");
	strcat(string, "EXJKI\0");
	char key[80] = "12345678901234567890123456789012345678901234567890123456789012345678901234567890";
	for (int i = 0; i < strlen(string); i++) {
		string[i] ^= key[i];
	}
	string[67] = '\0';
	printf("%s", string);
	_printf();
	exit(0);
}

void _printf(void) {

    _puts();
    printf("%s\n\n", flag);
}
void _puts(void){
	char fileName[10] = "N\x04\x0f\x05\x02";
	char key[10] = "abcde";
	for (int i = 0; i < strlen(fileName); i++) {
		fileName[i] ^= key[i];
	}
	fflush(stdout);
    FILE *file = fopen(fileName, "r");

    if (file == NULL) {
    	char tmpflag[24] = "w~rsNrrzlwv{}sjp{y~M";
    	char key[24] = "12345678901234567890";
    	for (int i = 0; i < strlen(tmpflag); i++) {
    		tmpflag[i] ^= key[i];
		}	
        strcpy(flag, tmpflag);
        return;
    }

    fread(flag, sizeof(flag), 1, file);
    fclose(file);
}

void getShell(void) {
	system("/bin/sh");
}

void createData(struct totalDataStore *block) {
	block->a = 0;
	block->b = 'A';
	memset(block->c, 0, sizeof(block->c));
	block->d = NULL;
}