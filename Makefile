CC=gcc
CFLAGS=-Wextra -std=c11 -pedantic
OBJS=main.o

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

expenses: main.o
	$(CC) $(CFLAGS) -o $@ $(OBJS)
