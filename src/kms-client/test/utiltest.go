package test

import (
"fmt"
"math/rand"
"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func TestRandomNumber() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomNum := random(00001, 99999)
		fmt.Println("Random Num:", randomNum)
	}
}