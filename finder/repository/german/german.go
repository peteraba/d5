package german

import (
	"gopkg.in/mgo.v2"

	german "github.com/peteraba/d5/lib/german"
)

type Repo struct {
}

func (r Repo) fetchCollection(mgoSession *mgo.Session, databaseName, collectionName string, query map[string]string) ([]german.Superword, error) {
	var (
		collection *mgo.Collection
		err        error
		result     = []german.Superword{}
	)

	collection = mgoSession.DB(databaseName).C(collectionName)

	err = collection.Find(query).All(&result)

	return result, err
}

func (r Repo) CreateDictionary(mgoSession *mgo.Session, dbName, collectionName string, query map[string]string) (interface{}, error) {
	var (
		err          error
		searchResult []german.Superword
		dictionary   german.Dictionary
	)

	searchResult, err = r.fetchCollection(mgoSession, dbName, collectionName, query)
	if err != nil {
		return dictionary, err
	}

	dictionary = german.SuperwordsToDictionary(searchResult)

	return dictionary, err
}
