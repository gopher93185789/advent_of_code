package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func horizontal(data [][]rune) bool {
	return false
}
func vertical(data [][]rune) bool {
	return false
}
func diagonal(data [][]rune) bool {
	return false
}

func parseData(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	m1 := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		m2 := make([]rune, 0)
		for _, v := range txt {
			m2 = append(m2, v)
		}

		m1 = append(m1, m2)
	}

	return m1
}

func main() {

	m1 := parseData("../testdata.txt")

	for _, l := range m1 {
		for _, r := range l {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}