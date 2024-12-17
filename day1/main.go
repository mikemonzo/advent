package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// read a file nad create 2 list of numbers
// input file: numbers.txt
// output: list1, list2
func read_file(filepath string) ([]int, []int, error) {
	// open file to	read file
	file, err := os.Open(filepath)

	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var list1, list2 []int
	// read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// split line by space or tab
		parts := strings.Fields(line)
		if len(parts) == 2 {
			// add first number to list1
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error: ", err1, err2)
				return nil, nil, fmt.Errorf("error opening file: %v", err)
			}
			list1 = append(list1, num1)
			list2 = append(list2, num2)
			// add second number to list2
			// close file
		}
	}
	return list1, list2, nil
}

func calculate_distance(list1 []int, list2 []int) (int, error) {
	if len(list1) != len(list2) {
		return 0, fmt.Errorf("list1 and list2 should have same length")
	}

	distance := 0
	for i := 0; i < len(list1); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		distance += diff
	}
	return distance, nil
}

func calculate_similarity_score(list1 []int, list2 []int) (int, error) {
	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num]++
	}

	similarity_score := 0

	for _, num := range list1 {
		similarity_score += num * counts[num]
	}

	return similarity_score, nil
}

func main() {
	list1, list2, err := read_file("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	sort.Ints(list1)
	sort.Ints(list2)

	distance, err := calculate_distance(list1, list2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	similarity_score, err := calculate_similarity_score(list1, list2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// create a list
	// fmt.Println("List1: ", list1)
	// fmt.Println("List2: ", list2)
	fmt.Println("Distance: ", distance)
	fmt.Println("Similarity Score: ", similarity_score)
}
