#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <assert.h>
#include <string.h>

#include "geelang.h"

extern struct variable* global_variables[MAX_VARS];

static int INSTRUCTION_POINTER = -1;

// Default print function for int variables
void print_int(char* var_name, int var_value) {
    printf("%s - %d\n", var_name, var_value);
}

/* ================== */


static struct variable* create_variable_from_string(char* var, int init_value) {
    struct variable* variable = malloc(sizeof(struct variable));
    assert(variable != NULL);

    variable->value = init_value;
    variable->print_variable = print_int;
    strncpy(variable->name, var, 16);

    int i = 0;
    for(struct variable** vars = global_variables; *vars != NULL; vars++, i++) {
    }

    if (i >= MAX_VARS) {
        error("Maximum amount of variabled exceeded");
    }

    global_variables[i] = variable;
    return variable;
}

static int get_variable_index(char* var) {
    int i = 0;
    for(struct variable** vars = global_variables; *vars != NULL; vars++, i++) {
        struct variable* variable = *vars;
        if (strcmp(variable->name, var) == 0) {
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

    int value = atoi(ip->arg2);
    struct variable* newvar = create_variable_from_string(ip->arg1, value);

    assert(newvar != NULL);
}

static void set(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    
    if (arg1_index == -1) {
        error("Variable not found!");
    }

    int value = atoi(ip->arg2);
    global_variables[arg1_index]->value = value;
}

static void inc(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    
    if (arg1_index == -1) {
        error("Variable not found!");
    }

    global_variables[arg1_index]->value += 1;
}

static void dec(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    
    if (arg1_index == -1) {
        error("Variable not found!");
    }

    global_variables[arg1_index]->value -= 1;
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

    global_variables[arg1_index]->value = global_variables[arg2_index]->value;
}

static void add(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);
    
    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    global_variables[arg1_index]->value += global_variables[arg2_index]->value;
}

static void sub(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);
    
    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    global_variables[arg1_index]->value -= global_variables[arg2_index]->value;
}

static void mul(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);
    
    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    global_variables[arg1_index]->value *= global_variables[arg2_index]->value;
}

static void divi(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    int arg2_index = get_variable_index(ip->arg2);
    
    if (arg1_index == -1 || arg2_index == -1) {
        error("Variable not found!");
    }

    if (global_variables[arg2_index]->value == 0) {
        error("Div by 0 err");
    }

    global_variables[arg1_index]->value /= global_variables[arg2_index]->value;
}

static void print(struct instruction *ip) {
    int arg1_index = get_variable_index(ip->arg1);
    
    if (arg1_index == -1) {
        error("Variable not found!");
    }
    struct variable* var = global_variables[arg1_index];

    var->print_variable(var->name, var->value);
}

static void jmpz(struct instruction *ip, struct instruction* program[MAX_INSTRS]) {
    int arg1_index = get_variable_index(ip->arg1);
    
    if (arg1_index == -1) {
        error("Variable not found!");
    }

    if (global_variables[arg1_index]->value != 0) {
        return;
    }


    int jump_to = atoi(ip->arg2);
    if (jump_to >= 0 && jump_to <= MAX_INSTRS && program[jump_to] != 0) {
        INSTRUCTION_POINTER = jump_to - 2;
    }
}

static void jmpnz(struct instruction *ip, struct instruction* program[MAX_INSTRS]) {
    int arg1_index = get_variable_index(ip->arg1);
    
    if (arg1_index == -1) {
        error("Variable not found!");
    }

    if (global_variables[arg1_index]->value == 0) {
        return;
    }


    int jump_to = atoi(ip->arg2);
    if (jump_to >= 0 && jump_to <= MAX_INSTRS && program[jump_to] != 0) {
        INSTRUCTION_POINTER = jump_to - 2;
    }
}




static void run_instruction(struct instruction* ip, struct instruction* program[MAX_INSTRS]) {
    switch(ip->instr_type) {
        case INT: 
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


