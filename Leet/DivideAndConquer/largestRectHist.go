// file largestRectHist.go

// hints: use stack, dynamic programming

package main

import (
    "fmt"
)

// Naive solution
func getMin(nums []int) int {
    var min int = math.MaxInt64;
    for _, val := range nums {
        if (val < min) {
            min = val;
        }
    }
    return min;
}


func largestAreaHist(heights []int, start int, end int, memo *[][]int, largest *int) {
    if (end <= start) {
        return;
    }
    if ((*memo)[start][end] == 0) {
        min := getMin(heights[start:end]);
        if (*largest < min*(end-start)) {
            *largest = min*(end-start);
        }
        (*memo)[start][end] = min*(end-start);
        largestAreaHist(heights, start+1, end, memo, largest);
        largestAreaHist(heights, start, end-1, memo, largest);
        largestAreaHist(heights, start+1, end-1, memo, largest);
    }
}

func largestRectangleArea(heights []int) int {
    memo := make([][]int, len(heights));
    for i := range memo {
        memo[i] = make([]int, len(heights)+1);
    }
    largest := 0;
    largestAreaHist(heights, 0, len(heights), &memo, &largest);
    return largest;
}


// sol'n with divide and conquer

func getMinIndex(nums []int, start int, end int) int {
    var min int = math.MaxInt64;
    var minIndex int;
    for i := start; i <= end; i++ {
        if (nums[i] < min) {
            min = nums[i];
            minIndex = i;
        }
    }
    return minIndex;
}

func largestAreaHist(heights []int, start int, end int, largest* int) {
    if (end < start) {
        return;
    }
    if (start == end) {
        if (heights[start] > *largest) {
            *largest = heights[start]
        }
        return;
    }
    minIndex := getMinIndex(heights, start, end);
    if (heights[minIndex]*(end-start+1) > *largest) {
        *largest = heights[minIndex]*(end-start+1);
    }
    largestAreaHist(heights, start, minIndex-1, largest); // split by minIndex
    largestAreaHist(heights, minIndex+1, end, largest);
}

func largestRectangleArea(heights []int) int {
    largest := 0;
    largestAreaHist(heights, 0, len(heights)-1, &largest);
    return largest;
}

func main() {
    heights := []int{3,4,6,7,4,5,7,8,1};
    fmt.Println(largestRectangleArea(heights));
}
