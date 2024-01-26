package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("C:\\Users\\grosy\\Desktop\\17.txt")
	if err != nil {
		panic(err)
	}
	s := strings.Fields(string(f))
	var list_int []int
	for i := 0; i < len(s); i++ {
		a, err := strconv.Atoi(s[i])
		if err != nil {
			panic(err)
		}
		list_int = append(list_int, a)

	}
	count := 0
	mx_pair := 0
	for i := 0; i < len(list_int)-1; i++ {
		for j := i + 1; j < len(list_int); j++ {
			if ((list_int[i]%7 == 0) || (list_int[j]%7 == 0)) && ((list_int[i] % 160) != list_int[j]%160) {
				count++
				if list_int[i]+list_int[j] > mx_pair {
					mx_pair = list_int[i] + list_int[j]
				}
			}
		}
	}
	fmt.Println(count, "\n", mx_pair)

}
