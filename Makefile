.PHONY: build run docker-build docker-run docker-compose-up

build:
	go build -o main ./cmd

run: build
	./main

docker-build:
	docker build -t http-proxy-server .

docker-run:
	docker run -p 8080:8080 http-proxy-server

docker compose up:
	docker compose up --build
