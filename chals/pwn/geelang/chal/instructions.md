Simple programming language

Online compiler, so we can teach our interns how to code...

Variable naming convention: 
- Regex: [a-zA-Z][a-zA-Z0-9]{0, 14}
- English: Must start with letter, can contain letters or numbers, max size 15 chars.

Statements must end in a newline

boxed vs unboxed

Statements: 
# Comment
INT x <val>     # Create's a unboxed int variable named X, value <value> (Must be 32 bit signed int)
BOX x <val>     # Create's a boxed int variable named X, value <value> (Must be 32 bit signed int)
SET x <val>     # Reassigns variable x to value <value>. Variable x must exist
MOV x y         # move variable y into x
DEL x           # Deletes variable x and frees resources associated to x

INC x
DEC x

ADD x y         # Adds variable x and y, stores value in x
SUB x y         # Subtracts variable y from x, stores value in x
MUL x y         # Multiplies variable x and y, stores value in x
DIV x y         # Divides variable x by y, stores value in x
PRINT x         # Prints variable x as a decimal

JMPZ x <line>     # Jmp to line if x is 0
JMPNZ x <line>    # Jmp to line if x is not 0

**All Programs must end with END**

# Usage

`$ geelang-compiler <file.geelang>`
or
`geelang-compiler` and enter program to stdin

