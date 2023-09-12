.PHONY: build build-no-cache up down ps

# docker-compose関連のコマンド

build:
	docker compose build

build-no-cache:
	docker compose build --no-cache

up:
	docker compose up

down:
	docker compose down

ps:
	docker compose ps -a
