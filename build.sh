#! /bin/bash

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

