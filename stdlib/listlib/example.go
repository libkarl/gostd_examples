package listlib

import (
	"container/list"
	"fmt"
)

func ListF() {
	// Create a new list and put some numbers in it.
	l := list.New()
	// přidá 4 na konec listu
	e4 := l.PushBack(4)
	// přidá 1 na začátek listu
	e1 := l.PushFront(1)
	// přidej 3 před e4 proměnnou v listu
	l.InsertBefore(3, e4)
	// přidej 2 po e1 proměnné v listu
	l.InsertAfter(2, e1)
	e5 := l.PushBack(5)
	l.InsertAfter(6, e5)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
