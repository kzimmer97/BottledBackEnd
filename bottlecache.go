package main

import (
	"fmt"
)

func (b *BottleBase) UpdateBottleCache(bottleID int, bottle Bottle) {
	b.bottleCache[bottleID] = bottle
}

func (b BottleBase) HasBottleCached(bottleID int) bool {
	_, ok := b.bottleCache[bottleID]

	return ok
}

//all unique ids
func (b BottleBase) GetUserIDs() map[int]User {
	var id = make(map[int]User)

	fmt.Printf("\n%v\n", b.bottleCache)
	for k, v := range b.bottleCache {
		id[v.senderID] = b.userCache[v.senderID]
		fmt.Println(k)
	}
	return id
}

func (b *BottleBase) CreateNewBottle(x Bottle) {
	b.ArchiveBottle(x.senderID, x.secret, x.life)
	x.bottleID = b.GetBottleID()
	b.UpdateBottleCache(x.bottleID, x)
}

func (b *BottleBase) buildBottleCache() {
	sqlStatement := fmt.Sprintf(`SELECT outgoingbottles.bottleid, outgoingbottles.sendid, outgoingbottles.message, users.name, users.lat, users.long, users.locationEnabled FROM outgoingbottles INNER JOIN users ON outgoingbottles.sendid=users.userid`)

	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.

	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	var name, message string
	var x, y float64
	var bottleID, sendID int
	var locationEnabled bool

	for rows.Next() {
		err := rows.Scan(&bottleID, &sendID, &message, &name, &x, &y, &locationEnabled)
		if err != nil {
			panic(err)
		}

		//build cache on program start to save database load
		bottle := Bottle{
			bottleID: bottleID,
			senderID: sendID,
			secret:   message,
			location: Point{
				lat:          x,
				long:         y,
				userID:       bottleID,
				locationless: !locationEnabled,
			},
		}

		b.bottleCache[bottle.bottleID] = bottle
		fmt.Println(bottle)
	}
}
