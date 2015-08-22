package game

import (
	"github.com/peteraba/d5/lib/german"
	"github.com/peteraba/d5/lib/german/entity"
	"gopkg.in/mgo.v2/bson"
)

type Game interface {
	GetQuestion() string
	GetId() string
}

type GameOption struct {
	GameDebug
	Question string `json:"question"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2,omitempty"`
	Option3  string `json:"option3,omitempty"`
	Option4  string `json:"option4,omitempty"`
	Id       string `json:"id"`
}

func NewGameOption() GameOption {
	game := GameOption{}

	return game
}

type GameAnswer struct {
	GameDebug
	Question string `json:"question"`
	Id       string `json:"id"`
}

func NewGameAnswer() GameAnswer {
	game := GameAnswer{}

	return game
}

type GameDebug struct {
	Error      string             `json:"error,omitempty"`
	Dictionary *german.Dictionary `json:"dictionary,omitempty"`
	Word       entity.Word        `json:"word,omitempty"`
	Query      *bson.M            `json:"options,omitempty"`
	Options    *bson.M            `json:"options,omitempty"`
	Right      []string           `json:"right,omitempty"`
	English    []entity.Meaning   `json:"english,omitempty"`
	Third      []entity.Meaning   `json:"third,omitempty"`
}

func (g GameOption) GetQuestion() string {
	return g.Question
}

func (g GameOption) GetId() string {
	return g.Question
}

func (g GameAnswer) GetQuestion() string {
	return g.Question
}

func (g GameAnswer) GetId() string {
	return g.Question
}
