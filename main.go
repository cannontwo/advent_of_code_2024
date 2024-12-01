package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var DAY_COMMANDS = map[int]func(){
	1: run_day_one,
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Please provide an advent of code problem to run")
	}

	problem, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Could not convert %v to number: %v", os.Args[1], err)
	}

	fmt.Printf("Running day %v\n", problem)
	df, ok := DAY_COMMANDS[problem]
	if !ok {
		log.Fatalf("No function registered for day %v", problem)
	}

	df()
}
