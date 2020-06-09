package benchmark

import (
	"math/rand"
	"testing"
)

const dim = 5000

func construct2DArray() ([][]uint8, [][]uint8) {
	a := make([][]uint8, dim)
	b := make([][]uint8, dim)
	for i := range a {
		a[i] = make([]uint8, dim)
		b[i] = make([]uint8, dim)
		for j := 0; j < dim; j++ {
			a[i][j] = uint8(rand.Uint64() / 256)
		}
	}
	return a, b
}

func BenchmarkMissCacheline(t *testing.B) {
	a, b := construct2DArray()
	for test := 0; test < t.N; test++ {
		for i := 0; i < dim; i++ {
			for j := 0; j < dim; j++ {
				b[j][i] = a[j][i]
			}
		}
	}

}

func BenchmarkHitCacheline(t *testing.B) {
	a, b := construct2DArray()
	for test := 0; test < t.N; test++ {
		for i := 0; i < dim; i++ {
			for j := 0; j < dim; j++ {
				b[i][j] = a[i][j]
			}
		}
	}
}
