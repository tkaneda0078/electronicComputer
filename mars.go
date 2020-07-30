package main

import "fmt"

func main() {
	fmt.Print("火星での体重は")
	fmt.Print(63 * 0.3783) // 火星での体重は地球での体重の37.83%
	fmt.Print("キロ。年齢は")
	fmt.Print(28 * 365 / 687) // 火星が太陽の周りを一周するのに地球の日付で687日ほどかかる
	fmt.Print("歳になるでしょう。")
}
