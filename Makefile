APP_NAME=quero2pay
VERSION?=0.0.1
GOCMD=go
CGO_ENABLED=0

.PHONY: all

all: build

clean: clean-bin
	$(GOCMD) clean

clean-bin:
	rm -rf ./bin/*

build:
	$(GOCMD) build -o ./bin ./...

get:
	$(GOCMD) get ./...
	$(GOCMD) mod verify
	$(GOCMD) mod tidy

test:
	$(GOCMD) test -v -race ./...

req-all-endpoints: business-req-endpoints employee-req-endpoints

business-req-endpoints: business-req-create business-req-read-one business-req-read-all business-req-delete

business-create:
	./scripts/business/create-curl.sh

business-read-all:
	./scripts/business/read-all-curl.sh

business-read-one:
	./scripts/business/read-one-curl.sh

business-delete:
	./scripts/business/delete-curl.sh

employee-req-endpoints: employee-req-create employee-req-read-one employee-req-read-all employee-req-delete

employee-create:
	./scripts/employee/create-curl.sh

employee-read-all:
	./scripts/employee/read-all-curl.sh

employee-read-one:
	./scripts/employee/read-one-curl.sh

employee-delete:
	./scripts/employee/delete-curl.sh

lint:
	golangci-lint run

up:
	mkdir -p ./bin
	docker-compose up --build -d

down:
	docker-compose down
