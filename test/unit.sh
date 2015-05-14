#! /bin/bash

go test github.com/peteraba/d5/lib/german/entity -cover
echo ""

go test github.com/peteraba/d5/lib/german/dict -cover
echo ""

go test github.com/peteraba/d5/lib/german/util -cover
echo ""

go test github.com/peteraba/d5/lib/util -cover
echo ""
