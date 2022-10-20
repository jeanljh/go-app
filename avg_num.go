package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// output information
	fmt.Println("Average Number Calculator")
	fmt.Println("-------------------------")

	// prompt for user input
	r := bufio.NewReader(os.Stdin)
	fmt.Println(`Enter number(s) with each number separated by a space. e.g. "1 2 3": `)
	input, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// calculate and print the average number from a list of input numbers
	fmt.Println("Average Number:", calAvgNum(input))
}

// function to calculate average number
func calAvgNum(input string) float64 {
	numbers := strings.Split(strings.TrimSpace(input), " ")

	sum := 0.0
	for _, numStr := range numbers {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			panic(err)
		}
		sum += num
	}

	avg := sum / float64(len(numbers))
	ratio := math.Pow(10, float64(2))
	return math.Round(avg*ratio) / ratio
}
