package main

import (
	"strconv"
	"fmt"
)


func run_loop(code string, list []int, pointer int) {
	fmt.Println(list)
	for {
		if list[pointer] <= 0 {
			return
		} else {
			interpret(code,list,pointer)
		}
	}
}


func interpret(code string, list []int, pointer int) ([]int) {
	for i:=0; i < len(code);i++ {

		switch string(code[i]) {
			
		case "+":
			//plus 1
			list[pointer] = list[pointer] + 1
			
		case "-":
			//minus 1
			list[pointer] = list[pointer] - 1
			
		case ">":
			// shift right
			if pointer > len(list) {
				pointer = pointer + 1
			} else {
				pointer = pointer + 1
				list = append(list, 0)
			}

		case "<":
			// shift left
			if pointer > 0 {
				pointer = pointer - 1
			} else {
				pointer = pointer - 1
				list = append([]int{0}, list...)
			}

		case "[":
			//startloop
			fmt.Println(code[(i+1):])
			run_loop(code[(i+1):],list,pointer)

		case "]":
			break
			//endloop

		case ".":
			//show ascii
			fmt.Println(string(code[i]))

		case ",":
			// get_input
			var input string
			fmt.Scanln(&input)
			if i, err := strconv.Atoi(input); err == nil {
				list[pointer] = list[pointer] + int(i)
			} else {
				// add ascii value
				runes := []rune(input)
				list[pointer] = list[pointer] + int(runes[0])
			}
		}
	}
	return list
}

func main(){
	fmt.Println("Type Code:")
	var bf_code string
	pointer := 0
	list := []int{0}
	fmt.Scanln(&bf_code)
	executed := interpret(bf_code, list, pointer)
	fmt.Println(executed)
}

