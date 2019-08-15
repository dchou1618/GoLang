// file roseBloom.go

package main
import (
  "fmt"
)
/*
Given an array of roses. roses[i] means rose i will bloom on day roses[i].
Also given an int k, which is the minimum number of adjacent bloom roses
required for a bouquet, and an int n, which is the number of bouquets we need.
Return the earliest day that we can get n bouquets of roses.
*/
// Google interview problem
// Design plan - iteration while approach
func getNumBouquets(roses []int, day int, bouquetSize int) int {
    bouquets := 0;
    currSize := 0;
    var bouquetStart bool = false;
    for _, rose := range roses {
        if (rose <= day) {
            if (!bouquetStart) {
                bouquetStart = true;
            }
            if (bouquetStart){
                currSize++;
            }
        } else {
            if (bouquetStart) {
                bouquetStart = false;
                if (currSize >= bouquetSize) {
                    bouquets++;
                }
                currSize = 0;
            }
        }
    }
    if (bouquetStart) {
        bouquets++;
    }
    return bouquets;
}

func max(roses []int) int {
    maximum := 0;
    for _, val := range roses {
        if (maximum < val) {
            maximum = val;
        }
    }
    return maximum;
}

func minDaysBloom(roses []int, k int, n int) (int) {
    minDay, maxDay := 0, max(roses);
    midDay := (minDay+maxDay)/2;
    for (minDay <= maxDay) {
        currBouquets := getNumBouquets(roses, midDay, k);
        fmt.Println(currBouquets, midDay);
        if (currBouquets < n) {
            minDay = midDay+1;
        } else if (currBouquets > n){
            maxDay = midDay-1;
        } else {
            return midDay;
        }
        midDay = (minDay+maxDay)/2
    }
    return -1;
}

func main() {
  var roses []int = []int{1,2,4,9,3,4,1,3,4,6,5,2,7,6,5};
  var k int = 2;
  var n int = 3;
  fmt.Println(minDaysBloom(roses,k,n));
}
