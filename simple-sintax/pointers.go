package main

import (
	"fmt"
)

func change(i *int) int {
	*i++
	return *i
}
func nochange(a int) int {
	a++
	return a
}

func Pointcheck(a int) {
	fmt.Println(a)
	fmt.Println(nochange(a))
	fmt.Println(a)

	fmt.Println(a)
	fmt.Println(change(&a))
	fmt.Println(a)

}
