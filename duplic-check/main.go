package main

import (
	"bufio"
	"fmt"
	"os"
)

// -----------

func main() {
	counts := make(map[string]int)
	files := os.Args[1:] //читаем файлы из cmd
	if len(files) == 0 {
		countlines(os.Stdin, counts) //Если файлов нет, то функции скармиливаем текст из ввода в cmd
	} else {
		for _, arg := range files { // пробегаемся по файлам
			f, err := os.Open(arg) // открываем файлы (f - это сам файл, err - показывает что все прочиатлось нормально
			if err != nil {        // если равно nil, то все ок, если нет, то не ок
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countlines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countlines(f *os.File, counts map[string]int) {
	var input = bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++ // тоже самое что: line := input.Text(); counts[line] := counts[line]+1
	}
}
