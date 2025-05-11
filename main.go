package main

import "fmt"

func bazz() {
	fmt.Println("bazz")
}

func init() {
	bazz()
}

func main() {

	bazz()
}
