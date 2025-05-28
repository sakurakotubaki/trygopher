package main

import "fmt"

func logger() {
	fmt.Println("logger")
}

func loop() {
	// 3 for loop
	for i := 0; i < 3; i++ {
		logger()
	}
}

func main() {
  logger()
}
