D5 - Deutschodoro - Take 5
==========================

Excel formats
-------------

Using the excel format is only necessary when the the default Excel parsers are used, but it is the recommended and the only documented input schema. Spreadsheet header is not used by the system.

### General

Excel sheets are supposed to have 6 columns and should have a utf-8 encoding. File types can be .csv, .ods or .xlsx

 - *German:* German expression to be learned. Doesn't have to be unique, but there should be only one German word in each column
 - *English:* English meaning / description of the German word. Synonims can be separated by comma, different meanings should be separated by semicolons. Different meanings can have further explanations in paranthases. Since these should help giving context to the translation, synonims can not have their own explanations.
 - *Third:* Can be used for non-english translations, usually in the native language of the learner (Optional)
 - *Category:* While practically any category can be provided, some are used to mark special word types. Typical categories are `noun`, `verb`, `adj`, `exp`, `idiom`, `prep`, `adv`, `init`, `prefix`, `pron`, `conj`
 - *Date:* Date is used to mark the date a word was learned
 - *Score:* 1-10 integer that indicates the importance of the word. 10 should be used for the most useful words and 1 for least important ones.


| German      | English               | Third               | Category | Date       | Score  |
|-------------|-----------------------|---------------------|----------|------------|--------|
| passt schon | just fits; never mind | pont passzol; hagyd | exp      | 2014-05-01 | 5      |

### Verbs



### Nouns


### Adjectives


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

