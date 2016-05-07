package util

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadStdInput() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return ioutil.ReadAll(reader)
}
