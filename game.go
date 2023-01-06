package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxIntervalBetweenGenerations = 9
	minIntervalBetweenGenerations = time.Millisecond * 100
)

type Input struct {
	Number      int
	Seed        int64
	Generations int
	Speed       int
}

func getInput() Input {
	var number int
	var seed int64
	var generations int
	var speed int
	fmt.Print("Enter the size of the square: ")
	fmt.Scanln(&number)
	fmt.Print("Enter a random number: ")
	fmt.Scanln(&seed)
	fmt.Print("Enter the number of generations: ")
	fmt.Scanln(&generations)
	fmt.Printf("Enter the speed [1 - %d]: ", maxIntervalBetweenGenerations)
	fmt.Scanln(&speed)
	return Input{
		Number:      number,
		Seed:        seed,
		Generations: generations,
		Speed:       speed,
	}
}

func getFirstArray(number int) (square [][]int) {
	square = getEmptyArray(number, number)
	for i := 0; i < number; i++ {
		for j := 0; j < number; j++ {
			square[i][j] = rand.Intn(2)
		}
	}
	return
}

func printHeader(array [][]int, generation int) {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("Generation #%d\n", generation)
	fmt.Printf("Alive: %d\n", getNumAliveCells(array))
}

func printArray(array [][]int) {
	for i := range array {
		for j := range array[i] {
			if array[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func getPositiveModulus(dividend, divisor int) (modulo int) {
	modulo = dividend % divisor
	if modulo < 0 {
		modulo += divisor
	}
	return
}

func getLiveNeighbors(array [][]int, row, col int) (count int) {
	rows := len(array)
	cols := len(array[0])
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			idx := getPositiveModulus(row+i, rows)
			jdx := getPositiveModulus(col+j, cols)
			count += array[idx][jdx]
		}
	}
	return
}

func getNextArray(array [][]int) (result [][]int) {
	result = getEmptyArray(len(array), len(array[0]))
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			liveNeighbors := getLiveNeighbors(array, i, j)
			switch liveNeighbors {
			case 3:
				result[i][j] = 1
			case 2:
				result[i][j] = array[i][j]
			default:
				result[i][j] = 0
			}
		}
	}
	return
}

func getEmptyArray(rows, cols int) (array [][]int) {
	array = make([][]int, rows)
	for i := range array {
		array[i] = make([]int, cols)
	}
	return
}

func getNumAliveCells(array [][]int) (aliveCells int) {
	for _, row := range array {
		for _, val := range row {
			aliveCells += val
		}
	}
	return
}

func slowDown(speed int) {
	for i := maxIntervalBetweenGenerations; i >= speed; i-- {
		time.Sleep(minIntervalBetweenGenerations)
	}
}

func main() {
	input := getInput()
	rand.Seed(input.Seed)
	currentArray := getFirstArray(input.Number)
	var nextArray [][]int
	for i := 0; i < input.Generations; i++ {
		printHeader(currentArray, i+1)
		printArray(currentArray)
		slowDown(input.Speed)
		nextArray = getNextArray(currentArray)
		currentArray = nextArray
	}
}
