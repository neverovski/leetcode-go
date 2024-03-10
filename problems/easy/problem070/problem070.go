package problem070

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}

	oneStepBefore, twoStepBefore := 1, 0

	for i := 1; i <= n; i++ {
		temp := oneStepBefore

		oneStepBefore = oneStepBefore + twoStepBefore
		twoStepBefore = temp
	}

	return oneStepBefore
}
