D5 - Deutschodoro (Take 5)
==========================

Excel formats
-------------

Using the excel format is only necessary when the the default Excel parsers are used, but it is the recommended and the only documented input schema. Spreadsheet header is not used by the system.

### General

Excel sheets are supposed to have 6 columns and should have a utf-8 encoding. File types can be .csv, .ods or .xlsx

 - **German:** German expression to be learned. Doesn't have to be unique, but there should be only one German word in each column
 - **English:** English meaning / description of the German word. Synonims can be separated by comma, different meanings should be separated by semicolons. Different meanings can have further explanations in paranthases. Since these should help giving context to the translation, synonims can not have their own explanations.
 - **Third:** Can be used for non-english translations, usually in the native language of the learner (Optional)
 - **Category:** While practically any category can be provided, some are used to mark special word types. Typical categories are `noun`, `verb`, `adj`, `exp`, `idiom`, `prep`, `adv`, `init`, `prefix`, `pron`, `conj`
 - **Date:** Date is used to mark the date a word was learned
 - **Score:** 1-10 integer that indicates the importance of the word. 10 should be used for the most useful words and 1 for least important ones.


| German      | English                | Third               | Category | Date       | Score  |
|-------------|------------------------|---------------------|----------|------------|--------|
| passt schon | no problem; never mind | pont passzol; hagyd | exp      | 2014-05-01 | 5      |

### Verbs

Verb is the most complex type in d5. Has many rules:

 * There are 5 types based on the regularity of verbs.
    *# Regular verbs
    *# Verbs with irregular past tense
    *# Verbs with irregular conjugation in second/third person in present tense
    *# Very irregular verbs
 * Verbs must have an auxiliary, haben (h) and/or sein (s) as the first word.
 * Verbs may contain a sign of prefix separation by using a pipe in the first conjugation.
 * Any conjugation but the main dictionary form can have multiple versions separated by / characters.
 * Verbs may have any number of arguments, each might have in parantheses a notion of case in which the following part must be (**A:** acustative, **D:** dative, **G:** genitive).
 * The first argument may be an indication that the verb requires a sich for which cases can only be acusative or dative.

#### Regular verbs

| German                            | English                    | Third                                    | Category | Date       | Score  |
|-----------------------------------|----------------------------|------------------------------------------|----------|------------|--------|
| h ausprobieren                    | to try out                 | kipróbálni                               | verb     | 2015-03-29 | 5      |
| h/s durch|drehen                  | to freak out               | megőrülni; meghülyülni                   | verb     | 2014-06-04 | 5      |
| h diskutieren + über (A)          | to discuss sth             | megvitatni vmit; megbeszélni vmit        | verb     | 2014-05-08 | 5      |
| h beeilen + sich (A)              | to hurry                   | sietni                                   | verb     | 2014-08-20 | 5      |
| h Sorgen machen + sich (A) + über | to worry about sth.        | aggódni vmi miatt                        | verb     | 2014-05-10 | 5      |
| h verlassen + sich (A) + auf (A)  | to trust, to rely on       | elhagyni                                 | verb     | 2014-05-05 | 5      |
| h vor|machen + (D)                | to fool, to deceive (coll) | megtéveszteni vkit                       | verb     | 2014-08-31 | 5      |
| h vor|machen + sich (D)           | to lie to oneself          | beképzelni magának vmit, hazudni magának | verb     | 2014-08-31 | 5      |

#### Verbs with irregular past tense

| German                                                   | English  | Third      | Category | Date       | Score  |
|----------------------------------------------------------|----------|------------|----------|------------|--------|
| h schreiben, schrieben, geschrieben                      | to write | írni       | verb     | 2014-10-12 | 5      |
| h/s sport treiben, sport trieben, sport getrieben        | to sport | sportolni  | verb     | 2014-05-01 | 5      |
| h treffen, trafen, getroffen, triffst, trifft + sich (A) | to meet  | találkozni | verb     | 2014-10-08 | 5      |

#### Verbs with irregular second/third person in present

| German                                                  | English  | Third     | Category | Date       | Score  |
|---------------------------------------------------------|----------|-----------|----------|------------|--------|
| h verlassen, verließ, verlassen, verlässt, verlässt     | to leave | elhagyni  | verb     | 2014-05-05 | 5      |
| h verwenden, verwendeten/verwandten, verwendet/verwandt | to use   | használni | verb     | 2014-05-01 | 5      |

#### Very irregular verbs

| German                                             | English | Third           | Category | Date       | Score  |
|----------------------------------------------------|---------|-----------------|----------|------------|--------|
| h tun, tue, tust, tut, tun, tut, tun, taten, getan | to do   | tenni, csinálni | verb     | 2014-05-01 | 5      |


### Nouns

| German                                   | English                      | Third                     | Category   | Date       | Score  |
|------------------------------------------|------------------------------|---------------------------|------------|------------|--------|
| e Entzündung,~en                         | inflammation                 | gyulladás                 | noun       | 2015-03-04 | 5      |
| e Trauma,Traumata/Traumen                | trauma                       | trauma                    | noun       | 2015-04-18 | 5      |
| e Vereinigten Staaten von Amerika,- (pl) | The United States of America | Amerikai Egyesült Államok | noun       | 2014-10-16 | 5      |


### Adjectives

| German                     | English       | Third           | Category | Date       | Score  |
|----------------------------|---------------|-----------------|----------|------------|--------|
| klug,⍨er,⍨sten             | smart         | okos, értelmes  | adj      | 2014-05-01 | 5      |
| hochschwanger,-            | very pregnant | terhes (nagyon) | adj      | 2014-05-01 | 5      |
| schmal,~er/⍨er,~sten/⍨sten | narrow        | keskeny, szűk   | adj      | 2014-05-01 | 5      |



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

