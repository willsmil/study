package leetcode

import (
	"context"
	"fmt"
	"testing"
)

func TestParseToken(t *testing.T)  {
	s := GrayRuleService{}
	p := "1&!(4|(!4&3))"
	t.Log("pattern:", p)
	t.Log(s.parseRule2(context.Background(), 1, p))
}
func TestRuleLegal(t *testing.T)  {
	fmt.Println(ruleLegal("12&!(2|4)"))
}