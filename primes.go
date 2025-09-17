// Package primes provides functions for finding
// prime numbers.
package primes

import (
	"math"
)

// IsPrime reports whether n is a prime number.
// It implements the [Trial division] method.
//
// [Trial division]: https://en.wikipedia.org/wiki/Trial_division
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

// NextPrime returns the next prime number given
// n. If n is a prime number it is returned.
func NextPrime(n int) int {

	for i := n; true; i++ {
		if IsPrime(i) {
			return i
		}
	}

	return 0

}

// PrimesInRange returns a slice of all prime
// numbers in a given range from start to end.
func PrimesInRange(start int, end int) []int {

	var primes []int

	for i := start; i <= end; i++ {

		if IsPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}
