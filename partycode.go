package partycode

// TODO(fmrsn): Write documentation.

// Rand is a source of random numbers.
type Rand interface {
	// Intn returns, as an int, a non-negative random number in the half-open
	// interval [0,n).
	Intn(n int) int
}

type Generator struct {
	rand   Rand
	chars  []byte
	digits int
	limit  int
}

func New(rand Rand, chars []byte, digits int) Generator {
	base := len(chars)
	limit := base
	for i := 0; i < digits; i++ {
		limit *= base
	}
	return Generator{rand, chars, digits, limit}
}

func (g *Generator) Generate() string {
	b := make([]byte, g.digits)
	g.generate(b)
	return string(b)
}

func (g *Generator) generate(b []byte) {
	n := g.rand.Intn(g.limit)
	chars, base := g.chars, len(g.chars)
	for i := g.digits - 1; i >= 0; i-- {
		b[i] = chars[n%base]
		n /= base
	}
}
