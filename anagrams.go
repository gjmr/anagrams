package main

import (
	"fmt"
	"sort"
	"strings"

	load_dictionary "anagrams/load"
)

var word_list = load_dictionary.Load("2of4brif.txt")

func main() {
	var anagram_list []string
	name := "Foster"
	fmt.Println("Input name = ", name)
	name = strings.ToLower(name)
	fmt.Println("Using name = ", name)

	name_sorted := sorted(name)

	for _, word := range word_list {
		word = strings.ToLower(word)
		if word != name {
			if sorted(word) == name_sorted {
				anagram_list = append(anagram_list, word)
			}
		}
	}

	if len(anagram_list) == 0 {
		fmt.Println("You need a larger dictionary or a new name!")
	} else {
		fmt.Println("Anagrams = ", anagram_list)
	}
}

func sorted(word string) string {
	word_split := strings.Split(word, "")
	sort.Strings(word_split)
	word_sorted := strings.Join(word_split, "")
	return word_sorted
}
