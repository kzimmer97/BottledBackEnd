package main

type Bottle struct {
	bottleID int
	senderID int

	life  int
	broke bool

	secret       string
	senderName   string
	senderLetter string

	location Point
}

//clear bottle from the system to save Ram
func (b *Bottle) BreakBottle() {
	b.broke = true

	b.senderID = 0
	b.secret = ""
	b.senderName = ""
	b.location = Point{}
}

func (b *Bottle) DamageBottle() {
	b.life = b.life - 1

	if b.life == 0 {
		b.BreakBottle()
	}
}
