/* productArrayExceptSelf.go file */


package main

import (
    "fmt"
)
/*
  Planning - 1 [2] 3 4 - partition by value? divide and conquer
  multiplied 1 and 3*4
  T(N) = aT(N/b) + [work for merge];
    - b parts

*/

func maximum(a int, b int) int {
    if (a > b) {
        return a;
    } else {
        return b;
    }
}

func getProduct(nums []int, left int, right int, output *[]int) int {
    if (left == right) {
        return nums[left];
    } else if (left > right) {
        return 1;
    }
    var mid int = (left+right)/2;

    leftProduct := getProduct(nums, left, mid, output);
    rightProduct := getProduct(nums, mid+1, right, output);
    for i := left; i <= mid; i++ {
        (*output)[i] *= rightProduct;
    }
    for i := mid+1; i <= right; i++ {
        (*output)[i] *= leftProduct;
    }
    return leftProduct*rightProduct;
}

func productDAC(nums []int, left int, right int, leftMax *int, rightMax *int, max *int, product *int) {
    if (left == right) {
        *leftMax = nums[left];
        *rightMax = nums[left];
        *max = nums[left];
        *product = nums[left];
    } else {
        var mid int = (left+right)/2;
        var productLeft, productRight int;
        var maxLeftLeft, maxLeftRight int;
        var maxRightLeft, maxRightRight int;
        var maxLeft, maxRight int;
        productDAC(nums, left, mid, &maxLeftLeft, &maxLeftRight, &maxLeft, &productLeft);
        productDAC(nums, mid+1, right, &maxRightLeft, &maxRightRight, &maxRight, &productRight);
        *max = maximum(maximum(maxLeft, maxRight), maxLeftRight*maxRightLeft);
        *leftMax = maximum(maxLeftLeft, productLeft*maxRightLeft);
        *rightMax = maximum(maxRightRight, productRight*maxLeftRight);
        *product = productLeft*productRight;
    }
}

func maxProductArray(nums []int) int {
    var maxProduct int;
    var product int = 1;
    var leftMax, rightMax int;
    productDAC(nums, 0, len(nums)-1, &leftMax, &rightMax, &maxProduct, &product);
    return maxProduct;
}

func _productExceptSelf(nums []int) []int {
    output := make([]int, len(nums));
    for i,_ := range output {
        output[i] = 1;
    }
    getProduct(nums, 0, len(nums)-1, &output);
    return output;
}

/* O(N) sol'n with left and right products (at any index i,
    when left/right meet the product will be product of values surrounding i) */
func productExceptSelf(nums []int) []int {
    var left int = 0;
    var currLeftProduct = 1;
    var right int = len(nums)-1;
    var currRightProduct = 1;
    output := make([]int, len(nums));
    for i,_ := range output {
        output[i] = 1;
    }
    for (right >= 0 && left < len(nums)) {
        output[left] *= currLeftProduct;
        output[right] *= currRightProduct;
        currLeftProduct *= nums[left];
        currRightProduct *= nums[right];
        left++;
        right--;
    }
    return output;
}


func main() {
    nums := []int{2,-5,-2,-4,3};
    fmt.Println(maxProductArray(nums));
}
