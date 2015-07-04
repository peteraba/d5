package german

import (
	"gopkg.in/mgo.v2"

	german "github.com/peteraba/d5/lib/german"
)

type Repo struct {
	Db *mgo.Database
}

func (r *Repo) fetchCollection(collectionName string, query map[string]string) ([]german.Superword, error) {
	var (
		collection *mgo.Collection
		err        error
		result     = []german.Superword{}
	)

	collection = r.Db.C(collectionName)

	err = collection.Find(query).All(&result)

	return result, err
}

func (r Repo) CreateDictionary(collectionName string, query map[string]string) (interface{}, error) {
	var (
		err          error
		searchResult []german.Superword
		dictionary   german.Dictionary
	)

	searchResult, err = r.fetchCollection(collectionName, query)
	if err != nil {
		return dictionary, err
	}

	dictionary = german.SuperwordsToDictionary(searchResult)

	return dictionary, err
}

func (r *Repo) SetDb(db *mgo.Database) {
	r.Db = db
}
