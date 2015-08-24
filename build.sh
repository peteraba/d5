#! /bin/bash

cd "$GOPATH"/src/github.com/peteraba/d5

# On ubuntu
#sudo apt-get install pip3 libxml2 libxml2-dev libxslt1 libxslt1-dev

# Install prerequisite
#sudo pip3 install lxml ezodf openpyxl

# Install parser
echo "Install parser."
cd parser
go install

# Install persister
echo "Install persister."
cd ../persister
go install

# Install finder
echo "Install finder."
cd ../finder
go install

# Install scorer
echo "Install scorer."
cd ../scorer
go install

# Install derdiedas
echo "Install derdiedas."
cd ../game/derdiedas
go install

# Install conjugate
echo "Install conjugate."
cd ../conjugate
go install

echo "Build done."
echo ""
