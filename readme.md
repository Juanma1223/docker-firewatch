# Instructions

## Build docker image with

docker build -t docker-firewatch .

## Run it

docker run -v /var/run/docker.sock:/var/run/docker.sock --net=host --name docker-firewatch docker-firewatch