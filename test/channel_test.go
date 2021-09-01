package test

import (
	"fmt"
	"testing"
)

func TestChannelClose(t *testing.T)  {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	v, ok := <- ch
	fmt.Println("receive from closed channel:", v, ok)
}