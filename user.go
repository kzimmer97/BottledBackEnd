package main

import "bottled/utils"

type User struct {
	userID          int
	name            string
	letter          string
	locationEnabled bool
	burntBottles    map[int]Bottle
	stashIDs        []int
	M               utils.MinHeap
	Intitiated      int
	Point
}
