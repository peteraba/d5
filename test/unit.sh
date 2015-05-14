#! /bin/bash

for f in $(find ../lib -type d -printf '%d\t%P\n' | sort -r -nk1 | cut -f2-); 
do 
	echo $(go test github.com/peteraba/d5/lib/$f -cover)
	echo ""
done
