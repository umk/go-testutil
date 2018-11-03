package testrandom

// A 28-bit linear-feedback shift register which generates an m-sequence. Use
// this to produce a pseudo-random sequence of numbers to implement reproducible
// tests.
//
// The register must be initialized with a value, which is greater to zero and
// less than Lfsr28Max value.
type Lfsr28 uint32

// The exclusive max value, produced by Lfsr28.
const Lfsr28Max uint32 = 0x10000000

// Next returns the next pseudo-random number.
func (r *Lfsr28) Next() uint32 {
	n := *r
	t := (n ^ (n >> 13)) & 1
	*r = ((n >> 1) | (t << 27))
	return uint32(*r)
}

// DefaultLfsr28 creates a new linear-feedback shift register with the initial
// value 1.
func DefaultLfsr28() *Lfsr28 {
	r := Lfsr28(1)
	return &r
}
