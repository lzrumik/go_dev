package queue

//An FIFO queue
type Queue []int

// Pushes the element into the queue
// 		e.g. q.Push(123)
func (q *Queue) Push (v int){
	*q = append(*q,v)
}

// Pops element from head .
func (q *Queue)Pop()int{
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty or not
func (q *Queue) IsEmpty()bool{
	return len(*q) == 0
}


/**
进入目录
go doc
go doc queue
go doc IsEmpty
godoc -http :6060

访问 http://localhost:6060
http://localhost:6060/pkg/ch42_go_doc/01queue/
 */