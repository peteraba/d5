#! /bin/bash

cd "$(dirname "$0")"


rm -f output/*

export D5_DB_HOST="mongo.local"
export D5_DB_NAME="d5_test"

export D5_GAME_TYPE="german"
export D5_COLLECTION_DATA_GENERAL="general"
export D5_COLLECTION_DATA_GERMAN="german"
export D5_COLLECTION_RESULT="result"

user_name="john_doe"

error=0

solche_id=""
game_id=""

source util.sh

parser_port=10110
persister_port=10120
finder_port=10210
router_port=10220
scorer_port=10230
admin_port=10310
derdiedas_port=10410
conjugate_port=10420

parser_host='parser.local'
persister_host='persister.local'
finder_host='finder.local'
router_host='router.local'
scorer_host='scorer.local'
admin_host='admin.local'
derdiedas_host='derdiedas.local'
conjugate_host='conjugate.local'

function kill_app_listening_on_port()
{
	ps -ef | grep "$1" | egrep -v grep | awk '{print $2}' | xargs kill -9
}

function test_convert_ods_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.ods ]; then
		touch output/ods.json
		../spreadsheet/ods ../spreadsheet/fixture/gerdict.ods 8 | python -m json.tool > output/ods.json
	fi
}

function test_convert_xlsx_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.xlsx ]; then
		touch output/ods.json
		../spreadsheet/xlsx ../spreadsheet/fixture/gerdict.xlsx 8 | python -m json.tool > output/xlsx.json
	fi
}

function test_convert_csv_to_json()
{
	if [ -f ../spreadsheet/fixture/gerdict.csv ]; then
		touch output/ods.json
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

		if [ "${csv_xlsx_diff}" != "" ]; then
			test_error
			print_error ".xlsx and .csv files are different"
			print_error "${csv_xlsx_diff}"
			error=1
		else
			test_success
			print_output ".xlsx and .csv files are the same"
		fi
		if [ "${csv_ods_diff}" != "" ]; then
			test_error
			print_error ".ods and .csv files are different"
			print_error "${csv_ods_diff}"
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
		cat output/csv.json | parser "--user=${user_name}" -d > output/parsed.json
	else
		test_error
		print_error "parser is missing"
		error=1
	fi
}

function test_insert_into_db()
{
	if [ -f ../persister/persister.go ]; then
		cat output/parsed.json | persister -d
	else
		test_error
		print_error "persister is missing"
		error=1
	fi
}

function test_find_annehmbar()
{
	local result=""
	local search_expression="limit=2&query={\"word.german\": \"annehmbar\",\"word.user\": \"${user_name}\"}"

	result=$(echo "${search_expression}" | finder -d)
	
	if [[ "${result}" == *"acceptable"* ]]; then
		test_success
		print_output "Word 'annehmbar' and its translation were found."
	else
		test_error
		print_error "Word 'annehmbar' was not found or translation 'acceptable' was missing"
		print_error "Result: ${result}"
		error=1
	fi
}

function test_find_aufbauen()
{
	local result=""
	local search_expression="limit=2&query={\"word.german\": \"aufbauen\",\"word.user\": \"${user_name}\"}"

	result=$(echo "${search_expression}" | finder -d)

	if [[ "${result}" == *"build"* ]]; then
		test_success
		print_output "Word 'aufbauen' and its translation were found."
	else
		test_error
		print_error "Word 'aufbauen' was not found or translation 'build' was missing"
		print_error "Result: ${result}"
		error=1
	fi
}

function test_find_solche_via_server()
{
	local result=""

	finder --server "--port=${finder_port}" -d &

	result=$(curl --data "limit=2&query={\"word.german\":\"solche\",\"word.user\":\"${user_name}\"}" "http://${finder_host}:${finder_port}/" 2>&1 )

	kill_app_listening_on_port "${finder_port}"
	
	if [[ "${result}" == *"such"* ]]; then
		solche_id=$(echo "${result}" | grep -o "[0-9a-f\-]\{24\}")

		test_success
		print_output "Word 'solche' and its translation were found."
	else
		test_error
		print_error "Word 'solche' was not found or translation 'such' was missing"
		print_error "Result: ${result}"
		error=1
	fi
}

function test_score_solche()
{
	local result=""

	if [ "${solche_id}" != "" ]; then
		$(echo "wordId=${solche_id}&score=6" | scorer -d)

		local search_expression="limit=2&query={\"word.german\": \"solche\",\"word.user\": \"${user_name}\"}"

		result=$(echo "${search_expression}" | finder -d)

		if [[ "${result}" == *"\"result\": 6,"* ]]; then
			test_success
			print_output "Score 6 was found."
		else
			test_error
			print_error "Score 6 was not found."
			print_error "Result: ${result}"
			error=1
		fi
	else
		test_error
		print_error "Id for word solche is empty."
		error=1
	fi
}

function test_score_solche_via_server()
{
	local result=""

	if [ "${solche_id}" != "" ]; then
		(scorer --server "--port=${scorer_port}" -d & )

		result=$(curl --data "wordId=${solche_id}&score=7" "http://${scorer_host}:${scorer_port}/" 2>&1 )

		if [[ "${result}" == *"true"* ]]; then
			local search_expression="limit=2&query={\"word.german\": \"solche\",\"word.user\": \"${user_name}\"}"

			result=$(echo "${search_expression}" | finder -d)

			if [[ "${result}" == *"\"result\": 7,"* ]]; then
				test_success
				print_output "Score 7 was found."
			else
				test_error
				print_error "Score 7 was not found."
				print_error "Result: ${result}"
				error=1
			fi
		else
			test_error
			print_error "Setting the score failed."
			print_error "Result: ${result}"
			error=1
		fi

		kill_app_listening_on_port "${scorer_port}"

	else
		test_error
		print_error "Id for word solche is empty."
		error=1
	fi
}

function test_play_derdiedas()
{
	local result=''
	local word_id=''
	local german=''
	local result1=''
	local result2=''
	local result3=''
	local search_expression=''

	finder --server "--port=${finder_port}" -d &
	scorer --server "--port=${scorer_port}" -d &
	derdiedas -d "--port=${derdiedas_port}" "--finder=http://${finder_host}:${finder_port}/" "--scorer=http://${scorer_host}:${scorer_port}/" > /dev/null 2>&1 &

	sleep 0.1
	result=$(curl "http://${derdiedas_host}:${derdiedas_port}/game/${user_name}" 2>&1 )

	word_id=$(echo "${result}" | grep -o "[0-9a-f\-]\{24\}")
		
	if [[ "${result}" == *"question"* ]]; then
		result1=$(curl --data "id=${word_id}&answer=1" "http://${derdiedas_host}:${derdiedas_port}/answer/${user_name}" 2>&1 )
		result2=$(curl --data "id=${word_id}&answer=2" "http://${derdiedas_host}:${derdiedas_port}/answer/${user_name}" 2>&1 )
		result3=$(curl --data "id=${word_id}&answer=3" "http://${derdiedas_host}:1${derdiedas_port}/answer/${user_name}" 2>&1 )

		search_expression="limit=2&query={\"__id\": \"${word_id}\",\"word.user\": \"${user_name}\"}"

		result=$(echo "${search_expression}" | finder -d)
		german=$(echo "${result}" | grep -o "\"german\":\"[a-zA-ZäÄöÖüÜß -]*\"")
		german="${german:10:-1}"

		if [[ "${result}" == *"\"result\": 10,"* ]]; then
			test_success
			print_output "Score 10 was found."
			print_output "Word: ${german}, Id: ${word_id}"
		else
			test_error
			print_error "Score 10 was not found."
			print_error "Result: ${result}"
			error=1
		fi
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: ${result}"
		error=1
	fi

	kill_app_listening_on_port "${finder_port}"
	kill_app_listening_on_port "${scorer_port}"
	kill_app_listening_on_port "${derdiedas_port}"
}

function test_play_conjugate()
{
	local result=""
	local result_id=""
	local word_id=""
	local german=""
	local mongo=""
	local result1=""
	local search_expression=''

	finder --server "--port=${finder_port}" -d &
	scorer --server "--port=${scorer_port}" -d &
	conjugate -d --server "--port=${conjugate_port}" "--finder=http://${finder_host}:${finder_port}/" "--scorer=http://${scorer_host}:${scorer_port}/" > /dev/null 2>&1 &

	sleep 0.1
	result=$(curl "http://${conjugate_host}:${conjugate_port}/game/${user_name}" 2>&1 )

	if [[ "${result}" == *"question"* ]]; then
		result_id=$(echo "${result}" | grep -o "[0-9a-f\-]\{36\}")
		
		mongo=$(mongo ${D5_DB_NAME} --eval "db.${D5_COLLECTION_DATA_GERMAN}.find({\"_id\":\"${result}_id\"}).shellPrint()")

		word_id=$(echo "${mongo}" | grep -o "[0-9a-f]\{24\}")

		result=$(echo "${mongo}" | grep -o "\"right\" \: \[ \"[a-zA-ZäÄöÖüÜß \-]\{4,\}\" \]")
		result=${result:13:-3}

		result1=$(curl --data "id=${result}_id&answer=${result}" "http://${conjugate_host}:${conjugate_port}/answer/${user_name}" 2>&1 )

		search_expression="limit=2&query={\"__id\": \"${word_id}\",\"word.user\": \"${user_name}\"}"

		result=$(echo "${search_expression}" | finder -d)
		german=$(echo "${result}" | grep -o "\"german\":\"[a-zA-ZäÄöÖüÜß -]*\"")
		german=${german:10:-1}

		if [[ "${result}" == *"\"result\": 10,"* ]]; then
			test_success
			print_output "Score 10 was found."
			print_output "Word: ${german}, Id: ${word_id}"
		else
			test_error
			print_error "Score 10 was not found."
			print_error "Result: ${result}"
			error=1
		fi
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: ${result}"
		error=1
	fi

	kill_app_listening_on_port "${finder_port}"
	kill_app_listening_on_port "${scorer_port}"
	kill_app_listening_on_port "${conjugate_port}"
}

function test_create_game()
{
	admin "--port=${admin_port}" -d &

	result=$(curl --data "name=Der%20die%20das&route=derdiedas&url=http://${derdiedas_host}:${derdiedas_port}/&is-system=0" "http://${admin_host}:${admin_port}/game" 2>&1 )

	if [[ "${result}" == *"OK"* ]]; then
		print_output "Admin responded with OK."
		
		game_id=""

		result=$(curl "http://${admin_host}:${admin_port}/game/${game_id}" 2>&1 )

		if [[ "${result}" == *"OK"* ]]; then
			test_success
			print_output "Admin responded with OK."
		else
			test_error
			print_error "Initialising a game failed."
			print_error "Result: ${result}"
			error=1
		fi
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: ${result}"
		error=1
	fi
}

function test_update_game()
{
	admin "--port=${admin_port}" -d &

	result=$(curl --data "name=Der%20die%20das&route=derdiedas&url=http://${derdiedas_host}:${derdiedas_port}/&is-system=0" -X PATCH "http://${admin_host}:${admin_port}/game/${game_id}" 2>&1 )

	if [[ "${result}" == *"OK"* ]]; then
		test_success
		print_output "Admin responded with OK."
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: ${result}"
		error=1
	fi

	kill_app_listening_on_port "${admin_port}"
	# kill_app_listening_on_port "${derdiedas_port}"
}

function test_delete_game()
{
	admin "--port=${admin_port}" -d &

	result=$(curl -X "DELETE" "http://${admin_host}:${admin_port}/game/${game_id}" 2>&1 )

	if [[ "${result}" == *"OK"* ]]; then
		test_success
		print_output "Admin responded with OK."
	else
		test_error
		print_error "Initialising a game failed."
		print_error "Result: ${result}"
		error=1
	fi

	kill_app_listening_on_port "${admin_port}"
	# kill_app_listening_on_port "${derdiedas_port}"
}

function test_create_user()
{
	admin "--port=${admin_port}" -d &

	kill_app_listening_on_port "${admin_port}"
}

function test_update_user()
{
	admin "--port=${admin_port}" -d &

	kill_app_listening_on_port "${admin_port}"
}

function test_delete_user()
{
	admin "--port=${admin_port}" -d &

	kill_app_listening_on_port "${admin_port}"
}

function run_task()
{
	local task="$(echo $1 | tr "[:upper:]" "[:lower:]" | sed 's/ /_/g')"
	local max_time=$2

	get_time
	local start_time="${last_time}"

	print_title "Starting test: $1"
	test_"${task}"
	
	get_time
	local end_time="${last_time}"

	local delta_time=$((${end_time} - ${start_time}))

	if [ "${delta_time}" -gt "${max_time}" ]; then
		test_error
		print_error "Finished in ${delta_time} ms. (Max: ${max_time} ms)\n"
		error=1
	else
		test_success
		print_output "Finished in ${delta_time} ms.\n"
	fi
	test_end
}

function run_tests()
{
	run_task "convert ods to json" 2000
	run_task "convert csv to json" 400
	run_task "convert xlsx to json" 2000
	run_task "check json sizes" 200
	run_task "parse json" 500
	run_task "insert into db" 2000
	run_task "find annehmbar" 200
	run_task "find aufbauen" 200
	run_task "find solche via server" 200
	run_task "score solche" 200
	run_task "score solche via server" 200
	run_task "play derdiedas" 500
	run_task "play conjugate" 1000
	# run_task "create_game" 300
	# run_task "update_game" 300
	# run_task "delete_game" 300
	# run_task "create_user" 300
	# run_task "update_user" 300
	# run_task "delete_user" 300
}

function main()
{
	../build.sh

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
