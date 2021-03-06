// Copyright 2019 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package generated implements field arithmetics with field modulus q =
//
// 52068286795115451399219983130826905585508868087773948221737060012681594112383
//
// Code generated by goff DO NOT EDIT
// Element04 are assumed to be in Montgomery form in all methods
package generated

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
	"math/bits"
	"sync"

	"unsafe"
)

// Element04 represents a field element stored on 4 words (uint64)
// Element04 are assumed to be in Montgomery form in all methods
type Element04 [4]uint64

// Element04Limbs number of 64 bits words needed to represent Element04
const Element04Limbs = 4

// Element04Bits number bits needed to represent Element04
const Element04Bits = 255

// SetUint64 z = v, sets z LSB to v (non-Montgomery form) and convert z to Montgomery form
func (z *Element04) SetUint64(v uint64) *Element04 {
	z[0] = v
	z[1] = 0
	z[2] = 0
	z[3] = 0
	return z.ToMont()
}

// Set z = x
func (z *Element04) Set(x *Element04) *Element04 {
	z[0] = x[0]
	z[1] = x[1]
	z[2] = x[2]
	z[3] = x[3]
	return z
}

// SetZero z = 0
func (z *Element04) SetZero() *Element04 {
	z[0] = 0
	z[1] = 0
	z[2] = 0
	z[3] = 0
	return z
}

// SetOne z = 1 (in Montgomery form)
func (z *Element04) SetOne() *Element04 {
	z[0] = 13428944898394483970
	z[1] = 10320482683557142871
	z[2] = 10268145133511654010
	z[3] = 1856830769744930999
	return z
}

// Neg z = q - x
func (z *Element04) Neg(x *Element04) *Element04 {
	if x.IsZero() {
		return z.SetZero()
	}
	var borrow uint64
	z[0], borrow = bits.Sub64(2508899587657533823, x[0], 0)
	z[1], borrow = bits.Sub64(13286502731930980180, x[1], borrow)
	z[2], borrow = bits.Sub64(4089299470098948802, x[2], borrow)
	z[3], _ = bits.Sub64(8294956651982310308, x[3], borrow)
	return z
}

// Div z = x*y^-1 mod q
func (z *Element04) Div(x, y *Element04) *Element04 {
	var yInv Element04
	yInv.Inverse(y)
	z.Mul(x, &yInv)
	return z
}

// Equal returns z == x
func (z *Element04) Equal(x *Element04) bool {
	return (z[3] == x[3]) && (z[2] == x[2]) && (z[1] == x[1]) && (z[0] == x[0])
}

// IsZero returns z == 0
func (z *Element04) IsZero() bool {
	return (z[3] | z[2] | z[1] | z[0]) == 0
}

// field modulus stored as big.Int
var _element04ModulusBigInt big.Int
var onceelement04Modulus sync.Once

func element04ModulusBigInt() *big.Int {
	onceelement04Modulus.Do(func() {
		_element04ModulusBigInt.SetString("52068286795115451399219983130826905585508868087773948221737060012681594112383", 10)
	})
	return &_element04ModulusBigInt
}

