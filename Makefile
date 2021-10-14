.PHONY: run
run:
	go build && ./go-twi

.PHONY: build for AWS lambda
build:
	GOOS=linux CGO_ENABLED=0 go build && zip function.zip go-twi

