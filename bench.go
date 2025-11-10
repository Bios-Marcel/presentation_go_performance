package benchmarks

import (
	"io"
	"os"
	"path/filepath"
)

func PutUint32(b []byte, v uint32) {
	// Eliminiert Bounds-Check
	_ = b[3]
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func PutUint32_naive(b []byte, v uint32) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func fileSize_naive(dir string) int {
	var size int

	abs, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	if entries, err := os.ReadDir(abs); err != nil {
		panic(err)
	} else {
		for _, entry := range entries {
			file, err := os.Open(filepath.Join(abs, entry.Name()))
			if err != nil {
				panic(err)
			}
			bytes, err := io.ReadAll(file)
			if err != nil {
				panic(err)
			}

			size += len(bytes)
		}
	}

	return size
}

func fileSize(dir string) int {
	buffer := make([]byte, 4096)
	var size int

	abs, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	if entries, err := os.ReadDir(abs); err != nil {
		panic(err)
	} else {
		for _, entry := range entries {
			file, err := os.Open(filepath.Join(abs, entry.Name()))
			if err != nil {
				panic(err)
			}

			sizeRead, err := file.Read(buffer)
			if err != nil {
				panic(err)
			}

			size += sizeRead
		}
	}

	return size
}

func generateIntSlice_naive(size int) []int {
	var arr []int
	for i := range size {
		arr = append(arr, i)
	}

	return arr
}

func generateIntSlice(size int) []int {
	arr := make([]int, size)
	for i := range size {
		arr[i] = i
	}
	return arr
}

type StructA struct {
	A int8
	B int32
}

type StructB struct {
	A int16
	B int32
}
