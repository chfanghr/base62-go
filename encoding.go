package base62_go

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const base = 62

type Encoding struct {
	encode  string
	padding int
}

// NewEncoding returns a new Encoding defined by the given alphabet
func NewEncoding(encoder string) *Encoding {
	return &Encoding{
		encode: encoder,
	}
}

// EncodeBytes returns the base62 encoding of b
func (e *Encoding) EncodeBytes(b []byte) string {
	n := new(big.Int)
	n.SetBytes(b)
	return e.EncodeBigInt(n)
}

// EncodeInt64 returns the base62 encoding of n
func (e *Encoding) EncodeInt64(n int64) string {
	var (
		b   = make([]byte, 0)
		rem int64
	)

	// Progressively divide by base, store remainder each time
	// Prepend as an additional character is the higher power
	for n > 0 {
		rem = n % base
		n = n / base
		b = append([]byte{e.encode[rem]}, b...)
	}

	s := string(b)
	if e.padding > 0 {
		s = e.pad(s, e.padding)
	}

	return s
}

// EncodeBigInt returns the base62 encoding of an arbitrary precision integer
func (e *Encoding) EncodeBigInt(n *big.Int) string {
	var (
		b    = make([]byte, 0)
		rem  = new(big.Int)
		bse  = new(big.Int)
		zero = new(big.Int)
	)
	bse.SetInt64(base)
	zero.SetInt64(0)

	// Progressively divide by base, until we hit zero
	// store remainder each time
	// Prepend as an additional character is the higher power
	for n.Cmp(zero) == 1 {
		n, rem = n.DivMod(n, bse, rem)
		b = append([]byte{e.encode[rem.Int64()]}, b...)
	}

	s := string(b)
	if e.padding > 0 {
		s = e.pad(s, e.padding)
	}

	return s
}

// DecodeToBytes returns a byte array from a base62 encoded string
func (e *Encoding) DecodeToBytes(s string, padding ...int) []byte {
	nBytes := e.DecodeToBigInt(s).Bytes()
	if len(padding) > 0 && padding[0] > 0 && len(nBytes) < padding[0] {
		paddingBytes := make([]byte, padding[0]-len(nBytes))
		nBytes = append(paddingBytes, nBytes...)
	}
	return nBytes
}

// DecodeToInt64 decodes a base62 encoded string
func (e *Encoding) DecodeToInt64(s string) int64 {
	var (
		n     int64
		c     int64
		idx   int
		power int
	)

	for i, v := range s {
		idx = strings.IndexRune(e.encode, v)

		// Work downwards through powers of our base
		power = len(s) - (i + 1)

		// Calculate value at this position and add
		c = int64(idx) * int64(math.Pow(float64(base), float64(power)))
		n = n + c
	}

	return int64(n)
}

// DecodeToBigInt returns an arbitrary precision integer from the base62 encoded string
func (e *Encoding) DecodeToBigInt(s string) *big.Int {
	var (
		n = new(big.Int)

		c     = new(big.Int)
		idx   = new(big.Int)
		power = new(big.Int)
		exp   = new(big.Int)
		bse   = new(big.Int)
	)
	bse.SetInt64(base)

	// Run through each character to decode
	for i, v := range s {
		// Get index/position of the rune as a big int
		idx.SetInt64(int64(strings.IndexRune(e.encode, v)))

		// Work downwards through exponents
		exp.SetInt64(int64(len(s) - (i + 1)))

		// Calculate power for this exponent
		power.Exp(bse, exp, nil)

		// Multiplied by our index, gives us the value for this character
		c = c.Mul(idx, power)

		// Finally add to running total
		n.Add(n, c)
	}

	return n
}

// pad a string to a minimum length with zero characters
func (e *Encoding) pad(s string, minlen int) string {
	if len(s) >= minlen {
		return s
	}

	format := fmt.Sprint(`%0`, strconv.Itoa(minlen), "s")
	return fmt.Sprintf(format, s)
}

// SetEncodePadding sets the padding of encoded string
func (e *Encoding) SetEncodePadding(n int) {
	e.padding = n
}
