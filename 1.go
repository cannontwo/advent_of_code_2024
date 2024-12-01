package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func calculate_list_sim_score(left []int, right []int) int {
	var sim_score int
	scores_memo := make(map[int]int)

	var start_idx, iter_idx int

	for _, elem := range left {
		score, ok := scores_memo[elem]
		if ok {
			sim_score += score
		} else {
			// Calculate score
			var right_count int

			for ; start_idx < len(right)-1 && right[start_idx] < elem; start_idx++ {
			}

			for iter_idx = start_idx; iter_idx < len(right); iter_idx++ {
				if right[iter_idx] == elem {
					right_count += 1
				} else {
					break
				}
			}

			start_idx = iter_idx

			elem_score := elem * right_count
			scores_memo[elem] = elem_score
			sim_score += elem_score
		}
	}

	return sim_score
}

func calculate_list_distance(left []int, right []int) int {
	if len(left) != len(right) {
		panic("Lists different lengths")
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var distance int
	for i := range len(left) {
		diff := right[i] - left[i]
		if diff < 0 {
			diff *= -1
		}

		distance += diff
	}

	return distance
}

func run_day_one() {
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

	var left_list []int
	var right_list []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		left_list = append(left_list, left)

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		right_list = append(right_list, right)
	}

	fmt.Printf("Distance: %v\n", calculate_list_distance(left_list, right_list))
	fmt.Printf("Similarity Score: %v\n", calculate_list_sim_score(left_list, right_list))
}
