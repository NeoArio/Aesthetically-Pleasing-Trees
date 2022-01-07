package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func main(){
	var garden Garden

	err := json.Unmarshal([]byte(os.Args[1]), &garden)
	if err != nil {
		log.Fatal("incorrect input. Please try again and get help from README")
	}

	start := time.Now()
	answer := garden.SolveByFirstSolution()
	elapsed := time.Since(start)
	log.Printf("First solution answer: %v", answer)
	log.Printf("First solution execution time: %v", elapsed)

	start = time.Now()
	answer = garden.SolveBySecondSolution()
	elapsed = time.Since(start)
	log.Printf("Second solution answer: %v", answer)
	log.Printf("Second solution execution time: %v", elapsed)
}


