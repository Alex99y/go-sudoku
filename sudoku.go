package main

import "fmt"

func isValidSudoku(sudokuArray [9][9]int) bool {
	var i, j int
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if sudokuArray[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func subArray(value int, row int, col int, sudokuArray [9][9]int) bool {
	var iCol, jCol, iRow, jRow int
	if col < 3 && row < 3 {
		// Check first quadrant
		iRow = 0; jRow = 3; iCol = 0; jCol = 3;
	} else if col >= 3 && col < 6 && row < 3 {
		// Check second quadrant
		iRow = 3; jRow = 6; iCol = 0; jCol = 3;
	} else if col >= 6 && col < 9 && row < 3 {
		// Check thrid cuadrant
		iRow = 6; jRow = 9; iCol = 0; jCol = 3;
	} else if col < 3 && row >= 3 && row < 6 {
		// Check fourth cuadrant
		iRow = 0; jRow = 3; iCol = 3; jCol = 6;
	} else if col >= 3 && col < 6 && row >= 3 && row < 6 {
		// Check fifth quadrant
		iRow = 3; jRow = 6; iCol = 3; jCol = 6;
	} else if col >= 6 && col < 9 && row >= 3 && row < 6 {
		// Check sixth quadrant
		iRow = 6; jRow = 9; iCol = 3; jCol = 6;
	} else if col < 3 && row >= 6 && row < 9 {
		// Check seventh quadrant
		iRow = 0; jRow = 3; iCol = 6; jCol = 9;
	} else if col >= 3 && col < 6 && row >= 6 && row < 9 {
		// Check eighth quadrant
		iRow = 3; jRow = 6; iCol = 6; jCol = 9;	
	} else if col >= 6 && col < 9 && row >= 6 && row < 9 {
		// Check ninth quadrant
		iRow = 6; jRow = 9; iCol = 6; jCol = 9;	
	}
	var i, j int
	for j = iCol; j < jCol; j++ {
		for i= iRow; i < jRow; i++ {
			if sudokuArray[j][i] == value {
				return false
			}
		}
	}

	return true
}

func isValidPosition(value int, row int, col int, sudokuArray [9][9]int) bool {
	var i, j int
	// Check current position
	if sudokuArray[row][col] != 0 {
		return false
	}
	// Check horizontal position
	for i = 0; i < 9; i++ {
		if sudokuArray[row][i] == value {
			return false
		}
	}
	// Check vertical position
	for j = 0; j < 9; j++ {
		if sudokuArray[j][col] == value {
			return false
		}
	}
	// Check subArrays
	return subArray(value, row, col, sudokuArray)
}

func findNextPosition(sudokuArray [9][9]int) (int, int, bool) {
	var row, col int
	for row = 0; row < 9; row++ {
		for col = 0; col < 9; col++ {
			if sudokuArray[row][col] == 0 {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}


func sudokuBacktracking(sudokuArray [9][9]int) bool {
	if (isValidSudoku(sudokuArray)) {
		fmt.Print(sudokuArray)
		return true;
	}
	var i, j, e int
	var b bool
	i, j, b = findNextPosition(sudokuArray);
	if (!b) {
		return true
	}
	for e = 1; e < 10; e++ {
		if isValidPosition(e, i, j, sudokuArray) {
			sudokuArray[i][j] = e
			if sudokuBacktracking(sudokuArray) {
				return true
			}
			sudokuArray[i][j] = 0
		}
	}

	return false
}

func main() {
	var sudokuArray = [9][9]int {
		{0, 0, 0, 0, 2, 4, 3, 0, 7},
		{9, 0, 0, 0, 6, 0, 0, 0, 0},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 0, 7, 9, 4},
		{0, 6, 5, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 8, 0, 3, 0, 0, 0},
		{0, 5, 4, 0, 0, 0, 0, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 5, 9},
		{0, 0, 3, 4, 0, 8, 0, 0, 0},
	}
	sudokuBacktracking(sudokuArray)
}