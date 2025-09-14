package primes

import "testing"

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
