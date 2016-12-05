package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	germanEntity "github.com/peteraba/d5/lib/german/entity"
	"github.com/peteraba/d5/lib/server"
	"github.com/peteraba/d5/lib/util"
)

const name = "parser"
const version = "0.1"
const usage = `
Parser supports CLI and Server mode.

In CLI mode it expects input data on standard input, in server mode as raw POST body

Usage:
  parser [--server] [--port=<n>] [--debug] [--user=<s>]
  parser -h | --help
  parser -v | --version

Options:
  -s, --server    run in server mode
  -p, --port=<n>  port to open (server mode only) [default: 10010]
  -d, --debug     skip ticks and generate fake data concurrently
  -v, --version   show version information
  -h, --help      show help information
  -u, --user=<s>  user the data belongs to (cli mode only)

Accepted input data:
  - Raw JSON data to parse
`

/**
 * MAIN
 */

func main() {
	cliArguments := util.GetCliArguments(usage, name, version)
	isServer, port, isDebug := util.GetServerOptions(cliArguments)
	user, _ := cliArguments["--user"].(string)

	if isServer {
		startServer(port, isDebug)
	} else {
		serveCli(isDebug, user)
	}
}

/**
* DOMAIN
 */
func logParseErrors(isDebug bool, parseErrors []string) {
	if !isDebug {
		return
	}

	for _, word := range parseErrors {
		log.Printf("Failed: %v\n", word)
	}
}

func parseDictionary(dictionary [][8]string, user string) ([]germanEntity.Word, []string) {
	var (
		words       = []germanEntity.Word{}
		parseErrors = []string{}
	)

	for _, rawWord := range dictionary {
		word, german := createWord(rawWord, user)

		words = append(words, word)

		if german == "" {
			continue
		}

		parseErrors = append(parseErrors, german)
	}

	return words, parseErrors
}

func createWord(word [8]string, user string) (germanEntity.Word, string) {
	var (
		w                  germanEntity.Word
		articleOrAuxiliary = word[0]
		german             = word[1]
		english            = word[2]
		third              = word[3]
		category           = word[4]
		learned            = word[5]
		score              = word[6]
		tags               = word[7]
	)

	if english == "" {
		return w, german
	}

	switch category {
	case "adj":
		w = germanEntity.NewAdjective(german, english, third, user, learned, score, tags)
		break
	case "noun":
		if germanEntity.NounRegexp.MatchString(german) {
			w = germanEntity.NewNoun(articleOrAuxiliary, german, english, third, user, learned, score, tags)
		}
		break
	case "verb":
		if germanEntity.VerbRegexp.MatchString(german) {
			w = germanEntity.NewVerb(articleOrAuxiliary, german, english, third, user, learned, score, tags)
		}
		break
	default:
		w = germanEntity.NewAny(german, english, third, category, user, learned, score, tags, []string{})
	}

	if w != nil {
		return w, ""
	}

	w = germanEntity.NewAny(german, english, third, category, user, learned, score, tags, []string{"Parsing failed."})

	return w, german
}

/**
 * CLI
 */

func serveCli(isDebug bool, user string) {
	err := cliHandler(isDebug, user)

	util.LogFatalErr(err, isDebug)
}

func cliHandler(isDebug bool, user string) error {
	input, err := util.ReadStdInput()
	if err != nil {
		return err
	}

	words, parseErrors := getParserData(input, user)

	logParseErrors(isDebug, parseErrors)

	b, err := json.Marshal(words)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

/**
 * SERVER
 */

func startServer(port int, isDebug bool) {
	s := server.MakeServer(port, nil, isDebug)

	s.AddHandler("/", parseHandle, server.PostOnly)

	s.Start()
}

func parseHandle(w http.ResponseWriter, r *http.Request, mgoDb *mgo.Database, isDebug bool) error {
	words, parseErrors := getServerParserData(r)
	logParseErrors(isDebug, parseErrors)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(w).Encode(words)

	return err
}

func getServerParserData(r *http.Request) ([]germanEntity.Word, []string) {
	rawBody, _ := ioutil.ReadAll(r.Body)
	user := r.FormValue("user")

	return getParserData(rawBody, user)
}

/**
 * INPUT PARSING
 */

func getParserData(rawInput []byte, user string) ([]germanEntity.Word, []string) {
	var dictionary = [][8]string{}

	err := json.Unmarshal(rawInput, &dictionary)
	if err != nil {
		return []germanEntity.Word{}, []string{err.Error()}
	}

	words, parseErrors := parseDictionary(dictionary, user)

	return words, parseErrors
}
