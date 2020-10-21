package rsa

import "math/rand"

type PublicKey struct {
	n int
	e int
}

type PrivateKey struct {
	n int
	d int
}

func isCoPrime(a int, b int) bool {
	for a > 0 && b > 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	return a+b == 1
}

func NewRSAKeyPair(p int, q int) (PublicKey, PrivateKey) {
	n, totient := p*q, (p-1)*(q-1)
	e := totient
	for ; ; e = rand.Intn(totient-2) + 2 {
		if isCoPrime(e, totient) {
			break
		}
	}
	d := -1
	for k := 1; k < n; k++ {
		x := 1 + k*totient
		if x%e == 0 {
			d = x / e
		}
	}
	return PublicKey{n, e}, PrivateKey{n, d}
}

func PowerModulo(val int, times int, modulo int) int {
	if times == 0 {
		return 1
	}
	result := PowerModulo(val, times/2, modulo)
	result = (result * result) % modulo
	if times%2 == 1 {
		result = (result * val) % modulo
	}
	return result
}
