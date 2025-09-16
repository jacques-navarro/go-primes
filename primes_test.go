package primes

import (
	"slices"
	"testing"
)

var testcases = []struct {
	name string
	want bool
	n    int
}{
	{"2 is prime", true, 2},
	{"3 is prime", true, 3},
	{"5 is prime", true, 5},
	{"7 is prime", true, 7},
	{"11 is prime", true, 11},
	{"13 is prime", true, 13},
	{"17 is prime", true, 17},
	{"19 is prime", true, 19},
	{"23 is prime", true, 23},
	{"29 is prime", true, 29},
	{"31 is prime", true, 31},
	{"1,000,003 is prime", true, 1_000_003},
	{"4 is not prime", false, 4},
	{"6 is not prime", false, 6},
	{"8 is not prime", false, 8},
	{"9 is not prime", false, 9},
	{"10 is not prime", false, 10},
	{"12 is not prime", false, 12},
	{"14 is not prime", false, 14},
	{"15 is not prime", false, 15},
	{"16 is not prime", false, 16},
	{"18 is not prime", false, 18},
	{"20 is not prime", false, 20},
}

func TestIsPrimes(t *testing.T) {

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPrime(tc.n)

			want := tc.want

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}

}

var nextPrimeTestCases = []struct {
	name string
	want int
	n    int
}{
	{"3 is next prime after 3", 3, 3},
	{"5 is next prime after 4", 5, 4},
	{"5 is next prime after 5", 5, 5},
	{"7 is next prime after 6", 7, 6},
	{"11 is next prime after 8", 11, 8},
	{"11 is next prime after 9", 11, 9},
	{"11 is next prime after 10", 11, 10},
	{"11 is next prime after 11", 11, 11},
	{"13 is next prime after 12", 13, 12},
	{"17 is next prime after 14", 17, 14},
	{"17 is next prime after 15", 17, 15},
	{"17 is next prime after 16", 17, 16},
	{"19 is next prime after 18", 19, 18},
}

func TestNextPrime(t *testing.T) {

	for _, tc := range nextPrimeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NextPrime(tc.n)

			want := tc.want

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}

}

var primesInRangeTestCases = []struct {
	name  string
	want  []int
	start int
	end   int
}{
	{"primes in range 10 to 20", []int{11, 13, 17, 19}, 10, 20},
	{"primes in range 21 to 30", []int{23, 29}, 21, 30},
	{"primes in range 31 to 40", []int{31, 37}, 31, 40},
	{"primes in range 41 to 50", []int{41, 43, 47}, 41, 50},
	{"primes in range 10 to 50", []int{11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}, 10, 50},
	{"primes in range 20 to 22", []int{}, 20, 22},
	{"primes in range 38 to 40", []int{}, 38, 40},
}

func TestPrimesInRange(t *testing.T) {

	for _, tc := range primesInRangeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PrimesInRange(tc.start, tc.end)

			want := tc.want

			if !slices.Equal(got, want) {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}

}
