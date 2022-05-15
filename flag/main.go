package main

import (
	"flag"
	"fmt"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func init() {
	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "default", "string flag value")
}

func main() {
	// 在main方法中调用flag.Parse从os.Args[1:]中解析选项。因为os.Args[0]为可执行程序路径，会被剔除。
	flag.Parse()

	fmt.Println("int flag:", intflag)
	fmt.Println("bool flag:", boolflag)
	fmt.Println("string flag:", stringflag)
}

// 编译
// go build -o main.exe main.go

// 执行
// ./main.exe -intflag=12 -boolflag=1 -stringflag=test

// 结果
//int flag: 12
//bool flag: true
//string flag: test
