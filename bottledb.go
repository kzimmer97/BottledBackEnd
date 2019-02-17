package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//BottleBase is used to access the database-backend
type BottleBase struct {
	DB *sql.DB
}

func (b *BottleBase) CreateUser(name string, x float64, y float64, locationEnabled bool) error {
	sqlStatement := `
	INSERT INTO users (name, locx, locy, locationenabled)
	VALUES ($1, $2, $3,$4)`
	_, err := b.DB.Exec(sqlStatement, name, x, y, locationEnabled)

	if err != nil {
		panic(err)
	}

	return nil
}

func (b *BottleBase) GetUsers(userid int) string {
	sqlStatement := fmt.Sprintf(`SELECT name FROM users WHERE userid=%d`, userid)
	var name string
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	rows, _ := b.DB.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		return name
	}
	return "-1"
}

func (b *BottleBase) UpdateLocation(id int, x float64, y float64) {
	sqlStatement := `
UPDATE users
SET locx = $2, locy = $3
WHERE userid = $1;`
	_, err := b.DB.Exec(sqlStatement, id, x, y)
	if err != nil {
		panic(err)
	}

}

func (b *BottleBase) CreateBottleMessage(sendid int, message string) {
	sqlStatement := `
	INSERT INTO outgoingbottles (sendid, message)
	VALUES ($1, $2)`
	_, err := b.DB.Exec(sqlStatement, sendid, message)

	if err != nil {
		panic(err)
	}
}

type Bottle struct {
	senderID   int
	secret     string
	senderName string
}

func (b *BottleBase) DrinkBottleMessage(p Point) Bottle {
	sqlStatement := fmt.Sprintf(`SELECT outgoingbottles.bottleid, outgoingbottles.sendid, users.locx, users.locy, users.locationEnabled FROM outgoingbottles INNER JOIN users ON outgoingbottles.sendid=users.userid`)

	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	rows, _ := b.DB.Query(sqlStatement)

	defer rows.Close()

	var x, y float64
	var bottleID, sendID int
	var locationEnabled bool

	for rows.Next() {
		err := rows.Scan(&bottleID, &sendID, &x, &y, &locationEnabled)
		if err != nil {
			panic(err)
		}

		//construct q to be evaluated by p
		q := Point{
			lat:          x,
			long:         y,
			id:           bottleID,
			locationless: !locationEnabled,
		}

		if p.closestID == 0 {
			p.closestID = q.id

			p.EvaluateQ(q)
			continue
		}

		p.EvaluateQ(q)

	}

	bottle := Bottle{}

	return bottle
	//get you closest message
	//or the hundreth bottle
}

func (b *BottleBase) CreateOutgoingChatMessage(sendid int, receiveid int, message string, counter int) {
	sqlStatement := `
	INSERT INTO outgoingchats (sendid, receiveid, message, counter)
	VALUES ($1, $2, $3, $4)`
	_, err := b.DB.Exec(sqlStatement, sendid, receiveid, message, counter)

	if err != nil {
		panic(err)
	}
}

//pullbottlemessage
//getchats

func (b *BottleBase) ToggleLocationEnabled(userID int, enabled bool) {
	sqlStatement := `
	UPDATE users
	SET locationenabled = $2
	WHERE userid = $1;`
	_, err := b.DB.Exec(sqlStatement, userID, enabled)
	if err != nil {
		panic(err)
	}
}

func (b *BottleBase) Ping() error {
	err := b.DB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Happy Ping")
	return nil
}
