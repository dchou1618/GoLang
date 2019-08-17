// file distinctSubsequences.go


func subsequences(S string, memo map[string]bool, sub string, pos int, sequences *int) {
    if (pos == len(S)) {
        if (!memo[sub] && sub != "") {
            (*sequences)++;
            memo[sub] = true;
        }
    } else {
        subsequences(S, memo, sub+string(S[pos]), pos+1, sequences);
        subsequences(S, memo, sub, pos+1, sequences);
    }
}

func distinctSubseqII(S string) int {
    subs := 0;
    memo := make(map[string]bool);
    subsequences(S, memo, "", 0, &subs);
    return subs%1000000007;
}
