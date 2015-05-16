#! /bin/bash

cd "$(dirname "$0")"

error=0

for f in $(find ../lib -type d -printf '%d\t%P\n' | sort -r -nk1 | cut -f2-); 
do 
	line=$(go test github.com/peteraba/d5/lib/$f -cover)
	if [[  ${line:0:2} -ne "ok" ]]; then
		error=1
	fi
	echo $line
done

if [ "$error" -eq "1" ]; then
	exit 1
fi
