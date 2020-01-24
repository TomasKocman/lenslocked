.PHONY: postgresql bin build run

postgresql:
	docker rm $$(docker stop postgresql_lenslocked) 2> /dev/null || true
	docker run \
	--name postgresql_lenslocked \
	-e POSTGRES_DB=lenslocked_dev \
	-p 5432:5432 \
	-d postgres:latest

init:
	go mod download && go mod verify

bin:
	go build -o lenslocked.com -i ./cmd/lenslocked/*

build:
	go build ./...

run:
	./lenslocked.com
