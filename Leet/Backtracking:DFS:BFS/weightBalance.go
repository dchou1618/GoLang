// file weightBalance.go


package main

import (
	"fmt"
)

func balanceWrapper(target int, weights []int, currIndex int) (bool) {
    if (target == 0) {
        return true;
    } else if (currIndex == len(weights)) {
        return false;
    } else {
        // either add on the right or left side (-weights[currIndex] and weights[currIndex])
        return balanceWrapper(target-weights[currIndex], weights, currIndex+1) ||
        balanceWrapper(target+weights[currIndex], weights, currIndex+1) ||
        balanceWrapper(target, weights, currIndex+1);
    }
}

func canBalance(target int, weights []int) (bool) {
    index := 0;
    return balanceWrapper(target, weights, index);
}

func main() {
	target := 5;
	weights := []int{1,2,6};
	fmt.Println(canBalance(target,weights));
}
