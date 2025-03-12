package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Sudoku struct {
	board [9][9]int
}

func isSafe(sudoku *Sudoku, y int, x int, value int) bool {
	for i := range 9 {
		if sudoku.board[y][i] == value {
			return false
		}
	}

	for i := range 9 {
		if sudoku.board[i][x] == value {
			return false
		}
	}

	startX := x - (x % 3)
	startY := y - (y % 3)

	for i := range 3 {
		for j := range 3 {
			if sudoku.board[i+startY][j+startX] == value {
				return false
			}
		}
	}

	return true
}

func solveSudoku(sudoku *Sudoku, y int, x int) bool {
	if y == 9 {
		return true
	}

	if x == 9 {
		return solveSudoku(sudoku, y+1, 0)

	}

	if sudoku.board[y][x] != 0 {
		return solveSudoku(sudoku, y, x+1)
	}

	for value := 1; value <= 9; value++ {
		if isSafe(sudoku, y, x, value) {
			sudoku.board[y][x] = value
			if solveSudoku(sudoku, y, x+1) {
				return true
			}
			sudoku.board[y][x] = 0
		}
	}

	return false
}

func solveSudokuFR(sudoku *Sudoku) {
	solveSudoku(sudoku, 0, 0)

}

var previousParams = ""

func main() {
	var sudoku Sudoku

	router := http.NewServeMux()
	router.HandleFunc("/sudoku", func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		currentParams := fmt.Sprintf("%v", queries)

		if currentParams != previousParams {
			for i := range 9 {
				for j := range 9 {
					sudoku.board[i][j] = 0
				}
			}
		}

		for pos, values := range queries {

			x, err := strconv.Atoi(string(pos[3]))
			if err != nil {
				panic(err)
			}
			x--

			y, err := strconv.Atoi(string(pos[1]))
			if err != nil {
				panic(err)
			}
			y--

			for _, value := range values {
				v, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}

				sudoku.board[x][y] = v

				w.Header().Set("Content-Type", "text/plain")

			}
		}

		fmt.Fprintf(w, "Unsolved puzzle: \n")

		for x := range 9 {
			for y := range 9 {
				fmt.Fprintf(w, "%d ", sudoku.board[x][y])
			}
			fmt.Fprintf(w, "\n")
		}

		solveSudokuFR(&sudoku)

		fmt.Fprintf(w, "\nSolved puzzle: \n")

		for x := range 9 {
			for y := range 9 {
				fmt.Fprintf(w, "%d ", sudoku.board[x][y])
			}
			fmt.Fprintf(w, "\n")
		}

		previousParams = currentParams

	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
