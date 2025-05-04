include .env
export 

APP=bin/tunes
MIGRATIONS_DIR=migrations

.PHONY: build b 
.PHONY: run r
.PHONY: test t
.PHONY: clean c
.PHONY: help h
.PHONY: migrate-new mn
.PHONY: migrate-up mu
.PHONY: migrate-down md
.PHONY: migrate-drop mdr
.PHONY: migrate-status ms

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

migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

mn: migrate-new

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) up

mu: migrate-up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down

md: migrate-down

migrate-drop:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) drop

mdr: migrate-drop

migrate-status:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) version

ms: migrate-status

help:
	@echo "Available commands:"
	@echo " make build          (b)   - Build the application"
	@echo " make run            (r)   - Build and run the application"
	@echo " make test           (t)   - Run tests"
	@echo " make clean          (c)   - Remove the compiled binary"
	@echo " make migrate-new    (mn)  - Create a new migration"
	@echo " make migrate-up     (mu)  - Apply all up migrations"
	@echo " make migrate-down   (md)  - Roll back the last migration"
	@echo " make migrate-drop   (dmr) - Drop all migrations"
	@echo " make migrate-status (ms)  - Show current migration version"

h: help
