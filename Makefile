.PHONY: build clean deploy

build: install
	GOOS=linux GOPATH=$(shell pwd)  go build -ldflags="-s -w" -o bin/handler

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

install:
	GOPATH=$(shell pwd) go get github.com/aws/aws-lambda-go/events
	GOPATH=$(shell pwd) go get github.com/aws/aws-lambda-go/lambda
