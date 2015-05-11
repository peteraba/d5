#! /bin/bash

cd "$(dirname "$0")"

rm output/*

error=0

# Output start message of a check
#
# @param string $1 check name
function test_start()
{
	printf "\n\e[1;33m%-40s" "• Running $1"
}

# Output a nice [✔] with colors for a successfully check
function test_success()
{
	echo -en "  \e[1;33m[\e[1;32m✔\e[1;33m]\e[0m"
}

# Output a nice [✘] with colors for a unsuccessfully check
function test_error()
{
	echo -en "  \e[1;33m[\e[1;31m✘\e[1;33m]\e[0m"
}

function test_end()
{
	echo -en "\e[0m"
}

# @param string $1 output of the command
function print_output()
{
	printf "%s\n" "$(echo -e "$1" | awk '{printf "\t%s\n", $0}')"
}

# @param string $1 output of the command
function print_error()
{
	printf "%s\n" "$(echo -e "$1" | awk '{printf "\t%s\n", $0}')"
}

# @param string $1 output of the command
function print_title()
{
	printf "\033[1m%s\033[0m\n" "$(echo -e "$1" | awk '{printf "\t%s\n", $0}')"
}

function test_convert_ods_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.ods ]; then
		../spreadsheet/ods ../spreadsheet/fixture/gerdict.ods 8 > output/ods.json
	fi
}

function test_convert_xlsx_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.xlsx ]; then
		../spreadsheet/xlsx ../spreadsheet/fixture/gerdict.xlsx 8 > output/xlsx.json
	fi
}

function test_convert_csv_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.csv ]; then
		../spreadsheet/csv ../spreadsheet/fixture/gerdict.csv 8 > output/csv.json
	fi
}

function test_check_json_sizes()
{
	if [ -f output/csv.json ]; then
		csv_size=$(wc -c output/csv.json | cut -f 1 -d ' ')
		xlsx_size=$(wc -c output/xlsx.json | cut -f 1 -d ' ')
		ods_size=$(wc -c output/ods.json | cut -f 1 -d ' ')

		if [ "$csv_size" -ne "$xlsx_size" ]; then
			test_error
			print_error ".xlsx and .csv files are different"
			error=1
		else
			test_success
			print_output ".xlsx and .csv files are the same in size"
		fi
		if [ "$csv_size" -ne "$ods_size" ]; then
			test_error
			print_error ".ods and .csv files are different"
			error=1
		else
			test_success
			print_output ".ods and .csv files are the same in size"
		fi
	fi
}

function test_parse_json()
{
	if [ -f ../parser/parser.go ]; then
		cat output/csv.json | go run ../parser/parser.go -user=peteraba > output/parsed.json
	fi
}

function test_insert_into_db()
{
	if [ -f ../persister/persister.go ]; then
		cat output/parsed.json | go run ../persister/persister.go -host=localhost -db=test -coll=words
	fi

}

function run_task()
{
	local task="$(echo $1 | tr "[:upper:]" "[:lower:]" | sed 's/ /_/g')"
	local max_time=$2

	local start_time=`date +%s%N | cut -b1-13`

	print_title "Starting test: $1"
	test_"$task"
	
	local end_time=`date +%s%N | cut -b1-13`
	local delta_time=$(($end_time - $start_time))

	if [ "$delta_time" -gt "$max_time" ]; then
		test_error
		print_error "Finished in $delta_time ms. (Max: $max_time ms)\n"
		error=1
	else
		test_success
		print_output "Finished in $delta_time ms.\n"
	fi
	test_end
}

function run_tests()
{
	run_task "convert ods to json" 2000
	run_task "convert csv to json" 100
	run_task "convert xlsx to json" 2000
	run_task "check json sizes" 50
	run_task "parse json" 500
	run_task "insert into db" 2000
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
