package game

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/peteraba/d5/lib/mongo"
	"github.com/peteraba/d5/lib/util"
	"gopkg.in/mgo.v2"
)

const usageBase = `
%s supports CLI and Server mode.

In CLI mode it expects input data on standard input, in server mode as raw POST body

Usage:
  %s [--server] [--port=<n>] [--debug] [--user=<s>] [--action=<s>] [--finder=<s>] [scorer=<s>]
  %s -h | --help
  %s -v | --version

Options:
  -s, --server      run in server mode
  -p, --port=<n>    port to open (server mode only) [default: %s]
  -d, --debug       skip ticks and generate fake data concurrently
  -u, --user=<s>    user the data belongs to (cli mode only)
  -a, --action=<s>  action (cli mode only)
  -f, --finder=<s>  url to use to access finder (optional)
  -c, --scorer=<s>  url to use to access scorer (optional)
  -v, --version     show version information
  -h, --help        show help information

API
  - game
    - CLI:  ?action=game&user=john_doe
    - HTTP: /game/john_doe
  - answer
    - CLI:  ?action=answer&user=john_doe
    - HTTP: /answer/john_doe
`

type GameServer interface {
	MakeGameHandle(finderUrl string, mgoCollection *mgo.Collection, isDebug bool) func(*gin.Context)
	MakeCheckAnswerHandle(finderUrl, scorerUrl string, mgoCollection *mgo.Collection, isDebug bool) func(*gin.Context)
}

func Main(name, version, defaultPort string, gameServer GameServer) {
	nameLower := strings.ToLower(name)
	usage := fmt.Sprintf(usageBase, name, nameLower, nameLower, nameLower, defaultPort)

	cliArguments := util.GetCliArguments(usage, name, version)
	isServer, port, isDebug := util.GetServerOptions(cliArguments)
	finderUrl, scorerUrl := util.GetGameOptions(cliArguments)

	mgoDb := mongo.CreateMgoDbFromEnvs()
	mgoCollection := mgoDb.C(mongo.ParseResultCollection())

	if isServer {
		startServer(gameServer, port, mgoCollection, finderUrl, scorerUrl, isDebug)
		return
	}

	user, _ := cliArguments["--user"].(string)
	action, _ := cliArguments["--action"].(string)
	serveCli(gameServer, mgoCollection, user, action, finderUrl, scorerUrl, isDebug)
}

/**
 * CLI
 */

func serveCli(gameServer GameServer, mgoCollection *mgo.Collection, user, action, finderUrl, scorerUrl string, isDebug bool) {
	log.Println("CLI support is not available yet.")
}

/**
 * SERVER
 */

func startServer(gameServer GameServer, port int, mgoCollection *mgo.Collection, finderUrl, scorerUrl string, isDebug bool) {
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/game/:user", gameServer.MakeGameHandle(finderUrl, mgoCollection, isDebug))
	router.POST("/answer/:user", gameServer.MakeCheckAnswerHandle(finderUrl, scorerUrl, mgoCollection, isDebug))

	router.Run(fmt.Sprintf(":%d", port))
}
