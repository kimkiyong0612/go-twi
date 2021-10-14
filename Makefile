.PHONY: run
run:
	go build && ./api

.PHONY: build for AWS lambda
build:
	GOOS=linux CGO_ENABLED=0 go build main.go && zip function.zip main
