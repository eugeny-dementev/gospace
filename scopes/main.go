package main

import "fmt"

type Entity struct {
	prop string
}

func (e *Entity) String() string {
	return fmt.Sprintf("String from entity: %v", e.prop)
}

func main() {
	e := Entity{"Hello world"}
	fmt.Printf("Entity: %v\n", e)

  fmt.Println(ReturnFuncString(e.String))
}

func ReturnFuncString(stringer func() string) string {
  return stringer()
}
