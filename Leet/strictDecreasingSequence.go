// file strictDecreasingSequence.go
// construct a max heap so all children will be less than parent

package main

import (
    "fmt"
)

type MaxHeap struct {
    elts [][]int
    size int
}

func makeHeap() MaxHeap {
    this := new(MaxHeap);
    this.elts = [][]int{[]int{}};
    this.size = 0;
    return *this;
}

func (this *MaxHeap) addToHeap(val []int) {
    this.elts = append(this.elts, val);
    var index int = this.size;
    this.size++;
    for (index/2 > 0 && this.elts[index][0] > this.elts[index/2][0]) {
        var temp []int = this.elts[index];
        this.elts[index] = this.elts[index/2];
        this.elts[index/2] = temp;
        index /= 2;
    }
}

func decreasingSubsequence(sequence []int) (int, [][]int) {
    newHeap := makeHeap();
    subsequences := [][]int{};
    decreasingSeqs := 0;
    for i, val := range sequence {
        newHeap.addToHeap([]int{val, i});
    }
    //
    return decreasingSeqs, subsequences;
}

/* Only O(n^2) efficiency */
func greedyDecreasingSubsequence(sequence []int) (int, [][]int) {
    subsequences := [][]int{};
    seqs := 0;
    for _, val := range sequence {
        addedToSeq := false;
        for i, seq := range subsequences {
            if (seq[len(seq)-1] > val) {
                subsequences[i] = append(subsequences[i], val);
                addedToSeq = true;
                break;
            }
        }
        if (!addedToSeq) {
            subsequences = append(subsequences, []int{val});
            seqs++;
        }
    }
    return seqs, subsequences;
}



func main() {
    seq := []int{2, 9, 12, 13, 4, 7, 6, 5, 10};
    fmt.Println(greedyDecreasingSubsequence(seq));
}
