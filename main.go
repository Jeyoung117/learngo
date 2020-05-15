package main

import (
	"fmt"

	"github.com/Jeyoung117/lerango/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "First")
	word1, err1 := dictionary.Search(baseWord)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(word1)

	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
