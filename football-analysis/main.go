package main

import (
	"fmt"
	"os"
)

func main() {
	str, _ := os.Getwd()
	fmt.Println("当前目录：", str)
}
