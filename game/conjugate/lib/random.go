package lib

import (
	"math/rand"

	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2/bson"
)

func GetRandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	switch rand.Int31n(6) {
	case 0:
		return GetS2RandomPieces(user)
	case 1:
		return GetPastRandomPieces(user)
	case 2:
		return GetPastParticleRandomPieces(user)
	}

	return GetGeneralRandomPieces(user)
}

func GetS2RandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	var (
		pp    entity.PersonalPronoun
		query = getBaseQuery(user)
	)

	query["s2"] = bson.M{"$regex": bson.RegEx{".*", "s"}}

	switch rand.Int31n(2) {
	case 0:
		pp = entity.S2
		break
	case 1:
		pp = entity.S3
		break
	}

	return query, pp, entity.Present
}

func GetPastRandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	var (
		pp    entity.PersonalPronoun
		tense entity.Tense
		query bson.M
	)

	query, pp, tense = GetGeneralRandomPieces(user)

	query["preterite"] = bson.M{"$regex": bson.RegEx{".*", "s"}}

	switch rand.Int31n(3) {
	case 0:
	case 1:
		tense = entity.PastParticiple
		break
	case 2:
		tense = entity.Preterite
		break
	}

	return query, pp, tense
}

func GetPastParticleRandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	var (
		pp    entity.PersonalPronoun
		query = getBaseQuery(user)
	)

	query, pp, _ = GetGeneralRandomPieces(user)

	query["auxiliary"] = "s"

	return query, pp, entity.PastParticiple
}

func GetGeneralRandomPieces(user string) (bson.M, entity.PersonalPronoun, entity.Tense) {
	var (
		tense entity.Tense
		query = getBaseQuery(user)
	)

	switch rand.Int31n(3) {
	case 0:
		tense = entity.Present
		break
	case 1:
		tense = entity.Preterite
	case 2:
		tense = entity.PastParticiple
	}

	switch rand.Int31n(6) {
	case 0:
		return query, entity.S1, tense
	case 1:
		return query, entity.S2, tense
	case 2:
		return query, entity.S3, tense
	case 3:
		return query, entity.P1, tense
	case 4:
		return query, entity.P2, tense
	}

	return query, entity.P3, tense
}

func getBaseQuery(user string) bson.M {
	var (
		query = bson.M{}
	)

	query["word.category"] = "verb"
	query["word.user"] = user

	return query
}
