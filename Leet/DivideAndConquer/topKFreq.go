// file topKFreq.go



package main

import (
  "fmt"
)

func topKFrequent(nums []int, k int) []int {
    counts := make(map[int]int);
    var max int;
    for _, val := range nums {
        counts[val]++;
        if (counts[val] > max) {
            max = counts[val];
        }
    }
    kFrequent := make([][]int, max+1);
    for elem, freq := range counts {
        //fmt.Println(elem, kFrequent);
        kFrequent[freq] = append(kFrequent[freq], elem);
    }
    topFreq := []int{};
    kItems := 0;
    for freq := max; freq >= 0; freq-- {
        for _, val := range kFrequent[freq] {
            topFreq = append(topFreq, val);
            kItems++;
            if (kItems == k) {
                kFrequent = nil;
                return topFreq;
            }
        }
    }
    kFrequent = nil;
    return topFreq;
}

func main() {
  
}
