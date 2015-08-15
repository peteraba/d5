package general

import (
	"math/rand"
	"time"
)

const (
	// 60 * 60
	oneHour = 3600
	// 24 * oneHour
	oneDay = 86400
	// one month in seconds: 365.25 / 12 * oneDay
	oneMonth = 2629800
	// two months in seconds: 2 * twoMonth
	twoMonths = 5259600
)

func GetLearnedAtScore(learnedAt time.Time) int64 {
	var timeDelta int64

	timeDelta = time.Now().Unix() - learnedAt.Unix()

	if timeDelta/twoMonths > 6 {
		return 6
	}

	return int64(timeDelta / twoMonths)
}

func GetProgressScore(scores []*Score) int64 {
	var (
		learnedScore, progressScore float64
		now, timeDelta, hours, days int64
		currentScore                int64
	)

	now = time.Now().Unix()

	for _, score := range scores {
		timeDelta = now - score.LearnedAt.Unix()
		hours = int64(timeDelta / oneHour)
		days = int64(timeDelta / oneDay)
		currentScore = int64(score.Result - 5)

		if timeDelta < 300 {
			learnedScore = 100
		} else if hours < 5 {
			learnedScore = float64(currentScore * (5 - hours))
		} else if days < 10 {
			learnedScore = float64(currentScore*(10-days)) * 0.1
		} else {
			learnedScore = float64(currentScore) * 0.1
		}

		progressScore -= learnedScore
	}

	return int64(progressScore)
}

func GetRandomScore() int64 {
	return rand.Int63n(10)
}
