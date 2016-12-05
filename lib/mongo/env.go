package mongo

import "os"

const (
	envDbHost            = "D5_DB_HOST"
	envDbName            = "D5_DB_NAME"
	envGameType          = "D5_GAME_TYPE"
	envCollectionGeneral = "D5_COLLECTION_DATA_GENERAL"
	envCollectionGerman  = "D5_COLLECTION_DATA_GERMAN"
	envCollectionResult  = "D5_COLLECTION_RESULT"
)

const (
	german  = "german"
	general = "general"
)

func ParseDbEnvs() (string, string) {
	dbHost := os.Getenv(envDbHost)
	dbName := os.Getenv(envDbName)

	return dbHost, dbName
}

func ParseGameType() string {
	gameType := os.Getenv(envGameType)

	return gameType
}

func IsGameGerman(collectionType string) bool {
	return collectionType == ParseGameType()
}

func ParseDataCollection() string {
	var collectionName string

	gameType := os.Getenv(envGameType)

	switch gameType {
	case german:
		collectionName = os.Getenv(envCollectionGerman)
		break
	case general:
		collectionName = os.Getenv(envCollectionGeneral)
		break
	}

	return collectionName
}

func ParseGeneralCollection() string {
	collectionName := os.Getenv(envCollectionGeneral)

	return collectionName
}

func ParseGermalCollection() string {
	collectionName := os.Getenv(envCollectionGerman)

	return collectionName
}

func ParseResultCollection() string {
	collectionName := os.Getenv(envCollectionResult)

	return collectionName
}
