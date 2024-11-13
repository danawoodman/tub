.DEFAULT_GOAL := dev  
.PHONY: dev
dev:
	@cng -ik '*.go' -- go run .

.PHONY: build
build:
	@go build .
