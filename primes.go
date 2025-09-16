package primes

import (
	"math"
)

func IsPrime(n int) bool {

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	midPoint := int(math.Sqrt(float64(n)))

	for i := 3; i <= midPoint+1; i += 2 {

		if n%i == 0 {
			return false
		}
	}

	return true

}

func NextPrime(n int) int {

	for i := n; true; i++ {
		if IsPrime(i) {
			return i
		}
	}

	return 0

}

func PrimesInRange(start int, end int) []int {

	var primes []int

	for i := start; i <= end; i++ {
		
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}
