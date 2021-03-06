test:
	gcc -c splay.c -o splay.o
	g++ test.cc splay.o -lgtest -lpthread -o output
	./output
	gcc -c splay-topdown.c -o splay.o
	g++ test-topdown.cc splay.o -lgtest -lpthread -o output
	./output