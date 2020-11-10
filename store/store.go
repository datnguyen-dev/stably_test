package store

import (
	"math"
)

func checkIsPrime(num int) bool {
	if (num % 2) != 0 {
		fNum := float64(num)
		sqrt := math.Sqrt(fNum)
		for i := 2; i <= int(sqrt); i++ {
			if (num % i) == 0 {
				return false
			}
		}
		return true
	}
	return false
}

//FindPrime - Find prime number lower input
func FindPrime(input int) int {
	numcheck := 0
	if (input % 2) == 0 {
		numcheck = input - 1
	} else {
		numcheck = input - 2
	}
	for i := numcheck; i > 1; i = i - 2 {
		if checkIsPrime(i) {
			return i
		}
	}
	return 0
}
