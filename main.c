#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define UNIMPLEMENTED \
    do { \
        fprintf(stderr, "'%s' has not been implemented yet!", \
                __func__); \
        abort(); \
    } while(0); \

struct Expense {
    char name[30];
    double cost;
    unsigned int year;
    unsigned int month;
    unsigned int day;
};

void print_expense(struct Expense exp) {
    printf("%s: €%.2f on %d.%d.%d\n",
            exp.name, exp.cost, exp.day, exp.month, exp.year);
}

double total_months_expenses(FILE *fp) {
    time_t current_time = time(NULL);
    struct tm tm = *localtime(&current_time);

    unsigned int target_year = tm.tm_year + 1900;
    unsigned int target_month = tm.tm_mon + 1;

    char line[500];
    double total = 0.0;

    while((fgets(line, 500, fp)) != NULL) {
        char *token;
        char *rest = line;
        int section = 0;
        unsigned int year = 0;
        unsigned int month = 0;
        double number;

        while((token = strtok(section == 0 ? rest : NULL, ";")) != NULL) {
            if(section == 1) {
                number = atof(token);
            } else if(section == 3) {
                month = atoi(token);
            } else if(section == 4) {
                year = atoi(token);
            }
            section++;
        }

        if(year == target_year && month == target_month) {
            total += number;
        }
    }

    return total;
}

void remove_expense(FILE *fp, char *name, double cost) {
    UNIMPLEMENTED;
}

struct Expense create_expense() {
    struct Expense new_exp;
    time_t current_time = time(NULL);
    struct tm tm = *localtime(&current_time);

    printf("Expense name: ");

    if (fgets(new_exp.name, sizeof(new_exp.name), stdin) != NULL) {
        size_t len = strlen(new_exp.name);

        if (len > 0 && new_exp.name[len - 1] == '\n') {
            new_exp.name[len - 1] = '\0';
        }
    }

    printf("Expense cost: ");
    scanf("%lf", &new_exp.cost);

    new_exp.year = (unsigned int)(tm.tm_year + 1900);
    new_exp.month = (unsigned int)(tm.tm_mon + 1);
    new_exp.day = (unsigned int)(tm.tm_mday);

    return new_exp;
}

void add_expense(FILE *fp, struct Expense exp) {
    fprintf(fp, "%s;%.2lf;%d;%d;%d\n",
            exp.name, exp.cost, exp.day, exp.month, exp.year);

    print_expense(exp);
}

void list_months_expenses(FILE *fp) {
    UNIMPLEMENTED;
}

void help(void) {
    puts("-f -> filename");
    puts("-h -> this message");
    puts("-a -> add an expense");
    puts("-r -> remove an expense");
    puts("-l -> list all expenses this month");
    puts("-t -> show current total spent this month");
}

int main(int argc, char **argv) {
    if(argc < 3) {
        fprintf(stderr, "You must pass in the location of your expenses file and a command!\n");
        help();
        return 1;
    }

    char *expenses_file_name;
    FILE *expenses_file;

    for(int i = 1; i < argc; i++) {
        if(strcmp(argv[i], "-f") == 0) {
            expenses_file_name = argv[i + 1];
        } else if(strcmp(argv[i], "-h") == 0) {
            help();
        } else if(strcmp(argv[i], "-a") == 0) {
            expenses_file = fopen(expenses_file_name, "a");

            if(expenses_file == NULL) {
                fprintf(stderr, "There was an error opening the file!\n");
                return 1;
            }

            struct Expense expense = create_expense();
            add_expense(expenses_file, expense);
        } else if(strcmp(argv[i], "-r") == 0) {
            remove_expense(NULL, NULL, 0.0);
        } else if(strcmp(argv[i], "-l") == 0) {
            list_months_expenses(expenses_file);
        } else if(strcmp(argv[i], "-t") == 0) {
            expenses_file = fopen(expenses_file_name, "r");

            if(expenses_file == NULL) {
                fprintf(stderr, "There was an error opening the file!\n");
                return 1;
            }

            printf("Total spent this month: €%.2lf\n", total_months_expenses(expenses_file));
        }
    }

    fclose(expenses_file);

    return 0;
}
