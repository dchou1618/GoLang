
package main

import (
  "map"
  "fmt"
)

func twoSum(nums []int, target int) []int {
    indices := []int{0,0};
    vals := make(map[int]int);
    for k := 0; k < len(nums); k++ {
        vals[nums[k]] += 1;
    }
    for i := 0; i < len(nums); i++ {
        if ((vals[target-nums[i]] == 1 && target-nums[i] != nums[i]) || (vals[target-nums[i]] > 1 && target-nums[i] == nums[i])) {
            indices[0] = i;
            break;
        }
    }
    for j := 0; j < len(nums); j++ {
        if (nums[j] == target-nums[indices[0]] && j != indices[0]) {
            indices[1] = j;
            break;
        }
    }
    return indices;
}

func main() {
  nums := []int{2,3,6,8,5,-2};
  target := 9;
  fmt.Println(twoSum(twoSum, target));
}
