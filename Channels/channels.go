package Channels

import ("time"
"math/rand")

func PrintAfter(ch chan int) {
	n :=  rand.Intn(3000)
	time.Sleep(time.Duration(n)* time.Millisecond)
	ch <- n
}