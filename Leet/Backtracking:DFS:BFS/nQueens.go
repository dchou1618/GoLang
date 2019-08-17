// file nQueens.go
package main

import (
  "fmt"
)


func validNQueens(board []string, row int, col int, piece byte) (bool){
    originalRow, originalCol := row, col;
    if (!(0 <= row && row < len(board) && 0 <= col && col < len(board))) {
        return false;
    }
    for (0 <= row && row < len(board) && 0 <= col && col < len(board)) {
        if (board[row][col] == piece) {
            return false;
        }
        row++;
        col++;
    }
    row, col = originalRow, originalCol;
    for (0 <= row && row < len(board) && 0 <= col && col < len(board)) {
        if (board[row][col] == piece) {
            return false;
        }
        row--;
        col--;
    }
    row, col = originalRow, originalCol;
    for (0 <= row && row < len(board) && 0 <= col && col < len(board)) {
        if (board[row][col] == piece) {
            return false;
        }
        row++;
        col--;
    }
    row, col = originalRow, originalCol;
    for (0 <= row && row < len(board) && 0 <= col && col < len(board)) {
        if (board[row][col] == piece) {
            return false;
        }
        row--;
        col++;
    }
    col = originalCol;
    for r := 0; r < len(board); r++ {
        if (board[r][col] == piece) {
            return false;
        }
    }
    row = originalRow;
    for c := 0; c < len(board); c++ {
        if (board[row][c] == piece) {
            return false;
        }
    }
    return true;
}
func nQueensWrapper(n int, board []string, boards *[][]string, row int, col int, used int){
    if (used == 0) {
        temp := make([]string, len(board));
        copy(temp,board);
        *boards = append(*boards,temp);
    } else if (row >= len(board)) {
        return;
    } else {
        if (validNQueens(board,row,col,'Q')) {
            board[row] = board[row][:col] + "Q" + board[row][col+1:];
            if (col == len(board[0])-1) {
                nQueensWrapper(n,board,boards,row+1,0,used-1);
            } else {
                nQueensWrapper(n,board,boards,row,col+1,used-1);
            }
            board[row] = board[row][:col] + "." + board[row][col+1:];
            if (col == len(board[0])-1) {
                nQueensWrapper(n,board,boards,row+1,0,used);
            } else {
                nQueensWrapper(n,board,boards,row,col+1,used);
            }
        } else {
            if (col == len(board[0])-1) {
                nQueensWrapper(n,board,boards,row+1,0,used);
            } else {
                nQueensWrapper(n,board,boards,row,col+1,used);
            }
        }
    }
}
func solveNQueens(n int) [][]string {
    board := []string{};
    for row := 0; row < n; row++ {
        currRow := "";
        for col := 0; col < n; col++ {
            currRow += ".";
        }
        board = append(board,currRow);
    }
    boards := [][]string{};
    used := n;
    row,col := 0,0;
    nQueensWrapper(n, board, &boards, row, col, used);
    return boards;
}



func main() {
  
}
