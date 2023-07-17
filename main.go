package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{"This code was written by Abulfazl for company members to learn", "I love coding", "I am proud of golang"}
	mostWordsFount(sentences)
}

func mostWordsFount(sentences []string) {
	max := 0

	for _, word := range sentences {
		lenght := len(strings.Split(word, " "))

		if lenght > max {
			max = lenght
		}
	}
	fmt.Println(max)

}
