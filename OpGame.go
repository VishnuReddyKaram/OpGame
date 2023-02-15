package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Solve finds all solutions to the given numbers to result in the target.
// nums is an array of elements to solve for.

func solve(nums []int, i int, acc int, target int, path string, operation string) []string {
	var nextPath string
	// Append the current number to the path string.
	if i == 0 {
		nextPath = ""
	} else if i == 1 {
		num := nums[i-1]
		nextPath = fmt.Sprintf("%d", num)
	} else if i < len(nums)+1 {
		num := nums[i-1]
		nextPath = fmt.Sprintf("%s %s %d", path, operation, num)
	}

	if i == len(nums) {

		// check If the accumulator matches the target, return the solution.
		if acc == target {
			return []string{nextPath}
		}
		// here no solutions found.
		return nil
	}

	// Get the current number from the array.
	num := nums[i]

	// All results found in this recursive run.
	var solutions []string

	// Trying addition.
	if res := solve(nums, i+1, acc+num, target, nextPath, "+"); res != nil {
		solutions = append(solutions, res...)
	}

	// Trying subtraction.
	if res := solve(nums, i+1, acc-num, target, nextPath, "-"); res != nil {
		solutions = append(solutions, res...)
	}

	// Trying multiplication.
	if res := solve(nums, i+1, acc*num, target, nextPath, "*"); res != nil {
		solutions = append(solutions, res...)
	}

	// Trying division, but only if num is not 0
	// evenly with num.
	if num != 0 && acc%num == 0 {
		if res := solve(nums, i+1, acc/num, target, nextPath, "/"); res != nil {
			solutions = append(solutions, res...)
		}
	}

	return solutions
}

// This function takes a slice of integers and the index of the current integer in the slice and the target value, and check the solution to get the target.
func main() {
	scanner := bufio.NewScanner(os.Stdin) // Initialize a scanner object
	for scanner.Scan() {
		line := scanner.Text() // Storing the user's input as string
		nums := make([]int, 0) // Creating empty slice of type int
		for _, s := range regexp.MustCompile("\\s+").Split(line, -1) {
			num, err := strconv.Atoi(s) // converting string to int
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: ", err)
				os.Exit(1)
			}
			nums = append(nums, num) // Append the int to nums slice
		}
		if len(nums) < 2 { // If the number of items in nums is lower than expected
			fmt.Fprintln(os.Stderr, "Please Enter a valid input")
			os.Exit(1)
		}
		target := nums[len(nums)-1]                                   // Last elements of the input will be set as target
		solutions := solve(nums[:len(nums)-1], 0, 0, target, "", "+") // Call solve function
		if len(solutions) == 0 {
			fmt.Println()
		} else {
			for _, solution := range solutions { // Iterate through solutions
				fmt.Println(solution, "=", target)
			}
		}
	}
}
