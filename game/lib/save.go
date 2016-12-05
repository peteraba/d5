package game

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

func SaveAnswer(id string, wordIds, right []string, mgoCollection *mgo.Collection) error {
	var (
		err     error
		tracker Tracker
	)

	tracker = Tracker{}
	tracker.Id = id
	tracker.WordIds = wordIds
	tracker.Right = right
	tracker.CreatedAt = time.Now()

	err = mgoCollection.Insert(tracker)

	return err
}

func FindAnswer(id string, mgoCollection *mgo.Collection) (Tracker, error) {
	var (
		trackers = []Tracker{}
		err      error
		errorMsg string
	)

	err = mgoCollection.FindId(id).All(&trackers)
	if err != nil {
		log.Printf("Error while finding answer with id: %s, err: %v\n", id, err)

		return Tracker{}, err
	}

	if len(trackers) < 1 {
		errorMsg = fmt.Sprintf("Failed to find answer with id: %s", id)

		log.Println(errorMsg)

		return Tracker{}, errors.New(errorMsg)
	}

	return trackers[0], nil
}
