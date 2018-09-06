package main

import "fmt"

func main() {
	fmt.Println(checker("hhaacckkekraraannk"))
}

func checker(word string) string {
	hackerrank := "hackerrank"
	if len(word) < len(hackerrank) {
		return "NO"
	}
	var j int

	for i := 0; i < len(word); i++ {
		if j < len(hackerrank) && string(word[i]) == string(hackerrank[j]) {
			j++
		}
	}

	if  j == len(hackerrank) {
		return "YES"
	}
	return "NO"
}


