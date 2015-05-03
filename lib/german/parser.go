package german

import "encoding/json"

type superword struct {
	DefaultWord    `bson:"word" json:"word"`
	Auxiliary      []string  `bson:"auxiliary" json:"auxiliary"`
	Noun           string    `bson:"noun" json:"noun"`
	Adjective      string    `bson:"adjective" json:"adjective"`
	PastParticiple []string  `bson:"pastParticiple" json:"pastParticiple"`
	Preterite      []string  `bson:"preterite" json:"preterite"`
	S1             []string  `bson:"s1" json:"s1"`
	S2             []string  `bson:"s2" json:"s2"`
	S3             []string  `bson:"s3" json:"s3"`
	P1             []string  `bson:"p1" json:"p1"`
	P2             []string  `bson:"p2" json:"p2"`
	P3             []string  `bson:"p3" json:"p3"`
	Reflexive      Reflexive `bson:"reflexive" json:"reflexive"`
	Arguments      []string  `bson:"arguments" json:"arguments"`
	Articles       []Article `bson:"article" json:"article"`
	Plural         []string  `bson:"plural" json:"plural"`
	Comparative    []string  `bson:"comparative" json:"comparative"`
	Superlative    []string  `bson:"superlative" json:"superlative"`
}

func ParseWords(input []byte) ([]Word, error) {
	var (
		superwords = []superword{}
		words      = []Word{}
		word       Word
		err        error
	)

	if err = json.Unmarshal(input, &superwords); err != nil {
		return words, err
	}

	for _, superword := range superwords {
		switch superword.Category {
		case "verb":
			word = &Verb{
				superword.DefaultWord,
				superword.Auxiliary,
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
			word = &Noun{
				superword.DefaultWord,
				superword.Articles,
				superword.Plural,
			}
			break
		case "adjective":
			word = &Adjective{
				superword.DefaultWord,
				superword.Comparative,
				superword.Superlative,
			}
			break
		default:
			word = &Any{
				superword.DefaultWord,
			}
			break
		}

		words = append(words, word)
	}

	return words, nil
}
