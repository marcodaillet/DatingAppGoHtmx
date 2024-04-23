build:
	go build -o bin/tinder src/*.go

clean:
	rm -r bin

.PHONY: run clean build