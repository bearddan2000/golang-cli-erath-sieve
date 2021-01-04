package main

import "fmt"

func sieveOfEratosthenes(n int) {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	prime := make([]bool, n+1, 2*(n+1))
	for i := 0; i < n; i++ {
		prime[i] = true
	}

	p := 2

	for p*p <= n {
		// If prime[p] is not changed, then it is a prime
		if prime[p] == true {
			// Update all multiples of p
			for i := p * p; i < n; i += p {
				prime[i] = false
			}
		}
		p += 1
	}

	// Print all prime numbers
	fmt.Print("[OUTPUT] ")
	for i := 2; i < n; i++ {
		if prime[i] == true {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}
func main() {
	input := 20
	fmt.Printf("[INPUT] %d\n", input)
	sieveOfEratosthenes(input)
}
