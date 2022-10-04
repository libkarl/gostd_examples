package ringstd

import (
	"container/ring"
	"fmt"
)

func RingEx1Do() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	// volá funci s argumentem p nad každým elementem v rings v v pořadí od začátku do konce
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

}

func RingLen() {
	// Create a new ring of size 4
	r := ring.New(4)

	// Print out its length
	fmt.Println(r.Len())

}

func RingMove() {
	// Create a new ring of size 5
	r := ring.New(6)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Move the pointer forward by three steps
	// přesune všechno od dvojky nahoru včetně
	r = r.Move(2)

	// Iterate through the ring and print its contents
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

}

func RingNext() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	for j := 0; j < n; j++ {
		fmt.Println(r.Value)
		r = r.Next()
	}

}

func RingPrev() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring backwards and print its contents
	for j := 0; j < n; j++ {
		r = r.Prev()
		fmt.Println(r.Value)
	}

}

func RingUnlink() {
	// Create a new ring of size 6
	r := ring.New(6)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Unlink three elements from r, starting from r.Next()
	r.Unlink(3)

	// Iterate through the remaining ring and print its contents
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

}
