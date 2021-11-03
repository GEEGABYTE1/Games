package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is a sample")

	rand.Seed(time.Now().UnixNano())

}
