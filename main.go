package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 100
	fmt.Printf("数字,保留0位小数,保留1位小数,保留2位小数,保留3位小数,保留4位小数,保留5位小数\n")
	for i := 0; i < count; i++ {
		var r = float32(rand.Intn(100)) + rand.Float32()
		fmt.Println(fmt.Sprintf("%.6f", r))
	}
}