// Inverse z = x^-1 mod q
// Algorithm 16 in "Efficient Software-Implementation of Finite Fields with Applications to Cryptography"
// if x == 0, sets and returns z = x
func (z *Element04) Inverse(x *Element04) *Element04 {
	if x.IsZero() {
		return z.Set(x)
	}

	// initialize u = q
	var u = Element04{
		2508899587657533823,
		13286502731930980180,
		4089299470098948802,
		8294956651982310308,
	}

	// initialize s = r^2
	var s = Element04{
		11987161547385592410,
		17714078913370961707,
		17074890819590993293,
		2162564708293599491,
	}

	// r = 0
	r := Element04{}

	v := *x

	var carry, borrow, t, t2 uint64
	var bigger, uIsOne, vIsOne bool

	for !uIsOne && !vIsOne {
		for v[0]&1 == 0 {

			// v = v >> 1
			t2 = v[3] << 63
			v[3] >>= 1
			t = t2
			t2 = v[2] << 63
			v[2] = (v[2] >> 1) | t
			t = t2
			t2 = v[1] << 63
			v[1] = (v[1] >> 1) | t
			t = t2
			v[0] = (v[0] >> 1) | t

			if s[0]&1 == 1 {

				// s = s + q
				s[0], carry = bits.Add64(s[0], 2508899587657533823, 0)
				s[1], carry = bits.Add64(s[1], 13286502731930980180, carry)
				s[2], carry = bits.Add64(s[2], 4089299470098948802, carry)
				s[3], _ = bits.Add64(s[3], 8294956651982310308, carry)

			}

			// s = s >> 1
			t2 = s[3] << 63
			s[3] >>= 1
			t = t2
			t2 = s[2] << 63
			s[2] = (s[2] >> 1) | t
			t = t2
			t2 = s[1] << 63
			s[1] = (s[1] >> 1) | t
			t = t2
			s[0] = (s[0] >> 1) | t

		}
		for u[0]&1 == 0 {

			// u = u >> 1
			t2 = u[3] << 63
			u[3] >>= 1
			t = t2
			t2 = u[2] << 63
			u[2] = (u[2] >> 1) | t
			t = t2
			t2 = u[1] << 63
			u[1] = (u[1] >> 1) | t
			t = t2
			u[0] = (u[0] >> 1) | t

			if r[0]&1 == 1 {

				// r = r + q
				r[0], carry = bits.Add64(r[0], 2508899587657533823, 0)
				r[1], carry = bits.Add64(r[1], 13286502731930980180, carry)
				r[2], carry = bits.Add64(r[2], 4089299470098948802, carry)
				r[3], _ = bits.Add64(r[3], 8294956651982310308, carry)

			}

			// r = r >> 1
			t2 = r[3] << 63
			r[3] >>= 1
			t = t2
			t2 = r[2] << 63
			r[2] = (r[2] >> 1) | t
			t = t2
			t2 = r[1] << 63
			r[1] = (r[1] >> 1) | t
			t = t2
			r[0] = (r[0] >> 1) | t

		}

		// v >= u
		bigger = !(v[3] < u[3] || (v[3] == u[3] && (v[2] < u[2] || (v[2] == u[2] && (v[1] < u[1] || (v[1] == u[1] && (v[0] < u[0])))))))

		if bigger {

			// v = v - u
			v[0], borrow = bits.Sub64(v[0], u[0], 0)
			v[1], borrow = bits.Sub64(v[1], u[1], borrow)
			v[2], borrow = bits.Sub64(v[2], u[2], borrow)
			v[3], _ = bits.Sub64(v[3], u[3], borrow)

			// r >= s
			bigger = !(r[3] < s[3] || (r[3] == s[3] && (r[2] < s[2] || (r[2] == s[2] && (r[1] < s[1] || (r[1] == s[1] && (r[0] < s[0])))))))

			if bigger {

				// s = s + q
				s[0], carry = bits.Add64(s[0], 2508899587657533823, 0)
				s[1], carry = bits.Add64(s[1], 13286502731930980180, carry)
				s[2], carry = bits.Add64(s[2], 4089299470098948802, carry)
				s[3], _ = bits.Add64(s[3], 8294956651982310308, carry)

			}

			// s = s - r
			s[0], borrow = bits.Sub64(s[0], r[0], 0)
			s[1], borrow = bits.Sub64(s[1], r[1], borrow)
			s[2], borrow = bits.Sub64(s[2], r[2], borrow)
			s[3], _ = bits.Sub64(s[3], r[3], borrow)

		} else {

			// u = u - v
			u[0], borrow = bits.Sub64(u[0], v[0], 0)
			u[1], borrow = bits.Sub64(u[1], v[1], borrow)
			u[2], borrow = bits.Sub64(u[2], v[2], borrow)
			u[3], _ = bits.Sub64(u[3], v[3], borrow)

			// s >= r
			bigger = !(s[3] < r[3] || (s[3] == r[3] && (s[2] < r[2] || (s[2] == r[2] && (s[1] < r[1] || (s[1] == r[1] && (s[0] < r[0])))))))

			if bigger {

				// r = r + q
				r[0], carry = bits.Add64(r[0], 2508899587657533823, 0)
				r[1], carry = bits.Add64(r[1], 13286502731930980180, carry)
				r[2], carry = bits.Add64(r[2], 4089299470098948802, carry)
				r[3], _ = bits.Add64(r[3], 8294956651982310308, carry)

			}

			// r = r - s
			r[0], borrow = bits.Sub64(r[0], s[0], 0)
			r[1], borrow = bits.Sub64(r[1], s[1], borrow)
			r[2], borrow = bits.Sub64(r[2], s[2], borrow)
			r[3], _ = bits.Sub64(r[3], s[3], borrow)

		}
		uIsOne = (u[0] == 1) && (u[3]|u[2]|u[1]) == 0
		vIsOne = (v[0] == 1) && (v[3]|v[2]|v[1]) == 0
	}

	if uIsOne {
		z.Set(&r)
	} else {
		z.Set(&s)
	}

	return z
}

