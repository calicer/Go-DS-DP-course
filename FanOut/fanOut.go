package FanOut

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(ch chan<- int) {
	for {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		n := rand.Intn(50)
		ch <- n
	}
}

func Consumer(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf("ch %s -> %d \n",name, n)
	}

}

func FanOut(chA <-chan int, chB, chC chan<- int) {
	var n int
	for {
		select {
		case n = <-chA:
			if n < 25 {
				chB <- n
				fmt.Printf("chB -> %d \n", n)
			} else {
				chC <- n
				fmt.Printf("chC -> %d \n", n)
			}
		}
	}
}
