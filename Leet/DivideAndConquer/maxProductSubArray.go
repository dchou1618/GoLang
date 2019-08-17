// file maxProductSubArray.go

package main

import (
  "fmt"
)

// DP sol'n
// Maximum Subarray Problem has been tackled by Kadane, a statistics
// professor from CMU, who solved it in O(n).
// 

// Sol'n MLE due to recursive calls to the stack
func product(nums[]int) int{
    prod := 1;
    for i := range nums {
        prod *= nums[i];
    }
    return prod;
}

func productMax(nums []int, memo [][]int, max *int, start int, end int) {
    if (start >= end) {
        return;
    } else if (memo[start][end] != math.MinInt64) {
        return;
    } else {
        prod := product(nums[start:end]);
        if (prod > *max) {
            *max = prod;
        }
        memo[start][end] = prod;
        productMax(nums, memo, max, start+1, end-1);
        productMax(nums, memo, max, start+1, end);
        productMax(nums, memo, max, start, end-1);
    }
}

func maxProduct(nums []int) int {
    max := math.MinInt64;
    memo := make([][]int,len(nums)+1);
    for r := 0; r < len(memo); r++ {
        for c := 0; c < len(memo); c++ {
            memo[r] = append(memo[r], math.MinInt64);
        }
    }
    productMax(nums, memo, &max, 0, len(nums));
    memo = nil;
    return max;
}
