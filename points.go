package main

import (
	"math"
	"math/rand"
)

const (
	pi float64 = .017453292519943295
)

type Point struct {
	id int

	locationless bool

	lat  float64
	long float64

	closestDistance float64
	closestID       int
	closestEst      int

	oldestIndex int
	oldestAge   int
}

func (p *Point) EvaluateQ(q Point) {
	//calculate distance
	distance := p.calcDistance(q)

	if p.closestDistance > distance {
		p.closestDistance = distance
		p.closestID = q.id
		p.closestEst = int(distance)
	}

	//run age check
	if p.oldestAge > q.oldestAge {
		p.oldestAge = q.oldestAge
		p.oldestIndex = q.id
	}
}

/* https://stackoverflow.com/questions/27928/calculate-distance-between-two-latitude-longitude-points-haversine-formula
function distance(lat1, lon1, lat2, lon2) {
  var p = 0.017453292519943295;    // Math.PI / 180
  var c = Math.cos;
  var a = 0.5 - c((lat2 - lat1) * p)/2 +
          c(lat1 * p) * c(lat2 * p) *
          (1 - c((lon2 - lon1) * p))/2;

  return 12742 * Math.asin(Math.sqrt(a)); // 2 * R; R = 6371 km
}
*/

func (p *Point) calcDistance(q Point) float64 {
	a := 0.5 - math.Cos((q.lat-p.lat)*pi)/2 + math.Cos(p.lat*pi)*math.Cos(q.lat*pi)*(1-math.Cos((q.long-p.long)*pi))/2
	a = 12742 * math.Asin(math.Sqrt(a))

	return a
}

func (p Point) HasGoodLocationBottle(maxDistance int) bool {
	//if you don't have location enabled, you're not even considered for this
	if p.locationless {
		return false
	}

	if p.closestEst <= maxDistance {
		return true
	}

	return false
}

func (p Point) EvaluateBottleOptions() int {
	if p.locationless {
		return p.oldestIndex
	}

	r := rand.Intn(10)

	//roll dice... there's a 90% chance we give you someone based on location
	//this leaves a 10% chance the oldest chat in the bottle gets viewed

	if r == 0 {
		return p.oldestIndex
	}

	return p.closestID
}
