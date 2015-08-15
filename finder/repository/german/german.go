package german

import (
	"sort"

	"gopkg.in/mgo.v2"

	"github.com/peteraba/d5/lib/general"
	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
)

type Repo struct {
	Db         *mgo.Database
	lastResult []entity.Word
}

type ByLearned []entity.Word

func (a ByLearned) Len() int {
	return len(a)
}
func (a ByLearned) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByLearned) Less(i, j int) bool {
	return getScore(a[i]) < getScore(a[j])
}

func getScore(w entity.Word) int64 {
	var score int64

	score += general.GetLearnedAtScore(w.GetLearned())
	score += general.GetProgressScore(w.GetScores())
	score += general.GetRandomScore()

	return score
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

func (r *Repo) FetchDictionary(collectionName string, query map[string]string) (interface{}, error) {
	var (
		err          error
		searchResult []german.Superword
	)

	searchResult, err = r.fetchCollection(collectionName, query)
	if err != nil {
		return []entity.Word{}, err
	}

	r.lastResult = german.SuperwordsToWords(searchResult)

	return r.lastResult, err
}

func (r *Repo) FilterDictionary(limit int) (interface{}, error) {
	sort.Sort(ByLearned(r.lastResult))

	if limit > 0 && limit < len(r.lastResult) {
		r.lastResult = r.lastResult[:limit]
	}

	return r.lastResult, nil
}

func (r *Repo) SetDb(db *mgo.Database) {
	r.Db = db
}
