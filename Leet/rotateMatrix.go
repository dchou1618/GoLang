// file rotateMatrix.go

// rotates matrix in-place

package main

import (
    "fmt"
)

func swap(matrix *[][]int, r1 int, c1 int, r2 int, c2 int) {
    var temp int = (*matrix)[r1][c1];
    (*matrix)[r1][c1] = (*matrix)[r2][c2];
    (*matrix)[r2][c2] = temp;
}

func rotate(matrix [][]int)  {
    for row := 0; row < len(matrix)/2; row++ {
        for col := row; col < len(matrix)-row-1; col++ {
            swap(&matrix, row, col, len(matrix)-col-1, row);
            swap(&matrix, len(matrix)-col-1, row, len(matrix)-row-1, len(matrix)-col-1);
            swap(&matrix, len(matrix)-row-1, len(matrix)-col-1, col, len(matrix)-row-1);
        }
    }
}
