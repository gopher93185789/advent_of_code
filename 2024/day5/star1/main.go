package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../testdata.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.Contains(txt, "|"){
			parts := strings.Split(txt, "|")
			n1, err := strconv.Atoi(parts[0])
			n2, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalln(err)
			}

			m1[n1] = n2
			m2[n2] = n1
		}
	}

	fmt.Println(m1,m2)


}