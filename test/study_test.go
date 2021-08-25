package test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
	"time"
)

func TestChanSelect(t *testing.T)  {
	ch := make(chan int)
	go func(){
		for {
			select {
			case x, ok := <-ch:
				fmt.Println("select...", x, ok)
			case <-time.After(time.Second):
				fmt.Println("sleep...")
			}
		}
	}()
	time.Sleep(2*time.Second)
}

func TestRangeAppend(t *testing.T)  {
	var arr []int
	for i:=0; i<10; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	for _, i := range arr {
		arr = append(arr, i)
	}
	fmt.Println(arr)
}

func TestAssert(t *testing.T)  {
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
	fmt.Println(1<<4)
	fmt.Printf("%x\n", 16)
}

func TestMapIter(t *testing.T) {
	m := make(map[int]int, 100)
	ch := make(chan int)
	for i:=0; i<100; i++ {
		m[i] = i
	}
	go func() {
		for k, v := range m {
			fmt.Println(k,":", v)
		}
		ch <- 1
	}()
	time.Sleep(50)
	delete(m, 20)
	fmt.Println("delete 20")
	<- ch
}
