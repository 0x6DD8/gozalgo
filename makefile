build:
	mkdir bin
	go build -o ./bin/gozalgo

install:
	make
	chmod +x ./bin/gozalgo
	mv ./bin/gozalgo /usr/local/bin/gozalgo
	make clean

uninstall:
	rm /usr/local/bin/gozalgo

clean:
	rm -rf ./bin

run:
	go run .