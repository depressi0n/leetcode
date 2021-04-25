package question

func isHappy202_1(n int) bool {
	v := 0
	cur := n
	see := make(map[int]struct{})
	see[n] = struct{}{}
	for {
		v = 0
		for cur != 0 {
			v += (cur % 10) * (cur % 10)
			cur /= 10
		}
		if v == 1 {
			return true
		}
		if _, ok := see[v]; ok {
			return false
		}
		see[v] = struct{}{}
		cur = v
	}
}

//TODO:快慢指针
func isHappy202_2(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
