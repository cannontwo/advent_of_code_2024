package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func is_line_safe(parts []string) bool {
	prev_num := -1
	increasing := true
	safe := true
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		if i == 1 {
			if num > prev_num {
				increasing = true
			} else if num < prev_num {
				increasing = false
			} else {
				safe = false
				break
			}
		}

		if i > 0 {
			var diff int
			if increasing {
				diff = num - prev_num
			} else {
				diff = prev_num - num
			}

			if diff < 1 || diff > 3 {
				safe = false
				break
			}
		}

		prev_num = num
	}

	return safe
}

func is_line_dampener_safe(parts []string) bool {
	// Could speed up by recording the number of unsafe jumps during initial safety scan
	for i := range parts {
		sublist := slices.Concat(parts[0:i], parts[i+1:])
		safe := is_line_safe(sublist)
		if safe {
			return true
		}
	}

	return false
}

func run_day_two() {
	if len(os.Args) < 3 {
		log.Fatalf("Please provide a filename to read as input for day 1")
	}

	filename := os.Args[2]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var safe_line_count int
	var safe_dampener_line_count int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		safe := is_line_safe(parts)

		if safe {
			safe_line_count += 1
		}

		if !safe {
			dampener_safe := is_line_dampener_safe(parts)
			if dampener_safe {
				safe_dampener_line_count += 1
			}

		}

	}

	fmt.Printf("%v reports were safe\n", safe_line_count)
	fmt.Printf("%v reports were safe (with dampener)\n", safe_line_count+safe_dampener_line_count)
}
