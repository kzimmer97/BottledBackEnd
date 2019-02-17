package main

import (
	"bottled/utils"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//BottleBase is used to access the database-backend
type BottleBase struct {
	db          *sql.DB
	bottleCache map[int]Bottle
	userCache   map[int]User
}

func NewBottleBase(database *sql.DB) *BottleBase {
	b := BottleBase{
		db: database,
	}

	b.bottleCache = make(map[int]Bottle)
	b.buildBottleCache()

	b.userCache = make(map[int]User)
	b.buildUserCache()

	return &b
}

func (b *BottleBase) GrabSomeBottles(userID int) []Bottle {
	var outgoingBottles []Bottle

	user := b.userCache[userID]

	if user.M.Initiated == 0 {
		user.M = *utils.NewHeap(15)
	}

	m := b.GetUserIDs()

	//if we've got gps enabled
	user.locationEnabled = true

	if user.locationEnabled {
		fmt.Print(len(m))
		for k, _ := range m {
			u := b.userCache[k]

			if user.userID != k {
				q := u.Point
				n := utils.NewNode(userID, user.CalcDistance(q), (utils.Node{}))

				user.M.Insert(n)
			}
		}
	}

	x := user.M.Traverse()

	for _, e := range x {
		user := b.userCache[e]

		ihh := b.GrabBottleID(e)

		t := b.bottleCache[ihh]

		t.senderName = user.name

		outgoingBottles = append(outgoingBottles, t)
	}

	return outgoingBottles
}

func (b *BottleBase) PourBottle(bottleID int) Bottle {
	bottle := b.bottleCache[bottleID]

	defer b.PossibleCleanup(bottle)

	return bottle
}

//each bottle can only be sipped from 5 times, then we get rid of it.
func (b *BottleBase) PossibleCleanup(bottle Bottle) {
	bottle.DamageBottle()

	if bottle.broke {
		delete(b.bottleCache, bottle.bottleID)
	}
}

//given point p, check if location enabled. if it is, give 70/30 randomness with other non-locationers
//otherwise, find closest
