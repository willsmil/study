package wrapper

import (
	"fmt"
	"testing"
)

func Test_Wrapper(t *testing.T) {
	e := NewEspresso()
	fmt.Println(e.GetDescription(), "$", e.Cost())

	x := NewWhip(NewMocha(NewMocha(&e)))
	fmt.Println(x.GetDescription(), "$", x.Cost())
}
