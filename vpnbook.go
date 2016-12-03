package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"io"
	"github.com/kardianos/osext"
	"log"
	"path"
)

func main() {
	type Cred struct {
		Path string `json:"credPath"`
	}

	folderPath, err5 := osext.ExecutableFolder()
	if err5 != nil {
		log.Fatal(err5)
	}
	configFile:= path.Join(folderPath, "parser-config.json")

	data, err1 := ioutil.ReadFile(configFile)
	if err1 != nil {
		panic(err1)
	}
	var cred Cred
	err := json.Unmarshal(data, &cred)
	if err != nil {
		panic(err)
	}

	filename := cred.Path
	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_TRUNC, os.FileMode(0666))
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocument("http://www.vpnbook.com/freevpn")
	if err != nil {
		fmt.Println(err)
	}
	username := doc.Find(".disc li:nth-child(7) strong").First().Text()
	password := doc.Find(".disc li:nth-child(8) strong").First().Text()

	n, err := io.WriteString(file, username+"\r\n"+password)
	if err != nil {
		fmt.Println(n, err)
	}

	file.Close()
}
