package main

import (
	helper111 "github.com/xeviar/myniceprograme/helper"
	"log"
)

const numPool = 10
func CalculateValue(intChan chan int) {
	randomNumber := helper111.RandomNumber(numPool)

	intChan <- randomNumber
}

func main(){
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	value := <-intChan

	log.Println("测试：received from channel", value)



}
