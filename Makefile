NAME     := sendslack
VERSION  := 0.3.0
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS  := -ldflags="-X \"main.version=$(VERSION)\" -X \"main.revision=$(REVISION)\""

bin/$(NAME): format deps
	go build $(LDFLAGS) -o bin/$(NAME)

linux: format deps
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(NAME)

deps:
	glide install

update:
	glide update

format:
	go fmt $(SRCS)

clean:
	rm -rf bin/*

install:
	go install $(LDFLAGS)

zip:
	tar zcvf bin/sendslack_v$(VERSION).tar.gz bin/$(NAME)

digest:
	openssl dgst -sha256 bin/sendslack_v$(VERSION).tar.gz

.PHONY: linux deps update format clean install zip digest
