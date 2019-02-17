package main

import "fmt"

func (b *BottleBase) UpdateUserCache(bottleID int, bottle Bottle) {
	b.bottleCache[bottleID] = bottle
}

func (b BottleBase) HasUserCached(userID int) bool {
	_, ok := b.userCache[userID]

	return ok
}

func (b BottleBase) GetUser(userID int) *User {
	user := b.userCache[userID]

	return &user
}

//builds cache of useful info about users
func (b *BottleBase) buildUserCache() {
	sqlStatement := fmt.Sprintf(`SELECT userid, name, lat, long, locationEnabled FROM users`)

	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	var name string
	var x, y float64
	var userID int
	var locationEnabled bool

	for rows.Next() {
		err := rows.Scan(&userID, &name, &x, &y, &locationEnabled)
		if err != nil {
			panic(err)
		}

		//build user cache on program start to save database load
		user := User{
			userID: userID,
			name:   name,
			letter: string([]rune(name)[0]),
			Point: Point{
				lat:          x,
				long:         y,
				userID:       userID,
				locationless: !locationEnabled,
			},
		}

		b.userCache[user.userID] = user
	}
}
