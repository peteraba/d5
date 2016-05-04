D5 - Deutschodoro (Take 5)
==========================

Excel formats
-------------

Using the excel format is only necessary when the default Excel parsers are used, but it is the recommended and the only documented input schema. The first row of the spreadsheet is considered to be the header is not used by the system at this point.

German dictionary
-----------------

Deutschodoro will be able to handle simple, two-column dictionaries in medium term, but it's primarly focus is dealing with German dictionaries.

### General

German dictionaries are supposed to have 6 columns and they should be utf-8 encoded. File types can be .csv, .ods or .xlsx

The columns need to be in the right order and they need to contain the following information

 - **Article/Auxiliary:** Only used for nouns and verbs. Article notions for nouns, auxiliary for verbs.
 - **German:** German expression to be learned. Doesn't have to be unique but the whole content of the column will be interpreted as one word. It must *not* contain synonims.
 - **English:** English meaning(s) of the German word. Synonims are separated by comma, different meanings are separated by semicolons. Different meanings can have further explanations in paranthases. Since these should help giving context to the translation, synonims can not have their own explanations.
 - **Third:** Can be used for non-english translations, usually in the native language of the learner (Optional)
 - **Category:** While practically any category can be provided, some are used to mark special word types. Typical categories are `noun`, `verb`, `adj`, `exp`, `idiom`, `prep`, `adv`, `init`, `prefix`, `pron`, `conj`
 - **Date:** Date is used to mark the date a word was learned
 - **Score:** 1-10 integer that indicates the importance of the word. 10 should be used for the most useful words and 1 for least important ones.


| A/A | German      | English                | Third               | Category | Date       | Score  |
|-----|-------------|------------------------|---------------------|----------|------------|--------|
|     | passt schon | no problem; never mind | pont passzol; hagyd | exp      | 2014-05-01 | 5      |

### Verbs

Verb is the most complex type in d5. Verb content must follow the following rules:

 * There are 5 types based on the regularity of verbs.
    * Regular verbs
    * Verbs with irregular past tense
    * Verbs with irregular conjugation in second/third person in present tense
    * Very irregular verbs
 * Verbs must have at least one auxiliary verb, haben (`h`) and/or sein (`s`), as the first word, multiple auxiliary verbs must be separated by a `/` character (without spaces)
 * Separable verbs may contain a pipe sign (`|`) a sign of the place prefix separation using a pipe in the first conjugation. Note that the system is able to note most separable verbs, but in cases of prefixes that may or may not be separable, it will always assume that the verb is inseparable.
 * Any conjugation of a verb but the first one (dictionary form) can have multiple versions separated by / characters. This can indicate that multiple version of a form is in use.
 * Verbs may have any number of arguments, each should be separated from the rest of the definition by a `+` sign. Each argument may have in parantheses a notion of case in which the following part must be (**A:** acustative, **D:** dative, **G:** genitive). In some cases there will only be a notion of case.
 * Reflexive verbs must have as first argument `sich (A)` or `sich (D)` depending on the case in which the reflexive part is in.

#### Regular verbs

Regular verbs define only 1 word which is the plural, first person conjugation (us/wir).

| A/A | German                              | English                    | Third                                    | Category | Date       | Score  |
|-----|-------------------------------------|----------------------------|------------------------------------------|----------|------------|--------|
| h   | ausprobieren                        | to try out                 | kipróbálni                               | verb     | 2015-03-29 | 5      |
| h/s | durch|drehen                        | to freak out               | megőrülni; meghülyülni                   | verb     | 2014-06-04 | 5      |
| h   | diskutieren + über (A)              | to discuss sth             | megvitatni vmit; megbeszélni vmit        | verb     | 2014-05-08 | 5      |
| h   | beeilen + sich (A)                  | to hurry                   | sietni                                   | verb     | 2014-08-20 | 5      |
| h   | Sorgen machen + sich (A) + über (A) | to worry about sth.        | aggódni vmi miatt                        | verb     | 2014-05-10 | 5      |
| h   | verlassen + sich (A) + auf (A)      | to trust, to rely on       | elhagyni                                 | verb     | 2014-05-05 | 5      |
| h   | vor|machen + (D)                    | to fool, to deceive (coll) | megtéveszteni vkit                       | verb     | 2014-08-31 | 5      |
| h   | vor|machen + sich (D)               | to lie to oneself          | beképzelni magának vmit, hazudni magának | verb     | 2014-08-31 | 5      |

#### Verbs with irregular past tense

Verbs with irregular past tense, but regular second and third persons in present tense are defined by 3 conjugations.

3 conjugations in order:

 * dictionary form (present form for we)
 * preterite
 * past particle

| A/A | German                                                 | English  | Third      | Category | Date       | Score  |
|-----|--------------------------------------------------------|----------|------------|----------|------------|--------|
| h   | schreiben, schrieben, geschrieben                      | to write | írni       | verb     | 2014-10-12 | 5      |
| h/s | Sport treiben, trieben, getrieben                      | to sport | sportolni  | verb     | 2014-05-01 | 5      |
| h   | treffen, trafen, getroffen, triffst, trifft + sich (A) | to meet  | találkozni | verb     | 2014-10-08 | 5      |

#### Verbs with irregular second/third person in present

