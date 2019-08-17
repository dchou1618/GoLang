/* distributedMedianServer.go file asked by Amazon */
package main

import (
    "fmt"
)

func partitionTowardsMedian(medians *[]int, left int, right int, pivot int) int {
    // returns new pivot index.
    var end int = (*medians)[pivot];
    (*medians)[pivot] = (*medians)[right];
    (*medians)[right] = end;
    newPivot := left; // keeping track of values below pivot value
    for i := left; i < right; i++ {
        if ((*medians)[i] < end) {
            var temp int = (*medians)[i];
            (*medians)[i] = (*medians)[newPivot];
            (*medians)[newPivot] = temp;
            newPivot++;
        }
    }
    // swap pivot position with the end
    var temp2 int = (*medians)[newPivot];
    (*medians)[newPivot] = (*medians)[right];
    (*medians)[right] = temp2;

    return newPivot;
}

func getMedian(medians []int, left int, right int, index int) int {
    if (left == right) {
        return medians[left];
    } else {
        pivotIndex := (left+right)/2;
        pivotIndex = partitionTowardsMedian(&medians, left, right, pivotIndex);
        if (pivotIndex < index) {
            return getMedian(medians, pivotIndex+1, right, index);
        } else if (pivotIndex > index) {
            return getMedian(medians, left, pivotIndex-1, index);
        } else {
            return medians[pivotIndex]
        }
    }
}

/* given n sorted server lists, obtain distributed median */
func _distributedMedianServer(servers [][]int) float64 {
    var serverValues []int;
    for _, server := range servers { // O(s).
        for _, val := range server { // O(n/s)
            serverValues = append(serverValues, val);
        }
    }
    var grandMedian float64;
    if (len(serverValues)%2 == 0) {
        var left int = getMedian(serverValues, 0, len(serverValues)-1, (len(serverValues)/2)-1);
        var right int = getMedian(serverValues, 0, len(serverValues)-1, len(serverValues)/2);
        grandMedian = float64(left + right)/float64(2);
    } else {
        grandMedian = float64(getMedian(serverValues, 0, len(serverValues)-1, len(serverValues)/2));
    }
    return grandMedian;
}


func main() {
    servers := [][]int{[]int{1,3,4,5,6},[]int{2,3,5,6},[]int{1,12,14,16,27,38}};
    fmt.Println(_distributedMedianServer(servers));
}
