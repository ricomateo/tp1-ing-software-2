include .env

server_image: 
	docker build . -f Dockerfile -t server

test_image:
	docker build . -f Dockerfile.test -t test

start_server:
	docker run --rm -it \
		--network classconnect \
		--name server \
		-e HOST=$(HOST) \
		-e PORT=$(PORT) \
		-e ENVIRONMENT=$(ENVIRONMENT) \
		-p $(PORT):$(PORT) server

docker_network:
	docker network create classconnect

tests:
	docker compose up --abort-on-container-exit
	docker compose down
