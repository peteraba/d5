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
	docker rm --force peteraba/d5 docker/d5

build:
	$(MAKE) build_go
	$(MAKE) copy_binaries_for_docker
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
	cd $(PWD)/docker/$(NAME); GOOS=linux go build -o ./bin/$(NAME) github.com/peteraba/d5/$(RELPATH)$(NAME)

build_docker:
	$(MAKE) build_docker_image NAME=python
	$(MAKE) build_docker_image NAME=spreadsheet
	$(MAKE) build_docker_image NAME=parser
	$(MAKE) build_docker_image NAME=persister
	$(MAKE) build_docker_image NAME=finder
	$(MAKE) build_docker_image NAME=scorer
	$(MAKE) build_docker_image NAME=admin
	$(MAKE) build_docker_image NAME=derdiedas
	$(MAKE) build_docker_image NAME=conjugate
	docker build --force-rm -t peteraba/d5 docker/d5

build_docker_image:
	docker build --force-rm -t peteraba/d5-$(NAME) docker/$(NAME)

copy_binaries_for_docker:
	cp $(PWD)/spreadsheet/csv $(PWD)/docker/spreadsheet/bin/
	cp $(PWD)/spreadsheet/xlsx $(PWD)/docker/spreadsheet/bin/
	cp $(PWD)/spreadsheet/ods $(PWD)/docker/spreadsheet/bin/
	cp $(PWD)/spreadsheet/spreadsheet $(PWD)/docker/spreadsheet/bin/
	
	cp $(PWD)/spreadsheet/csv $(PWD)/docker/d5/bin/
	cp $(PWD)/spreadsheet/xlsx $(PWD)/docker/d5/bin/
	cp $(PWD)/spreadsheet/ods $(PWD)/docker/d5/bin/
	cp $(PWD)/spreadsheet/spreadsheet $(PWD)/docker/d5/bin/
	
	cp $(PWD)/docker/parser/bin/parser $(PWD)/docker/d5/bin/
	cp $(PWD)/docker/persister/bin/persister $(PWD)/docker/d5/bin/
	cp $(PWD)/docker/finder/bin/finder $(PWD)/docker/d5/bin/
	cp $(PWD)/docker/scorer/bin/scorer $(PWD)/docker/d5/bin/
	cp $(PWD)/docker/admin/bin/admin $(PWD)/docker/d5/bin/
	cp $(PWD)/docker/derdiedas/bin/derdiedas $(PWD)/docker/d5/bin/
	cp $(PWD)/docker/conjugate/bin/conjugate $(PWD)/docker/d5/bin/

push:
	docker push peteraba/d5
	$(MAKE) push_docker_image NAME=spreadsheet
	$(MAKE) push_docker_image NAME=parser
	$(MAKE) push_docker_image NAME=persister
	$(MAKE) push_docker_image NAME=finder
	$(MAKE) push_docker_image NAME=scorer
	$(MAKE) push_docker_image NAME=admin
	$(MAKE) push_docker_image NAME=derdiedas
	$(MAKE) push_docker_image NAME=conjugate

push_docker_image:
	docker push peteraba/d5-$(NAME)

