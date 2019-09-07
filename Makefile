.PHONY: postgresql bin build run

postgresql:
	docker rm $$(docker stop postgresql_lenslocked) 2> /dev/null || true
	docker run \
	--name postgresql_lenslocked \
	-e POSTGRES_DB=lenslocked_dev \
	-p 5432:5432 \
	-d postgres:latest

bin:
	go build -i -o bin/lenslocked cmd/lenslocked.go

build:
	go build ./...

run:
	bin/lenslocked
