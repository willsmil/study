package test

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func TestSetFunc(t *testing.T) {
	fmt.Println(strings.Index("abc", "b"))
}

func TestSplit(t *testing.T) {
	s := " hello  word"
	ss := strings.Split(s, " ")
	t.Log(ss)
}

var a, b int

func TestGoroutine(t *testing.T) {
	go func() {
		a++
		fmt.Println("a2:", a)
	}()
	// time.Sleep(10)
	fmt.Println("a1:", a)
	b = a + 1
	fmt.Println("b:", b)
	time.Sleep(time.Second)
}

type S struct {
	A int
	B string
}

func (s *S) Set(a int) {
	s.A = a
}

var SPoll = sync.Pool{
	New: func() interface{} {
		return new(S)
	},
}

func TestPoll(t *testing.T) {
	s := SPoll.Get().(*S)
	fmt.Println(s)
	s.A = 1
	s.B = "xx"
	fmt.Println(s)
}

func TestPtr(t *testing.T) {
	type S struct {
		a int
		b string
	}
	a := new(S)
	fmt.Println(&a, &a.a, &a.b)
	pa := unsafe.Pointer(uintptr(unsafe.Pointer(a)) + unsafe.Offsetof(a.a))
	pb := unsafe.Pointer(uintptr(unsafe.Pointer(a)) + unsafe.Offsetof(a.b))
	fmt.Println(unsafe.Offsetof(a.a), unsafe.Offsetof(a.b))
	fmt.Println(pa, pb)
	*(*int)(pa) = 1
	*(*string)(pb) = "aaa"
	fmt.Println(a)
}

func TestChanSelect(t *testing.T) {
	ch := make(chan int)
	go func() {
		for {
			select {
			case x, ok := <-ch:
				fmt.Println("select...", x, ok)
			case <-time.After(time.Second):
				fmt.Println("sleep...")
			}
		}
	}()
	time.Sleep(2 * time.Second)
}

func TestRangeAppend(t *testing.T) {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	for _, i := range arr {
		arr = append(arr, i)
	}
	fmt.Println(arr)
}

func TestAssert(t *testing.T) {
	var s interface{} = "sss"
	// x, ok := s.(int)
	fmt.Println(s.(string))
}

func TestStr2ByteSlice(t *testing.T) {
	a := "aaa"
	b := []byte(a)
	fmt.Println(a, []byte(a))
	fmt.Println(&a, &b)

	addr := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	byt := *(*[]byte)(unsafe.Pointer(&addr))
	fmt.Println(byt)
}

func TestSlice(t *testing.T) {
	s := make([]int, 2, 10)
	fmt.Println(s)
	appendSlice(s, 2, 3)
	fmt.Println(s)
	var x []int
	x = append(x, 1)
	x = nil
	x = append(x, 2)
	fmt.Println(x)
}
func appendSlice(s []int, x ...int) []int {
	return append(s, x...)
}

func TestString(t *testing.T) {
	s := "123456789"
	fmt.Println(&s)
	s = "12312312312312"
	fmt.Println(&s)
}

func TestSizeOf(t *testing.T) {
	a := [0]int{}
	b := [0]float64{}
	fmt.Printf("%p: %d, %p: %d, %v, %v\n", &a, unsafe.Sizeof(a), &b, unsafe.Sizeof(b), reflect.TypeOf(a), reflect.TypeOf(b))
}

func TestMapSpace(t *testing.T) {
	s := make([]int, 0, 10)
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	fmt.Println(len(s), cap(s))
	s = append(s, 2)
	fmt.Println(len(s), cap(s))
}

func TestMapOperation(t *testing.T) {
	m1 := make(map[int]int, 1)
	fmt.Println(m1)
	appendMap(m1, 1, 1)
	appendMap(m1, 2, 2)
	appendMap(m1, 3, 3)
	fmt.Println(m1)
}
func appendMap(m map[int]int, key, value int) {
	m[key] = value
}

func TestBitOperation(t *testing.T) {
	fmt.Println(1 << 4)
	fmt.Printf("%x\n", 16)
}

func TestMapIter(t *testing.T) {
	m := make(map[int]int, 100)
	ch := make(chan int)
	for i := 0; i < 100; i++ {
		m[i] = i
	}
	go func() {
		for k, v := range m {
			fmt.Println(k, ":", v)
		}
		ch <- 1
	}()
	time.Sleep(50)
	delete(m, 20)
	fmt.Println("delete 20")
	<-ch
}
