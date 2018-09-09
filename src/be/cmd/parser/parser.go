package main

import (
	"be/common/log"
	"be/options"
	"be/parser"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: parser [xxx.txt]")
		os.Exit(1)
	}

	// 从命令行、配置文件初始化配置
	options.Options.InitOptions()

	// 初始化Log
	log.InitLog()

	fileName := os.Args[1]
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err.Error())
	}
	content := string(fd)

	p := parser.NewParser()
	err = p.Parser(content)
	if err != nil {
		panic(err.Error())
	}
	parsedResult := p.GetResult()
	b, err := json.MarshalIndent(parsedResult, "", "    ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(b))
}
