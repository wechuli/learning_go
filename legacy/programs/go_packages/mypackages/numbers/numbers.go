package numbers

func Even(n uint) bool {
	return n%2 == 0
}

func odd(n uint) bool {
	return n%2 != 0
}

func isPrime(num uint) bool {
	for i := uint(2); i < num/2; i++ {
		if num%i == uint(0) {
			return false
		}
	}
	return true
}

func OddAndPrime(n uint) bool {
	return odd(n) && isPrime(n)
}
