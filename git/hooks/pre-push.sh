#!/bin/bash

cd "$(dirname "$0")"

source ../../test/util.sh


print_title  "Integration tests"
print_output "================="
echo ""
../../test/integration.sh
if [[ $? -ne 0 ]]; then
	test_error
	print_error "Integration tests failed."
	exit 1
fi
