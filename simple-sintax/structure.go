package main

import "fmt"

type id struct {
	year  int16
	month string
	date  int8
}

func calend() {
	first_date := id{year: 2020, month: "Feb", date: 2}
	first_sex := id{year: 2020, month: "Feb", date: 9}
	fmt.Println(first_date)
	fmt.Println(first_sex)

	declar_of_love := id{year: 2020, month: "May", date: 30}
	declar_of_love.month = "March"
	declar_of_love.date = 26
	fmt.Println(declar_of_love)

}
