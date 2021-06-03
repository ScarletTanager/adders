package adders

type Digit uint8

const (
	Zero Digit = iota
	One
)

func (d Digit) Boolean() bool {
	if d == One {
		return true
	}

	return false
}

func OR(p, q Digit) Digit {
	if p.Boolean() || q.Boolean() {
		return One
	}
	return Zero
}

func AND(p, q Digit) Digit {
	if p.Boolean() && q.Boolean() {
		return One
	}
	return Zero
}

func NOT(p Digit) Digit {
	return 3 % (p + 2)
}

func HalfAdder(p, q Digit) (carry, sum Digit) {
	carry = AND(p, q)
	sum = AND(OR(p, q), NOT(carry))

	return carry, sum
}

func FullAdder(p, q, r Digit) (carry, sum Digit) {
	c1, s1 := HalfAdder(p, q)
	c2, sum := HalfAdder(r, s1)
	carry = OR(c1, c2)

	return carry, sum
}
