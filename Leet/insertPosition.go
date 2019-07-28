// file insertPosition.go

package main

import (
  "fmt"
)

// 4 ms - 92.01%, 3.1 mb - 99.17%
// binary search implementation
func searchInsert(nums []int, target int) int {
    lowerBound, upperBound := 0, len(nums)-1;
    mid := (lowerBound+upperBound)/2;
    for (lowerBound <= upperBound) {
        if (nums[mid] == target) {
            return mid;
        } else if (nums[mid] < target) {
            lowerBound = mid+1;
        } else {
            upperBound = mid-1;
        }
        mid = (lowerBound+upperBound)/2;
    }
    return lowerBound;
}
