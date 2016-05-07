package util

import (
	"flag"
	"os"
)

const (
	dbhost_env = "D5_DBHOST"
	dbname_env = "D5_DBNAME"
)

const (
	coll_type_default = "default"
	coll_type_german  = "german"
)

func ParseFlags() (bool, int, string, string, bool, map[string]string) {
	isServer := flag.Bool("server", false, "Starts a server")
	port := flag.Int("port", 17172, "Port for server")

	collectionName := flag.String("coll", "german", "Port for server")
	collectionType := flag.String("type", coll_type_german, "Type of collection (german, anything else)")

	debug := flag.Bool("debug", false, "Enables debug logs")

	json := flag.String("data", "{}", "")

	flag.Parse()

	data, _ := JsonToStringMap([]byte(*json))

	return *isServer, *port, *collectionName, *collectionType, *debug, data
}

func ParseEnvs() (string, string) {
	// Mongo database host
	dbHost := os.Getenv(dbhost_env)

	// Mongo database name
	dbName := os.Getenv(dbname_env)

	return dbHost, dbName
}

func IsGerman(collectionType string) bool {
	if collectionType == "" {
		return false
	}

	if collectionType == coll_type_default {
		return false
	}

	return true
}
