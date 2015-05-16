#!/bin/bash

cd "$(dirname "$0")"

echo "Unit tests"
../../test/unit.sh

echo ""
echo "Integration tests"
../../test/integration.sh

