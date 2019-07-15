// file
package main

import (
  "fmt"
)


func binarySearch(lst [][]int, target []int) (int) {
    low, high := 0, len(lst)-1;
    mid := (low+high)/2;
    for (low <= high) {
        if (lst[mid][0] == target[0]) {
            fmt.Println(lst[mid],target);
            if (lst[mid][1] == target[1]) {
                return mid;
            } else if (lst[mid][1] > target[1]) {
                high = mid-1;
            } else {
                low = mid+1;
            }
        } else if (lst[mid][0] < target[0]) {
            low = mid+1;
        } else {
            high = mid-1;
        }
        mid = (low+high)/2;
    }
    return low;
}
func sortLst(intervals [][]int) [][]int {
    finalLst := [][]int{};
    for i := 0; i < len(intervals); i++ {
        index := binarySearch(finalLst,intervals[i]);
        newLst := [][]int{};
        for j := 0; j < index; j++ {
            newLst = append(newLst,finalLst[j]);
        }
        newLst = append(newLst,intervals[i]);
        for k := index; k < len(finalLst); k++ {
            newLst = append(newLst,finalLst[k]);
        }
        finalLst = newLst;
    }
    return finalLst;
}

func merge(intervals [][]int) [][]int {
    intervals = sortLst(intervals);
    i := 0;
    for (i < len(intervals)-1) {
        if (intervals[i][1] >= intervals[i+1][0]) {
            newLst := [][]int{};
            for j := 0; j < i; j++ {
                newLst = append(newLst,intervals[j]);
            }
            if (intervals[i][1] > intervals[i+1][1]) {
                newLst = append(newLst, []int{intervals[i][0],intervals[i][1]});
            } else {
                newLst = append(newLst, []int{intervals[i][0],intervals[i+1][1]});
            }
            for k := i+2; k < len(intervals); k++ {
                newLst = append(newLst,intervals[k]);
            }
            intervals = newLst;
            i--;
        }
        i++;
    }
    return intervals;
} // sorting inplace can allow for O(nlogn) time complexity
// although 
