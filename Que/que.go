package Que


type Que struct{
	Item chan int
}

func (q *Que) EnQue(i int){
	q.Item <- i
}

func (q *Que) DeQue() int {
	return <- q.Item
}