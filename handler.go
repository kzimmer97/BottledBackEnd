package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	B BottleBase
}

func EncodeB(bo Bottle) []byte {

	user := &bo
	b, _ := json.Marshal(user)

	return b
}

func EncodeC(bo Message) []byte {

	user := &bo
	b, _ := json.Marshal(user)

	fmt.Println(string(b))

	return b
}

func (h Handler) CreateBottle(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("secret")
	i, _ := strconv.Atoi(r.FormValue("senderID"))

	bottle := Bottle{
		secret:   s,
		senderID: i,
		life:     3,
	}

	h.B.CreateNewBottle(bottle)
}

func (h Handler) MakeName(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("name")

	x, _ := strconv.ParseFloat(r.FormValue("x"), 64)

	jj := r.FormValue("y")
	y, _ := strconv.ParseFloat(jj, 64)
	e := r.FormValue("e")

	var v bool

	if e == "1" {
		v = true
	}
	if e == "0" {
		v = false
	}

	h.B.CreateUser(n, x, y, v)

	rr := h.B.GetUserID()

	w.Write([]byte(fmt.Sprintf("%d", rr)))
}

func (h Handler) GetBottles(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.FormValue("userID"))

	bs := h.B.GrabSomeBottles(n)

	z := EncodeB(bs[0])

	w.Write([]byte(z))
	//COME BACK HERE
}

func (h Handler) SendChat(w http.ResponseWriter, r *http.Request) {
	to, _ := strconv.Atoi(r.FormValue("from"))
	from, _ := strconv.Atoi(r.FormValue("to"))
	m := r.FormValue("message")
	c, _ := strconv.Atoi(r.FormValue("count"))

	h.B.CreateOutgoingChatMessage(to, from, m, c)

	w.Write([]byte("received"))

}

func (h Handler) GetNewMessages(w http.ResponseWriter, r *http.Request) {

	n, _ := strconv.Atoi(r.FormValue("id"))

	m := h.B.PullNewMessages(n)

	for _, u := range m {
		fmt.Println(EncodeC(u))
	}

	w.Write([]byte("welcome"))
}
