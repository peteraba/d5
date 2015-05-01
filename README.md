D5 - Deutschodoro - Take 5
==========================

Utils
-----

### Convert Excel to Plain JSON

```bash
./spreadsheet/ods spreadsheet/fixture/gerdict.ods
./spreadsheet/xlsx spreadsheet/fixture/gerdict.xlsx
./spreadsheet/csv spreadsheet/fixture/gerdict.csv
```

### Convert Plain JSON to Parsed JSON

```bash
cat parser/test.json | go run parser/parser.go peteraba false > persister/test.json
cat parser/test.json | parser peteraba false > persister/test.json
```

```bash
cat parser/test.json | go run parser/parser.go peteraba true
cat parser/test.json | parser peteraba true
```

### Persist Parsed JSON

```bash
cat persister/test.json | go run persister.go localhost test words
cat persister/test.json | persister localhost test words
```

