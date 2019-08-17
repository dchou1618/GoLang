// letterCombos.go file
// simple exhaustive search of letterCombinations

package letterMain

import (
  "fmt"
)

func combineLetters(digits string, letters []string,
                    combos *[]string, curr string) {
    if (len(digits) == 0) {
        *combos = append(*combos, curr);
    } else {
        possibleLetters := letters[digits[0] - '1'];
        for i := 0; i < len(possibleLetters); i++ {
            combineLetters(digits[1:],letters, combos,
                           curr+string(possibleLetters[i]));
        }
    }
}

func letterCombinations(digits string) []string {
    if (len(digits) == 0) {
        return []string{};
    } else {
        letters := []string{"","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"};
        combos := []string{};
        combineLetters(digits, letters, &combos, "");
        letters = nil;
        return combos;
    }
}
