package general

import (
	"fmt"
	"testing"
)

func TestNewScoreSuccess(t *testing.T) {
	score, err := NewScore(5)

	if err != nil {
		t.Fatalf("Error while creating score: %v", err)
	}

	if score.Result != 5 {
		t.Fatalf(fmt.Sprintf("Score result is wrong. Expected 5, got: %d.", score.Result))
	}
}

var newScoreFailureCases = []struct{ score int }{
	{-15},
	{15},
}

func TestNewScoreFailure(t *testing.T) {
	for num, testCase := range newScoreFailureCases {
		score, err := NewScore(testCase.score)

		if err == nil {
			t.Fatalf("No error while creating score #%d", num)
		}

		if score != nil {
			t.Fatalf("Score result is wrong. Expected nil.")
		}
	}
}
