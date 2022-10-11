package reflectex

import (
	"fmt"
	"reflect"
)

func ReflectExample1() {
	intSlice := make([]int, 0)
	sliceType := reflect.TypeOf(intSlice)
	fmt.Println(sliceType)

	intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)
	v := 10
	rv := reflect.ValueOf(v)
	fmt.Println(rv)
	// The next line will append rv in created intSliceReflect slice
	intSliceReflect = reflect.Append(intSliceReflect, rv)
	intSlice2 := intSliceReflect.Interface().([]int)
	fmt.Println(intSlice2)
}

// second example
func InvertSlice(args []reflect.Value) (result []reflect.Value) {
	// otočí pořadí prvků
	inSlice, n := args[0], args[0].Len()
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, n)
	for i := n - 1; i >= 0; i-- {
		element := inSlice.Index(i)
		outSlice = reflect.Append(outSlice, element)
	}
	return []reflect.Value{outSlice}
}

func Bind(p interface{}, f func([]reflect.Value) []reflect.Value) {
	invert := reflect.ValueOf(p).Elem()

	//Use of MakeFunc() method
	invert.Set(reflect.MakeFunc(invert.Type(), f))
}

// Main function
func ReflectExample2() {
	var invertInts func([]int) []int
	Bind(&invertInts, InvertSlice)
	fmt.Println(invertInts([]int{1, 2, 3, 4, 2, 3, 5}))
}

//
