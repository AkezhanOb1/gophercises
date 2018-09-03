package main

import (

	"fmt"
	"sync"
	"strings"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	var message string
	fmt.Println("Enter The Message")
	fmt.Scanln(&message)

	ch := make(chan string, 3)
	result := make(chan int)

	go splitter(ch, message)
	go receiver(ch, result)

	wg.Wait()
	fmt.Println(<-result)
}


func splitter(ch chan <-string, s string) {
	s = strings.ToLower(s)
	for _, letter := range s {
		ch <- string(letter)
	}
	close(ch)
	wg.Done()
}


func receiver(ch <-chan string, result chan <-int ) {
	var s string
	var counter int
	for letter := range ch {
		s += letter
		if len(s) > 2 {
			counter += checker(&s)
		}
	}
	wg.Done()
	result <- counter
}



func checker(s *string) int {
	var counter int
	word := *s
	if string(word[0]) != "s" {
		counter++
	}
	if string(word[1]) != "o" {
		counter++
	}
	if string(word[2]) != "s" {
		counter++
	}
	*s = ""
	return counter
}