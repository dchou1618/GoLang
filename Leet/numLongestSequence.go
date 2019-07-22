// file

package main;

import (
  "fmt"
)

func LIS(nums []int, currSeq []int, longest *int, numLongest *int, pos int) {
    if (pos >= len(nums)) {
        if (len(currSeq) > *longest) {
            *longest = len(currSeq);
            *numLongest = 1;
        } else if (len(currSeq) == *longest) {
            *numLongest++;
        }
        return;
    }
    if (len(currSeq) == 0)  {
        LIS(nums, append(currSeq,nums[pos]), longest, numLongest, pos+1);
    } else if (currSeq[len(currSeq)-1] < nums[pos]) {
        LIS(nums, append(currSeq,nums[pos]), longest, numLongest, pos+1);
    }
    LIS(nums, currSeq, longest, numLongest, pos+1);
}

func findNumberOfLIS(nums []int) int {
    if (len(nums) == 0) {
        return 0;
    }
    longest := 1;
    numLongest := 0;
    pos := 0;
    curr := []int{};
    LIS(nums, curr, &longest, &numLongest, pos);
    return numLongest;
}
