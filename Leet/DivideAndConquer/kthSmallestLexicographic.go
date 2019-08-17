// file kthSmallestLexicographic.go

package main

import (
  "fmt"
  "strconv"
)

// conduct binary insertion on lexicographically ordered numbers
// TLE
func compareLex(n1 int, n2 int) string {
    var longer int;
    var shorter int;
    if (len(strconv.Itoa(n1)) > len(strconv.Itoa(n2))) {
        longer = n1;
        shorter = n2;
    } else {
        longer = n2;
        shorter = n1;
    }
    newShort, _ := strconv.Atoi(strconv.Itoa(shorter)+
        strings.Repeat("0",len(strconv.Itoa(longer))-len(strconv.Itoa(shorter))));
    if (newShort > longer) {
        if (shorter == n2) {
            return "second";
        } else {
            return "first";
        }
    } else if (newShort == longer) {
        return "equal";
    } else {
        if (longer == n1) {
            return "first";
        } else {
            return "second";
        }
    }
}

func binaryInsertion(nums []int, k int, newK int) int {
    if (len(nums) >= k) {
        if (compareLex(newK, nums[k-1]) == "first") {
            return -1;
        }
    }
    lower := 0;
    higher := len(nums)-1;
    mid := (lower+higher)/2;
    for (lower <= higher) {
        if (compareLex(nums[mid], newK) == "first") {
            higher = mid - 1;
        } else {
            lower = mid + 1;
        }
        mid = (lower+higher)/2;
    }
    return lower;
}

func findKthNumber(n int, k int) int {
    lexicographic := []int{};
    for i := 1; i <= n; i++ {
        index := binaryInsertion(lexicographic, k, i);
        if (index != -1) {
            lexicographic = append(lexicographic[0:index],
                            append([]int{i},lexicographic[index:]...)...);
            if (len(lexicographic) > k) {
                lexicographic = lexicographic[0:len(lexicographic)-1];
                // -1 from length because push off last el't
            }
        }
    }
    return lexicographic[k-1];
}
