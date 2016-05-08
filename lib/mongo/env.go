package mongo

import "os"

const (
	env_dbhost   = "D5_DBHOST"
	env_dbname   = "D5_DBNAME"
	env_collname = "D5_COLLECTION_NAME"
	env_colltype = "D5_COLLECTION_TYPE"
)

const (
	coll_type_default = "default"
	coll_type_german  = "german"
)

func ParseDbEnvs() (string, string) {
	dbHost := os.Getenv(env_dbhost)
	dbName := os.Getenv(env_dbname)

	return dbHost, dbName
}

func ParseCollectionEnvs() (string, string) {
	collectionName := os.Getenv(env_collname)
	collectionType := os.Getenv(env_colltype)

	return collectionName, collectionType
}

func IsGerman(collectionType string) bool {
	return collectionType == coll_type_german
}
