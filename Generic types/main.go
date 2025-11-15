package main

import(
	"fmt"
)

type List[T any] struct{
	Next *List[T]
	Node T
}

func (l *List[T]) Add(node T) *List[T]{
	newNode := &List[T]{Node: node}

	current := l

	for (current.Next != nil){
		current= current.Next
	}

	current.Next = newNode

	return l

}

func (l *List[T]) toSlice() []T{
	var result []T
	current:= l

	for(current.Next != nil){
		result = append(result, current.Node)
		current = current.Next
	}

	result = append(result, current.Node)

	return result
}
func main(){

	intList := &List[int]{Node: 10}
	intList.Add(20)
	intList.Add(30)

	fmt.Println(intList.toSlice())
}