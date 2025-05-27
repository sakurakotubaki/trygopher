package main

import "fmt"

/*
シリコンバレー一流プログラマーが教える
Go大全
p60の学習
*/

func bazz() {
	fmt.Println("bazz")
}

func init() {
	bazz()
}

func main() {
	// Slice
	n := []int{1, 2, 3}
	fmt.Println(n)
	// 要素の追加
	n = append(n, 100)
	fmt.Println(n)
	// 要素の出力
	fmt.Println(n[3])
}
