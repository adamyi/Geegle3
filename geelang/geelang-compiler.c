#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#include "geelang.h"

void __attribute__((noreturn)) error(char* msg) {
    puts(msg);
    _exit(1);
    __builtin_unreachable();
}

// NULL Terminated arrays
struct variable* global_variables[MAX_VARS];
static struct instruction* program[MAX_INSTRS];

int main(int argc, char* argv[], char* envp[]) {
    load_program("./progname", program);
    run_program(program);
    
}

