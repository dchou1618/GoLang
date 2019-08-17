// file insertPosition.go

package main

import (
  "fmt"
  "math"
  "sort"
)

// 4 ms - 92.01%, 3.1 mb - 99.17%
// binary search implementation
func searchInsert(nums []int, target int) int {
    lowerBound, upperBound := 0, len(nums)-1;
    mid := (lowerBound+upperBound)/2;
    for (lowerBound <= upperBound) {
        if (nums[mid] == target) {
            return mid;
        } else if (nums[mid] < target) {
            lowerBound = mid+1;
        } else {
            upperBound = mid-1;
        }
        mid = (lowerBound+upperBound)/2;
    }
    return lowerBound;
}


func searchRange(nums []int, target int) []int {
    result := []int{-1,-1};
    var lower int = 0;
    var higher int = len(nums)-1;
    var mid int = (lower+higher)/2;
    for (lower <= higher) {
        if (nums[mid] == target) {
            bound1 := mid;
            bound2 := mid;
            for (bound1-1 >= 0 && nums[bound1-1] == target) {
                bound1--;
            }
            for (bound2+1 < len(nums) && nums[bound2+1] == target) {
                bound2++;
            }
            return []int{bound1,bound2};
        } else if (nums[mid] < target) {
            lower = mid + 1;
        } else {
            higher = mid - 1;
        }
        mid = (lower+higher)/2;
    }
    return result;
}

func minSwapsArray(nums []int) int {
    var sorted []int;
    sorted = append(sorted, nums...);
    sort.Ints(sorted);
    sortedIndices := make(map[int]int);
    for i, val := range sorted {
        sortedIndices[val] = i; // assuming values in nums are distinct
    }
    visited := make(map[int]bool); // visited indices
    minSwaps := 0;
    for i := range nums {
        temp := i;
        cycleLength := 0;
        for (!visited[temp]) {
            visited[temp] = true;
            temp = sortedIndices[nums[temp]];
            cycleLength++;
        }
        if (cycleLength != 0) {
            minSwaps += cycleLength-1;
        }
    }
    return minSwaps;
}

// with items {1,2,...,n}, how many moves to front are required
func minMovesToFront(nums []int) int {
    var moves int = 0;
    for i := range nums {
        if (nums[i] != i+1) {
            moves++;
        }
    }
    return moves;
}




// sort array with 3 el'ts in linear runtime. Constant memory?
// "sorting colors"

// idea is to grow the values of 2 from the top in the array
// and 0 values from the lower part of the array
// aside from using counting sort (overwrite the array with 0,1,2)
func sortThree(nums []int)  {
    var i int = 0;
    var j int = 0;
    var n int = len(nums)-1;
    for (j <= n) {
        if (nums[j] < 1) { // splitting values by mid value
            var temp int = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
            i++; // swapping lower values
            j++;
        } else if (nums[j] > 1) {
            var temp int = nums[j];
            nums[j] = nums[n];
            nums[n] = temp;
            n--;
        } else {
            j++; // maintaining i <= j
        }
    }
}

func sortTwoMids(arr []int, mid1 int, mid2 int) {
    var firstThird int = 0;
    var index int = 0;
    var thirdThird int = 0;
    for (index <= thirdThird) {
        if (arr[index] < 3) {
            var temp int = arr[index];
            arr[index] = arr[firstThird];
            arr[firstThird] = temp;
            index++;
            firstThird++;
        } else {

        }
        if (arr[index] > 3) {
            var temp int = arr[index];
            arr[index] = arr[thirdThird];
            arr[thirdThird] = temp;
            thirdThird--;
        } else {
            index++;
        }
    }
}

func minSubArrayLen(s int, nums []int) int {
    // using two pointer method where you trace one index
    minLength := math.MaxInt64;
    pointer1 := 0; // first el't
    sum := 0;
    for i := range nums { // using two pointer
        sum += nums[i];
        for (sum >= s) { // then sum too large -> reset
            if (i+1-pointer1 < minLength) {
                // length of current interval from first pointer to index i
                minLength = i+1-pointer1;
            }
            sum -= nums[pointer1];
            pointer1++; // moving pointer along;
        }
    }
    if (minLength == math.MaxInt64) {
        return 0;
    }
    return minLength;
}


func maxSubArraySumK(sum int, nums []int) int{
    maxLength := -1;
    currSum := 0;
    left := 0;
    for i := range nums {
        currSum += nums[i];
        for (currSum >= sum) {
            if (i+1-left > maxLength) {
                maxLength = i+1-left;
            }
            currSum -= nums[left];
            left++;
        }
    }
    return maxLength;
}


/* Plan: search for minimum value in the array
   indices -> (i+minIndex)%(len(nums)+1);

*/

func getMinIndex(nums []int) int {
    if (len(nums) == 0) {
        return 0;
    } else {
        var min int = nums[0];
        var minIndex int = 0;
        for i, val := range nums {
            if (val < min) {
                min = val;
                minIndex = i;
            }
        }
        return minIndex;
    }
}

func search(nums []int, target int) int {
    minIndex := getMinIndex(nums);
    var lower, upper int = minIndex, (len(nums)-1+minIndex);
    var mid int = (lower+upper)/2;
    for (lower <= upper) {
        if (nums[mid%len(nums)] == target) {
            return mid%len(nums);
        } else if (nums[mid%len(nums)] < target) { // look in upper half
            lower = mid+1;
        } else {
            upper = mid-1;
        }
        mid = (lower+upper)/2;
    }
    return -1;
}

// executing binary search on optimized fewer comparisons

func nextLowestTarget(nums []int, target int) int {
    lower, upper := 0, len(nums);
    var mid int;
    for (upper-lower > lower) {
        mid = lower+(upper-lower)/2;
        if (nums[mid] <= target) {
            lower = mid;
        } else {
            upper = mid;
        }
    }
    return nums[lower];
}

/* sorted array with duplicates - finding occurrences in list
   motivation is to find the leftmost and rightmost occurrence
*/

func numOccurrences(nums []int, target int) int {
    var lowerMax, upperMax int = 0, len(nums);
    var lowerMin, upperMin int = -1, len(nums)-1;
    var midMax int;
    var midMin int;
    for (upperMin-lowerMin > lowerMin) {
        midMin = lowerMin+(upperMin-lowerMin)/2;
        if (nums[midMin] <= target) {
            lowerMin = midMin;
        } else {
            upperMin = midMin;
        }
    }
    for (upperMax-lowerMax > lowerMax) {
        midMax = lowerMax+(upperMax-lowerMax)/2;
        if (nums[midMax] >= target) {
            upperMax = midMax;
        } else {
            lowerMax = midMax;
        }
    }
    if (lowerMin < 0 && upperMax < 0) {
        return 0;
    } else {
        return lowerMax-upperMin+1;
    }
}


func main() {
    nums := []int{2,3,1,7,4,6};
    sorted := []int{1,2,5,6,7,9,12,15};
    //occur := []int{1,1,2,2,3,4,4,4,5,5,5,6,6,7,8,8,8,8,8,8,9,9,9,10,10,10,13};
    fiveColors := []int{4,2,0,3,2,1,2,0,2,1,1,2,2,0,3,2,4,3,1,2,4};
    sortTwoMids(fiveColors, 1, 3)
    fmt.Println(fiveColors);
    fmt.Println(maxSubArraySumK(7, nums))
    fmt.Println(minSwapsArray(nums));
    fmt.Println(nextLowestTarget(sorted, 8));
    //fmt.Println(numOccurrences(occur, 1));
}
