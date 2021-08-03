package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/skip2/go-qrcode"
)

func main() {
	var filePath string
	var content string
	flag.StringVar(&filePath, "f", "", "file path")
	flag.StringVar(&content, "c", "", "content")
	flag.Parse()

	var qrContent string
	if filePath != "" {
		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
        qrContent = string(bytes)
	} else if content != "" {
		qrContent = content
	} else {
		fmt.Println("Please input -f xxx or -c xxx")
		return
	}

	q, err := qrcode.New(qrContent, qrcode.Medium)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q.ToSmallString(false))
}
