// file uniquePermutations.go


// 8ms (20.87%), 7 mb (41.51%).
func swap(nums *[]int, i int, j int) {
    temp := (*nums)[i];
    (*nums)[i] = (*nums)[j];
    (*nums)[j] = temp;
}

func permuteWrapper(nums []int, permutes *[][]int, pos int) {
    if (pos == len(nums)) {
        var newNums []int;
        newNums = append(newNums,nums...); // avoid aliasing
        *permutes = append(*permutes, newNums);
    } else {
        for i := pos; i < len(nums); i++ { // O(n)
            swap(&nums, pos, i); // idea is to swap every possible pair
            permuteWrapper(nums, permutes, pos+1);
            swap(&nums, pos, i);
        }
    }
}

// since distinct integers, permutations won't have repeats such as 1,1,5 or 3,4,4
func permute(nums []int) [][]int {
    permutations := [][]int{};
    permuteWrapper(nums, &permutations, 0);
    return permutations;
}

// 8 ms (96.09%), 8 mb (22.58%).

func noDuplicates(nums []int, start int, pos int) bool {
    for i := start; i < pos; i++ {
        if (nums[i] == nums[pos]) {
            return false;
        }
    }
    return true;
}

func permute(nums[]int, permutes *[][]int, pos int) {
    if (pos == len(nums)) {
        var arr []int;
        arr = append(arr, nums...);
        *permutes = append(*permutes, arr);
    } else {
        for i := pos; i < len(nums); i++ {
            if (noDuplicates(nums, pos, i)) {
                swap(&nums, pos, i);
                permute(nums, permutes, pos + 1);
                swap(&nums, pos, i);
            }
        }
    }
}

func permuteUnique(nums []int) [][]int {
    permutes := [][]int{};
    permute(nums, &permutes, 0);
    return permutes;
}
