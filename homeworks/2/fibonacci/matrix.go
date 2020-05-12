package fibonacci

import (
	"fmt"
	"math/big"
)

// Matrix2x2Uint is 2x2 matrix of uint64 numbers.
type Matrix2x2Uint [4]uint64

func (m *Matrix2x2Uint) Multiply(w Matrix2x2Uint) {
	temp := m[0]*w[0] + m[1]*w[2]
	m[1] = m[0]*w[1] + m[1]*w[3]
	m[0] = temp

	temp = m[2]*w[0] + m[3]*w[2]
	m[3] = m[2]*w[1] + m[3]*w[3]
	m[2] = temp
}

func (m *Matrix2x2Uint) Pow2() {
	isSymmetric := m[1] == m[2]

	if !isSymmetric {
		m.Multiply(*m)
		return
	}

	// optimization hack: cause matrix is symmetric x01 = x10
	x00 := m[0]*m[0] + m[1]*m[2]
	x01 := m[0]*m[1] + m[1]*m[3]

	// x11
	m[3] = m[2]*m[1] + m[3]*m[3]

	m[0] = x00
	m[1] = x01
	m[2] = x01
}

func (m *Matrix2x2Uint) Pow(n uint) {
	if n == 0 {
		m[0] = 1
		m[1] = 0
		m[2] = 0
		m[3] = 1

		return
	}

	if n == 1 {
		return
	}

	w := *m

	m[0] = 1
	m[1] = 0
	m[2] = 0
	m[3] = 1

	for n > 0 {
		// odd case
		if n&1 == 1 {
			n--

			// a^n = a * a^(n-1)
			m.Multiply(w)
		}

		// n now is even, divide by 2 for next step
		n >>= 1

		// a^n = a^(n/2) * a^(n/2) == (a^(n/2))^2
		w.Pow2()
	}
}

// Matrix2x2BigInt is 2x2 matrix of big int numbers.
type Matrix2x2BigInt [4]*big.Int

func (m *Matrix2x2BigInt) Multiply(w Matrix2x2BigInt) {
	m00 := new(big.Int)
	m01 := new(big.Int)
	m10 := new(big.Int)
	m11 := new(big.Int)

	temp1 := new(big.Int)
	temp2 := new(big.Int)

	temp1 = temp1.Mul(m[0], w[0])
	temp2 = temp2.Mul(m[1], w[2])

	m00 = m00.Add(temp1, temp2)

	temp1 = temp1.Mul(m[0], w[1])
	temp2 = temp2.Mul(m[1], w[3])

	m01 = m01.Add(temp1, temp2)

	temp1 = temp1.Mul(m[2], w[0])
	temp2 = temp2.Mul(m[3], w[2])

	m10 = m10.Add(temp1, temp2)

	temp1 = temp1.Mul(m[2], w[1])
	temp2 = temp2.Mul(m[3], w[3])

	m11 = m11.Add(temp1, temp2)

	m[0].Set(m00)
	m[1].Set(m01)
	m[2].Set(m10)
	m[3].Set(m11)
}

func (m *Matrix2x2BigInt) Pow2() {
	isSymmetric := m[1].Cmp(m[2]) == 0
	if !isSymmetric {
		m.Multiply(*m)
		return
	}

	// optimization hack: cause matrix is symmetric m01 = m10
	m00 := new(big.Int)
	m01 := new(big.Int)
	m11 := new(big.Int)

	temp1 := new(big.Int)
	temp2 := new(big.Int)

	temp1 = temp1.Mul(m[0], m[0])
	temp2 = temp2.Mul(m[1], m[2])

	m00 = m00.Add(temp1, temp2)

	temp1 = temp1.Mul(m[0], m[1])
	temp2 = temp2.Mul(m[1], m[3])

	m01 = m01.Add(temp1, temp2)

	temp1 = temp1.Mul(m[2], m[1])
	temp2 = temp2.Mul(m[3], m[3])

	m11 = m11.Add(temp1, temp2)

	m[0].Set(m00)
	m[1].Set(m01)
	m[2].Set(m01)
	m[3].Set(m11)
}

func (m *Matrix2x2BigInt) Equals(w *Matrix2x2BigInt) bool {
	for i := 0; i < 4; i++ {
		if m[i].Cmp(w[i]) != 0 {
			return false
		}
	}

	return true
}

func (m *Matrix2x2BigInt) Pow(n uint) {
	if n == 0 {
		m[0].Set(big.NewInt(1))
		m[1].Set(big.NewInt(0))
		m[2].Set(big.NewInt(0))
		m[3].Set(big.NewInt(1))

		return
	}

	if n == 1 {
		return
	}

	w := Matrix2x2BigInt{
		new(big.Int).Set(m[0]),
		new(big.Int).Set(m[1]),
		new(big.Int).Set(m[2]),
		new(big.Int).Set(m[3]),
	}

	m[0].Set(big.NewInt(1))
	m[1].Set(big.NewInt(0))
	m[2].Set(big.NewInt(0))
	m[3].Set(big.NewInt(1))

	for n > 0 {
		// odd case
		if n&1 == 1 {
			n--

			// a^n = a * a^(n-1)
			m.Multiply(w)
		}

		// n now is even, divide by 2 for next step
		n >>= 1

		// a^n = a^(n/2) * a^(n/2) = (a^(n/2))^2
		w.Pow2()
	}
}

func (m Matrix2x2BigInt) String() string {
	return fmt.Sprintf("%s %s %s %s", m[0].String(), m[1].String(), m[2].String(), m[3].String())
}
