// file singleElem.go

package main

import (
  "fmt"
)

func singleElem(nums []int) int {
    if (len(nums) <= 1) {
        panic("List too short");
    } else {
        for i := 0; i < len(nums); i+=2 {
            if (i+1 >= len(nums)) {
                return nums[i];
            } else {
                if (nums[i] != nums[i+1]) {
                    return nums[i];
                }
            }
        }
        panic("No singled el't");
    }
}



func main() {
    nums := []int{3, 3, 2, 3, 3};
    fmt.Println(singleElem(nums));
}
