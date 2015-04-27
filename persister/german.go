package main

import (
	"os"

	"gopkg.in/mgo.v2"
)

func main() {
	var (
		url = os.Args[1]
	)
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
}
