
docker_image: 
	docker build . -t classconnect

start_service:
	docker run --rm -it -e HOST=0.0.0.0 -e PORT=8080 -p 8080:8080 classconnect
