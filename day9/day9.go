package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("input.txt")
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}

func readInput(filename string) []string {
	var input []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func part1(input []string) int {
	sum := 0

	for _, line := range input {
		strNumbers := strings.Split(line, " ")
		numbers := []int{}
		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, num)
		}
		nextNumber := predictNextValue(numbers)
		sum += nextNumber
	}

	return sum
}

func part2(input []string) int {
	sum := 0

	for _, line := range input {
		strNumbers := strings.Split(line, " ")
		numbers := []int{}
		for _, strNum := range strNumbers {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, num)
		}
		prevNumber := predictPreviousValue(numbers)
		sum += prevNumber
	}

	return sum
}

func predictNextValue(sequence []int) int {
	matrix := [][]int{
		sequence,
	}
	nonZeroFound := true
	for nonZeroFound {
		nonZeroFound = false
		next := []int{}

		matrixHistoryRow := matrix[len(matrix)-1]
		for i := 1; i < len(matrixHistoryRow); i++ {
			prev := matrixHistoryRow[i-1]
			curr := matrixHistoryRow[i]
			next = append(next, curr-prev)
			if next[len(next)-1] != 0 {
				nonZeroFound = true
			}
		}
		matrix = append(matrix, next)
	}

	ans := 0
	for _, row := range matrix {
		ans += row[len(row)-1]
	}

	return ans
}

func predictPreviousValue(sequence []int) int {
	matrix := [][]int{
		sequence,
	}
	nonZeroFound := true
	for nonZeroFound {
		nonZeroFound = false
		next := []int{}

		matrixHistoryRow := matrix[len(matrix)-1]
		for i := 1; i < len(matrixHistoryRow); i++ {
			prev := matrixHistoryRow[i-1]
			curr := matrixHistoryRow[i]
			next = append(next, curr-prev)
			if next[len(next)-1] != 0 {
				nonZeroFound = true
			}
		}
		matrix = append(matrix, next)
	}

	ans := 0
	for r := len(matrix) - 1; r >= 0; r-- {
		ans = matrix[r][0] - ans
	}

	return ans
}
