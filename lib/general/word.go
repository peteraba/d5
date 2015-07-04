package general

import (
	"errors"
	"fmt"
	"time"
)

type Score struct {
	Result    int       `bson:"result" json:"result"`
	LearnedAt time.Time `bson:"learned_at" json:"learned_at"`
}

func NewScore(result int) (*Score, error) {
	if result < -10 || result > 10 {
		return nil, errors.New(fmt.Sprintf("Result must be between -10 and 10, got %d", result))
	}

	return &Score{result, time.Now()}, nil
}

type Word interface {
	GetScores() []Score
}
