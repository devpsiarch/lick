all:compile 
compile:lick.c lick.h
	gcc lick.c -lz -o lick

init:
	./lick init

clean:
	rm lick
	rm -r .lick

