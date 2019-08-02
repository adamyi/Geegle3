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
    assert(argc == 2);
    load_program(argv[1], program);
    run_program(program);

}

