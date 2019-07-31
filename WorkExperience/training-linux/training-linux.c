#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

char flag[64] = {0};

void load_flag();

#define MAX_LEN 100

int main(int argc, char * argv[]) {

	char ans[101] = {0};
	setbuf(stdout, NULL);

	printf("\nTo show the contents of the current directory, we use the command (no spaces): ");
	scanf("%100s", ans);
	for (int i = 0; i < strlen(ans); i++)
		ans[i] = toupper((unsigned char) ans[i]);
	if (strncmp(ans, "LS", 100) != 0) {
		printf("Oops, not quite. Try Googling!\n");\
		return 1;
	}
	for (int i= 0; i < strlen(ans); i++) ans[i] = 0;

	printf("\nTo enter a new directory, the command is (no spaces): ");
	scanf("%100s", ans);
	for (int i = 0; i < strlen(ans); i++)
		ans[i] = toupper((unsigned char) ans[i]);
	if (strncmp(ans, "CD", 100) != 0) {
		printf("Oops, not quite. Try Googling!\n");\
		return 1;
	}
	for (int i= 0; i < strlen(ans); i++) ans[i] = 0;

	// This block has some wild ugly stuff to handle reading in a space
	printf("\nTo go 'up' a directory level, the full command is: ");
	fflush(stdout);
	if (fflush(stdin)) {
		printf("Oops, something went wrong...\n This shouldn't ever print.\n");
		printf("This is not part of the challenge. Please try again.\n");
		return 1;
	}
	int c;
	while ((c = getchar()) != '\n' && c != EOF) { }
	fgets(ans, MAX_LEN, stdin);
	char *pos;
	if ((pos=strchr(ans, '\n')) != NULL)
    	*pos = '\0';
	for (int i = 0; i < strlen(ans); i++)
		ans[i] = toupper((unsigned char) ans[i]);
	if (strncmp(ans, "CD ..", 100) != 0) {
		printf("Oops, not quite. Try Googling!\n");\
		return 1;
	}
	for (int i= 0; i < strlen(ans); i++) ans[i] = 0;

	printf("\nDo you want to get to the next question [Y/n]? ");
	scanf("%100s", ans);
	for (int i = 0; i < strlen(ans); i++)
		ans[i] = toupper((unsigned char) ans[i]);
	if (strncmp(ans, "Y", 100) != 0) {
		printf("Ok, bye!\n");\
		return 1;
	}
	for (int i= 0; i < strlen(ans); i++) ans[i] = 0;

	printf("\nWhat command should I use to get Administrator permissions? (no spaces): ");
	scanf("%100s", ans);
	for (int i = 0; i < strlen(ans); i++)
		ans[i] = toupper((unsigned char) ans[i]);
	if (strncmp(ans, "SUDO", 100) != 0) {
		printf("Oops, not quite. Try Googling!\n");\
		return 1;
	}
	for (int i= 0; i < strlen(ans); i++) ans[i] = 0;

	printf("\nWhen I execute a program, information such as program variables and Return Pointers are stored in memory in a data structure known as a (no spaces): ");
	scanf("%100s", ans);
	for (int i = 0; i < strlen(ans); i++)
		ans[i] = toupper((unsigned char) ans[i]);
	if (strncmp(ans, "STACK", 100) != 0) {
		printf("Oops, not quite. Try Googling!\n");\
		return 1;
	}
	for (int i= 0; i < strlen(ans); i++) ans[i] = 0;

	printf("\nAre you ready to start some real challenges [Y/n]? ");
	scanf("%100s", ans);
	for (int i = 0; i < strlen(ans); i++)
		ans[i] = toupper((unsigned char) ans[i]);
	while (strncmp(ans, "Y", 1) != 0) {
		printf("Let's try that again...\n");
		printf("Are you ready to start some real challenges [Y/n]? ");
		scanf("%100s", ans);
	}

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
