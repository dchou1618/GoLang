// file isHoppable.go


package main

import (
  "fmt"
)

func isHoppable(nums []int) bool {
  index := 0;
  for (index < len(nums)) {
    if (nums[index] == 0) {
      return false;
    } else {
      index += nums[index];
    }
  }
  return true;
}

func main() {
  nums := []int{4,2,0,0,2,0};
  fmt.Println(isHoppable(nums));
}
