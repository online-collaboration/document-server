.PHONY: test

PWD := $(shell pwd)
CONTAINER := $(shell docker ps | grep document-server | awk '{print $$1}')
WORKDIR = /opt/document-server

# build project by alpine3.15 image
build:
	docker run --rm \
		-v $(PWD):$(WORKDIR) \
		-w /opt/document-server \
		-ti golang:1.17.5-alpine3.15 \
		go build -o document-server main.go

compose:
	docker-compose -f docker-compose/docker-compose.yml up

# startup in development mode
dev:
ifeq ($(CONTAINER) , $())
	echo "project init ..."
else
	docker rm -f $(CONTAINER)
endif
	docker build -f docker/dev/Dockerfile -t document-server:0.01 .
	docker run -v $(PWD):$(WORKDIR) -p 44800:44800 document-server:0.01

# kill docker container
kill:
	docker rm -f $(CONTAINER)

