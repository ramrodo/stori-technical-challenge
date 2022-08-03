.PHONY: build clean deploy

install:
	@go mod download

build: clean
	@export GO111MODULE=on
	@sh build.sh

clean:
	@rm -rf ./bin
