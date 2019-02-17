package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=kirk password=testing123 dbname=bottled sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	bottleDB := *NewBottleBase(db)

	handler := Handler{
		B: bottleDB,
	}

	r := chi.NewRouter()
	r.Post("/bottle", handler.CreateBottle) //
	r.Get("/bottle", handler.GetBottles)    //format to json
	r.Post("/send", handler.SendChat)       //
	r.Get("/new", handler.GetNewMessages)   //done
	r.Post("/name", handler.MakeName)       //

	http.ListenAndServe(":3000", r)

	//	bottleDB.Ping()
	/*bottleDB.CreateUser("Bruno", 4, 4.2, true)

	bottleDB.CreateUser("Kirk", 4.1, 4.2, true)

	bottleDB.CreateUser("Rob", 3.99, 4.2, true)

	bottleDB.CreateUser("Camilal", 4.1, 4.15, true)

	//bottleDB.UpdateLocation(16, 3, 3)
	//bottleDB.ToggleLocationEnabled(16, false)

	bottleDB.CreateUser("Chrissie", 4.05, 4.042, true)*/
	//bottleDB.CreateBottleMessage(21, "hey here's a secret")

	/*b := Bottle{
		senderID: 81,
		secret:   "xxxx",
		life:     3,
	}
	*/
	//	bottleDB.CreateNewBottle(b)

	//bottleDB.CreateBottleMessage(22, "hey here's a secret")

	//	bottleDB.ArchiveBottle(25, "hey here's another secret mannn")
	//bottleDB.CreateOutgoingChatMessage(13, 16, "hadsafjdgkljasdgjklasdgjkljkdlible lonely pers", 3)

	//collect the bottles
	//give requested bottles for the userID provided

	//	fmt.Print(bottleDB.GetUsers(77))
}
