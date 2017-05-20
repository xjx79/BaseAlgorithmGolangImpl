package pub

import (
	"fmt"
	"testing"
)

func Test_insertSort_1(t *testing.T) {
	var a []int = []int{0, 20, 10, 25, 15, 30, 28, 55, 432, 432432, 4234, 333,
		333, 21, 22, 3, 30, 8, 20, 2, 7, 9, 50, 80}
	fmt.Println(a)
	b := insertSort(a)
	fmt.Println(b)
}

func Test_insertSort_2(t *testing.T) {
	var a []int = []int{0, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 10}
	fmt.Println(a)
	b := insertSort(a)
	fmt.Println(b)
}

func Test_insertSort_3(t *testing.T) {
	var a []int = []int{20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
		20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20}
	fmt.Println(a)
	b := insertSort(a)
	fmt.Println(b)
}

func Test_insertSort_4(t *testing.T) {
	var a []int = []int{0}
	fmt.Println(a)
	b := insertSort(a)
	fmt.Println(b)
}

func Test_insertSort_5(t *testing.T) {
	var a []int = []int{}
	fmt.Println(a)
	b := insertSort(a)
	fmt.Println(b)
}
