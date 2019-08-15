/* using quickselect to find the kth largest el't
   in a data structure - practice with divide and conquer algorithms
*/

package main

import (
    "fmt"
)

func partitionByMid(nums []int, mid int) []int {
    var lastSmallestIndex int = 0;
    var lastLargestIndex int = len(nums)-1;
    for i := 0; i < len(nums); i++ {
        if (i == lastLargestIndex) {
            break;
        }
        if (nums[i] < mid) {
            var temp int = nums[i];
            nums[i] = nums[lastSmallestIndex];
            nums[lastSmallestIndex] = temp;
            lastSmallestIndex++;
        } else if (nums[i] > mid) {
            var temp int = nums[i];
            nums[i] = nums[lastLargestIndex];
            nums[lastLargestIndex] = temp;
            lastLargestIndex--;
            i--; // move index backwards if nums[i] swapped is less than mid.
        }
    }
    return nums;

}


/* attributed from quicksort's pivot selection
   implement quickselect - quicksort, but stopping at point of finding
   the kth smallest el't - pick a pivot, then partition around it
   - {3,2,1,5,6,4};
   - largest (min-heap - getting rid of smallest el'ts if size exceeds k O(nlog(k)))/
     smallest -> sorting/BST
*/

/* Not possible to find smallest element in sub-linear time unless
   data is stored in heap, bst, (etc.).
*/

func partitionBySmallest(nums *[]int, left int, right int, pivot int) int {
    // move pivot to end
    var temp int = (*nums)[pivot];
    (*nums)[pivot] = (*nums)[right];
    (*nums)[right] = temp;
    newPivot := left;
    for i := left; i < right; i++ {
        if ((*nums)[i] < temp) {
            var swap int = (*nums)[i];
            (*nums)[i] = (*nums)[newPivot];
            (*nums)[newPivot] = swap;
            newPivot++;
        }
    }
    var temp2 int = (*nums)[right];
    (*nums)[right] = (*nums)[newPivot];
    (*nums)[newPivot] = temp2;
    return newPivot;
}

func smallestWrapper(nums []int, left int, right int) int {
    if (left == right) {
        return nums[left];
    } else {
        pivotIndex := (left+right)/2;
        pivotIndex = partitionBySmallest(&nums, left, right, pivotIndex);
        if (pivotIndex == 0) {
            return nums[pivotIndex];
        } else {
            return smallestWrapper(nums, left, pivotIndex-1);
        }
    }
}

func smallest(nums []int) int{
    return smallestWrapper(nums, 0, len(nums)-1);
}


func partitionValues(nums *[]int, left int, right int, pivot int) int {
    var temp int = (*nums)[pivot];
    (*nums)[pivot] = (*nums)[right];
    (*nums)[right] = temp;
    newPivot := left;
    for i := left; i < right; i++ {
        if ((*nums)[i] < temp) {
          /* */
            var pivotSwap int = (*nums)[i];
            (*nums)[i] = (*nums)[newPivot];
            (*nums)[newPivot] = pivotSwap;
            newPivot++;
        }
    }
    /* generating section of length newPivot (values less than nums[pivot])
      pivot is at final position in sorted arrangement
    */
    var last int = (*nums)[right];
    (*nums)[right] = (*nums)[newPivot];
    (*nums)[newPivot] = last;
    return newPivot;
}

func kthLargestWrapper(nums *[]int, k int, left int, right int) int {
    if (left == right) {
        return (*nums)[left];
    } else {
        pivotIndex := (left+right)/2;
        pivotIndex = partitionValues(nums, left, right, pivotIndex);
        if (len(*nums)-k == pivotIndex) {
            return (*nums)[pivotIndex];
        } else if (len(*nums)-k < pivotIndex) {
            return kthLargestWrapper(nums, k, left, pivotIndex-1);
        } else {
            return kthLargestWrapper(nums, k, pivotIndex+1, right);
        }
    }
}

func findKthLargest(nums []int, k int) int {
    return kthLargestWrapper(&nums, k, 0, len(nums)-1);
}

func main() {
    nums := []int{2,7,-3,4,8,9,-12};
    fmt.Println(findKthLargest(nums, 3));
    nums2 := []int{-22,3,-4,7,8,4,-3,-2,-1,-5,33};
    fmt.Println(smallest(nums2));
    nums3 := []int{2,3,1,2,3,2,1,2,3,1,2,2,2,3,1,2};
    fmt.Println(partitionByMid(nums3, 2));
}
