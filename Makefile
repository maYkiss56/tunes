
APP=bin/tunes

.PHONY: build b 
.PHONY: run r
.PHONY: test t
.PHONY: clean c
.PHONY: help h

build: 
	go build -o $(APP) ./cmd/main.go

b: build

run: build
	./$(APP)

r: run

test: 
	go test -v ./...

t: test

clean:
	rm -rf ./bin || true

c: clean


help:
	@echo "Available commands:"
	@echo " make build           - Build the application"
	@echo " make run             - Build and run the application"
	@echo " make test            - Run tests"
	@echo " make clean           - Remove the compiled binary"

h: help
