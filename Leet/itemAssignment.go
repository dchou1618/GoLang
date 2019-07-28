// file itemAssignment.go

package main

import (
  "fmt"
)

func packItemsWrapper(items []int, bagSizes []int, packed [][]int, pos int) [][]int {
    if (pos >= len(items)) {
        return packed;
    } else {
        for i := 0; i < len(bagSizes); i++ {
            if (bagSizes[i]-items[pos] >= 0) {
                packed[i] = append(packed[i], items[pos]);
                bagSizes[i] -= items[pos];
                result := packItemsWrapper(items, bagSizes, packed, pos+1);
                if (len(result) > 0) {
                    return result;
                } else {
                  bagSizes[i] += items[pos];
                  packed[i] = packed[i][0:len(packed[i])-1];
                }
            }
        }
    }
    return [][]int{};
}

func packItems(items []int, bagSizes []int) [][]int{
    packed := make([][]int, len(bagSizes));
    pos := 0;
    return packItemsWrapper(items, bagSizes, packed, pos);
}

func isValid(n int, dist []int, pos int) bool {
    if (pos+n+1 >= len(dist)) {
        return false;
    } else {
        for _, val := range dist {
            if (val == n) {
                return false;
            }
        }
        if (dist[pos+n+1] == 0 && dist[pos] == 0) {
            return true;
        } else {
          return false;
        }
    }
}

func completeDist(dist []int) bool {
    for _, val := range dist {
        if (val == 0) {
          return false;
        }
    }
    return true;
}

func distWrapper(n int, dist []int, num int) []int {
    if (completeDist(dist)) {
        return dist;
    } else if (num > n) {
        return []int{};
    } else {
        for pos := 0; pos < len(dist); pos++ {
            if (isValid(num, dist, pos)) {
                dist[pos] = num;
                dist[pos+num+1] = num;
                distLst := distWrapper(n, dist, num+1); // once num has been placed
                // place num+1
                if (len(distLst) > 0) {
                    return distLst;
                } else {
                    dist[pos] = 0;
                    dist[pos+num+1] = 0;
                }
            }
        }
    }
    return []int{};
}

func distList(n int) []int {
    dist := make([]int, 2*n);
    return distWrapper(n, dist, 1);
}

func main() {
  items := []int{4, 8, 1, 4, 3};
  bags := []int{12, 3, 4, 3};
  fmt.Println(packItems(items,bags));
  fmt.Println(distList(8));
}
