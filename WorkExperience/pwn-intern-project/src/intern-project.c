#include <stdio.h>
#include <sys/time.h>
#include <string.h>
#include <stdlib.h>

#define KEY 0xd8

void init_random() {
    printf("Crypto requires genuine randomness... Please enter up to 16 characters to seed random number: ");
    
    char buffer[17] = {0};
    fgets(buffer, 17, stdin);
    
    int sum = 0;

    for (char* c = buffer; *c != 0; c++) {
        sum += (int) *c;
    }

    srand(sum); 
}

void encrypt(char* flag, int key) {
    for(char* c = flag; *c != 0; c++) {
        *c = *c ^ key;
    }
}

void decrypt(char* encrypted, int key) {
    for(char* c = encrypted; *c != 0; c++) {
        *c = *c ^ key;
    }
}


char* load_flag() {
    printf("*** loading flag ***\n");

    char* flag = malloc(64);
    FILE* file = fopen("flag", "r");
    
    if (file == NULL) {
        strncpy(flag, "FLAG{DEBUGGING_FLAG}", 64);
        encrypt(flag, KEY);
        return flag;
    }

    fread(flag, 64, 1, file);
    fclose(file);

    encrypt(flag, KEY);
    return flag;
}

void no_brute_forcing_please() {
    printf("L"); 
    sleep(1);
    printf("O"); 
    sleep(1);
    printf("A"); 
    sleep(1);
    printf("D"); 
    sleep(1);
    printf("I"); 
    sleep(1);
    printf("N"); 
    sleep(1);
    printf("G"); 
    sleep(1);
    printf("."); 
    sleep(1);
    printf("."); 
    sleep(1);
    printf(".\n"); 
    sleep(1);
}

int main(void) {
    setbuf(stdout, NULL);
    no_brute_forcing_please();

    init_random();
    
    puts("Cool now that we have enough randomness... Let's try to decrypt the flag.\n");
    printf("The flag is xor encrypted with the byte 0x%x\n", KEY);
    printf("Let's hope the next call to (rand() %% 0xFF) returns the decryption byte (0x%x)\n\n", KEY);

    int decrypt_key = rand() % 0xFF;

    printf("Decryption key generated. Key is 0x%x\n", decrypt_key);
    
    char* flag = load_flag();
    decrypt(flag, decrypt_key);

    printf("%s\n", flag);
}

