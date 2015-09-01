package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"os"
)

func main() {
	doc, err := goquery.NewDocument("http://www.vpnbook.com/freevpn")
	if err != nil {
		fmt.Println(err)
	}
	username := doc.Find(".disc li:nth-child(7) strong").First().Text()
	password := doc.Find(".disc li:nth-child(8) strong").First().Text()

	filename := "d:\\opt\\login.conf"
	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_TRUNC, os.FileMode(0666))
	if err != nil {
		panic(err)
	}
	n, err := io.WriteString(file, username+"\r\n"+password)
	if err != nil {
		fmt.Println(n, err)
	}

	file.Close()
}
