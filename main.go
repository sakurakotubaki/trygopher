package main

import "fmt"

func main() {
	// Slice
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	// 要素の追加
	n = append(n, 100)
	fmt.Println(n)
	// 要素の出力
	fmt.Println(n[3])
}
