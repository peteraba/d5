#! /bin/bash

cd "$(dirname "$0")"

error=0

source util.sh

function run_test()
{
  filename="$1"
	line=$(go test "github.com/peteraba/d5/lib/${filename}" -cover)
	if [ ${line:0:2} == "ok" ]; then
		line="Dir:\t\t$(echo -e "${line:2}" | sed -e 's/^[[:space:]]*//')"
		line="$(echo ${line} | sed -r 's/ 0/ \nTime:\t\t0/g' | sed -r 's/ coverage: /\nCoverage:\t/g')"
		if [[ "${line}" == *"100.0%"* ]]; then
			test_success
		else
			test_warning
		fi
		print_output "${line}\n"
	else
		test_error
		print_error "${line}\n"
		error=1
	fi
}

function run_tests()
{
	for filename in $(find ../lib -type d -printf '%d\t%P\n' | sort -r -nk1 | cut -f2-);
	do
		count=`ls -1 "../lib/${filename}/*.go" 2>/dev/null | wc -l`
		if [ ${count} != 0 ]
		then
			print_title "Starting test: ${filename}"
			run_test "${filename}"
		fi
	done
}

function main()
{
	get_time
	local start_time="${last_time}"

	run_tests

	get_time
	local end_time="${last_time}"
	
	local delta_time=$((${end_time} - ${start_time}))
	
	if [ ${error} -ne 0 ]; then
		test_error
		error=1
	else
		test_success
	fi

	print_output "All tests finished in ${delta_time} ms."

	echo ""

	if [ ${error} -ne 0 ]; then
		exit 1
	fi
}

main
