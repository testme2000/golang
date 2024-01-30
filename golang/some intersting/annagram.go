package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Annagram example in go")
	result := is_annagram2("I am Meku", "am I Meku")
	fmt.Println(result)
}

func is_annagram(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	one := strings.Split(first, "")
	sort.Strings(one)
	final_one := strings.Join(one, "")

	two := strings.Split(second, "")
	sort.Strings(two)
	final_second := strings.Join(two, "")
	if final_one != final_second {
		return false
	}
	return true
}

func is_annagram2(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	char_count := make(map[string]int)
	for _, char_detail := range first {
		_, status := char_count[string(char_detail)]
		if status {
			char_count[string(char_detail)] += 1
		} else {
			char_count[string(char_detail)] = 1
		}
	}
	fmt.Println(char_count, first)
	for _, char_detail := range second {
		_, status := char_count[string(char_detail)]
		if !status {
			return false
		}
		char_count[string(char_detail)] -= 1
	}
	fmt.Println(char_count, second)
	return true
}
