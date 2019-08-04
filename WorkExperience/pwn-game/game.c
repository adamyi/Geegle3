#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <ctype.h>
#include <time.h>

#define NUM_COINS 3
#define TICKET_COST 1
#define SUPER_TICKET_COST 5
#define MAX_TICKETS NUM_COINS / TICKET_COST

#define FALSE 0
#define TRUE 1

#define UNDEFINED -1
#define TICKET 1
#define SUPER_TICKET 2

#define NAME_LEN 256

struct ticket {
	int type;
	int valid;
	char name[NAME_LEN];
};

void gameHandler(void);
void printMenu(void);
char getCharWNewline(void);
void createReward(void);
void giveReward(void);

char reward[256] = {0};

int main(int argc, char * argv[]) {

	if (NUM_COINS >= SUPER_TICKET_COST) {
		printf("Invalid game: Coin allocations are incorrect.\n");
		return EXIT_FAILURE;
	}

	printf("Welcome to the totally legal not-gambling game, where you pay money to use surprise mechanics that maybe give you a prize!!!\n");
	printf("Ready to play? [Y/n] ");
	char ans = getCharWNewline();
	if (toupper(ans) == 'Y') {
		gameHandler();
	} else {
		printf("Come back when you want to risk valuable things for prizes!\n");
	}

	return EXIT_SUCCESS;
}

void gameHandler(void) {
	srand(time(NULL));

	int coins = NUM_COINS;
	int tickets = 0;
	int supers = 0;

	struct ticket myTickets[MAX_TICKETS];
	for (int i = 0; i < MAX_TICKETS; i++) {
		myTickets[i].type = UNDEFINED;
		myTickets[i].valid = FALSE;
		memset(myTickets[i].name, 0, NAME_LEN);
	}

	while (coins > 0) {
		printf("\nYou have:\n %d coin(s)\n %d ticket(s)\n %d super tickets\n\tin your hand.\nWhat do you do?\n", coins, tickets, supers);
		printMenu();
		char ansC = getCharWNewline();
		if (ansC < '1' || ansC > '3') {
			printf("'%c' is an invalid response! Try again\n", ansC);
			continue;
		}
		int ans = atoi(&ansC);
		switch (ans) {
			case 1:
				if (coins < TICKET_COST) {
					printf("Not enough coins!\n");
					continue;
				} else {
					coins -= TICKET_COST;
					tickets++;
					int i = MAX_TICKETS - 1;
					while (myTickets[i].type != UNDEFINED && i > 0) i--;
					myTickets[i].type = TICKET;
					myTickets[i].valid = TRUE;
					printf("Name for ticket: ");
					scanf("%s", myTickets[i].name);
					getchar();
					printf("Your name on this ticket is '%s'\n", myTickets[i].name);
				}
				break;
			case 2:
				if (coins < SUPER_TICKET_COST) {
					printf("Not enough coins!\n");
					continue;
				} else {
					coins -= SUPER_TICKET_COST;
					supers++;
					int i = MAX_TICKETS - 1;
					while (myTickets[i].type != UNDEFINED && i > 0) i--;
					myTickets[i].type = SUPER_TICKET;
					myTickets[i].valid = TRUE;
					printf("Name for ticket: ");
					scanf("%s", myTickets[i].name);
					getchar();
					printf("Your name on this ticket is '%s'\n", myTickets[i].name);
				}
				break;
			case 3:
			game:
				printf("\nPlaying Game!\n\n");
				int i = MAX_TICKETS - 1;
				while (myTickets[i].valid == TRUE && i >= 0) {
					int type;
					if (myTickets[i].type == SUPER_TICKET)type = SUPER_TICKET;
					else type = TICKET;
					printf("First ticket (Name: '%s') is a ", myTickets[i].name);
					if (type == SUPER_TICKET) printf("Super ");
					printf("ticket. Rolling...\n");
					int randInt = 10000;
					if (type != SUPER_TICKET) {
						randInt = rand() % 10000 + 1;
					}
					if (randInt == 10000) {
						printf("Winner!\nYour prize: ");
						giveReward();
						printf("Quitting.\n");
						return;
					} else {
						printf("Aww, you didn't win. Moving on.\n");
					}
					i--;
				}
				return;
			default:
				printf("Invalid response received!\n");
				continue;
		}
		if (coins < 1) goto game;
	}

}

char getCharWNewline(void) {
	char c = getchar(); getchar();
	return c;
}

void printMenu(void) {
	printf("\n");
	printf("[1]: Buy a ticket\t\tCost: %d coin\n", TICKET_COST);
	printf("[2]: Buy a super ticket\t\tCost: %d coin\n", SUPER_TICKET_COST);
	printf("[3]: Play with the current tickets\n");
	printf("You decide to choose: ");
}

void giveReward(void) {

    createReward();
    printf("%s\n", reward);
}

void createReward(){
    FILE *file = fopen("/flag", "r");

    if (file == NULL) {
        strcpy(reward, "FLAG{DEBUGGING_FLAG}");
        return;
    }

    fread(reward, sizeof(reward), 1, file);
    fclose(file);
}
