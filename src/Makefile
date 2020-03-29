libapp.a:app.o
	ar rs $@ $^

libapp.so:app.o
	gcc app.o -shared -o libapp.so
	cp libapp.so ..

app.o:app.c app.h
	gcc -m64 -O2 -fPIC -Wno-attributes -I . -I/opt/intel/sgxsdk/include -c app.c -o app.o

.PHONY: clean
clean:
	rm *.o *.so
