package _8queue

type Queue []int

func (q *Queue) Puse (v int){
	*q = append(*q,v)
}

func (q *Queue)Pop()int{
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty()bool{
	return len(*q) == 0
}
