# Set dir of Makefile to a variable to use later
MAKEPATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MAKEDIR := $(dir $(MAKEPATH))
GOPATH := $(GOPATH)

UNAME := $(shell uname)

.PHONY: help clean install_deps install all

help:
	@ printf "all\n\
	clean\n\
	install\n\
	build\n\
	push\n\
	"

all:
	@ $(MAKE) clean
	@ $(MAKE) install_deps
	@ $(MAKE) build
	@ $(MAKE) push

clean:
	@ # Coming soon...

install:
	@ ifeq ($(UNAME), Darwin)
	@   install_deps_brew
	@ endif

install_darwin:
	@ ifeq (1, 0)
	@   /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	@ endif
	@ brew update
	@ brew install docker docker-compose go

install_linux:
	@ # Pull requests welcome :)

docker_build:
	@ docker build --force-rm --no-cache -t peteraba/d5-$(NAME) docker/$(NAME)

docker_push:
	@ docker push peteraba/d5-$(NAME)

go_build:
	@ echo "export GOOS=linux; cd $(PWD)/docker/$(NAME); go build github.com/peteraba/d5/$(RELPATH)$(NAME)"
	@ export GOOS=linux; cd $(PWD)/docker/$(NAME); go build github.com/peteraba/d5/$(RELPATH)$(NAME)

build:
	@ $(MAKE) build_go
	@ $(MAKE) build_copy_binaries
	@ $(MAKE) build_docker

build_go:
	@ $(MAKE) go_build NAME=parser
	@ $(MAKE) go_build NAME=persister
	@ $(MAKE) go_build NAME=finder
	@ $(MAKE) go_build NAME=scorer
	@ $(MAKE) go_build NAME=admin
	@ $(MAKE) go_build NAME=derdiedas RELPATH=game/
	@ $(MAKE) go_build NAME=conjugate RELPATH=game/

build_copy_binaries:
	@ cp $(PWD)/spreadsheet/csv $(PWD)/docker/spreadsheet/
	@ cp $(PWD)/spreadsheet/xlsx $(PWD)/docker/spreadsheet/
	@ cp $(PWD)/spreadsheet/ods $(PWD)/docker/spreadsheet/
	@ cp $(PWD)/spreadsheet/spreadsheet $(PWD)/docker/spreadsheet/
	@
	@ cp $(PWD)/spreadsheet/csv $(PWD)/docker/d5/
	@ cp $(PWD)/spreadsheet/xlsx $(PWD)/docker/d5/
	@ cp $(PWD)/spreadsheet/ods $(PWD)/docker/d5/
	@ cp $(PWD)/spreadsheet/spreadsheet $(PWD)/docker/d5/
	@
	@ cp $(PWD)/docker/parser/parser $(PWD)/docker/d5/
	@ cp $(PWD)/docker/persister/persister $(PWD)/docker/d5/
	@ cp $(PWD)/docker/finder/finder $(PWD)/docker/d5/
	@ cp $(PWD)/docker/scorer/scorer $(PWD)/docker/d5/
	@ cp $(PWD)/docker/admin/admin $(PWD)/docker/d5/
	@ cp $(PWD)/docker/derdiedas/derdiedas $(PWD)/docker/d5/
	@ cp $(PWD)/docker/conjugate/conjugate $(PWD)/docker/d5/

build_docker:
	@ $(MAKE) docker_build NAME=python
	@ $(MAKE) docker_build NAME=spreadsheet
	@ $(MAKE) docker_build NAME=parser
	@ $(MAKE) docker_build NAME=persister
	@ $(MAKE) docker_build NAME=finder
	@ $(MAKE) docker_build NAME=scorer
	@ $(MAKE) docker_build NAME=admin
	@ $(MAKE) docker_build NAME=derdiedas
	@ $(MAKE) docker_build NAME=conjugate
	@ docker build --force-rm --no-cache -t peteraba/d5 docker/d5

push:
	@ $(MAKE) docker_push NAME=spreadsheet
	@ $(MAKE) docker_push NAME=parser
	@ $(MAKE) docker_push NAME=persister
	@ $(MAKE) docker_push NAME=finder
	@ $(MAKE) docker_push NAME=scorer
	@ $(MAKE) docker_push NAME=admin
	@ $(MAKE) docker_push NAME=derdiedas
	@ $(MAKE) docker_push NAME=conjugate
	@ docker push peteraba/d5

