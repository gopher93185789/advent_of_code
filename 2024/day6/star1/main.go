package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getBoard(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	board := [][]rune{}
	for scanner.Scan() {
		row := []rune{}
		for _, v := range scanner.Text() {
			row = append(row, v)
		}
		board = append(board, row)
	}

	return board
}

func moveUp(board [][]rune, x, y int) (newx int, newy int, end bool, dirchange rune) {
	if y == 0 {
		return 0, 0, true, '^'
	}
	
	
	if board[y-1][x] != '#' {
		board[y][x] = 'X'
		board[y-1][x] = '^'
	}else {
		board[y][x] = '>'
		return x, y, false, '>'
	}

	return x, y-1, false, '^'
}

func printBoard(board [][]rune) {
	for _, row := range board {
		for _, char := range row {
			print(string(char))
		}
		println()
	}
}



func main() {
	board := getBoard("../testdata.txt")


	var (
		x int
		y int
		p rune
	)
	for y1, row := range board {
		for x1, char := range row {
			if char == '^' || char == '>' || char == '<' || char == 'v' {
				x = x1
				y = y1
				p = char
			}
		}
	}

	printBoard(board)
	fmt.Println()

	for {
		switch p {
		case '^':
			nx, ny, end, dc := moveUp(board, x, y)
			x = nx
			y = ny
			p = dc
			if end {
				break
			}
		default:
			break
		}

	}

	printBoard(board)
}