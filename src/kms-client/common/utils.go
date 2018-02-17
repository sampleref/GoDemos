package common

import (
	"time"
	"math/rand"
	"fmt"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GetRandomNumberId() (randomNum float64) {
	rand.Seed(time.Now().UnixNano())
	randomNum = float64(random(00001, 99999))
	fmt.Println("Random Id Generated: ", randomNum)
	return
}
