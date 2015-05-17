#! /bin/bash

cd "$(dirname "$0")"

error=0

source util.sh

function run_test()
{
	line=$(go test github.com/peteraba/d5/lib/$f -cover)
	if [ ${line:0:2} == "ok" ]; then
		test_success
		# remove "ok", trim left
		line="Dir:\t\t$(echo -e "${line:2}" | sed -e 's/^[[:space:]]*//')"
		line="$(echo $line | sed -r 's/ 0/ \nTime:\t\t0/g' | sed -r 's/ coverage: /\nCoverage:\t/g')"
		print_output "$line\n"
	else
		test_error
		print_error "$line\n"
		error=1
	fi
}

function run_tests()
{
	for f in $(find ../lib -type d -printf '%d\t%P\n' | sort -r -nk1 | cut -f2-); 
	do 
		print_title "Starting test: $f"
		run_test $f
	done
}

function main()
{
	local start_time=`date +%s%N | cut -b1-13`

	run_tests

	local end_time=`date +%s%N | cut -b1-13`
	
	local delta_time=$(($end_time - $start_time))
	
	if [ $error -ne 0 ]; then
		test_error
		error=1
	else
		test_success
	fi

	print_output "All tests finished in $delta_time ms."

	if [ $error -ne 0 ]; then
		exit 1
	fi
}

main
