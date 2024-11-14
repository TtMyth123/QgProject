package main

import (
	"fmt"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/sysInit"
	"os"
	//"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/sysInit"
)

func main() {
	str, _ := os.Getwd()
	fmt.Println("当前目录：", str)
	sysInit.Init()
	fmt.Println("结束:", str)
}