Verbs with irregular second and third persons in present have 5 words after the auxiliary and before arguments.

5 words in order:

 * dictionary form (present form for wir)
 * preterite
 * past particle
 * present form for du
 * present form for er

| A/A | German                                                | English  | Third     | Category | Date       | Score  |
|-----|-------------------------------------------------------|----------|-----------|----------|------------|--------|
| h   | verlassen, verließ, verlassen, verlässt, verlässt     | to leave | elhagyni  | verb     | 2014-05-05 | 5      |
| h   | verwenden, verwendeten/verwandten, verwendet/verwandt | to use   | használni | verb     | 2014-05-01 | 5      |

#### Very irregular verbs

Some verbs are truly irregular in present tense and have 9 words after the auxiliary and before arguments.

9 words in order:

 * dictionary form (not used for word generation)
 * present form for I
 * present form for you
 * present form for he
 * present form for we
 * present form for you
 * present form for they
 * preterite
 * past particle

| A/A | German                                               | English | Third           | Category | Date       | Score  |
|-----|------------------------------------------------------|---------|-----------------|----------|------------|--------|
| s   | sein, bin, bist, ist, sind, seid, sind, war, gewesen | to be   | lenni, létezni  | verb     | 2014-05-01 | 5      |


### Nouns

Nouns are composed of an article (**r:** der, **e:** die, **s:** das), a main word and a notion for plural and optionally for genitive. Plural and genitive notions can be full words or extensions of the main word. If notions are extensions only than they are prefixed by `~` or `⍨`. Latter notes that the base word gets an äumlaut in the given form.

| A/A | German                                 | English                      | Third                     | Category   | Date       | Score  |
|-----|----------------------------------------|------------------------------|---------------------------|------------|------------|--------|
| e   | Entzündung,~en                         | inflammation                 | gyulladás                 | noun       | 2015-03-04 | 5      |
| e   | Trauma,Traumata/Traumen                | trauma                       | trauma                    | noun       | 2015-04-18 | 5      |
| e   | Vereinigten Staaten von Amerika,- (pl) | The United States of America | Amerikai Egyesült Államok | noun       | 2014-10-16 | 5      |


### Adjectives

Adjectives are composed of a main word, optionally a notion for comparative and optionally for superlative. Comparative and suparlative notions can be full words or extensions of the main word. If notions are extensions only than they are prefixed by `~` or `⍨`. Latter notes that the base word gets an äumlaut in the given form.

| A/A | German                     | English       | Third           | Category | Date       | Score  |
|-----|----------------------------|---------------|-----------------|----------|------------|--------|
|     | klug,⍨er,⍨sten             | smart         | okos, értelmes  | adj      | 2014-05-01 | 5      |
|     | hochschwanger,-            | very pregnant | terhes (nagyon) | adj      | 2014-05-01 | 5      |
|     | schmal,~er/⍨er,~sten/⍨sten | narrow        | keskeny, szűk   | adj      | 2014-05-01 | 5      |



Utils
-----

### Preparations

```bash
export D5_DBHOST="localhost"
export D5_DBNAME="d5"
```


### Convert Excel to Plain JSON

Convert excel file into a json file.

```bash
# .ods to .json, 8 columns processed, file saved into parser/fixture/gerdict.json
./spreadsheet/ods spreadsheet/fixture/gerdict.ods 8 > parser/fixture/gerdict.json

# .xlsx to .json, 8 columns processed, file saved into parser/fixture/gerdict.json
./spreadsheet/xlsx spreadsheet/fixture/gerdict.xlsx 8 > parser/fixture/gerdict.json

# .csv to .json, 8 columns processed, file saved into parser/fixture/gerdict.json
./spreadsheet/csv spreadsheet/fixture/gerdict.csv 8 > parser/fixture/gerdict.json
```


### Convert Plain JSON to Parsed JSON

```bash
cat parser/fixture/gerdict.json | parser --user=peteraba
```


### Persist Parsed JSON

```bash
cat persister/fixture/gerdict.json | persister --coll=german
```


### Run the import chain at once

```bash
./spreadsheet/csv spreadsheet/fixture/gerdict.csv 8 | parser --user=peteraba | persister --coll=german
```


### Finder

Used to find words

```bash
echo "limit=2&query={\"word.german\": \"solche\", \"word.user\": \"peteraba\"}" | finder --coll=german
```


### Scorer

Used to store game scores on words

```bash
scorer --coll=german --wordId=55c72270288a2139ea45f52e --score=7
```


### Router

Connects frontends ands games

```bash

```


Admin
=====


Games
=====

```bash
export GAME_HOSTNAME="localhost"
export GAME_DBNAME="d5"
```


Translate
---------

To be documented later.


Conjugate
---------

To be documented later.


Decline Noun
------------

Planned for 1.0.


Decline Adjective
-----------------

Planned for 1.0.


Frontends
=========

deutschodoro.com
----------------

Planned for 1.0.


Deutschodoro Chrome Plugin
--------------------------

Planned for 1.0.


Contributing to Deutschodoro
============================

Contributions to Deutschodoro of any kind is more than welcome including documentation, themes, organization, tutorials, blog posts, bug reports, issues, feature requests, feature implementation, pull requests, etc.

