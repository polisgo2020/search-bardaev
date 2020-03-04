package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// В первом параметре передавать путь к файлу, во втором название выходного файла

func main()  {
	data, err := ioutil.ReadFile(os.Args[1])
	check(err)

	var str string = string(data)
	var s []string = strings.Split(str, "\n")
	docs := make([][]string, len(s))

	for i := 0; i < len(s); i++ {
		docs[i] = strings.Split(s[i], " ")
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

	fmt.Println(s)

	file, err := os.Create(os.Args[2])
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
		panic(e)
	}
}