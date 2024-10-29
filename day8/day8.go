package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cleanDest(input string) string {
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.TrimSpace(input)
	return input
}

func getDest(T []string, lr int) string {
	T[1] = cleanDest(T[1])
	dest := strings.Split(T[1], ",")
	return dest[lr]
}

func part1(input []string, instruction string) {
	index := 0
	steps := 0
	dest := "AAA"
	i := 0
	for dest != "ZZZ" {
		if input[i][:3] == dest {
			splitLine := strings.Split(input[i], "=")
			loc := strings.TrimSpace(splitLine[0])

			switch instruction[index] {
			case 'R':
				dest = getDest(splitLine, 1)
				fmt.Printf("Current location: %v, Destination: %v\n", loc, dest)
				steps++
			case 'L':
				dest = getDest(splitLine, 0)
				fmt.Printf("Current location: %v, Destination: %v\n", loc, dest)
				steps++
			}

			if index == len(instruction)-1 {
				index = 0
			} else {
				index++
			}
		}

		if i == len(input)-1 {
			i = 0
		} else {
			i++
		}
	}

	fmt.Println(steps)
}

func part2(input []string, instruction string) {
	var locations []string
	for _, line := range input {
		if line[2:3] == "A" {
			locations = append(locations, line)
		}
	}

	index := 0
	steps := 0
	dest0 := "11A"
	dest1 := "22A"
	// dest2 := "CMA"
	// dest3 := "MNA"
	// dest4 := "NJA"
	// dest5 := "RVA"
	cycle := 0
	i := 0

	// for dest0[2:3] != "Z" && dest1[2:3] != "Z" && dest2[2:3] != "Z" && dest3[2:3] != "Z" && dest4[2:3] != "Z" && dest5[2:3] != "Z" {
	for dest0[2:3] != "Z" && dest1[2:3] != "Z" {
		switch input[i][:3] {
		case dest0:
			splitLine := strings.Split(input[i], "=")
			loc := strings.TrimSpace(splitLine[0])

			switch instruction[index] {
			case 'R':
				dest0 = getDest(splitLine, 1)
				fmt.Printf("Current location: %v, Destination: %v\n", loc, dest0)
			case 'L':
				dest0 = getDest(splitLine, 0)
				fmt.Printf("Current location: %v, Destination: %v\n", loc, dest0)
			}
		case dest1:
			splitLine := strings.Split(input[i], "=")
			loc := strings.TrimSpace(splitLine[0])

			switch instruction[index] {
			case 'R':
				dest1 = getDest(splitLine, 1)
				fmt.Printf("Current location: %v, Destination: %v\n", loc, dest1)
			case 'L':
				dest1 = getDest(splitLine, 0)
				fmt.Printf("Current location: %v, Destination: %v\n", loc, dest1)
			}
			// case dest2:
			// 	splitLine := strings.Split(input[i], "=")
			// 	loc := strings.TrimSpace(splitLine[0])

			// 	switch instruction[index] {
			// 	case 'R':
			// 		dest2 = getDest(splitLine, 1)
			// 		fmt.Printf("Current location: %v, Destination: %v\n", loc, dest2)
			// 	case 'L':
			// 		dest2 = getDest(splitLine, 0)
			// 		fmt.Printf("Current location: %v, Destination: %v\n", loc, dest2)
			// 	}
			// case dest3:
			// 	splitLine := strings.Split(input[i], "=")
			// 	loc := strings.TrimSpace(splitLine[0])

			// 	switch instruction[index] {
			// 	case 'R':
			// 		dest3 = getDest(splitLine, 1)
			// 		fmt.Printf("Current location: %v, Destination: %v\n", loc, dest3)
			// 	case 'L':
			// 		dest3 = getDest(splitLine, 0)
			// 		fmt.Printf("Current location: %v, Destination: %v\n", loc, dest3)
			// 	}
			// case dest4:
			// 	splitLine := strings.Split(input[i], "=")
			// 	loc := strings.TrimSpace(splitLine[0])

			// 	switch instruction[index] {
			// 	case 'R':
			// 		dest4 = getDest(splitLine, 1)
			// 		fmt.Printf("Current location: %v, Destination: %v\n", loc, dest4)
			// 	case 'L':
			// 		dest4 = getDest(splitLine, 0)
			// 		fmt.Printf("Current location: %v, Destination: %v\n", loc, dest4)
			// 	}
			// case dest5:
			// 	splitLine := strings.Split(input[i], "=")
			// 	loc := strings.TrimSpace(splitLine[0])

			// 	switch instruction[index] {
			// 	case 'R':
			// 		dest5 = getDest(splitLine, 1)
			// 		fmt.Printf("Current location: %v, Destination: %v\n", loc, dest5)
			// 	case 'L':
			// 		dest5 = getDest(splitLine, 0)
			// 		fmt.Printf("Current location: %v, Destination: %v\n", loc, dest5)
			// 	}
		}

		if cycle == 2 {
			if index == len(instruction)-1 {
				index = 0
			} else {
				index++
			}
			steps++
			cycle = 0
		} else {
			cycle++
		}

		if i == len(input)-1 {
			i = 0
		} else {
			i++
		}
	}

	fmt.Println(steps)
}

func main() {
	file, err := os.Open("input.txt")
	var input []string
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	ins := strings.TrimSpace(input[0])
	fmt.Println(ins)

	// part1(input[2:], ins)
	part2(input[2:], ins)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
