package general

import "gopkg.in/mgo.v2"

type Repo struct {
	Db *mgo.Database
}

func (r *Repo) fetchCollection(collectionName string, query map[string]string) ([]interface{}, error) {
	var (
		collection *mgo.Collection
		err        error
		result     []interface{}
	)

	collection = r.Db.C(collectionName)

	err = collection.Find(query).All(&result)

	return result, err
}

func (r *Repo) CreateDictionary(collectionName string, query map[string]string) (interface{}, error) {
	return r.fetchCollection(collectionName, query)
}

func (r *Repo) SetDb(db *mgo.Database) {
	r.Db = db
}
