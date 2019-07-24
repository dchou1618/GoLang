import (
    "fmt"
)

type Node struct {
    val int
    left *Node
    right *Node
}

// // credit of code to CS DOJO.
// func countUnivals2(root *Node) (int) {
//     totalCount, _ := helper(root);
//     return totalCount;
// }
//
// func helper(root *Node) (int, bool) {
//     if (root == nil) {
//         return 0, true;
//     }
//     leftCount, isLeftUnival := helper(root.left);
//     rightCount, isRightUnival := helper(root.right);
//     isUnival := true;
//     if (!isLeftUnival || !isRightUnival) {
//         isUnival = false;
//     }
//     if (root.left != nil && root.left.val != root.val) {
//         isUnival = false;
//     }
//     if (root.right != nil && root.right.val != root.val) {
//         isUnival = false;
//     }
//     if (isUnival) {
//         return leftCount+rightCount+1, true;
//     } else {
//         return leftCount+rightCount, false;
//     }
// }
// could have sufficed without boolean values
// O(n) sol'n
func numUVT(root *Node) (int, bool) {
    if (root != nil) {

        leftCount, isLeftUnival := numUVT(root.left);
        rightCount, isRightUnival := numUVT(root.right);
        if (root.left == nil && root.right == nil) {
            return 1, true;
        } else if (root.left != nil && root.right != nil &&
          root.left.val == root.val && root.left.val == root.right.val
          && isLeftUnival && isRightUnival) {
            return leftCount+rightCount+1,true;
        } else {
            return leftCount+rightCount, false;
        }
    }
    return 0, true;
}

// tree must be balanced
func numUniversalTrees(root *Node) (int){
    if (root == nil) {
        return 0;
    } else {
        count, _ := numUVT(root);
        return count;
    }
}

func main() {
    root := &Node{1,nil,nil};
    root.left = &Node{1,nil,nil};
    root.right = &Node{1,nil,nil};
    root.right.left = &Node{1,nil,nil};
    root.right.left.left = &Node{1,nil,nil};
    root.right.left.right = &Node{1,nil,nil};
    root.right.right = &Node{1,nil,nil};
    root.right.right.right = &Node{1,nil,nil};
    root.right.right.left = &Node{1,nil,nil};
    fmt.Println(numUniversalTrees(root));
    fmt.Println(countUnivals2(root));
}
