package palindrome_partitioning

func Partition(s string) [][]string {
	var ctx context
	backtrack(&ctx, s)
	return ctx.result
}

type context struct {
	stack  []string
	result [][]string
}

func backtrack(ctx *context, s string) {
	if n := len(s); n > 0 {
		for i := 1; i <= n; i++ {
			sub := s[:i]
			if isPalindrome(sub) {
				ctx.stack = append(ctx.stack, sub)
				backtrack(ctx, s[i:])
				ctx.stack = ctx.stack[:len(ctx.stack)-1]
			}
		}
	} else {
		oneResult := make([]string, len(ctx.stack))
		copy(oneResult, ctx.stack)
		ctx.result = append(ctx.result, oneResult)
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
