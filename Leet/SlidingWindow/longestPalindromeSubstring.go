// file longestPalindromeSubstring.go

package main 

func isPalindrome(s string) (bool) {
    for i := 0; i < len(s)/2; i++ {
        if (s[i] != s[len(s)-1-i]) {
            return false;
        }
    }
    return true;
}
func longestPalinWrapper(s string, longest *string) {
    for size := 1; size <= len(s); size++ { // version of sliding window
        for pos := 0; pos <= len(s)-size; pos++ {
            if (isPalindrome(s[pos:pos+size])) {
                if (len(s[pos:pos+size]) > len(*longest)) {
                    *longest = s[pos:pos+size];
                }
            }
        }
    }
}
func longestPalindrome(s string) string {
    longest := "";
    longestPalinWrapper(s, &longest);
    return longest;
}
