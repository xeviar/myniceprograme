package helper111

import (
	"math/rand"
	"time"
)

type SomeStruct struct {
	Firstname string
	Lastname string
}


func RandomNumber(n int) int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(n)


	return value
}