#! /bin/bash

cd "$(dirname "$0")"

browser=$1

error=0

source util.sh

function run_test()
{
  filename=$1
	directory=$(dirname ${filename})
	mkdir -p "coverage/lib/${directory}"

	# Generate coverage file
	go test "github.com/peteraba/d5/lib/${filename}" -coverprofile=/tmp/coverage.out

	# Generate html output
	go tool cover -html=/tmp/coverage.out -o=coverage/lib/${filename}.html

	# Open file in browser
	if [[ -n ${browser} ]]; then
		if [[ ${browser} != "silent" ]]; then
			${browser} "${PWD}/coverage/lib/${filename}.html"
		fi
	else
		x-www-browser "${PWD}/coverage/lib/${filename}.html"
	fi
}

function main()
{
	rm -rf coverage/*

	for filename in $(find ../lib -type d -printf '%d\t%P\n' | sort -r -nk1 | cut -f2-);
	do 
		print_title "Generating coverage data: ${filename}"
		run_test "${filename}"
	done
}

main
