#! /bin/bash

cd "$(dirname "$0")"

rm output/*

export D5_DBHOST="localhost"
export D5_DBNAME="d5_test"

export GAME_DBHOST="localhost"
export GAME_DBNAME="d5_test"

game_dbname="d5_test"
german_test_collection="german_test"
result_test_collection="result_test"

error=0

solche_id=""
game_id=""

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
	local search_expression="limit=2&query={\"word.german\": \"annehmbar\",\"word.user\": \"peteraba\"}"

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
	local search_expression="limit=2&query={\"word.german\": \"aufbauen\",\"word.user\": \"peteraba\"}"

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

	result=$(curl --data 'limit=2&query={"word.german":"solche","word.user":"peteraba"}' http://localhost:11111/ 2>&1 )

	(fuser -k 11111/tcp > /dev/null 2>&1 & )
	
	if [[ "$result" == *"such"* ]]; then
		solche_id=$(echo "$result" | grep -o "[0-9a-f\-]\{24\}")

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

	if [ "$solche_id" != "" ]; then
		$(scorer --coll=$german_test_collection --wordId=$solche_id --score=6 )

		local search_expression="limit=2&query={\"word.german\": \"solche\",\"word.user\": \"peteraba\"}"

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

	if [ "$solche_id" != "" ]; then
		(scorer --coll=$german_test_collection --server=true --port=11112 & )

		result=$(curl --data "wordId=$solche_id&score=7" http://localhost:11112/ 2>&1 )

		if [[ "$result" == *"true"* ]]; then
			local search_expression="limit=2&query={\"word.german\": \"solche\",\"word.user\": \"peteraba\"}"

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

		(fuser -k 11112/tcp > /dev/null 2>&1 & )

	else
		test_err
		print_error "Id for word solche is empty."
		error=1
	fi
}

function test_play_derdiedas()
{
	local result=""
	local word_id=""
	local german=""
	local result1=""
	local result2=""
	local result3=""
	local search_expression

	(finder --coll=$german_test_collection --server=true --port=11121 & )
	(scorer --coll=$german_test_collection --server=true --port=11122 & )
	(derdiedas --debug=false --port=11123 --finder=http://localhost:11121/ --scorer=http://localhost:11122/ > /dev/null 2>&1 & )

	sleep 0.1
	result=$(curl http://localhost:11123/game/peteraba 2>&1 )

	word_id=$(echo "$result" | grep -o "[0-9a-f\-]\{24\}")
		
	if [[ "$result" == *"question"* ]]; then
		result1=$(curl --data "id=$word_id&answer=1" http://localhost:11123/answer/peteraba 2>&1 )
		result2=$(curl --data "id=$word_id&answer=2" http://localhost:11123/answer/peteraba 2>&1 )
		result3=$(curl --data "id=$word_id&answer=3" http://localhost:11123/answer/peteraba 2>&1 )

		search_expression="limit=2&query={\"__id\": \"$word_id\",\"word.user\": \"peteraba\"}"

		result=$(echo $search_expression | finder --coll=$german_test_collection )
		german=$(echo "$result" | grep -o "\"german\":\"[a-zA-ZäÄöÖüÜß -]*\"")
		german=${german:10:-1}

		if [[ "$result" == *"\"result\":10,"* ]]; then
			test_success
			print_output "Score 10 was found."
			print_output "Word: $german, Id: $word_id"
		else
			test_error
			print_error "Score 10 was not found."
			print_error "Result: $result"
			error=1
		fi
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: $result"
		error=1
	fi

	(fuser -k 11121/tcp > /dev/null 2>&1 & )
	(fuser -k 11122/tcp > /dev/null 2>&1 & )
	(fuser -k 11123/tcp > /dev/null 2>&1 & )
}

function test_play_conjugate()
{
	local result=""
	local result_id=""
	local word_id=""
	local german=""
	local mongo=""
	local result1=""
	local search_expression

	(finder --coll=$german_test_collection --server=true --port=11131 & )
	(scorer --coll=$german_test_collection --server=true --port=11132 & )
	(conjugate --debug=false --port=11133 --finder=http://localhost:11131/ --scorer=http://localhost:11132/ --coll=$result_test_collection > /dev/null 2>&1 & )

	sleep 0.1
	result=$(curl http://localhost:11133/game/peteraba 2>&1 )

	if [[ "$result" == *"question"* ]]; then
		result_id=$(echo "$result" | grep -o "[0-9a-f\-]\{36\}")
		
		mongo=$(mongo $game_dbname --eval "db.$result_test_collection.find({\"_id\":\"$result_id\"}).shellPrint()")
	
		word_id=$(echo "$mongo" | grep -o "[0-9a-f]\{24\}")

		result=$(echo "$mongo" | grep -o "\"right\" \: \[ \"[a-zA-ZäÄöÖüÜß \-]\{4,\}\" \]")
		result=${result:13:-3}

		result1=$(curl --data "id=$result_id&answer=$result" http://localhost:11133/answer/peteraba 2>&1 )

		search_expression="limit=2&query={\"__id\": \"$word_id\",\"word.user\": \"peteraba\"}"

		result=$(echo $search_expression | finder --coll=$german_test_collection )
		german=$(echo "$result" | grep -o "\"german\":\"[a-zA-ZäÄöÖüÜß -]*\"")
		german=${german:10:-1}

		if [[ "$result" == *"\"result\":10,"* ]]; then
			test_success
			print_output "Score 10 was found."
			print_output "Word: $german, Id: $word_id"
		else
			test_error
			print_error "Score 10 was not found."
			print_error "Result: $result"
			error=1
		fi
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: $result"
		error=1
	fi

	(fuser -k 11131/tcp > /dev/null 2>&1 & )
	(fuser -k 11132/tcp > /dev/null 2>&1 & )
	(fuser -k 11133/tcp > /dev/null 2>&1 & )
}

function test_create_game()
{
	(admin --port=11141 & )

	result=$(curl --data 'name=Der%20die%20das&route=derdiedas&url=http://localhost:12345/&is-system=0' http://localhost:11111/game 2>&1 )

	if [[ "$result" == *"OK"* ]]; then
		print_output "Admin responded with OK."
		
		game_id=""

		result=$(curl http://localhost:11111/game/$game_id 2>&1 )

		if [[ "$result" == *"OK"* ]]; then
			test_success
			print_output "Admin responded with OK."
		else
			test_error
			print_error "Initialising a game failed."
			print_error "Result: $result"
			error=1
		fi
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: $result"
		error=1
	fi

	(fuser -k 11141/tcp > /dev/null 2>&1 & )
}

function test_update_game()
{
	(admin --port=11141 & )

	result=$(curl --data "name=Der%20die%20das&route=derdiedas&url=http://localhost:12345/&is-system=0" -X "PATCH" http://localhost:11111/game/$game_id 2>&1 )

	if [[ "$result" == *"OK"* ]]; then
		test_success
		print_output "Admin responded with OK."
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: $result"
		error=1
	fi

	(fuser -k 11141/tcp > /dev/null 2>&1 & )
}

function test_delete_game()
{
	(admin --port=11141 & )

	result=$(curl -X "DELETE" http://localhost:11111/game/$game_id 2>&1 )

	if [[ "$result" == *"OK"* ]]; then
		test_success
		print_output "Admin responded with OK."
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: $result"
		error=1
	fi

	(fuser -k 11141/tcp > /dev/null 2>&1 & )
}

function test_create_user()
{
	(admin --port=11141 & )

	(fuser -k 11141/tcp > /dev/null 2>&1 & )
}

function test_update_user()
{
	(admin --port=11141 & )

	(fuser -k 11141/tcp > /dev/null 2>&1 & )
}

function test_delete_user()
{
	(admin --port=11141 & )

	(fuser -k 11141/tcp > /dev/null 2>&1 & )
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
	run_task "play derdiedas" 500
	run_task "play conjugate" 1000
	run_task "create_game" 300
	#run_task "update_game" 300
	#run_task "delete_game" 300
	#run_task "create_user" 300
	#run_task "update_user" 300
	#run_task "delete_user" 300
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
