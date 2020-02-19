package main

import "C"

import (
	"fmt"
	"math"
	"sort"
	"sync"

	"github.com/yuneejang/go-shared-object/implement"
)

//Go 코드를 공유 라이브러리로 컴파일하기 전에 수행해야 할 4 가지 요구 사항이 있는데요.
//1. 패키지는 반드시 main 패키지 여야 합니다. 컴파일러는 패키지와 모든 종속성을 단일 공유 객체 바이너리로 빌드합니다.
//2. 소스는 pseudo-package "C"를 import 해야 합니다.
//3. // export 주석을 사용하여 다른 언어에서 액세스 할 수 있도록 하려는 함수에 주석을 답니다.
// 함수와 동일한 이름을 적어주어야합니다.
//4. 빈 main 함수를 선언해야합니다.

//go build -o awesome.so -buildmode=c-shared awesome.go
var count int
var mtx sync.Mutex

//export Add
func Add(a, b int) int {
	return a + b
}

//export Sub
func Sub(a, b int) int {
	return implement.Sub(a, b)
}

//export Mul
func Mul(a, b int) int {
	return implement.Mul(a, b)
}

//export Cosine
func Cosine(x float64) float64 {
	return math.Cos(x)
}

//export Sort
func Sort(vals []int) {
	sort.Ints(vals)
}

//export SortPtr
func SortPtr(vals *[]int) {
	Sort(*vals)
}

//export Log
func Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

//export LogPtr
func LogPtr(msg *string) int {
	return Log(*msg)
}

func main() {}
