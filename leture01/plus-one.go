package leture01

func plusOne(digits []int) []int {
	for index := len(digits) - 1; index >= 0; index-- {
		digits[index]++
		digits[index] = digits[index] % 10
		if digits[index] != 0 {
			return digits
		}
	}
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}
