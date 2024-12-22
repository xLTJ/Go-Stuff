package collatzconjecture

import "errors"

// CollatzConjecture calculates the number of steps required to reach 1 using the collatz conjecture algorithm
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n cannot be negative")
	}

	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}

	return steps, nil
}
