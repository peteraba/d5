package german

import (
	"encoding/json"

	"github.com/peteraba/d5/lib/german/entity"
)

type superword struct {
	entity.DefaultWord `bson:"word" json:"word"`
	Auxiliary          []entity.Auxiliary `bson:"auxiliary" json:"auxiliary"`
	Prefix             entity.Prefix      `bson:"prefix" json:"prefix"`
	Noun               string             `bson:"noun" json:"noun"`
	Adjective          string             `bson:"adjective" json:"adjective"`
	PastParticiple     []string           `bson:"pastParticiple" json:"pastParticiple"`
	Preterite          []string           `bson:"preterite" json:"preterite"`
	S1                 []string           `bson:"s1" json:"s1"`
	S2                 []string           `bson:"s2" json:"s2"`
	S3                 []string           `bson:"s3" json:"s3"`
	P1                 []string           `bson:"p1" json:"p1"`
	P2                 []string           `bson:"p2" json:"p2"`
	P3                 []string           `bson:"p3" json:"p3"`
	Reflexive          entity.Reflexive   `bson:"reflexive" json:"reflexive"`
	Arguments          []entity.Argument  `bson:"arguments" json:"arguments"`
	Articles           []entity.Article   `bson:"article" json:"article"`
	Plural             []string           `bson:"plural" json:"plural"`
	Genitive           []string           `bson:"genitive" json:"genitive"`
	IsPluralOnly       bool               `bson:"plural_only" json:"plural_only"`
	Comparative        []string           `bson:"comparative" json:"comparative"`
	Superlative        []string           `bson:"superlative" json:"superlative"`
}

func ParseWords(input []byte) ([]entity.Word, error) {
	var (
		superwords = []superword{}
		words      = []entity.Word{}
		word       entity.Word
		err        error
	)

	if err = json.Unmarshal(input, &superwords); err != nil {
		return words, err
	}

	for _, superword := range superwords {
		switch superword.Category {
		case "verb":
			word = &entity.Verb{
				superword.DefaultWord,
				superword.Auxiliary,
				superword.Prefix,
				superword.Noun,
				superword.Adjective,
				superword.PastParticiple,
				superword.Preterite,
				superword.S1,
				superword.S2,
				superword.S3,
				superword.P1,
				superword.P2,
				superword.P3,
				superword.Reflexive,
				superword.Arguments,
			}
			break
		case "noun":
			word = &entity.Noun{
				superword.DefaultWord,
				superword.Articles,
				superword.Plural,
				superword.Genitive,
				superword.IsPluralOnly,
			}
			break
		case "adj":
			word = &entity.Adjective{
				superword.DefaultWord,
				superword.Comparative,
				superword.Superlative,
			}
			break
		default:
			word = &entity.Any{
				superword.DefaultWord,
			}
			break
		}

		words = append(words, word)
	}

	return words, nil
}
