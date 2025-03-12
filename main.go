package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/sudoku", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Accesable modes are with 9, 18 and 27 cues! All are accesible by adding number of cues like so /sudoku/numberOfCues"))
	})

	var sudoku [9][9]int

	for i := range 9 {
		for j := range 9 {
			sudoku[i][j] = 0
		}
	}

	router.HandleFunc("/sudoku/9", func(w http.ResponseWriter, r *http.Request) {

		queries := r.URL.Query()

		for pos, values := range queries {

			x, err := strconv.Atoi(string(pos[1]))
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(string(pos[3]))
			if err != nil {
				panic(err)
			}

			for _, value := range values {
				v, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}

				w.Write([]byte(fmt.Sprintf("X: %d, Y: %d, Value: %d \n", x, y, v)))

				sudoku[x][y] = v
				//TODO translate sudoku to string

				w.Write([]byte(fmt.Println(sudoku)))
			}

		}

	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
