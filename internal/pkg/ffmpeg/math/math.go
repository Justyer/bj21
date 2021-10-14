package math

// 最接近n的最小偶数
// eg： CeilOdd(41) is 40, CeilOdd(30) is also 30.
func CeilOdd(n int64) int64 {
	if n%2 == 0 {
		return n
	}
	return n - 1
}
