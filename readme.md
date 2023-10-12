# This sofware is intended to implement alerts and restart crashed docker containers automatically

# Instructions to run 

## Build docker image with

docker build -t docker-firewatch .

## Run it

docker run -e PORT=4200 -v /var/run/docker.sock:/var/run/docker.sock --net=host --name docker-firewatch docker-firewatch

# Developing

## You can create a .env file to fill variables needed
