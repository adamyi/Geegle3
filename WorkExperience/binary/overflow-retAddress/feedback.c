#include <stdio.h>
#include <stdlib.h>
#include <string.h>


char data[64] = {0};

void load_external_data();

int oldFuncToGiveMorePermissions(void);

int actuallyDoThings(void);

int sendFeedbackToServer(char *feedback);


int main(int argc, char *argv[]) {
    setbuf(stdout, NULL);
    return actuallyDoThings();
}


int actuallyDoThings(void) {

    setbuf(stdout, NULL);
    char feedback[2] = {0};
    
    printf("\nThank you for completing Work Experience with Geegle inc.!\nWe hope you've enjoyed your time with us.\n");
    printf("In order to ensure that our Work Experience placements continue to love working for us, we like to gather data on how you found your time.\n");
    printf("What feedback would you like to give us? (max. 2048 characters please)\n\n\tFeedback: ");

    gets(feedback);
    if (sendFeedbackToServer(feedback) == 0) {
        printf("\n\tThank you, your feedback has been saved!\n\n");
    } else {
        printf("\nhmm, something went wrong.Thank you anyway!\n");
    }


    return EXIT_SUCCESS;
}

int sendFeedbackToServer(char *feedback) {
    if (feedback == NULL) return -1;
    if (strlen(feedback) < 1) return -2;

    FILE *f = fopen("/dev/null", "w");
    fprintf(f, "%s", feedback);
    fclose(f);

    return 0;

}

// TODO - Please remove this as it's no longer used
int oldFuncToGiveMorePermissions(void) {

    load_external_data();
    printf("%s\n", data);

    return EXIT_SUCCESS;
}

void load_external_data(){
    FILE *file = fopen("/flag", "r");

    if (file == NULL) {
        strcpy(data, "FLAG{DEBUGGING_FLAG}");
        return;
    }

    fread(data, sizeof(data), 1, file);
    fclose(file);
}
