#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <assert.h>

#include "geelang.h"

void __attribute__((noreturn)) error(char* msg) {
    puts(msg);
    _exit(1);
    __builtin_unreachable();
}

// NULL Terminated arrays
struct variable* global_variables[MAX_VARS];
char global_variables_names[MAX_VARS][32];

static struct instruction* program[MAX_INSTRS];

int main(int argc, char* argv[], char* envp[]) {
    setbuf(stdout, NULL);
    FILE* file = stdin;
    if (argc == 2) {
        file = fopen(argv[1], "r");
        if (file == NULL) {
            puts("Failed to open file\n");
            exit(0);
        }
    } else {
        puts("Enter program to stdin:");
        puts("To end the program send the command END");
    }

    load_program(file, program);
    run_program(program);

}

