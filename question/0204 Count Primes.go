package question

func countPrimes204_1(n int) int {
	//isPrime:= func(n int) bool{
	//	if n<4{
	//		return true
	//	}
	//	for i:=2;i*i<=n;i++{
	//		if i*(n/i)==n{
	//			return false
	//		}
	//	}
	//	return true
	//}
	////枚举
	//cnt:=0
	//for i:=2;i<n;i++{
	//	if isPrime(i) {
	//		cnt++
	//	}
	//}
	//return cnt
	cnt := 0
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return cnt
}

//TODO:线性筛
func countPrimes204_2(n int) int {
	primes := []int{}
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}
