package main

import (
	"fmt"
	"math/rand"
	"time"
)

const generations = 10

type Input struct {
	Number int
	//Seed        int64
	Generations int
}

func getInput() Input {
	var number int
	//var seed int64
	//var generations int
	fmt.Scanln(&number)
	return Input{
		Number: number,
		//Seed:        seed,
		Generations: generations,
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

func main() {
	input := getInput()
	//rand.Seed(input.Seed)
	currentArray := getFirstArray(input.Number)
	nextArray := getEmptyArray(input.Number, input.Number)
	for i := 0; i < input.Generations; i++ {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation #%d\n", i+1)
		fmt.Printf("Alive: %d\n", getNumAliveCells(currentArray))
		printArray(currentArray)
		time.Sleep(time.Second / 3)
		nextArray = getNextArray(currentArray)
		currentArray = nextArray
	}
}
