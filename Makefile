CC=gcc
CFLAGS=-Wextra -std=c11 -pedantic
OBJS=main.o

PREFIX=/usr/local
INFODIR=/usr/share/expenses

all: expenses

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

expenses: main.o
	$(CC) $(CFLAGS) -o $@ $(OBJS)

clean:
	rm -f expenses *.o

install: all
	mkdir -p ${DESTDIR}${PREFIX}/bin
	mkdir -p ${INFODIR}
	cp -f expenses ${DESTDIR}${PREFIX}/bin

uninstall:
	rm -f ${DESTDIR}${PREFIX}/bin/expenses
