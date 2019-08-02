#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <assert.h>
#include <string.h>

#include "geelang.h"

extern struct variable* global_variables[MAX_VARS];
extern char global_variables_names[MAX_VARS][33];

static int INSTRUCTION_POINTER = -1;

// Default print function for int variables
void print_int(char* var_name, long int var_value) {
    printf("%s - %ld\n", var_name, var_value);
}

/* ================== */


static struct variable* create_variable_from_string(char* var, int init_value, int boxed) {
    struct variable* variable = malloc(sizeof(struct variable));
    assert(variable != NULL);

    if (boxed) {
        // Memory needs to be allocated
        variable->value = (long int) malloc(sizeof(long int));
        assert(variable->value != 0);
        *((long int*) variable->value) = init_value;
    } else {
        variable->value = init_value;
    }

    variable->print_variable = print_int;

    int i = 0;
    for(struct variable** vars = global_variables; *vars != NULL; vars++, i++) {
    }

    if (i >= MAX_VARS) {
        error("Maximum amount of variabled exceeded");
    }

    global_variables[i] = variable;
    strncpy(global_variables_names[i], var, 30);
    global_variables_names[i][32] = boxed;
    return variable;
}

static int get_variable_index(char* var) {
    int i = 0;
    for(struct variable** vars = global_variables; *vars != NULL; vars++, i++) {
        if (strcmp(global_variables_names[i], var) == 0) {
            return i;
        }
    }

    return -1;
}

static void initialise(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);

    if (arg1_index != -1) {
        error("Attempting to initalise variable already initialised");
    }

    long int value = strtol(ip->arg2, NULL, 10);
    struct variable* newvar = create_variable_from_string(ip->arg1, value, ip->instr_type == BOX ? 1 : 0);

    assert(newvar != NULL);
}

static void set_var(int variable_id, long int value) {
    if (global_variables_names[variable_id][32] == 1) {
        long int* ptr = global_variables[variable_id]->value;
        *ptr = value;
    }

    global_variables[variable_id]->value = value;
}

static long int get_var(int variable_id) {
    if (global_variables_names[variable_id][32] == 1) {
        return *((long int*) global_variables[variable_id]->value);
    }

    return global_variables[variable_id]->value;
}


static void set(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);

    if (arg1_index == -1) {
        error("Variable not found!");
    }

    long int value = strtol(ip->arg2, NULL, 10);
    set_var(arg1_index, value);
}

static void inc(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);

    if (arg1_index == -1) {
        error("Variable not found!");
    }

    set_var(arg1_index, get_var(arg1_index) + 1);
}

static void dec(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);

    if (arg1_index == -1) {
        error("Variable not found!");
    }

    set_var(arg1_index, get_var(arg1_index) - 1);
}


static void del(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);

    if (arg1_index == -1) {
        error("Variable not found!");
    }

    free(global_variables[arg1_index]);
}

static void move(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);

    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    set_var(arg1_index, get_var(arg2_index));
}

static void add(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);

    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    set_var(arg1_index, (long int) get_var(arg1_index) + (long int) get_var(arg2_index));
}

static void sub(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);

    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    set_var(arg1_index, get_var(arg1_index) - get_var(arg2_index));
}

static void mul(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);

    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    set_var(arg1_index, get_var(arg1_index) * get_var(arg2_index));
}

static void divi(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);

    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    if (get_var(arg2_index) == 0) {
        error("Div by 0 err");
    }

    set_var(arg1_index, get_var(arg1_index) / get_var(arg2_index));
}

static void print(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);

    if (arg1_index == -1) {
        error("Variable not found!");
    }
    struct variable* var = global_variables[arg1_index];

    var->print_variable(global_variables_names[arg1_index], get_var(arg1_index));
}

static void jmpz(struct instruction *ip, struct instruction* program[MAX_INSTRS]) {
    int arg1_index = get_variable_index(ip->arg1);

    if (arg1_index == -1) {
        error("Variable not found!");
    }

    if (get_var(arg1_index) != 0) {
        return;
    }


    long int jump_to = strtol(ip->arg2, NULL, 10);
    if (jump_to >= 0 && jump_to <= MAX_INSTRS && program[jump_to] != 0) {
        INSTRUCTION_POINTER = jump_to - 2;
    }
}

static void jmpnz(struct instruction *ip, struct instruction* program[MAX_INSTRS]) {
    int arg1_index = get_variable_index(ip->arg1);

    if (arg1_index == -1) {
        error("Variable not found!");
    }

    if (get_var(arg1_index) == 0) {
        return;
    }


    long int jump_to = strtol(ip->arg2, NULL, 10);
    if (jump_to >= 0 && jump_to <= MAX_INSTRS && program[jump_to] != 0) {
        INSTRUCTION_POINTER = jump_to - 2;
    }
}




static void run_instruction(struct instruction* ip, struct instruction* program[MAX_INSTRS]) {
    switch(ip->instr_type) {
        case INT:
        case BOX:
            initialise(ip);
            break;
        case SET:
            set(ip);
            break;
        case MOV:
            move(ip);
            break;
        case INC:
            inc(ip);
            break;
        case DEC:
            dec(ip);
            break;
        case DEL:
            del(ip);
            break;
        case ADD:
            add(ip);
            break;
        case SUB:
            sub(ip);
            break;
        case MUL:
            mul(ip);
            break;
        case DIV:
            divi(ip);
            break;
        case PRINT:
            print(ip);
            break;
        case JMPZ:
            jmpz(ip, program);
            break;
        case JMPNZ:
            jmpnz(ip, program);
            break;
    }
}

void run_program(struct instruction* program[MAX_INSTRS]) {
    for (INSTRUCTION_POINTER = 0; program[INSTRUCTION_POINTER] != NULL; INSTRUCTION_POINTER++) {
        run_instruction(program[INSTRUCTION_POINTER], program);
    }
}


