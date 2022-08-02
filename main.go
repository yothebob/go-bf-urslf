package main

import (
  "fmt"
)

type LinkedList struct {
	headNode int
	nextNode *LinkedList
	prevNode *LinkedList
	value int
}

func moveLeft(pi *LinkedList) {
	if pi.headNode == 1{
		pi.headNode = 0
		newHead := new(LinkedList)
		newHead.headNode = 1
		newHead.nextNode = pi
		return &newHead
	} else {
		return &pi.prevNode
	}
}

func moveRight(pi *LinkedList) {
	if pi.nextNode {
		newNode := new(LinkedList)
		pi.nextNode = newNode
		newNode.prevNode = pi
		return &newNode
	} else {
		return &pi.nextNode
	}
}

func addOne (pi *LinkedList) {
	return  pi.value + 1
}

func subOne (pi *LinkedList) {
	return  pi.value - 1 
}


func interpret(code string, pointer *LinkedList){
	for i:=0; i < len(code);i++ {
		fmt.Println(i) // index
		fmt.Println(string(code[i])) // char
		fmt.Println(code[i]) //ascii
		if string(code[i]) == ">" {
			move_right(pointer)
		}
		if string(code[i]) == "<" {
			move_left(pointer)
		}
		if string(code[i]) == "[" {
		}
		if string(code[i]) == "]" {
			return 
		}
		if string(code[i]) == "+" {
		}
		if string(code[i]) == "-" {
		}
		if string(code[i]) == "." {
		}
		if string(code[i]) == "," {
		}
	}
}

func main(){
	fmt.Println("Type Code:")
	var bf_code string
	pointer := new(LinkedList)
	fmt.Scanln(&bf_code)
	interpret(bf_code)
}

