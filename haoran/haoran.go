package haoran

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const READMEPATH = "haoran/lunyu/README.md"
const CONTENTPATH = "haoran/lunyu/lunyu.json"

type LunYu struct {
	Chapter    string   `json: chapter`
	Paragraphs []string `json: paragraphs`
}

var LUNYUALL []LunYu

func init() {
	content, err := readFile(CONTENTPATH)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(content), &LUNYUALL)
	if err != nil {
		log.Fatal(err)
	}
}

func Readme() (string, error) {
	return readFile(READMEPATH)
}

func Chapters() []string {
	var res []string
	for _, v := range LUNYUALL {
		res = append(res, v.Chapter)
	}
	return res
}

func Paragraphs() []string {
	var res []string
	for _, v := range LUNYUALL {
		res = append(res, v.Paragraphs...)
	}
	return res
}

func ReadContent(filePath string) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		line = strings.TrimSpace(line)
		fmt.Println(line)
		return nil
	}
}

func readFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	return string(content), err
}
