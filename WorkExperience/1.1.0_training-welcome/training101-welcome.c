#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char flag[64] = {0};

void load_flag();

int main(int argc, char * argv[]) {

	char ans[100] = {0};
	printf("In Linux we have 'Directories', whereas in Windows they are called: ");

	// read in and store in a var
	scanf("%100s", ans);

	if (strncmp(ans, "Folders", 100) != 0 && strncmp(ans, "folders", 100) != 0) {
		printf("Oops, not quite. Try Googling!\n");\
		return 1;
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
