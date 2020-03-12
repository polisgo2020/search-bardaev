package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// В первом параметре передавать название выходного файла, дальше названия входных файлов

func main()  {

	var args []string = os.Args
	if len(args) < 3 {
		log.Fatal("No input data")
	}

	var outFile string = args[1]
	var argLen int = len(args) - 2
	inputFile := make([]string, argLen)
	str := make([]string, argLen)

	for i := 0; i < argLen; i++ {
		inputFile[i] = args[i + 2]
	}

	for i := 0; i < argLen; i++ {
		data, err := ioutil.ReadFile(inputFile[i])
		check(err)
		str[i] = string(data)
	}

	docs := make([][]string, len(str))

	for i := 0; i < len(str); i++ {
		docs[i] = strings.Fields(str[i])
	}

	m := make(map[string]map[int]int)

	for i := 0; i < len(docs); i++ {
		for j := 0; j < len(docs[i]); j++ {
			if m[docs[i][j]] != nil {
				index := m[docs[i][j]]
				index[i] = i
			} else {
				m[docs[i][j]] = make(map[int]int)
				m[docs[i][j]][i] = i

			}
		}
	}

	fmt.Println(str)

	file, err := os.Create(outFile)
	check(err)
	defer file.Close()

	for key, value := range m {

		var temp string
		for k, _ := range value {
			temp = temp + strconv.Itoa(k) + " "
		}

		_, err := file.WriteString(key + " " + temp + "\n")
		check(err)
		e := file.Sync()
		check(e)
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}