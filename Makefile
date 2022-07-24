all: build run clean

build:
	@go build -o cc-example cmd/cc-example/main.go
run:
	@./cc-example
clean:
	@rm cc-example

tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor
clean-vendor:
	rm -rf vendor