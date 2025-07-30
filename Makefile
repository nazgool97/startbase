.PHONY: dev build up down migrate

dev:
	air -c .air.toml

build:
	go build -o bin/startbase ./cmd/server

up:
	docker-compose up --build

down:
	docker-compose down

migrate:
	psql ${DSN} -f migrations/001_add_role_and_reset.sql
