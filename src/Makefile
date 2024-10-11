# Makefile for ctrl
build:
	mkdir -p build
	go build -o build/ctrl main.go

run: build
	./build/ctrl

clean:
	rm -rf build

.PHONY: build run clean
