# Set dir of Makefile to a variable to use later
MAKEPATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MAKEDIR := $(dir $(MAKEPATH))
GOPATH := $(GOPATH)

UNAME := $(shell uname)

.PHONY: help clean install all

help:
	@ printf "all\n\
	clean\n\
	build\n\
	push\n\
	"

all:
	$(MAKE) clean
	$(MAKE) build
	@ # $(MAKE) push

clean:
	docker rm --force peteraba/d5 || true

build:
	$(MAKE) build_go
	$(MAKE) build_docker

build_go:
	$(MAKE) build_go_service NAME=parser
	$(MAKE) build_go_service NAME=persister
	$(MAKE) build_go_service NAME=finder
	$(MAKE) build_go_service NAME=scorer
	$(MAKE) build_go_service NAME=admin
	$(MAKE) build_go_service NAME=derdiedas RELPATH=game/
	$(MAKE) build_go_service NAME=conjugate RELPATH=game/

build_go_service:
	GOOS=linux go build -o ./$(RELPATH)$(NAME)/bin/$(NAME) github.com/peteraba/d5/$(RELPATH)$(NAME)

build_docker:
	$(MAKE) build_docker_image NAME=spreadsheet
	$(MAKE) build_docker_image NAME=parser
	$(MAKE) build_docker_image NAME=persister
	$(MAKE) build_docker_image NAME=finder
	$(MAKE) build_docker_image NAME=scorer
	$(MAKE) build_docker_image NAME=admin
	$(MAKE) build_docker_image NAME=derdiedas IMAGE_PREFIX=game- RELPATH=game/
	$(MAKE) build_docker_image NAME=conjugate IMAGE_PREFIX=game- RELPATH=game/

build_docker_image:
	docker build --force-rm -t peteraba/d5-$(IMAGE_PREFIX)$(NAME) $(RELPATH)$(NAME)
	
push:
	$(MAKE) push_docker_image NAME=spreadsheet
	$(MAKE) push_docker_image NAME=parser
	$(MAKE) push_docker_image NAME=persister
	$(MAKE) push_docker_image NAME=finder
	$(MAKE) push_docker_image NAME=scorer
	$(MAKE) push_docker_image NAME=admin
	$(MAKE) push_docker_image NAME=derdiedas IMAGE_PREFIX=game-
	$(MAKE) push_docker_image NAME=conjugate IMAGE_PREFIX=game-

push_docker_image:
	docker push peteraba/d5-$(IMAGE_PREFIX)$(NAME)
