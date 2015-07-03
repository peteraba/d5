package general

import "gopkg.in/mgo.v2"

type Repo struct {
}

func (r Repo) fetchCollection(mgoSession *mgo.Session, databaseName, collectionName string, query interface{}) ([]interface{}, error) {
	var (
		collection *mgo.Collection
		err        error
		result     []interface{}
	)

	collection = mgoSession.DB(databaseName).C(collectionName)

	err = collection.Find(query).All(&result)

	return result, err
}

func (r Repo) CreateDictionary(mgoSession *mgo.Session, dbName, collectionName string, query map[string]string) (interface{}, error) {
	return r.fetchCollection(mgoSession, dbName, collectionName, query)
}
