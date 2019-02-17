package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=kirk password=testing123 dbname=bottled sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	bottleDB := BottleBase{
		DB: db,
	}

	//	bottleDB.Ping()
	/*bottleDB.CreateUser("Bruno", 4, 4.2, true)

	bottleDB.CreateUser("Kirk", 4.1, 4.2, true)

	bottleDB.CreateUser("Rob", 3.99, 4.2, true)

	bottleDB.CreateUser("Camilal", 4.1, 4.15, true)
	*/

	//bottleDB.UpdateLocation(16, 3, 3)
	//bottleDB.ToggleLocationEnabled(16, false)

	bottleDB.CreateUser("Chrissie", 4.05, 4.042, true)
	//bottleDB.CreateBottleMessage(21, "hey here's a secret")

	//bottleDB.CreateBottleMessage(23, "hey here's a secret")

	//bottleDB.CreateBottleMessage(22, "hey here's a secret")

	bottleDB.CreateBottleMessage(25, "hey here's a secret")
	//bottleDB.CreateOutgoingChatMessage(13, 16, "hadsafjdgkljasdgjklasdgjkljkdlible lonely pers", 3)
	p := Point{
		long: 4.05,
		lat:  4.04,
	}

	bottleDB.DrinkBottleMessage(p)
	fmt.Print(bottleDB.GetUsers(9))
}
