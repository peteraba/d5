#!/bin/bash

cd "$(dirname "$0")"

source ../../test/util.sh


print_title  "Acceptance tests"
print_output "================"
echo ""
../../test/acceptance.sh
if [[ $? -ne 0 ]]; then
	test_error
	print_error "Acceptance tests failed."
	exit 1
fi
