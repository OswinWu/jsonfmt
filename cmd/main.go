package main

import (
	"bufio"
	"flag"
	"fmt"
	"jsonfmt/pkg/hilight"
	"os"

	"github.com/bytedance/sonic"
)

func main() {
	// get args from input
	if len(os.Args) != 2 {
		fmt.Println("Usage: json-pretty <file>")
		os.Exit(1)
	}

	var keyword, color string
	flag.StringVar(&keyword, "keyword", "", "highlight keyword")
	flag.StringVar(&color, "color", "Red", "hilight color")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputBytes := scanner.Bytes()
		str := prettyPrintJson(inputBytes)
		if str == "" {
			fmt.Println(inputBytes)
			continue
		}

		if keyword != "" {
			str = hilight.HighlightKeyword(keyword, str, color)
		}

		fmt.Println(str)
	}
}

func prettyPrintJson(json []byte) string {
	var row map[string]interface{}
	err := sonic.Unmarshal(json, &row)
	if err != nil {
		fmt.Println(err)
	}
	pretty, err := sonic.MarshalIndent(row, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	return string(pretty)
}
