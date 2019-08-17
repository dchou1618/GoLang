// file longestIncreasingPathDP.go

// topological sort attempt 
func isValid(matrix [][]int, row int, col int) bool {
    return (0 <= row && row < len(matrix) &&
            0 <= col && col < len(matrix[0]));
}

func genIndegree(matrix [][]int) [][]int{
    indegree := make([][]int, len(matrix));
    for r := range indegree {
        indegree[r] = make([]int, len(matrix[0]));
    }
    for r := 0; r < len(matrix); r++ {
        for c := 0; c < len(matrix[0]); c++ {
            if (isValid(matrix,r+1,c) && matrix[r][c] > matrix[r+1][c]) {
                indegree[r+1][c]++;
            }
            if (isValid(matrix,r-1,c) && matrix[r][c] > matrix[r-1][c]) {
                indegree[r-1][c]++;
            }
            if (isValid(matrix,r,c+1) && matrix[r][c] > matrix[r][c+1]) {
                indegree[r][c+1]++;
            }
            if (isValid(matrix,r,c-1) && matrix[r][c] > matrix[r][c-1]) {
                indegree[r][c-1]++;
            }
        }
    }
    return indegree;
}

func getQueue(indegree [][]int) [][]int {
    queue := [][]int{};
    for r := range indegree {
        for c := range indegree[r] {
            if (indegree[r][c] == 0) {
                queue = append(queue, []int{r,c});
            }
        }
    }
    return queue;
}
// extraneous memory from string conversion for key in neighbors dictionary
// func adjacencyLst(matrix [][]int) map[string][][]int {
//     neighbors := make(map[string][][]int);
//     for r := 0; r < len(matrix); r++ {
//         for c := 0; c < len(matrix[0]); c++ {
//             if (isValid(matrix,r+1,c) && matrix[r][c] > matrix[r+1][c]) {
//                 neighbors[string(r)+","+string(c)] = append(neighbors[string(r)+","+string(c)], []int{r+1,c});
//             }
//             if (isValid(matrix,r-1,c) && matrix[r][c] > matrix[r-1][c]) {
//                 neighbors[string(r)+","+string(c)] = append(neighbors[string(r)+","+string(c)], []int{r-1,c});
//             }
//             if (isValid(matrix,r,c+1) && matrix[r][c] > matrix[r][c+1]) {
//                 neighbors[string(r)+","+string(c)] = append(neighbors[string(r)+","+string(c)], []int{r,c+1});
//             }
//             if (isValid(matrix,r,c-1) && matrix[r][c] > matrix[r][c-1]) {
//                 neighbors[string(r)+","+string(c)] = append(neighbors[string(r)+","+string(c)], []int{r,c-1});
//             }
//         }
//     }
//     return neighbors;
// }

func longestIncreasingPath(matrix [][]int) int {
    if (len(matrix) == 0) {
        return 0;
    } else {
        indegree := genIndegree(matrix);
        queue := getQueue(indegree);
        if (len(queue) == 0) {
            return 0;
        } else {
            count := 0;
            for len(queue) != 0 {
                size := len(queue);
                for i := 0; i < size; i++ {
                    top := queue[0];
                    queue = queue[1:];
                    if (isValid(matrix,top[0]+1,top[1]) &&
                        matrix[top[0]][top[1]] > matrix[top[0]+1][top[1]]) {
                        indegree[top[0]+1][top[1]]--;
                        if (indegree[top[0]+1][top[1]] == 0) {
                            queue = append(queue, []int{top[0]+1,top[1]});
                        }
                    }
                    if (isValid(matrix,top[0]-1,top[1]) &&
                        matrix[top[0]][top[1]] > matrix[top[0]-1][top[1]]) {
                        indegree[top[0]-1][top[1]]--;
                        if (indegree[top[0]-1][top[1]] == 0) {
                            queue = append(queue, []int{top[0]-1,top[1]});
                        }
                    }
                    if (isValid(matrix,top[0],top[1]+1) &&
                        matrix[top[0]][top[1]] > matrix[top[0]][top[1]+1]) {
                        indegree[top[0]][top[1]+1]--;
                        if (indegree[top[0]][top[1]+1] == 0) {
                            queue = append(queue, []int{top[0],top[1]+1});
                        }
                    }
                    if (isValid(matrix,top[0],top[1]-1) &&
                        matrix[top[0]][top[1]] > matrix[top[0]][top[1]-1]) {
                        indegree[top[0]][top[1]-1]--;
                        if (indegree[top[0]][top[1]-1] == 0) {
                            queue = append(queue, []int{top[0],top[1]-1});
                        }
                    }
                }
                count++;
            }
            queue = nil;
            indegree = nil;
            return count;
        }
    }
}


