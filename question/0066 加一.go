package question

func plusOne(digits []int) []int {
	carry := 0
	digits[len(digits)-1]++
	if digits[len(digits)-1] >= 10 {
		digits[len(digits)-1] -= 10
		carry = 1
		index := len(digits) - 2
		for carry != 0 && index >= 0 {
			digits[index] = digits[index] + carry
			if digits[index] >= 10 {
				digits[index] -= 10
				index--
			} else {
				carry = 0
			}
		}
		if carry == 1 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
