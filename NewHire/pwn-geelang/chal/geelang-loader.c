#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <limits.h>
#include <assert.h>
#include <unistd.h>
#include <regex.h>
#include <ctype.h>

#include "geelang.h"

/** Code used for loading program from file into arrays  */
/* ===================================================== */
static int read_next_line(FILE* file, char* line) {
    if(fgets(line, 99, file) == NULL) {
        error("Reached EOF without END");
    }

    // Make sure the line has a newline
    char* newline = strchr(line, '\n');
    if (newline == NULL) {
        error("No newline in source");
        /** Doesn't return **/
    }

    *newline = 0; //remove newline


    char* comment = strchr(line, '#');
    if (comment != NULL) {
        *comment = 0;
    }

    if (strcmp(line, "END") == 0) {
        return 0;
    }

    return 1;
}

static enum INSTR_TYPE get_instr_type(char* oper, int len) {
    if (strncmp(oper, "INT", len) == 0) {
        return INT;
    }
    if (strncmp(oper, "BOX", len) == 0) {
        return BOX;
    }
    if (strncmp(oper, "SET", len) == 0) {
        return SET;
    }
    if (strncmp(oper, "INC", len) == 0) {
        return INC;
    }
    if (strncmp(oper, "DEC", len) == 0) {
        return DEC;
    }
    if (strncmp(oper, "DEL", len) == 0) {
        return DEL;
    }
    if (strncmp(oper, "MOV", len) == 0) {
        return MOV;
    }
    if (strncmp(oper, "ADD", len) == 0) {
        return ADD;
    }
    if (strncmp(oper, "SUB", len) == 0) {
        return SUB;
    }
    if (strncmp(oper, "MUL", len) == 0) {
        return MUL;
    }
    if (strncmp(oper, "DIV", len) == 0) {
        return DIV;
    }
    if (strncmp(oper, "PRINT", len) == 0) {
        return PRINT;
    }
    if (strncmp(oper, "JMPZ", len) == 0) {
        return JMPZ;
    }
    if (strncmp(oper, "JMPNZ", len) == 0) {
        return JMPNZ;
    }

    error("Unknown operator");
}

static void check_value(char* value) {
    if (value == NULL) {
        error("No value found");
    }
    char *endptr = NULL;
    long int int_value = strtol(value, &endptr, 10);

    if (endptr == NULL) {
        error("Incorrect int found");
    }

    if (*endptr == 0) {
        return; // VALID
    }

    error("Incorrect int found");
}

static void check_variable(char* value) {
    if (value == NULL) {
        error("No value found");
    }

    if (strlen(value) > 14) {
        error("Variable name too long");
    }

    if (strlen(value) < 1) {
        error("Variable name too short");
    }

    char first_char = value[0];
    if (toupper(first_char) < 'A' || toupper(first_char) > 'Z') {
        error("Variable name must start with letter");
    }

    regex_t regex;
    int ret;

    /* Match for non alphanumeric char */
    ret = regcomp(&regex, "[^a-zA-Z0-9]", 0);
    if (ret) {
        error("Could not compile regex");
    }

    /* If we find at least one, bad name */
    ret = regexec(&regex, value, 0, NULL, 0);

    if (!ret) {
        error("Invalid variable name");
    }

    regfree(&regex);
}

static struct instruction* decode_instruction(char* line) {
    struct instruction* instr = malloc(sizeof(struct instruction));
    assert(instr != NULL);

    char* oper = strtok(line, " ");
    if (oper == NULL) {
        error("Empty line");
    }

    char* temp = NULL;
    instr->instr_type = get_instr_type(oper, strlen(oper));
    // Read first argument
    temp = strtok(NULL, " ");
    check_variable(temp); // Make sure valid variable name

    instr->arg1 = strdup(temp);


    switch(instr->instr_type) {
        case INT:
        case BOX:
        case SET:
        case JMPZ:
        case JMPNZ:
            temp = strtok(NULL, " ");
            check_value(temp); // Make sure valid int

            instr->arg2 = strdup(temp);
            break;

        case ADD:
        case SUB:
        case MUL:
        case DIV:
        case MOV:
            temp = strtok(NULL, " ");
            check_variable(temp); // Make sure valid variable name

            instr->arg2 = strdup(temp);
            break;


        case DEL:
        case INC:
        case DEC:
        case PRINT:
            instr->arg2 = NULL;
            break;

        default:
            error("Bad instruction");
    }

    return instr;
}

void load_program(char* filename, struct instruction* program[MAX_INSTRS]) {
    FILE* file = fopen(filename, "r");
    char line[MAX_LINE_LEN] = {0};

    // Reach each line
    int count = 0;
    while(read_next_line(file, line)) {
        if (count >= MAX_INSTRS) {
            error("Too many instructions");
        }

        if (strlen(line) < 2) {
            continue;
        }

        struct instruction* instr = decode_instruction(line);
        program[count++] = instr;
    }

    program[count] = NULL;
}



