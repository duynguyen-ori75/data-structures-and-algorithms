package benchmark

import (
	"math/rand"
	"testing"
	"sync"
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
	t.ResetTimer()
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
	t.ResetTimer()
	for test := 0; test < t.N; test++ {
		for i := 0; i < dim; i++ {
			for j := 0; j < dim; j++ {
				b[i][j] = a[i][j]
			}
		}
	}
}

type Simple struct {
	val int64
}

type Padding struct {
	val int64
	_ [56]byte  // 64 bytes (cacheline) - 8 bytes (val size)
}


func BenchmarkSimpleStruct(t *testing.B) {
	a := Simple{val: 0}
	var wg sync.WaitGroup
	t.ResetTimer()
	for test := 0; test < t.N; test ++ {
		wg.Add(4)
		for thread := 0; thread < 4; thread ++ {
			go func() {
				for i := 0; i < 1000000; i ++ {
					a.val += int64(i)
				}
				wg.Done()
			} ()
		}
		wg.Wait()
	}
}

func BenchmarkPaddingStruct(t *testing.B) {
	a := Padding{val: 0}
	var wg sync.WaitGroup
	t.ResetTimer()
	for test := 0; test < t.N; test ++ {
		wg.Add(4)
		for thread := 0; thread < 4; thread ++ {
			go func() {
				for i := 0; i < 1000000; i ++ {
					a.val += int64(i)
				}
				wg.Done()
			} ()
		}
		wg.Wait()
	}
}