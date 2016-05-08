package mongo

import "os"

const (
	dbhost_env = "D5_DBHOST"
	dbname_env = "D5_DBNAME"
)

func ParseEnvs() (string, string) {
	// Mongo database host
	dbHost := os.Getenv(dbhost_env)

	// Mongo database name
	dbName := os.Getenv(dbname_env)

	return dbHost, dbName
}
