
.PHONY: build deps clean

build: blink

blink: deps
	go build -o blink .

deps:
	go get -v -d .

clean:
	rm -f blink

