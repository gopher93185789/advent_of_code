package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)


func getIndexes(data, delim string) []int {
    indexes := []int{}
    start := 0
    for {
        in := strings.Index(data[start:], delim)
        if in == -1 {
            break
        }
        indexes = append(indexes, start+in)
        start = start + in + len(delim)
    }
    return indexes
}

func editData(data []byte, donts, dos []int) []byte {
    // Loop through both donts and dos slices
    for _, i := range dos {
        for _, j := range donts {
            if i <= j {
                continue
            }

			fmt.Println(j, i)
			// fmt.Println(string(data[j:i]))
			break
        }
    }

    return data
}

func main() {
	data, err := os.ReadFile("../testdata.txt")
	if err != nil {
		log.Fatalln(err)
	}
	donts := getIndexes(string(data), "don't()")
	dos := getIndexes(string(data), "do()")
	slices.Sort(donts)
	slices.Sort(dos)
	editData(data, donts, dos)



	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(string(data), -1)

	total := 0

	for _, v := range matches {
		n1, err := strconv.Atoi(v[1])
		n2, err := strconv.Atoi(v[2])
		if err != nil {
			log.Fatalln(err)
		}

		total += n1 * n2
	}

	fmt.Println(total)
}