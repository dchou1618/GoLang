// file alphabetBoard.go


// shortest path - mvoement along board

func isValid(board []string, r int, pos int) bool {
    return 0 <= r && r < len(board) && 0 <= pos && pos < len(board[r]);
}

func alphabetBoardPath(target string) string {
    path := "";
    board := []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"};
    pos := []int{0,0};
    for i := 0; i < len(target); i++ {
        if (board[pos[0]][pos[1]] == target[i]) {
            path += "!";
        } else {
            newRow := (int(target[i])-int('a'))/5;
            newCol := (int(target[i])-int('a')+1)%5-1;
            if (newCol < 0) {
                newCol = 4;
            }
            diffCol := newCol-pos[1];
            diffRow := newRow-pos[0];
            if (board[pos[0]][pos[1]] == 'z') {
                if (diffRow < 0) {
                    path += strings.Repeat("U", -1*diffRow);
                } else {
                    path += strings.Repeat("D", diffRow);
                }
                if (diffCol < 0) {
                    path += strings.Repeat("L", -1*diffCol);
                } else {
                    path += strings.Repeat("R", diffCol)
                }
            } else {
                if (diffCol < 0) {
                    path += strings.Repeat("L", -1*diffCol);
                } else {
                    path += strings.Repeat("R", diffCol)
                }
                if (diffRow < 0) {
                    path += strings.Repeat("U", -1*diffRow);
                } else {
                    path += strings.Repeat("D", diffRow);
                }
            }
            path += "!";
            pos = []int{newRow,newCol};
        }
    }
    pos = nil;
    return path;
}
