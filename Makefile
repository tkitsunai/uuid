
.PHONY: install update test all

all: install test

test:
	go test -v ./...

install:
	glide install

update:
	glide update