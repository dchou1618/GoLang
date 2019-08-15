// file partitionSumK.go

// naive recursive sol'n

package main

import (
    "fmt"
)



func partitionSubsets(nums []int, k int, sumObj int, currSum int, pos int) bool {
    if (currSum > sumObj && sumObj > -1) {
        return false;
    }
    if (k == 0 && len(nums) == 0) {
        return true;
    } else if (k == 0) {
        return false;
    } else {
        if (pos >= len(nums)) {
            if (sumObj == -1) {
                sumObj = currSum;
                k--;
                pos = 0;
            } else {
                if (currSum != sumObj) {
                    return false;
                } else {
                    fmt.Println(currSum, nums, k);
                    pos = 0;
                    k--;
                }
            }
            return partitionSubsets(nums, k, sumObj, 0, pos);
        } else {
            val := nums[pos];
            items := []int{};
            items = append(items, nums...);
            return partitionSubsets(append(nums[:pos], nums[pos+1:]...), k, sumObj, currSum+val, pos) || partitionSubsets(items, k, sumObj, currSum, pos+1);
        }
    }
}

func canPartitionKSubsets(nums []int, k int) bool {
    if (k > len(nums)) {
        return false;
    } else if (len(nums) == 0) {
        return true;
    } else {
        sort.Ints(nums);
        return partitionSubsets(nums, k, -1, 0, 0);
    }
}


func main() {
    nums := []int{2,3,5,6,3,2,5,7,9,8,2};
    fmt.Println(canPartitionKSubsets(nums, 3));
}