// SetRandom sets z to a random element < q
func (z *Element04) SetRandom() *Element04 {
	bytes := make([]byte, 256)
	io.ReadFull(rand.Reader, bytes)
	z[0] = binary.BigEndian.Uint64(bytes[0:64])
	z[1] = binary.BigEndian.Uint64(bytes[64:128])
	z[2] = binary.BigEndian.Uint64(bytes[128:192])
	z[3] = binary.BigEndian.Uint64(bytes[192:256])
	z[3] %= 8294956651982310308

	// if z > q --> z -= q
	if !(z[3] < 8294956651982310308 || (z[3] == 8294956651982310308 && (z[2] < 4089299470098948802 || (z[2] == 4089299470098948802 && (z[1] < 13286502731930980180 || (z[1] == 13286502731930980180 && (z[0] < 2508899587657533823))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 2508899587657533823, 0)
		z[1], b = bits.Sub64(z[1], 13286502731930980180, b)
		z[2], b = bits.Sub64(z[2], 4089299470098948802, b)
		z[3], _ = bits.Sub64(z[3], 8294956651982310308, b)
	}

	return z
}

// Add z = x + y mod q
func (z *Element04) Add(x, y *Element04) *Element04 {
	var carry uint64

	z[0], carry = bits.Add64(x[0], y[0], 0)
	z[1], carry = bits.Add64(x[1], y[1], carry)
	z[2], carry = bits.Add64(x[2], y[2], carry)
	z[3], _ = bits.Add64(x[3], y[3], carry)

	// if z > q --> z -= q
	if !(z[3] < 8294956651982310308 || (z[3] == 8294956651982310308 && (z[2] < 4089299470098948802 || (z[2] == 4089299470098948802 && (z[1] < 13286502731930980180 || (z[1] == 13286502731930980180 && (z[0] < 2508899587657533823))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 2508899587657533823, 0)
		z[1], b = bits.Sub64(z[1], 13286502731930980180, b)
		z[2], b = bits.Sub64(z[2], 4089299470098948802, b)
		z[3], _ = bits.Sub64(z[3], 8294956651982310308, b)
	}
	return z
}

// AddAssign z = z + x mod q
func (z *Element04) AddAssign(x *Element04) *Element04 {
	var carry uint64

	z[0], carry = bits.Add64(z[0], x[0], 0)
	z[1], carry = bits.Add64(z[1], x[1], carry)
	z[2], carry = bits.Add64(z[2], x[2], carry)
	z[3], _ = bits.Add64(z[3], x[3], carry)

	// if z > q --> z -= q
	if !(z[3] < 8294956651982310308 || (z[3] == 8294956651982310308 && (z[2] < 4089299470098948802 || (z[2] == 4089299470098948802 && (z[1] < 13286502731930980180 || (z[1] == 13286502731930980180 && (z[0] < 2508899587657533823))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 2508899587657533823, 0)
		z[1], b = bits.Sub64(z[1], 13286502731930980180, b)
		z[2], b = bits.Sub64(z[2], 4089299470098948802, b)
		z[3], _ = bits.Sub64(z[3], 8294956651982310308, b)
	}
	return z
}

// Double z = x + x mod q, aka Lsh 1
func (z *Element04) Double(x *Element04) *Element04 {
	var carry uint64

	z[0], carry = bits.Add64(x[0], x[0], 0)
	z[1], carry = bits.Add64(x[1], x[1], carry)
	z[2], carry = bits.Add64(x[2], x[2], carry)
	z[3], _ = bits.Add64(x[3], x[3], carry)

	// if z > q --> z -= q
	if !(z[3] < 8294956651982310308 || (z[3] == 8294956651982310308 && (z[2] < 4089299470098948802 || (z[2] == 4089299470098948802 && (z[1] < 13286502731930980180 || (z[1] == 13286502731930980180 && (z[0] < 2508899587657533823))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 2508899587657533823, 0)
		z[1], b = bits.Sub64(z[1], 13286502731930980180, b)
		z[2], b = bits.Sub64(z[2], 4089299470098948802, b)
		z[3], _ = bits.Sub64(z[3], 8294956651982310308, b)
	}
	return z
}

// Sub  z = x - y mod q
func (z *Element04) Sub(x, y *Element04) *Element04 {
	var b uint64
	z[0], b = bits.Sub64(x[0], y[0], 0)
	z[1], b = bits.Sub64(x[1], y[1], b)
	z[2], b = bits.Sub64(x[2], y[2], b)
	z[3], b = bits.Sub64(x[3], y[3], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 2508899587657533823, 0)
		z[1], c = bits.Add64(z[1], 13286502731930980180, c)
		z[2], c = bits.Add64(z[2], 4089299470098948802, c)
		z[3], _ = bits.Add64(z[3], 8294956651982310308, c)
	}
	return z
}

// SubAssign  z = z - x mod q
func (z *Element04) SubAssign(x *Element04) *Element04 {
	var b uint64
	z[0], b = bits.Sub64(z[0], x[0], 0)
	z[1], b = bits.Sub64(z[1], x[1], b)
	z[2], b = bits.Sub64(z[2], x[2], b)
	z[3], b = bits.Sub64(z[3], x[3], b)
	if b != 0 {
		var c uint64
		z[0], c = bits.Add64(z[0], 2508899587657533823, 0)
		z[1], c = bits.Add64(z[1], 13286502731930980180, c)
		z[2], c = bits.Add64(z[2], 4089299470098948802, c)
		z[3], _ = bits.Add64(z[3], 8294956651982310308, c)
	}
	return z
}

// Exp z = x^e mod q
func (z *Element04) Exp(x Element04, e uint64) *Element04 {
	if e == 0 {
		return z.SetOne()
	}

	z.Set(&x)

	l := bits.Len64(e) - 2
	for i := l; i >= 0; i-- {
		z.Square(z)
		if e&(1<<uint(i)) != 0 {
			z.MulAssign(&x)
		}
	}
	return z
}

// FromMont converts z in place (i.e. mutates) from Montgomery to regular representation
// sets and returns z = z * 1
func (z *Element04) FromMont() *Element04 {

	// the following lines implement z = z * 1
	// with a modified CIOS montgomery multiplication
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 10821694688364185985
		C := madd0(m, 2508899587657533823, z[0])
		C, z[0] = madd2(m, 13286502731930980180, z[1], C)
		C, z[1] = madd2(m, 4089299470098948802, z[2], C)
		C, z[2] = madd2(m, 8294956651982310308, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 10821694688364185985
		C := madd0(m, 2508899587657533823, z[0])
		C, z[0] = madd2(m, 13286502731930980180, z[1], C)
		C, z[1] = madd2(m, 4089299470098948802, z[2], C)
		C, z[2] = madd2(m, 8294956651982310308, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 10821694688364185985
		C := madd0(m, 2508899587657533823, z[0])
		C, z[0] = madd2(m, 13286502731930980180, z[1], C)
		C, z[1] = madd2(m, 4089299470098948802, z[2], C)
		C, z[2] = madd2(m, 8294956651982310308, z[3], C)
		z[3] = C
	}
	{
		// m = z[0]n'[0] mod W
		m := z[0] * 10821694688364185985
		C := madd0(m, 2508899587657533823, z[0])
		C, z[0] = madd2(m, 13286502731930980180, z[1], C)
		C, z[1] = madd2(m, 4089299470098948802, z[2], C)
		C, z[2] = madd2(m, 8294956651982310308, z[3], C)
		z[3] = C
	}

	// if z > q --> z -= q
	if !(z[3] < 8294956651982310308 || (z[3] == 8294956651982310308 && (z[2] < 4089299470098948802 || (z[2] == 4089299470098948802 && (z[1] < 13286502731930980180 || (z[1] == 13286502731930980180 && (z[0] < 2508899587657533823))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 2508899587657533823, 0)
		z[1], b = bits.Sub64(z[1], 13286502731930980180, b)
		z[2], b = bits.Sub64(z[2], 4089299470098948802, b)
		z[3], _ = bits.Sub64(z[3], 8294956651982310308, b)
	}
	return z
}

// ToMont converts z to Montgomery form
// sets and returns z = z * r^2
func (z *Element04) ToMont() *Element04 {
	var rSquare = Element04{
		11987161547385592410,
		17714078913370961707,
		17074890819590993293,
		2162564708293599491,
	}
	return z.MulAssign(&rSquare)
}

// ToRegular returns z in regular form (doesn't mutate z)
func (z Element04) ToRegular() Element04 {
	return *z.FromMont()
}

// String returns the string form of an Element04 in Montgomery form
func (z *Element04) String() string {
	var _z big.Int
	return z.ToBigIntRegular(&_z).String()
}

// ToBigInt returns z as a big.Int in Montgomery form
func (z *Element04) ToBigInt(res *big.Int) *big.Int {
	bits := (*[4]big.Word)(unsafe.Pointer(z))
	return res.SetBits(bits[:])
}

// ToBigIntRegular returns z as a big.Int in regular form
func (z Element04) ToBigIntRegular(res *big.Int) *big.Int {
	z.FromMont()
	bits := (*[4]big.Word)(unsafe.Pointer(&z))
	return res.SetBits(bits[:])
}

// SetBigInt sets z to v (regular form) and returns z in Montgomery form
func (z *Element04) SetBigInt(v *big.Int) *Element04 {
	z.SetZero()

	zero := big.NewInt(0)
	q := element04ModulusBigInt()

	// copy input
	vv := new(big.Int).Set(v)

	// while v < 0, v+=q
	for vv.Cmp(zero) == -1 {
		vv.Add(vv, q)
	}
	// while v > q, v-=q
	for vv.Cmp(q) == 1 {
		vv.Sub(vv, q)
	}
	// if v == q, return 0
	if vv.Cmp(q) == 0 {
		return z
	}
	// v should
	vBits := vv.Bits()
	for i := 0; i < len(vBits); i++ {
		z[i] = uint64(vBits[i])
	}
	return z.ToMont()
}

// SetString creates a big.Int with s (in base 10) and calls SetBigInt on z
func (z *Element04) SetString(s string) *Element04 {
	x, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("Element04.SetString failed -> can't parse number in base10 into a big.Int")
	}
	return z.SetBigInt(x)
}

// Mul z = x * y mod q
func (z *Element04) Mul(x, y *Element04) *Element04 {

	var t [4]uint64
	var c [3]uint64
	{
		// round 0
		v := x[0]
		c[1], c[0] = bits.Mul64(v, y[0])
		m := c[0] * 10821694688364185985
		c[2] = madd0(m, 2508899587657533823, c[0])
		c[1], c[0] = madd1(v, y[1], c[1])
		c[2], t[0] = madd2(m, 13286502731930980180, c[2], c[0])
		c[1], c[0] = madd1(v, y[2], c[1])
		c[2], t[1] = madd2(m, 4089299470098948802, c[2], c[0])
		c[1], c[0] = madd1(v, y[3], c[1])
		t[3], t[2] = madd3(m, 8294956651982310308, c[0], c[2], c[1])
	}
	{
		// round 1
		v := x[1]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 10821694688364185985
		c[2] = madd0(m, 2508899587657533823, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13286502731930980180, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 4089299470098948802, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		t[3], t[2] = madd3(m, 8294956651982310308, c[0], c[2], c[1])
	}
	{
		// round 2
		v := x[2]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 10821694688364185985
		c[2] = madd0(m, 2508899587657533823, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13286502731930980180, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], t[1] = madd2(m, 4089299470098948802, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		t[3], t[2] = madd3(m, 8294956651982310308, c[0], c[2], c[1])
	}
	{
		// round 3
		v := x[3]
		c[1], c[0] = madd1(v, y[0], t[0])
		m := c[0] * 10821694688364185985
		c[2] = madd0(m, 2508899587657533823, c[0])
		c[1], c[0] = madd2(v, y[1], c[1], t[1])
		c[2], z[0] = madd2(m, 13286502731930980180, c[2], c[0])
		c[1], c[0] = madd2(v, y[2], c[1], t[2])
		c[2], z[1] = madd2(m, 4089299470098948802, c[2], c[0])
		c[1], c[0] = madd2(v, y[3], c[1], t[3])
		z[3], z[2] = madd3(m, 8294956651982310308, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	if !(z[3] < 8294956651982310308 || (z[3] == 8294956651982310308 && (z[2] < 4089299470098948802 || (z[2] == 4089299470098948802 && (z[1] < 13286502731930980180 || (z[1] == 13286502731930980180 && (z[0] < 2508899587657533823))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 2508899587657533823, 0)
		z[1], b = bits.Sub64(z[1], 13286502731930980180, b)
		z[2], b = bits.Sub64(z[2], 4089299470098948802, b)
		z[3], _ = bits.Sub64(z[3], 8294956651982310308, b)
	}
	return z
}

// MulAssign z = z * x mod q
func (z *Element04) MulAssign(x *Element04) *Element04 {

	var t [4]uint64
	var c [3]uint64
	{
		// round 0
		v := z[0]
		c[1], c[0] = bits.Mul64(v, x[0])
		m := c[0] * 10821694688364185985
		c[2] = madd0(m, 2508899587657533823, c[0])
		c[1], c[0] = madd1(v, x[1], c[1])
		c[2], t[0] = madd2(m, 13286502731930980180, c[2], c[0])
		c[1], c[0] = madd1(v, x[2], c[1])
		c[2], t[1] = madd2(m, 4089299470098948802, c[2], c[0])
		c[1], c[0] = madd1(v, x[3], c[1])
		t[3], t[2] = madd3(m, 8294956651982310308, c[0], c[2], c[1])
	}
	{
		// round 1
		v := z[1]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 10821694688364185985
		c[2] = madd0(m, 2508899587657533823, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13286502731930980180, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 4089299470098948802, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		t[3], t[2] = madd3(m, 8294956651982310308, c[0], c[2], c[1])
	}
	{
		// round 2
		v := z[2]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 10821694688364185985
		c[2] = madd0(m, 2508899587657533823, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], t[0] = madd2(m, 13286502731930980180, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], t[1] = madd2(m, 4089299470098948802, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		t[3], t[2] = madd3(m, 8294956651982310308, c[0], c[2], c[1])
	}
	{
		// round 3
		v := z[3]
		c[1], c[0] = madd1(v, x[0], t[0])
		m := c[0] * 10821694688364185985
		c[2] = madd0(m, 2508899587657533823, c[0])
		c[1], c[0] = madd2(v, x[1], c[1], t[1])
		c[2], z[0] = madd2(m, 13286502731930980180, c[2], c[0])
		c[1], c[0] = madd2(v, x[2], c[1], t[2])
		c[2], z[1] = madd2(m, 4089299470098948802, c[2], c[0])
		c[1], c[0] = madd2(v, x[3], c[1], t[3])
		z[3], z[2] = madd3(m, 8294956651982310308, c[0], c[2], c[1])
	}

	// if z > q --> z -= q
	if !(z[3] < 8294956651982310308 || (z[3] == 8294956651982310308 && (z[2] < 4089299470098948802 || (z[2] == 4089299470098948802 && (z[1] < 13286502731930980180 || (z[1] == 13286502731930980180 && (z[0] < 2508899587657533823))))))) {
		var b uint64
		z[0], b = bits.Sub64(z[0], 2508899587657533823, 0)
		z[1], b = bits.Sub64(z[1], 13286502731930980180, b)
		z[2], b = bits.Sub64(z[2], 4089299470098948802, b)
		z[3], _ = bits.Sub64(z[3], 8294956651982310308, b)
	}
	return z
}

// Square z = x * x mod q
func (z *Element04) Square(x *Element04) *Element04 {
	return z.Mul(x, x)
}
