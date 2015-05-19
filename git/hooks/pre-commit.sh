#!/bin/bash

cd "$(dirname "$0")"

source ../../test/util.sh


print_title  "Unit tests"
print_output "=========="
echo ""
../../test/unit.sh
if [[ $? -ne 0 ]]; then
	test_error
	print_error "Unit tests failed."
	exit 1
fi

