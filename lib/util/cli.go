package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/tj/docopt"
)

func ReadStdInput() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}

func GetCliArguments(usage, name, version string) map[string]interface{} {
	arguments, _ := docopt.Parse(usage, nil, true, fmt.Sprintf("%s %s", name, version), false)

	return arguments
}

func GetServerOptions(arguments map[string]interface{}) (bool, int, bool) {
	isServer, _ := arguments["--server"].(bool)

	rawPort, _ := arguments["--port"].(string)
	port64, _ := strconv.ParseInt(rawPort, 10, 64)

	isDebug, _ := arguments["--debug"].(bool)

	return isServer, int(port64), isDebug
}

func GetGameOptions(arguments map[string]interface{}) (string, string) {
	finder, _ := arguments["--finder"].(string)
	scorer, _ := arguments["--scorer"].(string)

	return finder, scorer
}
