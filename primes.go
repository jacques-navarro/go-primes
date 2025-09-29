// Package primes provides functions for finding
// prime numbers.
package primes

import (
	"fmt"
	"math"
)

type StartGreaterThanEndError struct {
	start int
	end int
}

// StartGreaterThanEndError is returned when a function that works
// with a range receives a end argument which is less than or equal
// to the start argument
func (e StartGreaterThanEndError) Error() string {
	return fmt.Sprintf("end %d, must be greater than start %d", e.end, e.start)
}

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

	for i := 3; i <= midPoint; i += 2 {

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

// SumOfPrimesInRange returns the sum of all prime numbers
// between a given range from start to end.
func SumOfPrimesInRange(start, end int) int {
	primes := PrimesInRange(start, end)

	var sum int

	for _, p := range primes {
		sum += p
	}

	return sum
}

// LastPrimeInRange returns the last prime number
// in a given range from start to end.
func LastPrimeInRange(start int, end int) int {
	var lastPrime int

	for n := end; n >= start; {
		if IsPrime(n) {
			lastPrime = n
			break
		}

		if n%2 != 0 {
			n -= 2
			continue
		}

		if n%2 == 0 {
			n--
			continue
		}
	}

	return lastPrime
}

// PrimeGap returns how many numbers exist between consecutive prime numbers.
// If the input is a prime number, PrimeGap returns the next prime number minus the input.
// When the input is 37, a prime number, then the prime gap is the next prime number (41) minus 37, which is 4.
//
// In cases where the input is not a prime number, the previous prime number is first found,
// and then the difference between the next prime number is returned minus the previous one.
// When the input is 32, a composite number. The previous prime number is first found (31). The next prime number after
// the inpur (37) is then found and the previous prime number is then subtracted. In this case 37 - 31 = 6.
func PrimeGap(input int) int {

	if IsPrime(input) {
		return NextPrime(input+1) - input
	}

	for i := input - 1; ; i-- {
		if IsPrime(i) {
			return NextPrime(input) - i
		}
	}

}
