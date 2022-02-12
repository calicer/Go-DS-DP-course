package Five

import ("fmt"
"time"
)


func PrintFive(ch chan int){
	for {

		time.Sleep(time.Duration(5000)* time.Millisecond)
		ch<- 5
	//	Read(ch)
	}
}

func PrintTen(ch chan int){
	for {

		time.Sleep(time.Duration(1000)* time.Millisecond)
		ch<- 10
	//	Read(ch)
	}
}

func Read(chA, chB <-chan int){
	var n int
	for {
	select {
	case n = <- chA:
		fmt.Printf("%d chanA\n",n)
	case n = <- chB:
		fmt.Printf("%d chanB\n",n)
	}
}
}