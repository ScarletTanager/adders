package adders_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ScarletTanager/adders"
)

var _ = Describe("Adders", func() {
	Describe("Digit", func() {
		var (
			p Digit
		)

		Describe("Boolean", func() {
			It("Returns false", func() {
				Expect(p.Boolean()).To(BeFalse())
			})

			Context("When the digit is a 1", func() {
				BeforeEach(func() {
					p = One
				})

				It("Returns true", func() {
					Expect(p.Boolean()).To(BeTrue())
				})
			})
		})
	})

	Describe("NOT", func() {
		var (
			p Digit
		)

		It("Returns 1", func() {
			Expect(NOT(p)).To(Equal(One))
		})

		Context("When the digit is 1", func() {
			BeforeEach(func() {
				p = One
			})

			It("Returns 0", func() {
				Expect(NOT(p)).To(Equal(Zero))
			})
		})
	})

	Describe("OR", func() {
		var (
			p, q Digit
		)
		It("Returns 0", func() {
			Expect(OR(p, q)).To(Equal(Zero))
		})

		Context("When one digit is a 1", func() {
			BeforeEach(func() {
				p = One
			})

			It("Returns 1", func() {
				Expect(OR(p, q)).To(Equal(One))
			})

			Context("When both digits are 1s", func() {
				BeforeEach(func() {
					q = One
				})

				It("Returns 1", func() {
					Expect(OR(p, q)).To(Equal(One))
				})
			})
		})
	})

	Describe("AND", func() {
		var (
			p, q Digit
		)

		It("Returns a 0", func() {
			Expect(AND(p, q)).To(Equal(Zero))
		})

		Context("When one digit is a 1", func() {
			BeforeEach(func() {
				p = One
			})

			It("Returns a 0", func() {
				Expect(AND(p, q)).To(Equal(Zero))
			})

			Context("When both digits are 1s", func() {
				BeforeEach(func() {
					q = One
				})

				It("Returns a 1", func() {
					Expect(AND(p, q)).To(Equal(One))
				})
			})
		})
	})

	Describe("HalfAdder", func() {
		var (
			p, q Digit
		)

		It("Returns a carry and sum of 0 (binary 00)", func() {
			carry, sum := HalfAdder(p, q)
			Expect(carry).To(Equal(Zero))
			Expect(sum).To(Equal(Zero))
		})

		Context("When one digit is a 1", func() {
			BeforeEach(func() {
				p = One
			})

			It("Returns a carry of 0, a sum of 1 (binary 01)", func() {
				carry, sum := HalfAdder(p, q)
				Expect(carry).To(Equal(Zero))
				Expect(sum).To(Equal(One))
			})

			Context("When both digits are 1", func() {
				BeforeEach(func() {
					q = One
				})

				It("Returns a carry of 1, a sum of 0 (binary 10)", func() {
					carry, sum := HalfAdder(p, q)
					Expect(carry).To(Equal(One))
					Expect(sum).To(Equal(Zero))
				})
			})
		})
	})

	Describe("FullAdder", func() {
		var (
			p, q, r Digit
		)

		It("Returns a carry and sum both set to 0 (binary 00)", func() {
			carry, sum := FullAdder(p, q, r)
			Expect(carry).To(Equal(Zero))
			Expect(sum).To(Equal(Zero))
		})

		Context("When one digit is a 1", func() {
			BeforeEach(func() {
				p = One
			})

			It("Returns a carry of 0 and a sum of 1 (binary 01)", func() {
				carry, sum := FullAdder(p, q, r)
				Expect(carry).To(Equal(Zero))
				Expect(sum).To(Equal(One))
			})

			Context("When two digits are 1s", func() {
				BeforeEach(func() {
					q = One
				})

				It("Returns a carry of 1 and a sum of 0 (binary 10)", func() {
					carry, sum := FullAdder(p, q, r)
					Expect(carry).To(Equal(One))
					Expect(sum).To(Equal(Zero))
				})

				Context("When all three digits are 1s", func() {
					BeforeEach(func() {
						r = One
					})

					It("Returns a carry and sum both set to 1 (binary 11)", func() {
						carry, sum := FullAdder(p, q, r)
						Expect(carry).To(Equal(One))
						Expect(sum).To(Equal(One))
					})
				})
			})
		})
	})
})
