.PHONY: build run

build:
	CGO_CFLAGS="-w" go build .

run:
	CGO_CFLAGS="-w" go run .
