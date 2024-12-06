package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
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

func moveUp(board [][]rune, x, y int) bool {
	if y == 0 {
		return true
	}
	
	board[y][x] = 'X'
	if board[y-1][x] != '#' {
		board[y-1][x] = '^'
	}else {
		board[y-1][x] = '>'
	}


	return false
}

func moveDown(board [][]rune, x,y int) bool {
	if y == len(board) - 1{
		return true
	}
	
	board[y][x] = 'X'
	if board[y+1][x] != '#' {
		board[y+1][x] = 'v'
	}else {
		board[y+1][x] = '<'
	}

	return false
}

func moveLeft(board [][]rune,x,y int) bool {
	if x == 0{
		return true
	}
	
	board[y][x] = 'X'
	if board[y][x-1] != '#' {
		board[y][x-1] = '<'
	}else {
		board[y+1][x] = '^'
	}

	return false
	
}

func moveRight(board [][]rune, x,y int) bool {
	if x == len(board[y]) - 1{
		return true
	}
	
	board[y][x] = 'X'
	if board[y][x+1] != '#' {
		board[y][x+1] = '>'
	}else {
		board[y+1][x] = 'v'
	}

	return false
	
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
	for {
		all:
		for y, row := range board {
			for x, char := range row {
				switch char {
				case '^':
					if moveUp(board, x, y) {
						break all
					}

					break
				case '>':
					if moveRight(board, x, y) {
						break all
					}

					break
	
				case '<':
					if moveLeft(board, x, y) {
						break all
					}

					break
	
				case 'v':
					if moveDown(board, x, y) {
						break all
					}

					break
				}

				printBoard(board)
				fmt.Println("")
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}

}