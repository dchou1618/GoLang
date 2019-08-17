/* minSwapsUniform.go file */


package main

import (
    "fmt"
)

func min(n1 int, n2 int) int {
    if (n1 < n2) {
        return n1;
    } else {
        return n2;
    }
}

func _uniformArr(nums1 []int, nums2 []int) int {
    if (len(nums1) < 1) {
        return -1;
    } else if (len(nums1) == 1) {
        return 0;
    } else {
        var uniform int = nums1[0];
        var swaps int = 0;
        for i := 1; i < len(nums1); i++ {
            if (nums2[i] == uniform && nums1[i] != uniform) {
                swaps++;
            } else if (nums2[i] != uniform && nums1[i] != uniform) {
                return -1;
            }
        }
        return swaps;
    }
}


func _minSwapsSame(a []int, b []int) int {
    return min(_uniformArr(a,b), _uniformArr(b,a));
}

func main() {
    a := []int{4, 2, 4, 2, 4, 4};
    b := []int{2, 4, 2, 4, 2, 2};
    fmt.Println(_minSwapsSame(a,b))
}
