package problem067

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	a := assert.New(t)

	a.Equal("100", addBinary("11", "1"))
	a.Equal("10110", addBinary("1011", "1011"))
	a.Equal("10101", addBinary("1010", "1011"))
	a.Equal("10000", addBinary("101", "1011"))
	a.Equal("10001", addBinary("1", "10000"))
	a.Equal("0", addBinary("0", "0"))
	a.Equal("1", addBinary("0", "1"))
	a.Equal("10", addBinary("1", "1"))
	a.Equal("110", addBinary("11", "11"))
	a.Equal("1000", addBinary("111", "1"))
}

func BenchmarkAddBinary(b *testing.B) {
	a := "1011"
	bin := "1010"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		addBinary(a, bin)
	}
}
