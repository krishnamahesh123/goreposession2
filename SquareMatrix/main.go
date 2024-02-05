package main

import "fmt"

func main() {
	var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rows := len(matrix)
	var columns int = 0
	for _, num := range matrix {
		columns = len(num)
		if columns != rows {
			fmt.Printf("Matrix is not a square")
			break
		}
	}
	sum := 0
	if rows == columns {
		fmt.Println("The given matrix is a Squarematrix")
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				if i == j {
					sum = sum + matrix[i][j]
				}
			}
		}
		fmt.Printf("The sum of diagonals of square matrix is %d", sum)
	}

}
