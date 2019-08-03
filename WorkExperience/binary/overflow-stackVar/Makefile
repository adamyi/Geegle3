CC=gcc
CFLAGS=-m32 -fno-stack-protector -O0 -no-pie -mpreferred-stack-boundary=2
LDFLAGS=-m32 -z execstack -no-pie -fno-pic

game: game.o
clean:
	rm *.o
