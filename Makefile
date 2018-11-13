GOCMD=GO111MODULE=on go

.PHONY: test

test:
	$(GOCMD) test -v ./...
