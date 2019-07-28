//file squarefulPermutations.go



func isSquareful(A []int) bool {
    for i := 0; i < len(A)-1; i++ {
        var x float64 = float64(A[i]+A[i+1]);
        if (math.Sqrt(x) != math.Floor(math.Sqrt(x))) {
            return false;
        }
    }
    return true;
}

func noDuplicates(A []int, start int, end int) bool {
    for i := start; i < end; i++ {
        if (A[i] == A[end]) {
            return false;
        }
    }
    return true;
}

func swap(nums *[]int, i int, j int) {
    temp := (*nums)[i];
    (*nums)[i] = (*nums)[j];
    (*nums)[j] = temp;
}

func permute(A []int, pos int, count *int) {
    if (pos == len(A)) {
        if (isSquareful(A)) {
            (*count)++;
        }
    } else {
        for i := pos; i < len(A); i++ {
            if (noDuplicates(A, pos, i)) {
                swap(&A, pos, i);
                permute(A, pos+1, count);
                swap(&A, pos, i);
            }
        }
    }
}

func numSquarefulPerms(A []int) int {
    var numSquareful int = 0;
    permute(A, 0, &numSquareful);
    return numSquareful;
}
