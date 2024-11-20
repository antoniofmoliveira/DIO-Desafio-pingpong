package main

import (
	"fmt"
	"time"
)

func pingpong(c chan int) {
	for v := range c {
		if v == 0 {
			fmt.Println("ping")
		} else {
			fmt.Println("pong")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan int)
	go pingpong(c)
	for i := 0; i < 10; i++ {
		c <- i % 2
	}
}
