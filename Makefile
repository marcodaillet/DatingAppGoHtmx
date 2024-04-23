build:
	go build -o bin/matcha src/*.go

clean:
	rm -r bin

re: clean build

.PHONY: run clean build