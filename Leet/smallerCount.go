// file smallerCount.go

// binary search

// range minimum query
// consider using self-balanced trees,
// particularly avl trees (aside from segment trees).

// new idea - binary search
// since appending from the end to the beginning,
// all el'ts less than the inserted value will be the # of values that
// inserted number is greater than

func binaryInsertion(nums []int, currSize int, target int) int {
    lower := 0;
    higher := currSize-1;
    mid := (lower+higher)/2;
    for (lower <= higher) {
        if (nums[mid] < target) {
            lower = mid+1;
        } else {
            higher = mid-1;
        }
        mid = (lower+higher)/2;
    }
    return lower;
}


// concept of [5,2,6,1] where el'ts
func countSmaller(nums []int) []int {
    if (len(nums) <= 1) {
        return make([]int, len(nums));
    } else {
        counts := make([]int, len(nums));
        sorted := []int{};
        currSize := 0;
        for i := len(nums)-1; i >= 0; i-- {
            index := binaryInsertion(sorted, currSize, nums[i]);
            counts[i] = index;
            sorted = append(sorted[0:index],append([]int{nums[i]},
                            sorted[index:]...)...);
            currSize++;
        }
        return counts;
    }
}


// attempt at binary search tree
type node struct {
    val []int
    left *node
    right *node
    count int
}

func countNodes(root * node) int {
    if (root != nil) {
        return 1 + countNodes(root.left) + countNodes(root.right);
    }
    return 0;
}

func editCounts(root *node, counts *[]int) {
    if (root != nil) {
        (*counts)[root.val[1]] = root.count;
        editCounts(root.left, counts);
        editCounts(root.right, counts);
    }
    return;
}

// upon insertion, counts can be predetermined
func insertTree(root *node, lst []int) int {
    if (root != nil) {
        count := 0;
        currNode := root;
        parent := root;
        for (currNode != nil) {
            parent = currNode;
            if (lst[0] == currNode.val[0]) {
                currNode.count++;
                return currNode.count;
            } else if (lst[0] > currNode.val[0]) {
                count += currNode.count; // add on all counts lower than the currNode
                currNode = currNode.right;
            } else {
                currNode.count++;
                currNode = currNode.left;

            }
        }
        if (lst[0] > parent.val[0]) {
            parent.right = &node{lst,nil,nil, count};
        } else {
            parent.left = &node{lst, nil, nil, count};
        }
        return count;
    } else {
        root = &node{lst, nil, nil, 0};
        return 0;
    }
}

func countSmaller(nums []int) []int {
    if (len(nums) <= 1) {
        return make([]int, len(nums));
    } else {
        // root := &node{[]int{nums[0],0}, nil, nil, 0};
        var root *node = nil;
        for i := 0; i < len(nums); i++ {
            nums[i] = insertTree(root, []int{nums[i], i});
        }
        //editCounts(root, &nums);

        return nums;
    }
}
