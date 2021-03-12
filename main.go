package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fibMemod(upTo int, memo map[int]int) int {
	if upTo == 0 {
		return 0
	}
	if upTo <= 2 {
		return 1
	}
	value, exists := memo[upTo]
	if exists == true {
		return value
	}
	memo[upTo] = fibMemod(upTo-1, memo) + fibMemod(upTo-2, memo)
	return memo[upTo]
}

func convertInput(raw string) int {
	formatted := strings.Replace(raw, "\n", "", -1)
	num, err := strconv.Atoi(formatted)
	if err != nil || num < 0 {
		return -1
	}
	return num
}

func getUserInput() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("--\nPlease enter a number to calculate the nth Fibonacci position for: ")
	raw, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Unable to read stdin")
	}
	return convertInput(raw)
}

func RecursivelyMemoized(upTo int) int {
	return fibMemod(upTo, make(map[int]int))
}

func main() {
	input := getUserInput()
	if input < 0 {
		log.Fatal("Invalid input, please try again...")
	}
	fmt.Printf("--\nans: %d\n", RecursivelyMemoized(input))
}
