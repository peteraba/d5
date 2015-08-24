package game

import "time"

type Tracker struct {
	Id        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	WordIds   []string  `json:"wordIds,omitempty" bson:"wordIds,omitempty"`
	Right     []string  `json:"right,omitempty" bson:"right,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
