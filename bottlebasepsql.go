package main

import "fmt"

func (b *BottleBase) CreateUser(name string, x float64, y float64, locationEnabled bool) {
	sqlStatement := `
	INSERT INTO users (name, lat, long, locationenabled)
	VALUES ($1, $2, $3,$4)`
	b.db.Exec(sqlStatement, name, x, y, locationEnabled)

}

func (b *BottleBase) GetUsers(userid int) string {
	sqlStatement := fmt.Sprintf(`SELECT name FROM users WHERE userid=%d`, userid)
	var name string
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	rows, _ := b.db.Query(sqlStatement)

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

func (b *BottleBase) GetBottleID() int {
	sqlStatement := fmt.Sprintf(`SELECT bottleid FROM outgoingbottles ORDER BY bottleID ASC`)
	var bottleID int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bottleID)
		if err != nil {
			panic(err)
		}
		return bottleID
	}

	return -1
}

func (b *BottleBase) GrabBottleID(s int) int {
	sqlStatement := fmt.Sprintf(`SELECT bottleid FROM outgoingbottles WHERE senderid=%d`, s)
	var bottleID int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&bottleID)
		if err != nil {
			panic(err)
		}
		return bottleID
	}

	return -1
}

func (b *BottleBase) GetUserID() int {
	sqlStatement := fmt.Sprintf(`SELECT userID FROM users ORDER BY userID ASC`)
	var userID int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			panic(err)
		}
		return userID
	}

	return -1
}

func (b *BottleBase) UpdateLocation(id int, x float64, y float64) {
	sqlStatement := `
UPDATE users
SET lat = $2, long = $3
WHERE userid = $1;`
	_, err := b.db.Exec(sqlStatement, id, x, y)
	if err != nil {
		panic(err)
	}

}

func (b *BottleBase) ArchiveBottle(sendid int, message string, hearts int) {
	sqlStatement := `
	INSERT INTO outgoingbottles (sendid, message, hearts)
	VALUES ($1, $2, $3)`
	_, err := b.db.Exec(sqlStatement, sendid, message, hearts)

	if err != nil {
		panic(err)
	}
}

func (b *BottleBase) PullNewMessages(receiveid int) []Message {
	sqlStatement := fmt.Sprintf("SELECT sendid, message, counter FROM ougoingMessages WHERE receiveid=%d", receiveid)

	var s int
	var m string
	var c int

	rows, _ := b.db.Query(sqlStatement)

	defer rows.Close()

	var mu []Message

	for rows.Next() {
		err := rows.Scan(&s, &m, &c)
		if err != nil {
			panic(err)
		}

		mess := Message{
			SenderID:   s,
			ReceiverID: receiveid,
			message:    m,
			counter:    c,
		}

		mu = append(mu, mess)
	}

	sqlStatement = fmt.Sprintf("Delete FROM outgoingchats where receiveid=%d", receiveid)
	_, err := b.db.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}

	return mu
}

type Message struct {
	SenderID   int
	ReceiverID int
	message    string
	counter    int
}

func (b *BottleBase) CreateOutgoingChatMessage(sendid int, receiveid int, message string, counter int) {
	sqlStatement := `
	INSERT INTO outgoingchats (sendid, receiveid, message, counter)
	VALUES ($1, $2, $3, $4)`
	_, err := b.db.Exec(sqlStatement, sendid, receiveid, message, counter)

	if err != nil {
		panic(err)
	}
}

//getchats

func (b *BottleBase) ToggleLocationEnabled(userID int, enabled bool) {
	sqlStatement := `
	UPDATE users
	SET locationenabled = $2
	WHERE userid = $1;`
	_, err := b.db.Exec(sqlStatement, userID, enabled)
	if err != nil {
		panic(err)
	}
}

func (b *BottleBase) Ping() error {
	err := b.db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Happy Ping")
	return nil
}
