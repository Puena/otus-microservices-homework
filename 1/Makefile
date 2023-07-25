IMAGE_NAME=go-health

# create docker image
docker-build: 
	docker build --platform linux/amd64 -t $(IMAGE_NAME) .

# run docker image
docker-run:
	docker run --name healthcontainer -d -p 8080:8080 $(IMAGE_NAME)
	
# log docker container
docker-log:
	docker logs -f healthcontainer

# stop docker container
docker-stop:
	docker stop healthcontainer
	
# remove docker image
docker-remove:
	docker rm -f healthcontainer
	docker rmi -f $(IMAGE_NAME)
	
# test health endpoint
test-health:
	curl http://localhost:8080/health -v