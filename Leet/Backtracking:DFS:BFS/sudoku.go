// basic sudoku.go sudoku solving implementation


func validBoard(board [][]byte, row int, col int, num byte) (bool) {
    for r := 0; r < len(board); r++ {
        if (board[r][col] == num) {
            return false;
        }
    }
    for c := 0; c < len(board[row]); c++ {
        if (board[row][c] == num) {
            return false;
        }
    }
    var i int = row/3;
    var j int = col/3;
    var topLeftRow = i*3;
    var topLeftCol = j*3;
    for r := 0; r < 3; r++ {
        for c := 0; c < 3; c++ {
            if (board[topLeftRow+r][topLeftCol+c] == num) {
                return false;
            }
        }
    }
    return true;
}

func solveWrapper(board [][]byte, row int, col int) (bool) {
    if (col == len(board)) {
        col = 0;
        row++;
        if (row == len(board)) {
            return true;
        }
    }
    if (board[row][col] == '.') {
        for num := byte('1'); num <= '9'; num++ {
            if (validBoard(board, row, col, num)) {
                board[row][col] = num;
                sudoku := solveWrapper(board, row, col+1);
                if (!sudoku) {
                    board[row][col] = '.';
                } else {
                    return true;
                }
            }
        }
    } else {
        return solveWrapper(board, row, col+1);
    }
    return false;
}

func solveSudoku(board [][]byte)  {
    solveWrapper(board, 0, 0);
}
