#! /bin/bash

cd "$(dirname "$0")"

browser=$1

error=0

source util.sh

function run_test()
{
	dir=$(dirname $1)
	mkdir -p coverage/lib/$dir

	# Generate coverage file
	go test github.com/peteraba/d5/lib/$1 -coverprofile=/tmp/coverage.out

	# Generate htmloutput
	go tool cover -html=/tmp/coverage.out -o=coverage/lib/$f.html

	# Open file in browser
	if [[ -n $browser ]]; then
		if [[ $browser != "silent" ]]; then
			$browser "$PWD"/coverage/lib/"$1".html
		fi
	else
		$(x-www-browser "$PWD"/coverage/lib/"$1".html)
	fi

}

function main()
{
	rm -rf coverage/*

	for f in $(find ../lib -type d -printf '%d\t%P\n' | sort -r -nk1 | cut -f2-); 
	do 
		print_title "Generating coverage data: $f"
		run_test $f
	done
}

main
