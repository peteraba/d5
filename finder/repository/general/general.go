package general

import "gopkg.in/mgo.v2"

type Repo struct {
	Db         *mgo.Database
	lastResult []interface{}
}

func (r *Repo) fetchCollection(collectionName string, query map[string]string) ([]interface{}, error) {
	var (
		collection *mgo.Collection
		err        error
	)

	collection = r.Db.C(collectionName)

	err = collection.Find(query).All(&r.lastResult)

	return r.lastResult, err
}

func (r *Repo) FetchDictionary(collectionName string, query map[string]string) (interface{}, error) {
	return r.fetchCollection(collectionName, query)
}

func (r *Repo) FilterDictionary(limit int) (interface{}, error) {
	return r.lastResult, nil
}

func (r *Repo) SetDb(db *mgo.Database) {
	r.Db = db
}
