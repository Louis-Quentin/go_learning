package main

import (
	"fmt"
	"io/ioutil"
)

func LineToCsv(line []byte) ([]byte) {
		for i:=0; i<len(line); i++ {
			if line[i] == ',' {
				line[i] = '\n'
			}
		}
		return line
}

func main() {
	var path = "/home/alouis-quentin/go/src/awesomeProject/Data/data"
	content, _ := read(path)
	finalCsv := LineToCsv(content)
	fmt.Printf("%s", finalCsv)
}

func read(path string) ([]byte, error){
	content, err := ioutil.ReadFile(path)
	if err != nil {
		print("read fail\n")
	}
	return content, err
}