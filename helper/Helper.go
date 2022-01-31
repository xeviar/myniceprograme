package helper111

import (
	"math/rand"
	"time"
	"fmt"
)

type SomeStruct struct {
	Firstname string
	Lastname string
}


func RandomNumber(n int) int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(n)
	fmt.Printf("return %d", value)

	return value
}
