package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type BinaryPuzzle struct {
	Matrix [][]rune
	Length int
}

func main() {
	binarypuzzle, err := ReadBinaryPuzzle(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	binarypuzzle.Print()
	start := time.Now()
	success := binarypuzzle.Solve()
	end := time.Now()
	if !success {
		log.Fatal("Couldn't solve")
	}
	binarypuzzle.Print()

	fmt.Println("Time elapsed: " + strconv.Itoa(end.Nanosecond()-start.Nanosecond()) + "ns")
}

// Checks whether the puzzle is successful so far
func (b BinaryPuzzle) IsCorrect() bool {
	// We check every row
	for row := 0; row < b.Length; row++ {
		// Is the sum of the row == (1/2)length?
		sum := 0
		containsOpenSquare := false
		for _, ch := range b.Matrix[row] {
			if ch == '?' {
				containsOpenSquare = true
				break
			}
			if ch == '1' {
				sum += 1
			}
		}

		if !containsOpenSquare {
			if sum != (b.Length / 2) {
				return false
			}
		}

		for column := 0; column < b.Length-2; column++ {
			if b.Matrix[row][column] == '?' {
				continue
			}
			if b.Matrix[row][column] == b.Matrix[row][column+1] && b.Matrix[row][column] == b.Matrix[row][column+2] {
				return false
			}
		}
	}

	// Now we check the columns
	for column := 0; column < b.Length; column++ {
		// Sum of column == (1/2)length?
		sum := 0
		containsOpenSquare := false
		for _, row := range b.Matrix {
			if row[column] == '?' {
				containsOpenSquare = true
				break
			}
			if row[column] == '1' {
				sum += 1
			}
		}

		if !containsOpenSquare {
			if sum != (b.Length / 2) {
				return false
			}
		}

		for row := 0; row < b.Length-2; row++ {
			if b.Matrix[row][column] == '?' {
				continue
			}
			if b.Matrix[row][column] == b.Matrix[row+1][column] && b.Matrix[row][column] == b.Matrix[row+2][column] {
				return false
			}
		}
	}

	return true
}

// Backtracking algorithm
func (b BinaryPuzzle) Solve() bool {
	success := b.IsCorrect()
	if !success { // It's unsolvable
		return false
	}
	if b.IsSolved() {
		return true
	}
	// Test all squares
	for row := 0; row < b.Length; row++ {
		for column := 0; column < b.Length; column++ {
			if b.Matrix[row][column] == '?' {
				// Now we start placing a number
				b.Matrix[row][column] = '1'
				success := b.Solve()
				// We're done
				if success {
					return true
				}

				// Other option, it's 0
				b.Matrix[row][column] = '0'
				success = b.Solve()

				// If it's correct we're done, if it's not correct we made a mistake somewhere
				if success {
					return true
				}
				b.Matrix[row][column] = '?'
				return false // If it's correct we're done, if it's not correct it's unsolveable
			}
		}
	}
	fmt.Println("what?")
	return true
}

func (b BinaryPuzzle) IsSolved() bool {
	for _, row := range b.Matrix {
		for _, char := range row {
			if char == '?' {
				return false
			}
		}
	}
	return true
}

func ReadBinaryPuzzle(r io.Reader) (BinaryPuzzle, error) {
	reader := bufio.NewReader(r)
	lengthstr, err := reader.ReadString('\n')
	if err != nil {
		return BinaryPuzzle{}, err
	}
	length, err := strconv.Atoi(lengthstr[:len(lengthstr)-1])
	if err != nil {
		return BinaryPuzzle{}, err
	}

	matrix := make([][]rune, length)

	for row := 0; row < length; row++ {
		line, err := reader.ReadString('\n')
		matrix[row] = make([]rune, length)
		if err != nil {
			return BinaryPuzzle{}, err
		}
		for column, ch := range line[:len(line)-1] {
			matrix[row][column] = ch
		}
	}

	return BinaryPuzzle{matrix, length}, nil

}

func (b BinaryPuzzle) Print() {
	fmt.Println("---")
	for _, row := range b.Matrix {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}
}
