#! /bin/bash

cd "$(dirname "$0")"

# Generate raw json from .ods
d1=`date +%s%N | cut -b1-13`
if [ -f ../spreadsheet/fixture/gerdict.ods ]
then
	echo "ods to json"
	../spreadsheet/ods ../spreadsheet/fixture/gerdict.ods 8 > output/ods.json
fi

# Generate raw json from .xlsx
d2=`date +%s%N | cut -b1-13`
if [ -f ../spreadsheet/fixture/gerdict.xlsx ]
then
	echo "xlsx to json"
	../spreadsheet/xlsx ../spreadsheet/fixture/gerdict.xlsx 8 > output/xlsx.json
fi

# Generate raw json from .csv
d3=`date +%s%N | cut -b1-13`
if [ -f ../spreadsheet/fixture/gerdict.csv ]
then
	echo "csv to json"
	../spreadsheet/csv ../spreadsheet/fixture/gerdict.csv 8 > output/csv.json
fi

# Parse raw json and output parsed json
d4=`date +%s%N | cut -b1-13`
if [ -f ../parser/parser.go ]
then
	echo "Parsing raw json"
	cat output/csv.json | go run ../parser/parser.go -user=peteraba > output/parsed.json
fi

# Save parsed json into mongodb collection
d5=`date +%s%N | cut -b1-13`
if [ -f ../persister/persister.go ]
then
	echo "Parsed json saved into db collection"
	cat output/parsed.json | go run ../persister/persister.go -host=localhost -db=test -coll=words
fi

d6=`date +%s%N | cut -b1-13`

d1t6=$(($d6 - $d1))

echo "Tests finished in $d1t6 microseconds."

