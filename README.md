# go-primes

## Overview

go-primes is a simple module for working with prime numbers.

> A prime number (or a prime) is a natural number greater than 1 that is not a product of two smaller natural numbers <sup>[[1]](#1)</sup>.

This module uses the Trial Division algorithm to determine if a number is prime <sup>[[2]](#2)</sup>.

While prime numbers are infinite, integer representation in Go is limited to 64 bits. That means the largest signed 64-bit integer in Go is 9,223,372,036,854,775,807 <sup>[[3]](#3)</sup>.

What is the largest signed 64-bit prime number?

<details>
  <summary>Spoiler warning</summary>
  
  9,223,372,036,854,775,783
  
</details>


## Features

- Determine if a given number is prime
- Given a number, determine what the next prime number is
- Find all primes in a given range
- Find the sum of all primes in a given range
- Find the last prime number in a given range

## Prerequisites
- Requires Go 1.24.3

## Usage

### Add import path

```
import "github.com/jacques-navarro/go-primes"
```

### Add dependency to module

```
go get github.com/jacques-navarro/go-primes
```

### Example code

```
package main

import (
    "fmt"

    // module imported under gp alias 
    gp "github.com/jacques-navarro/go-primes"
)

func main() {
    input := 31

    if prime := gp.IsPrime(input); prime {
        fmt.Printf("%d is a prime number.\n", input)
        return
    }

    fmt.Printf("%d is not a prime number.\n")

}
```

## Changelog

[CHANGELOG](CHANGELOG.md)

## License

go-primes is licensed under the [MIT License](LICENSE).

## References

<a id="1">[1]</a>
https://en.wikipedia.org/wiki/Prime_number

<a id="2">[2]</a>
https://en.wikipedia.org/wiki/Trial_division

<a id="3">[3]</a>
https://go.dev/ref/spec#Numeric_types