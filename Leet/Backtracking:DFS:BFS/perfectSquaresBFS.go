// file perfectSquaresBFS.go


// 72 ms (36.48%), 7.8 mb (25.81%).
func leastSquaresSum(n int, queue [][]int) int {
    used := make(map[int]bool);
    for (len(queue) > 0) {
        top := queue[0];
        queue = queue[1:];
        for num := 1; num <= int(math.Sqrt(float64(n-top[0]))); num++ {
            if (top[0]+(num*num) < n && !used[top[0]+(num*num)]) {
              // verifies all possibilities at each depth
                queue = append(queue, []int{top[0]+(num*num),top[1]+1});
                used[top[0]+(num*num)] = true;
            } else if (top[0]+(num*num) == n) {
                used = nil;
                queue = nil;
                return top[1]+1;
            }
        }
    }
    return 0;
}

func numSquares(n int) int {
    queue := [][]int{[]int{0,0}};
    return leastSquaresSum(n, queue);
}
