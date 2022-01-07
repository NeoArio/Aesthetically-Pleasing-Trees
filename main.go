package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func main(){
	start := time.Now()
	var garden Garden

	err := json.Unmarshal([]byte(os.Args[1]), &garden)
	if err != nil {
		log.Fatal("incorrect input. Please try again and get help from README")
	}

	answer := garden.Solve()
	log.Printf("answer: %v", answer)

	elapsed := time.Since(start)
	log.Printf("Execution time: %v", elapsed)

}


