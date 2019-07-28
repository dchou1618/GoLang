// file findKClosestElements.go

package main

import (
  "fmt"
)

// priority queue?
// k-closest el'ts maxheap (difference from largest to smallest).
// highest to lowest priority
type kClosest struct {
    closest []*closeElement
    size int
    capacity int
}

type closeElement struct {
    priority int
    value int
}

func abs(num int) int {
    if (num < 0) {
       num *= -1;
    }
    return num;
}

func (pq *kClosest) swap(pos1 int, pos2 int) {
    var elem *closeElement = pq.closest[pos1];
    pq.closest[pos1] = pq.closest[pos2];
    pq.closest[pos2] = elem;
}

func (pq *kClosest) dequeue() int {
    if (pq.size == 0) {
        return -1;
    }
    var front *closeElement = pq.closest[1];
    pq.closest[1] = pq.closest[pq.size];
    pq.closest[pq.size] = &closeElement{0, 0};
    pq.size--;
    currIndex := 1;
    for (2*currIndex <= pq.size) {
        if (2*currIndex+1 > pq.size) {
            if (pq.closest[2*currIndex].priority > pq.closest[currIndex].priority) {
                pq.swap(currIndex, 2*currIndex);
                currIndex *= 2;
            } else if (pq.closest[2*currIndex].priority == pq.closest[currIndex].priority) {
                if (pq.closest[2*currIndex].value < pq.closest[currIndex].value) {
                    pq.swap(currIndex, 2*currIndex);
                    currIndex *= 2;
                } else {
                    break;
                }
            } else {
                break;
            }
        } else {
            if (pq.closest[currIndex].priority > pq.closest[2*currIndex].priority &&
                pq.closest[currIndex].priority > pq.closest[2*currIndex+1].priority) {
                break;
            } else if (pq.closest[2*currIndex].priority == pq.closest[2*currIndex+1].priority){
                if (pq.closest[2*currIndex].value < pq.closest[2*currIndex+1].value) {
                    pq.swap(currIndex, 2*currIndex);
                    currIndex *= 2;
                } else {
                    pq.swap(currIndex, 2*currIndex+1);
                    currIndex = 2*currIndex + 1;
                }
            } else if (pq.closest[2*currIndex].priority > pq.closest[2*currIndex+1].priority) {
                pq.swap(currIndex, 2*currIndex);
                currIndex *= 2;
            } else {
                pq.swap(currIndex, 2*currIndex+1);
                currIndex = 2*currIndex + 1;
            }
        }
    }
    return front.value;
}

func (pq *kClosest) enqueue(elem *closeElement) {
    if (pq.size == pq.capacity - 1) { // if there are already k el'ts
        if (pq.closest[1].priority < elem.priority) {
            return;
        } else if (pq.closest[1].priority == elem.priority) {
            if (pq.closest[1].value > elem.value) {
                pq.closest[1] = elem;
            }
        } else {
            pq.dequeue();
            currIndex := pq.size+1;
            pq.closest[currIndex] = elem;
            for (currIndex > 1) {
                if (pq.closest[currIndex].priority > pq.closest[currIndex/2].priority) {
                    pq.swap(currIndex, currIndex/2);
                    currIndex /= 2;
                } else if (pq.closest[currIndex].priority == pq.closest[currIndex/2].priority) {
                    if (pq.closest[currIndex].value < pq.closest[currIndex/2].value) {
                        pq.swap(currIndex, currIndex/2);
                        currIndex /= 2;
                    } else {
                        break;
                    }
                } else {
                    break;
                }
            }
            pq.size++;
        }
    } else {
        currIndex := pq.size+1;
        pq.closest[currIndex] = elem;
        for (currIndex > 1) {
            if (pq.closest[currIndex].priority > pq.closest[currIndex/2].priority) {
                pq.swap(currIndex, currIndex/2);
                currIndex /= 2;
            } else if (pq.closest[currIndex].priority == pq.closest[currIndex/2].priority) {
                if (pq.closest[currIndex].value < pq.closest[currIndex/2].value) {
                    pq.swap(currIndex, currIndex/2);
                    currIndex /= 2;
                } else {
                    break;
                }
            } else {
                break;
            }
        }
        pq.size++;
    }
}

func (pq *kClosest) isEmpty() bool {
    return pq.size == 0;
}

func findClosestElements(arr []int, k int, x int) []int {
    pq := &kClosest{make([]*closeElement, k+1), 0, k+1};
    for i := 0; i < len(arr); i++ {
        pq.enqueue(&closeElement{abs(x-arr[i]), arr[i]});
    }
    foundClosest := []int{};
    for i := 1; i < len(pq.closest); i++ { // O(n) 2nd traversal
        foundClosest = append(foundClosest, pq.closest[i].value);
    }
    pq = nil;
    sort.Ints(foundClosest);
    return foundClosest;
    // order by priority = abs(x-), val = arr[i]; 0 <= i && i < len(arr);
}

// minor tweak to solution
func binaryInsertion(nums []int, k int, newK int, x int) int {
    lower := 0;
    higher := len(nums)-1;
    mid := (lower+higher)/2;
    for (lower <= higher) {
        if (abs(nums[mid]-x) > abs(newK-x)) {
            higher = mid - 1;
        } else {
            lower = mid + 1;
        }
        mid = (lower+higher)/2;
    }
    return lower;
}

func findClosestElements(arr []int, k int, x int) []int {
    var kthNumbers []int;
    for _,val := range arr {
        index := binaryInsertion(kthNumbers, k, val, x);
        if (index != -1) {
            kthNumbers = append(kthNumbers[0:index],
                            append([]int{val},kthNumbers[index:]...)...);
            if (len(kthNumbers) > k) {
                kthNumbers = kthNumbers[:len(kthNumbers)-1]; // remove the
                // kth furthest value.
                // -1 from length because push off last el't
            }
        }
    }
    found := []int{};
    for _, val := range kthNumbers {
        found = append(found, val);
    }
    sort.Ints(found);
    return found;
}
