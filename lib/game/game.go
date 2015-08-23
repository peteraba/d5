package game

import (
	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2/bson"
)

type Game interface {
	GetQuestion() string
	GetId() string
	SetDebugQuery(bool, *bson.M, *german.Dictionary, entity.Word, *bson.M)
	SetDebugResult(bool, []string, []entity.Meaning, []entity.Meaning)
}

type GameOption struct {
	GameDebug
	Question string `json:"question" bson:"question"`
	Option1  string `json:"option1,omitempty" bson:"option1,omitempty"`
	Option2  string `json:"option2,omitempty" bson:"option2,omitempty"`
	Option3  string `json:"option3,omitempty" bson:"option3,omitempty"`
	Option4  string `json:"option4,omitempty" bson:"option4,omitempty"`
	Id       string `json:"_id,omitempty" bson:"_id,omitempty"`
}

func NewGameOption() GameOption {
	game := GameOption{}

	return game
}

func (g GameOption) GetQuestion() string {
	return g.Question
}

func (g GameOption) GetId() string {
	return g.Question
}

type GameAnswer struct {
	GameDebug
	Question string `bson:"question" json:"question"`
	Id       string `json:"_id,omitempty" bson:"_id,omitempty"`
}

func NewGameAnswer() GameAnswer {
	game := GameAnswer{}

	return game
}

func (g GameAnswer) GetQuestion() string {
	return g.Question
}

func (g GameAnswer) GetId() string {
	return g.Question
}

type GameDebug struct {
	Error      string             `json:"error,omitempty" bson:"error,omitempty"`
	Query      *bson.M            `json:"query,omitempty" bson:"query,omitempty"`
	Dictionary *german.Dictionary `json:"dictionary,omitempty" bson:"dictionary,omitempty"`
	Word       entity.Word        `json:"word,omitempty" bson:"word,omitempty"`
	Options    *bson.M            `json:"options,omitempty" bson:"options,omitempty"`
	Right      []string           `json:"right,omitempty" bson:"right,omitempty"`
	English    []entity.Meaning   `json:"english,omitempty" bson:",omitempty"`
	Third      []entity.Meaning   `json:"third,omitempty" bson:"third,omitempty"`
}

func (g *GameDebug) SetDebugQuery(debug bool, query *bson.M, dictionary *german.Dictionary, word entity.Word, options *bson.M) {
	if !debug {
		return
	}

	g.Query = query
	g.Dictionary = dictionary
	g.Word = word
	g.Options = options
}

func (g *GameDebug) SetDebugResult(debug bool, right []string, english []entity.Meaning, third []entity.Meaning) {
	if !debug {
		return
	}

	g.Right = right
	g.English = english
	g.Third = third
}
