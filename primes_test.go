package primes

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
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

var nextPrimeIsNotTestCases = []struct {
	name     string
	dontWant int
	input    int
}{
	{"10 is not the next prime after 10", 10, 10},
	{"12 is not the next prime after 10", 12, 10},
	{"13 is not the next prime after 10", 13, 10},
	{"9 is not the next prime after 10", 9, 10},
	{"12 is not the next prime after 12", 12, 12},
	{"14 is not the next prime after 12", 14, 12},
	{"15 is not the next prime after 12", 15, 12},
	{"16 is not the next prime after 12", 16, 12},
	{"17 is not the next prime after 12", 17, 12},
	{"18 is not the next prime after 18", 18, 18},
	{"20 is not the next prime after 18", 20, 18},
	{"21 is not the next prime after 18", 21, 18},
}

func TestNextPrimeIsNot(t *testing.T) {
	for _, tc := range nextPrimeIsNotTestCases {
		t.Run(tc.name, func(t *testing.T) {

			got := NextPrime(tc.input)

			dontWant := tc.dontWant

			if diff := cmp.Diff(got, tc.dontWant); diff == "" {
				t.Errorf("got %d should not equal want %d", got, dontWant)
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
	{"primes in range 524 to 540", []int{}, 524, 540},
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

var primesInRangeAreNotEqualTestCases = []struct {
	name     string
	dontWant []int
	start    int
	end      int
}{
	{"primes in range 10 t0 19 are not", []int{}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{10, 11, 13, 17, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{11, 13, 17, 18, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{12, 13, 17, 18, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{13, 14, 17, 18, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{13, 16, 17, 18, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{13, 17, 18, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{11}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{11, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{11, 13, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{11, 17, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{13, 17, 19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{13}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{17}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{19}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{10, 12, 14, 16, 18}, 10, 19},
	{"primes in range 10 t0 19 are not", []int{10, 12, 14, 15, 16, 18}, 10, 19},
}

func TestPrimesInRangeAreNotEqual(t *testing.T) {

	for _, tc := range primesInRangeAreNotEqualTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PrimesInRange(tc.start, tc.end)

			dontWant := tc.dontWant

			if slices.Equal(got, dontWant) {
				t.Errorf("got %v shouldn't equal %v", got, dontWant)
			}
		})
	}

}

var sumOfPrimesInRangeCases = []struct {
	name  string
	want  int
	start int
	end   int
}{
	{"sum of primes in range 10 t0 19", 60, 10, 19},
	{"sum of primes in range 20 t0 29", 52, 20, 29},
	{"sum of primes in range 30 t0 39", 68, 30, 39},
	{"sum of primes in range 40 t0 49", 131, 40, 49},
	{"sum of primes in range 10 t0 49", 311, 10, 49},
	{"sum of primes in range 24 t0 428", 0, 24, 28},
	{"sum of primes in range 32 t0 36", 0, 32, 36},
	{"sum of primes in range 524 t0 540", 0, 524, 540},
}

func TestSumOfPrimesInRange(t *testing.T) {

	for _, tc := range sumOfPrimesInRangeCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumOfPrimesInRange(tc.start, tc.end)

			want := tc.want

			if got != want {
				t.Errorf("got %d want %d", got, tc.want)
			}
		})
	}

}

var sumOfPrimesInRangeIsNotCases = []struct {
	name     string
	dontWant int
	start    int
	end      int
}{
	{"sum of primes in range 10 t0 19 is not 0", 0, 10, 19},
	{"sum of primes in range 10 t0 19 is not 11", 11, 10, 19},
	{"sum of primes in range 10 t0 19 is not 13", 13, 10, 19},
	{"sum of primes in range 10 t0 19 is not 17", 17, 10, 19},
	{"sum of primes in range 10 t0 19 is not 19", 19, 10, 19},
	{"sum of primes in range 10 t0 19 is not 64", 64, 10, 19},
	{"sum of primes in range 10 t0 19 is not 24", 24, 10, 19},
	{"sum of primes in range 10 t0 19 is not 41", 41, 10, 19},
	{"sum of primes in range 10 t0 19 is not 10", 10, 10, 19},
	{"sum of primes in range 10 t0 19 is not 145", 145, 10, 19},
}

func TestSumOfPrimesInRangeIsNot(t *testing.T) {
	for _, tc := range sumOfPrimesInRangeIsNotCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumOfPrimesInRange(tc.start, tc.end)

			dontWant := tc.dontWant

			if got == dontWant {
				t.Errorf("got %d should not equal %d", got, dontWant)
			}
		})
	}

}

var lastPrimeInRangeCases = []struct {
	name  string
	want  int
	start int
	end   int
}{
	{"last prime in range 10 to 19", 19, 10, 19},
	{"last prime in range 20 to 29", 29, 20, 29},
	{"last prime in range 30 to 39", 37, 30, 39},
	{"last prime in range 40 to 49", 47, 40, 49},
	{"last prime in range 50 to 59", 59, 50, 59},
	{"last prime in range 10 to 59", 59, 10, 59},
	{"last prime in range 523 to 541", 541, 523, 541},
	{"last prime in range 523 to 540", 523, 523, 540},
}

func TestLastPrimeInRange(t *testing.T) {
	for _, tc := range lastPrimeInRangeCases {
		t.Run(tc.name, func(t *testing.T) {
			got := LastPrimeInRange(tc.start, tc.end)

			want := tc.want

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}

}

var lastPrimeInRangeIsNotTestCases = []struct {
	name     string
	dontWant int
	start    int
	end      int
}{
	{"last prime in range 10 to 19 is not 20", 20, 10, 19},
	{"last prime in range 10 to 19 is not 18", 18, 10, 19},
	{"last prime in range 10 to 19 is not 17", 17, 10, 19},
	{"last prime in range 10 to 19 is not 16", 16, 10, 19},
	{"last prime in range 10 to 19 is not 15", 15, 10, 19},
	{"last prime in range 10 to 19 is not 14", 14, 10, 19},
	{"last prime in range 10 to 19 is not 13", 13, 10, 19},
	{"last prime in range 10 to 19 is not 12", 12, 10, 19},
	{"last prime in range 10 to 19 is not 11", 11, 10, 19},
	{"last prime in range 10 to 19 is not 10", 10, 10, 19},
}

func TestLastPrimeInRangeIsNot(t *testing.T) {
	for _, tc := range lastPrimeInRangeIsNotTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := LastPrimeInRange(tc.start, tc.end)

			dontWant := tc.dontWant

			if got == dontWant {
				t.Errorf("got %d should not equal %d", got, dontWant)
			}
		})
	}

}

var primeGapTestCases = []struct {
	name  string
	want  int
	input int
}{
	{"prime gap with input 31 should be 6", 6, 31},
	{"prime gap with input 32 should be 6", 6, 32},
	{"prime gap with input 33 should be 6", 6, 33},
	{"prime gap with input 34 should be 6", 6, 34},
	{"prime gap with input 35 should be 6", 6, 35},
	{"prime gap with input 36 should be 6", 6, 36},
	{"prime gap with input 37 should be 4", 4, 37},
	{"prime gap with input 38 should be 4", 4, 38},
	{"prime gap with input 39 should be 4", 4, 39},
	{"prime gap with input 40 should be 4", 4, 40},
	{"prime gap with input 41 should be 2", 2, 41},
	{"prime gap with input 42 should be 2", 2, 42},
	{"prime gap with input 43 should be 4", 4, 43},
}

func TestPrimeGap(t *testing.T) {
	for _, tc := range primeGapTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PrimeGap(tc.input)

			want := tc.want

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}

}

var primeGapShouldNotBeTestCases = []struct {
	name     string
	dontWant int
	input    int
}{
	{"prime gap with input 31 should not be 0", 0, 31},
	{"prime gap with input 31 should not be 1", 1, 31},
	{"prime gap with input 31 should not be 2", 2, 31},
	{"prime gap with input 31 should not be 3", 3, 31},
	{"prime gap with input 31 should not be 4", 4, 31},
	{"prime gap with input 31 should not be 5", 5, 31},
	{"prime gap with input 32 should not be 0", 0, 32},
	{"prime gap with input 32 should not be 1", 1, 32},
	{"prime gap with input 32 should not be 2", 2, 32},
	{"prime gap with input 32 should not be 3", 3, 32},
	{"prime gap with input 32 should not be 4", 4, 32},
	{"prime gap with input 32 should not be 5", 5, 32},
	{"prime gap with input 33 should not be 0", 0, 33},
	{"prime gap with input 33 should not be 1", 1, 33},
	{"prime gap with input 33 should not be 2", 2, 33},
	{"prime gap with input 33 should not be 3", 3, 33},
	{"prime gap with input 33 should not be 4", 4, 33},
	{"prime gap with input 33 should not be 5", 5, 33},
	{"prime gap with input 34 should not be 0", 0, 34},
	{"prime gap with input 34 should not be 1", 1, 34},
	{"prime gap with input 34 should not be 2", 2, 34},
	{"prime gap with input 34 should not be 3", 3, 34},
	{"prime gap with input 34 should not be 4", 4, 34},
	{"prime gap with input 34 should not be 5", 5, 34},
	{"prime gap with input 35 should not be 0", 0, 35},
	{"prime gap with input 35 should not be 1", 1, 35},
	{"prime gap with input 35 should not be 2", 2, 35},
	{"prime gap with input 35 should not be 3", 3, 35},
	{"prime gap with input 35 should not be 4", 4, 35},
	{"prime gap with input 35 should not be 5", 5, 35},
	{"prime gap with input 36 should not be 0", 0, 36},
	{"prime gap with input 36 should not be 1", 1, 36},
	{"prime gap with input 36 should not be 2", 2, 36},
	{"prime gap with input 36 should not be 3", 3, 36},
	{"prime gap with input 36 should not be 4", 4, 36},
	{"prime gap with input 36 should not be 5", 5, 36},
	{"prime gap with input 37 should not be 0", 0, 37},
	{"prime gap with input 37 should not be 1", 1, 37},
	{"prime gap with input 37 should not be 2", 2, 37},
	{"prime gap with input 37 should not be 3", 3, 37},
}

func TestPrimeGapShouldNotBe(t *testing.T) {
	for _, tc := range primeGapShouldNotBeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PrimeGap(tc.input)

			dontWant := tc.dontWant

			if diff := cmp.Diff(got, dontWant); diff == "" {
				t.Errorf("got %d should not equal %d", got, dontWant)
			}
		})
	}

}
