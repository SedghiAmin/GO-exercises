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

func (l *List[T]) toSlice() []T {
	
    if l == nil {
        return nil
    }
    
    var result []T
    current := l

    for current != nil {
        result = append(result, current.Node)
        current = current.Next
    }

    return result
}

func main(){
	/* var emptyList *List[int]
    fmt.Println(emptyList.toSlice()) // [] or nil */

	intList := &List[int]{Node: 10}
	intList.Add(20)
	intList.Add(30)

	fmt.Println(intList.toSlice())
}