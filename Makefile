
server_image: 
	docker build . -f Dockerfile -t server

test_image:
	docker build . -f test.Dockerfile -t test

start_server:
	docker run --rm -it \
		--network classconnect \
		--name server \
		-e HOST=0.0.0.0 \
		-e PORT=8080 -p 8080:8080 server

docker_network:
	docker network create classconnect

tests:
	docker compose up --abort-on-container-exit
	docker compose down
