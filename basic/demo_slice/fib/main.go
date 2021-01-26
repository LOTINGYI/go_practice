package main


import "fmt"


func fib(n int) ([]uint64) {
	fbnSlice := make([]uint64,n)
	fbnSlice[0] = 1
	fbnSlice[1] =1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1]+fbnSlice[i-2]
	}
	return fbnSlice
}
func main() {
	fbnSlice := fib(20)
	fmt.Println(fbnSlice)
}