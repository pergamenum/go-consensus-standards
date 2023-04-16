package reflection

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Alpha struct {
	S string    `automap:"s"`
	I int       `automap:"i"`
	B bool      `automap:"b"`
	T time.Time `automap:"t"`
	N *Alpha    `automap:"n"`
}

type Foo struct {
	S string    `automap:"s"`
	I int       `automap:"i"`
	B bool      `automap:"b"`
	T time.Time `automap:"t"`
	N *Foo      `automap:"n"`
}

func manualMap(input Alpha) Foo {

	f := Foo{
		S: input.S,
		I: input.I,
		B: input.B,
		T: input.T,
	}

	if input.N != nil {
		temp := manualMap(*input.N)
		f.N = &temp
	}

	return f
}

func newAlpha(depth int) Alpha {

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randSeq := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	a := Alpha{
		S: randSeq(10),
		I: rand.Intn(1337),
		B: rand.Intn(2) == 0,
		T: time.Now(),
	}

	if (depth - 1) > 0 {
		temp := newAlpha(depth - 1)
		a.N = &temp
	}

	return a
}

func generateInputSet(length int, depth int) []Alpha {

	var as []Alpha

	for i := 0; i < length; i++ {

		a := newAlpha(depth)
		as = append(as, a)
	}

	return as
}

func Benchmark_Mapping(b *testing.B) {

	rand.Seed(time.Now().UnixNano())

	l := 1000
	d := 10

	as := generateInputSet(l, d)

	s := fmt.Sprintf("manualMap_Single_L[%v]_D[%v]", l, d)
	b.Run(s, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			manualMap(as[0])
		}
	})
	s = fmt.Sprintf("AutoMap_Single_L[%v]_D[%v]", l, d)
	b.Run(s, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = AutoMap[Foo](as[0])
		}
	})

	s = fmt.Sprintf("manualMap_Multiple_L[%v]_D[%v]", l, d)
	b.Run(s, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, a := range as {
				manualMap(a)
			}
		}
	})

	s = fmt.Sprintf("AutoMapping_Multiple_L[%v]_D[%v]", l, d)
	b.Run(s, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, a := range as {
				_, _ = AutoMap[Foo](a)
			}
		}
	})

	d = 0
	as = generateInputSet(l, d)

	s = fmt.Sprintf("manualMap_Single_L[%v]_D[%v]", l, d)
	b.Run(s, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			manualMap(as[0])
		}
	})
	s = fmt.Sprintf("AutoMap_Single_L[%v]_D[%v]", l, d)
	b.Run(s, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = AutoMap[Foo](as[0])
		}
	})

	s = fmt.Sprintf("manualMap_Multiple_L[%v]_D[%v]", l, d)
	b.Run(s, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, a := range as {
				manualMap(a)
			}
		}
	})

	s = fmt.Sprintf("AutoMapping_Multiple_L[%v]_D[%v]", l, d)
	b.Run(s, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, a := range as {
				_, _ = AutoMap[Foo](a)
			}
		}
	})
}
