#!/bin/bash

cd "$(dirname "$0")"

../../test/unit.sh

../../test/integration.sh

