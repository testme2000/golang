package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Maximum insert")
	result := max_insert(3234, 3)
	fmt.Println(result)
}

func max_insert(whole_num, digit_to_add int) int {
	var final_num int = 0

	is_negative := false
	string_val := strconv.Itoa(whole_num)
	digit_string := strconv.Itoa(digit_to_add)
	if whole_num < 0 {
		is_negative = true
	}

	if !is_negative {
		for index, val := range string_val {
			real_int, _ := strconv.Atoi(string(val))
			fmt.Println(index, real_int)
			if digit_to_add >= real_int {
				start_str := string_val[:index]
				end_str := string_val[index:]
				result1, err := strconv.Atoi(start_str + digit_string + end_str)
				if err == nil {
					final_num = result1
					break
				}
			}
		}
	} else {
		digit_inserted := false
		for index, val := range string_val {
			fmt.Println(index, val)
			if digit_to_add <= int(val) {
				start_str := string_val[:index]
				end_str := string_val[index:]
				result1, err := strconv.Atoi(start_str + digit_string + end_str)
				if err == nil {
					final_num = result1
					digit_inserted = true
					break
				}
			}
		}
		if !digit_inserted {
			string_val += digit_string
			result1, err := strconv.Atoi(string_val)
			if err == nil {
				final_num = result1
			}
		}
	}

	return final_num
}
