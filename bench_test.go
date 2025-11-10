package benchmarks

import (
	"math/rand/v2"
	"testing"
	"unsafe"
)

var random_needles []int
var random_stack []int

var matrix = make([][]float64, 10000)

func init() {
	random_stack = make([]int, 10000)
	for i, _ := range random_stack {
		random_stack[i] = rand.IntN(1000)
	}
	random_needles = make([]int, 100)
	for i, _ := range random_needles {
		random_needles[i] = rand.IntN(3000) + 1500
	}

	for i := range matrix {
		matrix[i] = make([]float64, 10000)
	}
}

func Benchmark_PutUint32(b *testing.B) {
	var buf = make([]byte, 4)
	b.Skip()
	for b.Loop() {
		PutUint32(buf, 32)
	}
}

func Benchmark_PutUint32_naive(b *testing.B) {
	var buf = make([]byte, 4)
	b.Skip()
	for b.Loop() {
		PutUint32_naive(buf, 32)
	}
}

var int_sink int

func Benchmark_fileSize(b *testing.B) {
	for b.Loop() {
		int_sink = fileSize("files")
	}
}

func Benchmark_fileSize_naive(b *testing.B) {
	for b.Loop() {
		int_sink = fileSize_naive("files")
	}
}

var int_slice_sink []int

func Test_intSlicesEqual(t *testing.T) {
	a := generateIntSlice_naive(2)
	b := generateIntSlice(2)
	if len(a) != len(b) {
		t.FailNow()
	}

	if a[0] != b[0] || a[1] != b[1] {
		t.FailNow()
	}
}

// Pre-Allocation (Still on Heap)

func Benchmark_generateIntSlice(b *testing.B) {
	for b.Loop() {
		int_slice_sink = generateIntSlice(256)
	}
}

func Benchmark_generateIntSlice_naive(b *testing.B) {
	for b.Loop() {
		int_slice_sink = generateIntSlice_naive(256)
	}
}

// Cache Locality

func Benchmark_matrixAccess(b *testing.B) {
	for b.Loop() {
		for j := range 10000 {
			for i := range 10000 {
				matrix[j][i] += 1
			}
		}
	}
}

func Benchmark_matrixAccess_naive(b *testing.B) {
	for b.Loop() {
		for j := range 10000 {
			for i := range 10000 {
				matrix[i][j] += 1
			}
		}
	}
}

func Test_StructAlignment(t *testing.T) {
	t.Log("A:", unsafe.Sizeof(StructA{}))
	t.Log("B:", unsafe.Sizeof(StructB{}))
}
