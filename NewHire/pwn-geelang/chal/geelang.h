#ifndef geelang
#define geelang

void __attribute__((noreturn)) error(char* msg);

// Max of 1024 variables or instructions allowed
#define MAX_VARS 1024
#define MAX_INSTRS 1024

// Max length of a single line
#define MAX_LINE_LEN 100

struct variable {
    long int value;

    // Functon called to print this variable
    void (*print_variable)(char* name, long int value);
};

enum INSTR_TYPE {
    INT,
    BOX,
    SET,
    MOV,
    INC,
    DEC,
    DEL,
    ADD,
    SUB,
    MUL,
    DIV,
    PRINT,
    JMPZ,
    JMPNZ
};


// <value> arguments are stored as variables with names
// starting with !. As this is an invalid variable name
// it is unique.
struct instruction {
    enum INSTR_TYPE instr_type;
    char* arg1; // Name of variable
    char* arg2; // Arg2 is null if only takes 1
};


void load_program(char* filename, struct instruction* program[MAX_INSTRS]);

void run_program(struct instruction* program[MAX_INSTRS]);
#endif
