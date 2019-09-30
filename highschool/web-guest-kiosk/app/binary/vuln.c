#include <stdio.h>
#include <string.h>

char* decode_qrcode(char* qrcode) {
    return qrcode;
}

int main(int argc, char** argv) {
    if (argc == 1) {
        return 1;
    }

    char* qrcode = argv[1];
    setbuf(stdout, NULL);

    char guestname[16] = "guest";
    char qrcode_decoded[32] = {0};

    strcpy(qrcode_decoded, decode_qrcode(qrcode));

    if (!strcmp(guestname, "root")) {
        printf("access granted");
        return 0;
    }
    printf("No access for '%s', you must be '%s'", guestname, "root");
}
