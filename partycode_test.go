package partycode_test

import (
	"strconv"
	"strings"
	"testing"

	"bsid.es/partycode"
)

type riggedRand struct {
	i int
}

func (r *riggedRand) Intn(n int) int {
	return r.i % n
}

func TestGenerate(t *testing.T) {
	const max = 9 * 9 * 9 * 9 * 9 * 9 // 9^6
	r := &riggedRand{}
	gen := partycode.New(r, []byte("123456789"), 6)
	for i := 0; i < max; i++ {
		r.i = i
		want := formatBase9With6Digits(i)
		got := gen.Generate()
		if want != got {
			t.Errorf("wrong party code\nwant: %s\ngot:  %s", want, got)
		}
	}
}

func BenchmarkGenerate(b *testing.B) {
	r := &riggedRand{i: 0}
	gen := partycode.Default(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen.Generate()
	}
}

func formatBase9With6Digits(i int) string {
	s := strconv.FormatInt(int64(i), 9)
	if s == "0" {
		return "111111"
	}
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < 6-len(s); i++ {
		b.WriteByte('1')
	}
	for i := 0; i < len(s); i++ {
		b.WriteByte(s[i] + 1)
	}
	s = b.String()
	return s
}
