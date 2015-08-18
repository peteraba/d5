#! /bin/bash

cd "$GOPATH"/src/github.com/peteraba/d5

# On ubuntu
#sudo apt-get install pip3 libxml2 libxml2-dev libxslt1 libxslt1-dev

# Install prerequisite
#sudo pip3 install lxml ezodf openpyxl

# Install parser
cd parser
go install

# Install persister
cd ../persister
go install

# Install finder
cd ../finder
go install

# Install scorer
cd ../scorer
go install

# Install derdiedas
cd ../game/derdiedas
go install

