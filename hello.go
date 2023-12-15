package main

import "fmt"

func main() {
	myList := List[int]{}		
	myList.push(1)
	myList.push(2)

	result := myList.getAll()

	fmt.Println("hasil :", result[0])

}

type List[T any] struct{
	head, tail *element[T]
}

type element[T any] struct{
	next *element[T]
	value T
}

func (ll *List[T]) push(val T){
	if ll.head == nil {
		ll.head = &element[T]{value: val}
		ll.tail = ll.head
	}else{
		ll.tail.next = &element[T]{value: val}
		ll.tail = ll.tail.next
	}	
}

func (ll *List[T]) getAll() []T{
	result := make([]T,0)
	for e:=ll.head; e != nil; e = e.next{
		result = append(result, e.value)

	}
	return result
}

