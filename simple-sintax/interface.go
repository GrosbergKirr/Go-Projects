package main

import (
	"fmt"
)

type gfys struct {
}

type swears interface {
	first()
	second()
	third()
	fourth()
}

func (n gfys) first() {
	fmt.Println("Go and f*ck youself!")
}

func (n gfys) second() {
	fmt.Println("S*ck my ba*s!")
}

func (n gfys) third() {
	fmt.Println("Shut up your a&h*le!")
}

func (n gfys) fourth() {
	fmt.Println("F*ck you stupid nibba!")

}

func curse() {
	var i swears
	var a gfys
	i = a
	i.second()
}
