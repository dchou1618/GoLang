// file intervalIntersection.go

package main

import (
  "fmt"
)

// Only O(N^2) with 48 ms (40.43%) runtime and 8.1 MB (66.67%) memory

func mergeLsts(A [][]int, B [][]int) [][]int { // mergesort idea
    merged := [][]int{};
    i,j := 0, 0;
    for (i < len(A) && j < len(B)) {
        if (A[i][0] > B[j][0]) {
            merged = append(merged,B[j]);
            j++;
        } else if (A[i][0] < B[j][0]) {
            merged = append(merged,A[i]);
            i++;
        } else {
            if (A[i][1] > B[j][1]) {
                merged = append(merged, B[j]);
                j++;
            } else if (A[i][1] < B[j][1]) {
                merged = append(merged, A[i]);
                i++;
            } else {
                merged = append(merged,A[i]);
                merged = append(merged,B[j]);
                i++;
                j++;
            }
        }
    }
    for (i < len(A)) {
        merged = append(merged,A[i]);
        i++;
    }
    for (j < len(B)) {
        merged = append(merged,B[j]);
        j++;
    }
    return merged;
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
    mergedLst := mergeLsts(A,B);
    merged := [][]int{};
    for i := 0; i < len(mergedLst)-1; i++ {
        for j := i+1; j < len(mergedLst); j++ {
            if (mergedLst[i][0] <= mergedLst[j][0] && mergedLst[j][0] <= mergedLst[i][1])    {
                if (mergedLst[j][1] <= mergedLst[i][1]) {
                    merged = append(merged, []int{mergedLst[j][0],mergedLst[j][1]});
                } else {
                    merged = append(merged, []int{mergedLst[j][0],mergedLst[i][1]});
                }
            } else {
                break;
            }
        }
    }
    return merged;
}
