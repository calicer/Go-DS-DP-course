package FanIn
import ("fmt"
"time"
"math/rand")

func Producer(ch chan<- int, name string){
	for {
		sleep()
		n :=  rand.Intn(100)

		fmt.Printf("Channel %s -> %d \n", name , n)
		ch <- n

	}

}

func sleep(){
	time.Sleep(time.Duration(rand.Intn(3000))* time.Millisecond)
}

func Consumer(ch <-chan int){
	for n := range ch {
		fmt.Printf("<- %d \n",n)
	}
}

func FanIn(chA, chB <-chan int, chC chan<- int){
	var n int 
	for {
		select {
			case n = <- chA:
				chC<-n
			case n = <- chB:
				chC<-n
		}
	}
}