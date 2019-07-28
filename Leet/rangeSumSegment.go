// file rangeSumSegment.go implements a segment tree


package main

import (
  "fmt"
)


// count of range sum - segment tree

// bottom-up tree

/*      __2__
       /[0,2] \
  ___3___      \
 / [0,1] |      \
-2       5       -1
[0,0]  [1,1]    [2,2]
*/

type node struct {
    min int
    max int
    count int // count of values within range of min and max sums
    // interval start/end not important
    left *node
    right *node
}

func inBounds(num int, lower int, upper int) bool {
    return (lower <= num && num <= upper);
}

// func printTree(root *node, depth int) {
//     if (root != nil) {
//         printTree(root.left, depth+1);
//         for i := 0; i < depth; i++ {
//             fmt.Print(" ");
//         }
//         fmt.Print("[", root.min,",",root.max,"]");
//         fmt.Println();
//         printTree(root.right, depth+1);
//     }
// }

// binary divide and construct
func construct(start int, end int, sums []int) *node {
    root := &node{sums[start],sums[end], 0, nil, nil}; // prefixed sums (start min), (end max)
    if (start == end) {
        return root;
    }
    mid := (start+end)/2;
    root.left = construct(start, mid, sums);
    root.right = construct(mid+1, end, sums);
    return root;
}

func sumInRange(root *node, sum int) {
    if (root != nil) {
        if (inBounds(sum, root.min, root.max)) { // if sum between root.min and root.max
            root.count++; // prefixed sum lies between min and max
            if (root.min == root.max) { // reached a leaf
                return
            } else {
                sumInRange(root.left, sum);
                sumInRange(root.right, sum);
            }
        }
    }
}

func getRangeSums(root *node, sum int, low int, up int) int {
    if (root != nil) {
        if (root.min-sum >= low && root.max-sum <= up) { // difference in sums between low and up
            return root.count;
        }
        return getRangeSums(root.left, sum, low, up) + getRangeSums(root.right, sum, low, up);
    }
    return 0;
}

func countRangeSum(nums []int, lower int, upper int) int {
    if (len(nums) == 0) {
        return 0;
    } else {
        sum := 0;
        set := make(map[int]bool);
        sums := []int{};
        for _, val := range nums { // tracing prefix sums
            sum += val;
            if (!set[sum]) {
                sums = append(sums, sum);
                set[sum] = true;
            }
        }
        set = nil;
        sort.Ints(sums);
        // using prefix sums, find sums[i]-sums[j] where it lies in the range
        root := construct(0, len(sums)-1, sums); // O(logn) construction
        sums = nil;
        count := 0;
        for i := len(nums)-1; i >= 0; i-- { // O(n)
            sumInRange(root, sum); // current prefix sum
            sum -= nums[i]; // not to subtract sums but
            // by each el't to get preceding prefix sums
            count += getRangeSums(root, sum, lower, upper);
        }
        return count;
    }
}


func main() {

}
