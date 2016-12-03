#! /bin/bash

cd "$(dirname "$0")"

error=0

source util.sh

function run_test()
{
	filename="$1"
	line=$(go test "github.com/peteraba/d5/lib/${filename}" -cover)

	if [ ${line:0:2} == "ok" ]; then
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
	local filenames=''

	filenames=$(find ../lib -type d | sort -r -nk1 | cut -f2-)

	for filename in ${filenames}
	do
		go_files=`ls -1 ${filename}/*.go 2> /dev/null | wc -l | tr -d '[:space:]'`
		if [ ${go_files} == 0 ]
		then
			continue
		fi

		test_files=`ls -1 ${filename}/*_test.go 2> /dev/null | wc -l | tr -d '[:space:]'`
    source_files=`expr ${go_files} - ${test_files}`
		print_title "Checking: ${filename} (${test_files} tests for ${source_files} source files)"

		if [ ${test_files} != 0 ]
		then
			run_test "${filename}"
		else
			echo
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
