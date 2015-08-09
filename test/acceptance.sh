#! /bin/bash

cd "$(dirname "$0")"

rm output/*

export D5_DBHOST="localhost"
export D5_DBNAME="d5_test"

german_test_collection="german_test"

error=0

solcheId=""

source util.sh

function test_convert_ods_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.ods ]; then
		../spreadsheet/ods ../spreadsheet/fixture/gerdict.ods 8 | python -m json.tool > output/ods.json
	fi
}

function test_convert_xlsx_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.xlsx ]; then
		../spreadsheet/xlsx ../spreadsheet/fixture/gerdict.xlsx 8 | python -m json.tool > output/xlsx.json
	fi
}

function test_convert_csv_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.csv ]; then
		../spreadsheet/csv ../spreadsheet/fixture/gerdict.csv 8 | python -m json.tool > output/csv.json
	fi
}

function test_check_json_sizes()
{
	local csv_ods_diff=""
	local csv_xlsx_diff=""

	if [ -f output/csv.json ]; then
		csv_xlsx_diff=$(diff output/csv.json output/xlsx.json)
		csv_ods_diff=$(diff output/csv.json output/ods.json)

		if [ "$csv_xlsx_diff" != "" ]; then
			test_error
			print_error ".xlsx and .csv files are different"
			print_error "$csv_xlsx_diff"
			error=1
		else
			test_success
			print_output ".xlsx and .csv files are the same"
		fi
		if [ "$csv_ods_diff" != "" ]; then
			test_error
			print_error ".ods and .csv files are different"
			print_error "$csv_ods_diff"
			error=1
		else
			test_success
			print_output ".ods and .csv files are the same"
		fi
	else
		test_error
		print_error "output/csv.json is missing"
		error=1
	fi
}

function test_parse_json()
{
	if [ -f ../parser/parser.go ]; then
		cat output/csv.json | parser --user=peteraba > output/parsed.json
	else
		test_error
		print_error "parser is missing"
		error=1
	fi
}

function test_insert_into_db()
{
	if [ -f ../persister/persister.go ]; then
		cat output/parsed.json | persister --coll=$german_test_collection
	else
		test_error
		print_error "persister is missing"
		error=1
	fi
}

function test_find_annehmbar()
{
	local result=""
	local search_expression="{\"word.german\": \"annehmbar\",\"word.user\": \"peteraba\"}"

	result=$(echo $search_expression | finder --coll=$german_test_collection )
	
	if [[ "$result" == *"acceptable"* ]]; then
		test_success
		print_output "Word 'annehmbar' and its translation were found."
	else
		test_error
		print_error "Word 'annehmbar' was not found or translation 'acceptable' was missing"
		print_error "Result: $result"
		error=1
	fi
}

function test_find_aufbauen()
{
	local result=""
	local search_expression="{\"word.german\": \"aufbauen\",\"word.user\": \"peteraba\"}"

	result=$(echo $search_expression | finder --coll=$german_test_collection )

	if [[ "$result" == *"build"* ]]; then
		test_success
		print_output "Word 'aufbauen' and its translation were found."
	else
		test_error
		print_error "Word 'aufbauen' was not found or translation 'build' was missing"
		print_error "Result: $result"
		error=1
	fi
}

function test_find_solche_via_server()
{
	local result=""

	(finder --coll=$german_test_collection --server=true --port=11111 & )

	result=$(curl --data 'query={"word.german":"solche","word.user":"peteraba"}' http://localhost:11111/ 2>&1 )

	killall finder
	
	if [[ "$result" == *"such"* ]]; then
		solcheId=$(echo "$result" | grep -o "[0-9a-f\-]\{12,\}")

		test_success
		print_output "Word 'solche' and its translation were found."
	else
		test_error
		print_error "Word 'solche' was not found or translation 'such' was missing"
		print_error "Result: $result"
		error=1
	fi
}

function test_score_solche()
{
	local result=""

	if [ "$solcheId" != "" ]; then
		$(scorer --coll=$german_test_collection --wordId=$solcheId --score=6 )

		local search_expression="{\"word.german\": \"solche\",\"word.user\": \"peteraba\"}"

		result=$(echo $search_expression | finder --coll=$german_test_collection )

		if [[ "$result" == *"\"result\":6,"* ]]; then
			test_success
			print_output "Score 6 was found."
		else
			test_error
			print_error "Score 6 was not found."
			print_error "Result: $result"
			error=1
		fi
	else
		test_err
		print_error "Id for word solche is empty."
		error=1
	fi
}

function test_score_solche_via_server()
{
	local result=""


	if [ "$solcheId" != "" ]; then
	        (scorer --coll=$german_test_collection --server=true --port=11112 & )

	        result=$(curl --data "wordId=$solcheId&score=7" http://localhost:11112/ 2>&1 )

		if [[ "$result" == *"true"* ]]; then
			local search_expression="{\"word.german\": \"solche\",\"word.user\": \"peteraba\"}"

			result=$(echo $search_expression | finder --coll=$german_test_collection )

			if [[ "$result" == *"\"result\":7,"* ]]; then
				test_success
				print_output "Score 7 was found."
			else
				test_error
				print_error "Score 7 was not found."
				print_error "Result: $result"
				error=1
			fi
		else
			test_error
			print_error "Setting the score failed."
			print_error "Result: $result"
			error=1
		fi

		killall scorer

	else
		test_err
		print_error "Id for word solche is empty."
		error=1
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
	run_task "convert csv to json" 400
	run_task "convert xlsx to json" 2000
	run_task "check json sizes" 50
	run_task "parse json" 500
	run_task "insert into db" 2000
	run_task "find annehmbar" 200
	run_task "find aufbauen" 200
	run_task "find solche via server" 200
	run_task "score solche" 200
	run_task "score solche via server" 200
}

function main()
{
	../build.sh

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

	echo ""

	if [ $error -ne 0 ]; then
		exit 1
	fi
}

main
