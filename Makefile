build:
	go build -o bin/tinder *.go

clean:
	rm -r bin

.PHONY: run clean build