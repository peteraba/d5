package general

import (
	"testing"
	"time"
)

var now = time.Now()

var progressScoreCases = []struct {
	scores        []Score
	expectedScore int64
}{
	{
		[]Score{},
		0,
	},
	// 1 point, now
	{
		[]Score{
			Score{1, now},
		},
		-100,
	},
	// 10 points, 299 second ago
	{
		[]Score{
			Score{10, now.Add(-299 * time.Second)},
		},
		-100,
	},
	// 10 points, 301 second ago
	{
		[]Score{
			Score{10, now.Add(-301 * time.Second)},
		},
		-25,
	},
	// 1 point, 2 hours ago
	{
		[]Score{
			// -4 * 3
			Score{1, now.Add(-2 * time.Hour)},
		},
		12,
	},
	// 1 point, 3 hours 57 minutes ago
	{
		[]Score{
			Score{1, now.Add(-3*time.Hour - 57*time.Minute)},
		},
		8,
	},
	// 1 point, -8 hours ago
	{
		[]Score{
			Score{1, now.Add(-8 * time.Hour)},
		},
		4,
	},
	// 1 point, 1 day and 1 hour ago
	{
		[]Score{
			Score{1, now.Add(-25 * time.Hour)},
		},
		3,
	},
	// 10 points, 3 days ago
	{
		[]Score{
			Score{10, now.Add(-3 * 24 * time.Hour)},
		},
		-3,
	},
	// 10 points, 5 weeks ago
	{
		[]Score{
			Score{10, now.Add(-5 * 7 * 24 * time.Hour)},
		},
		0,
	},
	// 10 * 10 points, 50 weeks ago
	{
		[]Score{
			// 5 * 0.1 = 0.5
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
			Score{10, now.Add(-50 * 7 * 24 * time.Hour)},
		},
		-5,
	},
	// 1 point, 30 hours ago
	// 10 points 2 hours ago
	{
		[]Score{
			// -4 * 9 * 0.1 = -3.6
			Score{1, now.Add(-30 * time.Hour)},
			// 5 * 3 = 15
			Score{10, now.Add(-2 * time.Hour)},
		},
		// 15 - 3.6 - 11.4
		-11,
	},
	{
		[]Score{
			// -4 * 9 0.1 = -3.6
			Score{1, now.Add(-30 * time.Hour)},
			// 100
			Score{10, now.Add(-1 * time.Second)},
		},
		-96,
	},
}

func TestGetProgressScore(t *testing.T) {
	var (
		actualScore   int64
		scorePointers []*Score
	)

	for num, testCase := range progressScoreCases {
		scorePointers = []*Score{}

		for num, _ = range testCase.scores {
			scorePointers = append(scorePointers, &testCase.scores[num])
		}

		actualScore = GetProgressScore(scorePointers)

		if testCase.expectedScore != actualScore {
			t.Fatalf(
				"#%d. Progress score is different from expected. Expected: %v, got: %v.\n",
				num,
				testCase.expectedScore,
				actualScore,
			)
		}
	}

	t.Log(len(progressScoreCases), "test cases")
}

func TestGetRandomScore(t *testing.T) {
	var rand = GetRandomScore()

	if rand < 0 || rand >= 10 {
		t.Fatalf("Random score should be an integer between 0 and 9, got %d\n", rand)
	}

	t.Log(1, "test case")
}

var learnedAtScoreCases = []struct {
	learnedAt     time.Time
	expectedScore int64
}{
	{time.Now().Add(-2 * 32 * 24 * time.Hour), 1},
	{time.Now().Add(-4 * 32 * 24 * time.Hour), 2},
	{time.Now().Add(-20 * 32 * 24 * time.Hour), 6},
}

func TestGetLearnedAtScore(t *testing.T) {
	var actualResult int64

	for num, testCase := range learnedAtScoreCases {
		actualResult = GetLearnedAtScore(testCase.learnedAt)

		if actualResult != testCase.expectedScore {
			t.Fatalf(
				"#%d. Learned at score is wrong. Expected: %d, got: %d\n",
				num,
				testCase.expectedScore,
				actualResult,
			)
		}
	}

	t.Log(len(learnedAtScoreCases), "test cases")
}
