package main

import "fmt"

func tenLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	tenLoop()
}
