package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	str string
	base
  b base
}

func structEmbExp() {
	cont := container{
		"inside container",
		base{2},
		base{5},
	}

	fmt.Println(cont.num, cont.base.num, cont.describe())
	fmt.Println(cont.b.num, cont.b.describe())

  type describer interface {
    describe() string
  }
  var c describer = cont
  fmt.Println(c.describe())
}