// memoization
// Runtime: 40 ms (46.39%), 6.3 MB (47.06%)
func isValid(matrix [][] int, row int, col int) bool {
    return (0 <= row && row < len(matrix) &&
            0 <= col && col < len(matrix[0]))
}

func dfsPath(matrix [][]int, memo [][]int, row int, col int) int {
    if (isValid(matrix, row, col)) {
        if (memo[row][col] != -1) {
            return memo[row][col];
        }
        maxDepth := 0;
        if (isValid(matrix,row-1,col) && matrix[row][col] > matrix[row-1][col]) {
            path := dfsPath(matrix, memo, row-1, col);
            if (path > maxDepth) {
                maxDepth = path;
            }
        }
        if (isValid(matrix,row+1,col) && matrix[row][col] > matrix[row+1][col]) {
            path := dfsPath(matrix, memo, row+1, col);
            if (path > maxDepth) {
                maxDepth = path;
            }
        }
        if (isValid(matrix,row, col-1) && matrix[row][col] > matrix[row][col-1]) {
            path := dfsPath(matrix, memo, row, col-1);
            if (path > maxDepth) {
                maxDepth = path;
            }
        }
        if (isValid(matrix, row, col+1) && matrix[row][col] > matrix[row][col+1]) {
            path := dfsPath(matrix, memo, row, col+1);
            if (path > maxDepth) {
                maxDepth = path;
            }
        }
        memo[row][col] = 1+maxDepth;
        return 1+maxDepth;
    }
    return 0;
}

func longestIncreasingPath(matrix [][]int) int {
    if (len(matrix) == 0) {
        return 0;
    } else {
        memo := [][]int{};
        for r := 0; r < len(matrix); r++ {
            vals := []int{};
            for c := 0; c < len(matrix[0]); c++ {
                vals = append(vals, -1);
            }
            memo = append(memo, vals);
        }
        optimalPath := 1;
        for row := 0; row < len(matrix); row++ {
            for col := 0; col < len(matrix[0]); col++ {
                if (memo[row][col] != -1) {
                    if (memo[row][col] > optimalPath) {
                        optimalPath = memo[row][col];
                    }
                } else {
                    path := dfsPath(matrix, memo, row, col);
                    if (path > optimalPath) {
                        optimalPath = path;
                    }
                }
            }
        }
        return optimalPath;
    }
}

// Exponential time?..
// Runtime: 100 ms (9%), 6.8 MB (29.41%)
func validMoves(matrix [][] int, row int, col int) [][]int {
    moves := [][]int{{0,1},{0,-1},{1,0},{-1,0}};
    finalMoves := [][]int{};
    for _, move := range moves {
        if (0 <= row+move[0] && row+move[0] < len(matrix) &&
            0 <= col+move[1] && col+move[1] < len(matrix[0]) &&
            matrix[row][col] > matrix[row+move[0]][col+move[1]]) {
            finalMoves = append(finalMoves, move);
        }
    }
    moves = nil;
    return finalMoves;
}

func dfsPath(matrix [][]int, memo [][]int, row int, col int, depth int) int {
    moves := validMoves(matrix, row, col);
    if (len(moves) != 0) {
        if (memo[row][col] != -1) {
            return memo[row][col];
        }
        maxDepth := 1;
        for _, move := range moves {
            path := dfsPath(matrix, memo, row+move[0], col+move[1], depth+1);
            if (path > maxDepth) {
                maxDepth = path;
            }
        }
        memo[row][col] = 1+maxDepth;
        return memo[row][col];
    }
    return 0;
}

func longestIncreasingPath(matrix [][]int) int {
    if (len(matrix) == 0) {
        return 0;
    } else {
        memo := [][]int{};
        for r := 0; r < len(matrix); r++ {
            vals := []int{};
            for c := 0; c < len(matrix[0]); c++ {
                vals = append(vals, -1);
            }
            memo = append(memo, vals);
        }
        optimalPath := 1;
        for row := 0; row < len(matrix); row++ {
            for col := 0; col < len(matrix[0]); col++ {
                if (memo[row][col] != -1) {
                    if (memo[row][col] > optimalPath) {
                        optimalPath = memo[row][col];
                    }
                } else {
                    path := dfsPath(matrix, memo, row, col, 1);
                    if (path > optimalPath) {
                        optimalPath = path;
                    }
                }
            }
        }
        return optimalPath;
    }
}
